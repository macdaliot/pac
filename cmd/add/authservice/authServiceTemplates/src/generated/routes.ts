/*
WARNING: DO NOTE MAKE ANY CHANGES TO THIS FILE
         THIS WILL JUST BE REGENERATED EVERYTIME
tslint:disable
*/

import { Controller, ValidationService, FieldErrors, ValidateError, TsoaRoute } from 'tsoa';
import { AuthenticationController } from './../authentication-controller';

import * as passport from 'passport';
var passportOptions = { session: false };
import {
    intersection,
    isNullOrUndefined,
    HttpException,
    iocContainer,
    expressAuthentication,
    authenticateMiddleware,
    Logger
} from '@pyramid-systems/core';
import { IUser } from '@pyramid-systems/domain';

import { Request, Response, NextFunction, Express } from 'express'

const models: TsoaRoute.Models = {
};
const validationService = new ValidationService(models);

export function RegisterRoutes(app: Express) {
    app.get('/api/auth/login',
        function(request: Request, response: Response, next: NextFunction) {
            const args = {
                request: { "in": "request", "name": "request", "required": true, "dataType": "object" },
            };

            let validatedArgs: any[] = [];
            try {
                validatedArgs = getValidatedArgs(args, request);
            } catch (err) {
                if (err instanceof ValidateError) {
                    return next(new HttpException(err.status, err.message));
                } else {
                    return next(err);
                }
            }

            const controller = iocContainer.get<AuthenticationController>(AuthenticationController);
            if (typeof controller['setStatus'] === 'function') {
                (<any>controller).setStatus(undefined);
            }


            const promise = controller.getLogin.apply(controller, validatedArgs as any);
            promiseHandler(controller, promise, response, next);
        });
    app.post('/api/auth/callback',
        function(request: Request, response: Response, next: NextFunction) {
            const args = {
                request: { "in": "request", "name": "request", "required": true, "dataType": "object" },
            };

            let validatedArgs: any[] = [];
            try {
                validatedArgs = getValidatedArgs(args, request);
            } catch (err) {
                if (err instanceof ValidateError) {
                    return next(new HttpException(err.status, err.message));
                } else {
                    return next(err);
                }
            }

            const controller = iocContainer.get<AuthenticationController>(AuthenticationController);
            if (typeof controller['setStatus'] === 'function') {
                (<any>controller).setStatus(undefined);
            }


            const promise = controller.login.apply(controller, validatedArgs as any);
            promiseHandler(controller, promise, response, next);
        });


    function isController(object: any): object is Controller {
        return 'getHeaders' in object && 'getStatus' in object && 'setStatus' in object;
    }

    function promiseHandler(controllerObj: any, promise: any, response: Response, next: NextFunction) {
        return Promise.resolve(promise)
            .then((data: any) => {
                let statusCode;
                if (isController(controllerObj)) {
                    const headers = controllerObj.getHeaders();
                    Object.keys(headers).forEach((name: string) => {
                        response.set(name, headers[name]);
                    });

                    statusCode = controllerObj.getStatus();
                }

                if (data || data === false) { // === false allows boolean result
                    response.status(statusCode || 200).json(data);
                } else {
                    response.status(statusCode || 204).end();
                }
            })
            .catch((error: any) => next(error));
    }

    function getValidatedArgs(args: any, request: any): any[] {
        const fieldErrors: FieldErrors = {};
        const values = Object.keys(args).map((key) => {
            const name = args[key].name;
            switch (args[key].in) {
                case 'request':
                    return request;
                case 'query':
                    return validationService.ValidateParam(args[key], request.query[name], name, fieldErrors);
                case 'path':
                    return validationService.ValidateParam(args[key], request.params[name], name, fieldErrors);
                case 'header':
                    return validationService.ValidateParam(args[key], request.header(name), name, fieldErrors);
                case 'body':
                    return validationService.ValidateParam(args[key], request.body, name, fieldErrors, name + '.');
                case 'body-prop':
                    return validationService.ValidateParam(args[key], request.body[name], name, fieldErrors, 'body.');
            }
        });
        if (Object.keys(fieldErrors).length > 0) {
            throw new ValidateError(fieldErrors, '');
        }
        return values;
    }
}