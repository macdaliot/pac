import { Router, Request } from 'express';
import * as passport from 'passport';
import * as jwt from 'jsonwebtoken';
export let authRouter = Router();

const jwtSecret = process.env.JWT_SECRET || "foobar";

const createJwt = (req: Request, callback: jwt.SignCallback): void => {
    let groups = [];
    let groupSource = req.user["http://schemas.xmlsoap.org/claims/Group"];
    if (groupSource && Array.isArray(groupSource)){
        for (let groupname of groupSource){
            groups.push(groupname);
        }
    }
    let token = {
        "name": req.user["http://schemas.auth0.com/nickname"],
        "sub": req.user["http://schemas.xmlsoap.org/ws/2005/05/identity/claims/upn"],
        "groups": groups
    };
    const options = {
        expiresIn: "10h",
        issuer: "urn:pacAuth" // APPEND SERVICE NAME HERE
    }
    jwt.sign(token, jwtSecret, options, callback);
}

// Perform the final stage of authentication and redirect to previously requested URL or '/user'
authRouter.post('/callback', 
    passport.authenticate('saml'),
    (req, res) => {
        createJwt(req, (err, encoded) => {
            if (err){
                res.status(500).send(err);
            }
            else {

                res.redirect(process.env.APP_ROOT + "/login?" + encoded);
            }
        });
    });
