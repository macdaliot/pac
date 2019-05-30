/*
WARNING: DO NOTE MAKE ANY CHANGES TO THIS FILE
         THIS WILL JUST BE REGENERATED EVERYTIME
tslint:disable
*/

{{#if canImportByAlias}}
import { Controller, ValidationService, FieldErrors, ValidateError, TsoaRoute } from 'tsoa';
{{else}}
import { Controller, ValidationService, FieldErrors, ValidateError, TsoaRoute } from '../../../src';
import { NextFunction } from 'connect';
{{/if}}
    {{#each controllers}}
    import { {{name}} } from '{{modulePath}}';
    {{/each}}

        import * as passport from 'passport';
        var passportOptions = { session: false };
        import {
            intersection,
            isNullOrUndefined,
            HttpException,
            authenticateMiddleware,
            Request,
            ILogger
        } from '@pyramid-systems/core';
        import { Container } from 'inversify';
        import { IUser } from '@pyramid-systems/domain';

        import { Response, NextFunction, Express } from 'express'

        const models: TsoaRoute.Models = {
        {{#each models}}
        "{{@key}}": {
            {{#if enums}}
            "enums": {{{json enums}}},
            {{/if}}
                {{#if properties}}
                "properties": {
                    {{#each properties}}
                    "{{@key}}": {{{json this}}},
                    {{/each}}
                    },
                        {{/if}}
                            {{#if additionalProperties}}
                            "additionalProperties": {{{json additionalProperties}}},
                            {{/if}}
                            },
                                {{/each}}
                                };
                                    const validationService = new ValidationService(models);

                                    export function RegisterRoutes(app: Express, iocContainer: Container) {
                                        const logger = iocContainer.get<ILogger>(ILogger);
                                        {{#each controllers}}
                                        {{#each actions}}
                                        app.{{method}}('{{fullPath}}',
                                            {{#if security.length}}
                                        authenticateMiddleware({{json security}}, logger),
                                        {{/if}}
                                            function (request: Request, response: Response, next: NextFunction) {
                                                const args = {
                                                {{#each parameters}}
                                                {{@key}}: {{{json this}}},
                                                {{/each}}
                                                };

                                                    let validatedArgs: any[] = [];
                                                    try {
                                                        validatedArgs = getValidatedArgs(args, request);
                                                    } catch (err) {
                                                        if(err instanceof ValidateError) {
                                                            request.log.error(err.message);
                                                            return next(new HttpException(err.status, err.fields));
                                                        } else {
                                                            request.log.error(JSON.stringify(err));
                                                            return next(err);
                                                        }
                                                    }


                                                    const controller = iocContainer.get<{{../name}}>({{../name}});
                                        if (typeof controller['setStatus'] === 'function') {
                                            (<any>controller).setStatus(undefined);
                                        }

                                        const promise = controller.{{name}}.apply(controller, validatedArgs as any);
                                        promiseHandler(controller, promise, response, next);
                                    });
{{/each}}
    {{/each}}

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
            const fieldErrors: FieldErrors  = {};
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
                const message = Object.keys(fieldErrors).reduce((acc, key) => {
                    return `${acc}\n[${key}] has validation error: ${fieldErrors[key].message}!`
                }, '')
                throw new ValidateError(fieldErrors, message);
            }
            return values;
        }
    }