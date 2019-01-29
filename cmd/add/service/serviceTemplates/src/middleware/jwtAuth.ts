import * as passport from 'passport'
import { Request, Response, NextFunction } from 'express'
var passportOptions = { session: false };

export const passportAuthenticator = (req: Request, res: Response, next: NextFunction) => {
    passport.authenticate('jwt', passportOptions,
        (err, user, info) => {
            if (err) {
                return next(err);
            } 
            else if (info) {
                return next(info);
            }
            else {
                req.login(user, passportOptions, (err) => {
                    if (err) {
                        return next(err);
                    }
                    return next();
                });
            }
        })(req, res, next);
}
