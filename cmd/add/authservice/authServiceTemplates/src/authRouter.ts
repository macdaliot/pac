import { Router, Request, Response } from 'express';
import * as passport from 'passport';
import * as jwt from 'jsonwebtoken';
import { errorHandler, generateRandomString } from './functions'

export let authRouter = Router();

const jwtSecret = process.env.JWT_SECRET || generateRandomString();

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
        issuer: "urn:pacAuth"
    }
    jwt.sign(token, jwtSecret, options, callback);
}
authRouter.use(passport.authenticate('saml', { session: false }));

authRouter.get('/login', (req, res) => { res.status(200).send("OK")});

// Perform the final stage of authentication and redirect to previously requested URL or '/user'
authRouter.post('/callback',
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
authRouter.use("*", errorHandler);