import { authenticateMiddleware, handlePassportCallback, handleExpressRequestLoginCallback } from '@pyramid-systems/core';
import { TsoaRoute } from 'tsoa';

jest.mock('passport', () => ({ authenticate: jest.fn((authType, options, callback) => () => { callback('This is an error', null); })}))
import * as passport from 'passport';

describe('Authentication Middleware', () => {
    it('should throw an exception given no security definitions defined', () => {
        const securities: TsoaRoute.Security[] = [{
            something: ['']
        }]

        const mockILogger = jest.fn() as any;
        const mockRequest = jest.fn() as any;
        const mockResponse = jest.fn() as any;
        const mockNext = jest.fn() as any;
        const curriedFunction = authenticateMiddleware(securities, mockILogger)
        try {
            curriedFunction(mockRequest, mockResponse, mockNext);
        }
        catch (e) {
            expect(e.message).toBe(`No security groups defined in your controller route. Did you do @Security('groups', ['{scope-names}']) above the controller class?`)
        }
    })
    it('should call authenticate given security definitions defined', () => {
        const securities: TsoaRoute.Security[] = [{
            groups: ['test']
        }]

        const mockILogger = jest.fn() as any;
        const mockRequest = jest.fn() as any;
        const mockResponse = jest.fn() as any;
        const mockNext = jest.fn() as any;
        var i = 0;
        //(passport as any).authenticate = () => i++;
        const curriedFunction = authenticateMiddleware(securities, mockILogger)
        //console.log(passport.authenticate);
        curriedFunction(mockRequest, mockResponse, mockNext);
        expect(passport.authenticate).toBeCalledTimes(1);
    })


    it('should call next with error if passport authenticate returns an error', () => {
        const securities: TsoaRoute.Security[] = [{
            groups: ['admin']
        }]


        const mockILogger = jest.fn() as any;
        const mockRequest = jest.fn() as any;
        const mockNext = jest.fn() as any;
        const mockUserFromJwt = {
            groups: ['admin'],
            name: 'bob',
            sub: ''
        }
        const error = { message: "Unauthenticated" };
        const info = undefined;
        handlePassportCallback(securities[0], mockRequest, mockNext, error, mockUserFromJwt, info, mockILogger)

        expect(mockNext).toBeCalledWith(error);
        expect(mockRequest).not.toBeCalled()
    })


    it('should call next with info if passport authenticate returns an info', () => {
        const securities: TsoaRoute.Security[] = [{
            groups: ['admin']
        }]


        const mockILogger = jest.fn() as any;
        const mockRequest = jest.fn() as any;
        const mockNext = jest.fn() as any;
        const mockUserFromJwt = {
            groups: ['admin'],
            name: 'bob',
            sub: ''
        }
        const error = undefined;
        const info = { message: "Unauthenticated" };
        handlePassportCallback(securities[0], mockRequest, mockNext, error, mockUserFromJwt, info, mockILogger)

        expect(mockNext).toBeCalledWith(info);
        expect(mockRequest).not.toBeCalled()
    })


    it('should call request login if passport has no errors or info returned', () => {
        const securities: TsoaRoute.Security[] = [{
            groups: ['admin']
        }]


        const mockILogger = jest.fn() as any;
        const mockRequest = jest.fn(() => ({ login: jest.fn() })) as any;
        const mockNext = jest.fn() as any;
        const mockUserFromJwt = {
            groups: ['admin'],
            name: 'bob',
            sub: ''
        }
        const error = undefined;
        const info = undefined;
        const mockRequestInstance = mockRequest();
        handlePassportCallback(securities[0], mockRequestInstance, mockNext, error, mockUserFromJwt, info, mockILogger)

        expect(mockNext).not.toBeCalled();
        expect(mockRequestInstance.login).toBeCalled()
    })


    it('should call next with info if passport authenticate returns an info', () => {
        const securities: TsoaRoute.Security[] = [{
            groups: ['admin']
        }]


        const mockILogger = jest.fn(() => ({ info: jest.fn() })) as any;
        const mockRequest = jest.fn() as any;
        const mockNext = jest.fn() as any;
        const mockUserFromJwt = {
            groups: ['admin'],
            name: 'bob',
            sub: ''
        }
        const error = { message: "Unauthenticated" };
        const mockLoggerInstance = mockILogger()
        handleExpressRequestLoginCallback(securities[0], mockRequest, mockNext, error, mockUserFromJwt, mockLoggerInstance)

        expect(mockNext).toBeCalledWith(error);
        expect(mockLoggerInstance.info).not.toBeCalled()
    })


    it('should call next if user is authenticated', () => {
        const securities: TsoaRoute.Security[] = [{
            groups: ['admin']
        }]


        const mockILogger = jest.fn(() => ({ info: jest.fn() })) as any;
        const mockRequest = jest.fn() as any;
        const mockNext = jest.fn() as any;
        const mockUserFromJwt = {
            groups: ['admin'],
            name: 'bob',
            sub: ''
        }
        const error = undefined;
        const mockLoggerInstance = mockILogger()
        handleExpressRequestLoginCallback(securities[0], mockRequest, mockNext, error, mockUserFromJwt, mockLoggerInstance)

        expect(mockNext).toBeCalled()
        expect(mockLoggerInstance.info).toBeCalledWith('bob has the following groups matched: admin')
    })


    it('should call response with 401 with a message', () => {
        const securities: TsoaRoute.Security[] = [{
            groups: ['normal user']
        }]
        const mockSend = jest.fn(() => ({ send: jest.fn() }))
        const mockILogger = jest.fn(() => ({ info: jest.fn() })) as any;
        const mockRequest = jest.fn(() => {
            return {
                res: {
                    status: jest.fn(() => {
                        return {
                            send: mockSend
                        }
                    }),
                }
            }
        }) as any;
        const mockNext = jest.fn() as any;
        const mockUserFromJwt = {
            groups: ['admin'],
            name: 'bob',
            sub: ''
        }
        const error = undefined;
        const mockLoggerInstance = mockILogger()
        const mockRequestInstance = mockRequest();
        handleExpressRequestLoginCallback(securities[0], mockRequestInstance, mockNext, error, mockUserFromJwt, mockLoggerInstance)

        expect(mockNext).not.toBeCalled()
        expect(mockLoggerInstance.info).toBeCalledWith('bob tried to access a protected resource undefined')
        expect(mockLoggerInstance.info).toBeCalledWith('bob has the following groups matched: ')
        expect(mockRequestInstance.res.status).toBeCalledWith(401)
        expect(mockSend).toBeCalledWith({ message: 'You are not authorized to do this.' })
    })
})