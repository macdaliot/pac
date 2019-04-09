import { Request } from "express";
import { generateRandomString } from "./utilities";
import { Strategy as SamlStrategy } from "passport-saml";
import * as jwt from "jsonwebtoken";
import {
    Strategy as JwtStrategy,
    StrategyOptions as JwtStrategyOptions,
    ExtractJwt
} from "passport-jwt";


const jwtSecret = process.env.JWT_SECRET || generateRandomString();


const opts: JwtStrategyOptions = {
    jwtFromRequest: ExtractJwt.fromAuthHeaderAsBearerToken(),
    secretOrKey: process.env.JWT_SECRET || generateRandomString(),
    issuer: process.env.JWT_ISSUER || generateRandomString()
};

export const jwtStrategy = new JwtStrategy(opts, (jwt_payload, done) => {
    done(null, jwt_payload);
});

const signingCert = process.env.SAML_SIGNER
    ? Buffer.from(process.env.SAML_SIGNER, "base64").toString("ascii")
    : generateRandomString();

export const samlStrategy = new SamlStrategy(
    {
        callbackUrl:
            process.env.SAML_CALLBACK || `http://localhost/${generateRandomString()}`,
        entryPoint:
            process.env.AUTH0_CLIENT_ID && process.env.AUTH0_DOMAIN
                ? `https://${process.env.AUTH0_DOMAIN}/samlp/${
                    process.env.AUTH0_CLIENT_ID
                    }`
                : `https://localhost/samlp/${generateRandomString()}`,
        issuer: process.env.AUTH0_DOMAIN
            ? "urn:" + process.env.AUTH0_DOMAIN + "-saml"
            : `urn:${generateRandomString()}`,
        cert: signingCert,
        acceptedClockSkewMs: 60000
    },
    function(user, done) {
        return done(null, user, null);
    }
);

export const createJwt = (req: Request, callback: jwt.SignCallback): void => {
    let groups = [];
    let groupSource = req.user["http://schemas.xmlsoap.org/claims/Group"];
    if (groupSource && Array.isArray(groupSource)) {
        for (let groupname of groupSource) {
            groups.push(groupname);
        }
    }
    let token = {
        name: req.user["http://schemas.auth0.com/nickname"],
        sub: req.user["http://schemas.xmlsoap.org/ws/2005/05/identity/claims/upn"],
        groups: groups
    };
    const options = {
        expiresIn: "10h",
        issuer: "urn:pacAuth"
    };
    jwt.sign(token, jwtSecret, options, callback);
};
