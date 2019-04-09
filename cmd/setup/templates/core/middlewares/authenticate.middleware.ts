import { TsoaRoute } from 'tsoa';
import { Response, NextFunction } from 'express';
import * as passport from 'passport';
import { IUser } from 'domain';
import { intersection, isNullOrUndefined, Request } from '@pyramidlabs/core';

const passportOptions = { session: false };
export function authenticateMiddleware(security: TsoaRoute.Security[] = []) {
  return (request: Request, response: Response, next: NextFunction) => {
    for (const protectedWith of security) {
      if (protectedWith.groups.length > 0) {
        passport.authenticate(
          'jwt',
          passportOptions,
          (err, user: IUser, info) => {
            if (err) {
              request.log.error(err);
              return next(err);
            } else if (info) {
              request.log.error(err);
              return next(info);
            } else {
              request.login(user, passportOptions, err => {
                if (err) {
                  return next(err);
                }
                const groupsMatched = intersection(
                  user.groups,
                  protectedWith.groups
                );
                request.log.info(
                  `${
                    user.name
                  } has the following groups matched: {groupsMatched}`
                );
                if (
                  !isNullOrUndefined(groupsMatched) &&
                  groupsMatched.length > 0
                ) {
                  return next();
                }

                request.log.info(
                  `${user.name} tried to access a protected resource ${
                    request.path
                  }`
                );
                response
                  .status(401)
                  .send({ message: 'You are not authorized to do this.' });
              });
            }
          }
        )(request, response, next);
      }
    }
  };
}
