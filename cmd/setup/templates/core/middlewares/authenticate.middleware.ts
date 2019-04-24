import { TsoaRoute } from 'tsoa';
import { Response, NextFunction } from 'express';
import * as passport from 'passport';
import { IUser } from '@pyramid-systems/domain';
import { intersection, isNullOrUndefined, Request, ILogger } from '@pyramid-systems/core';

const passportOptions = { session: false };

export const handleExpressRequestLoginCallback = (protectedWith: TsoaRoute.Security, request: Request, next: NextFunction, requestLoginError: any, user: IUser, logger: ILogger) => {
  if (requestLoginError) {
    return next(requestLoginError);
  }
  const groupsMatched = intersection(
    user.groups,
    protectedWith.groups
  );
  logger.info(`${user.name} has the following groups matched: ${groupsMatched}`);
  if (
    !isNullOrUndefined(groupsMatched) &&
    groupsMatched.length > 0
  ) {
    return next();
  }
  logger.info(`${user.name} tried to access a protected resource ${request.path}`);
  request.res
    .status(401)
    .send({ message: 'You are not authorized to do this.' });
}

export const handlePassportCallback = (protectedWith: TsoaRoute.Security, request: Request, next: NextFunction, passportError: any, user: IUser, passportInfo: any, logger: ILogger) => {
  if (passportError) {
    return next(passportError);
  } else if (passportInfo) {
    return next(passportInfo);
  } else {
    request.login(user, passportOptions, err => {
      handleExpressRequestLoginCallback(protectedWith, request, next, user, err, logger);
    });
  }
}

export const authenticateMiddleware = (security: TsoaRoute.Security[] = [], logger: ILogger) => {
  return (request: Request, response: Response, next: NextFunction) => {
    for (const protectedWith of security) {
      if (isNullOrUndefined(protectedWith.groups)) {
        throw new Error(`No security groups defined in your controller route. Did you do @Security('groups', ['{scope-names}']) above the controller class?`);
      }
      if (protectedWith.groups.length > 0) {
        passport.authenticate(
          'jwt',
          passportOptions,
          (err, user: IUser, info) => {
            handlePassportCallback(protectedWith, request, next, err, user, info, logger);
          }
        )(request, response, next);
      }
    }
  };
}