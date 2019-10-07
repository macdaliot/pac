import { AuthenticationController } from "../src/authentication-controller";
import { logMock } from "./common";
import * as express from 'express';
import * as jwt from "jsonwebtoken";

let error: Error = null;
let encoded: string = null;
jest.mock("@pyramid-systems/core", () => (
    { 
    createJwt: (req: express.Request, callback: jwt.SignCallback) => { 
        callback(error, encoded);
    },
    Injectable: () => {},
    ILogger: {}
}));

describe('authentication controller', () => {
    let result: any = null;
    let status: number = -1;
    let capturer = (output: any) => {result = output};
    let statusCapturer = (response: any) => (output: number) => {status = output; return response};

    beforeEach(() => {
        error = null;
        encoded = null;
    });

    it('should return 200 for login', () => {
        let cut = new AuthenticationController(logMock);
        let response = { send: capturer } as any as express.Response;
        response.status = statusCapturer(response);
        let request = { body: {}, res: response } as any as express.Request;
        cut.getLogin(request);
        expect(status).toEqual(200);
        expect(result).toEqual("OK");
    });

    it('should error with a bad token', () => {
        let cut = new AuthenticationController(logMock);
        let response = { send: capturer } as any as express.Response;
        response.status = statusCapturer(response);
        let request = { body: {}, res: response } as any as express.Request;
        error = new Error("oh no");
        cut.processCallback(request);
        expect(status).toEqual(500);
    });
    it('should redirect with a good token', () => {
        let cut = new AuthenticationController(logMock);
        let response = { send: capturer } as any as express.Response;
        response.status = statusCapturer(response);
        response.redirect = () => { response.status(302)};
        let request = { body: {}, res: response } as any as express.Request;
        encoded = "ThisIsAToken";
        cut.processCallback(request);
        expect(status).toEqual(302);
    });
});