import { createJwt } from './passport.strategies'
import { Request } from 'express';
let tokenResult: any;
let optionsResult: any;
jest.mock("jsonwebtoken", () => ({
    sign: (token: any, secret: any, options: any, callback: Function) => {
        tokenResult = token;
        optionsResult = options;
    }
}))
describe('Passport strategies', () => {
    it('should create a jwt given a parsed saml token', () => {
        let req = {
            user: {
                "http://schemas.auth0.com/nickname": "TestUser",
                "http://schemas.xmlsoap.org/ws/2005/05/identity/claims/upn": "TestSubject"
            }
        } as any as Request;
        createJwt(req, null);
        expect(tokenResult).toBeTruthy();
        expect(tokenResult.name).toEqual("TestUser");
        expect(tokenResult.sub).toEqual("TestSubject");
    }),
    it('should map groups', () => {
        let req = {
            user: {
                "http://schemas.auth0.com/nickname": "TestUser",
                "http://schemas.xmlsoap.org/ws/2005/05/identity/claims/upn": "TestSubject",
                "http://schemas.xmlsoap.org/claims/Group": ["TestGroup1", "TestGroup2"]
            }
        } as any as Request;
        createJwt(req, null);
        expect(tokenResult).toBeTruthy();
        expect(tokenResult.groups).toHaveLength(2);
    })
})
