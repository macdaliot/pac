import {Authentication, authenticationReducer} from "./Authentication";
import {createAction} from "@app/core/action";
import {JWT_RECEIVED, logoutActions} from "@app/redux/Actions/Authentication";
import {getUserFromToken} from "@app/core/token.helper";
import {IWebStorage, WebStorage} from "@app/config";

jest.mock('@app/config');

let StorageMock: jest.Mocked<IWebStorage> = WebStorage as any;


describe('authenticationReducer', () => {

    beforeEach(() => {
        jest.resetAllMocks();
    });

    it("should parse the jwt payload for JWT_RECEIVED action type", () => {
        const expectedState : any = {token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJ0b3B0YWwuY29tIiwiZXhwIjoxNDI2NDIwODAwLCJodHRwOi8vdG9wdGFsLmNvbS9qd3RfY2xhaW1zL2lzX2FkbWluIjp0cnVlLCJjb21wYW55IjoiVG9wdGFsIiwiYXdlc29tZSI6dHJ1ZX0" +
                ".yRQYnWzskCZUxPwaQupWkiUzKELZ49eM7oWxAQK_ZXw", user: {awesome: true, company: "Toptal", exp: 1426420800, "http://toptal.com/jwt_claims/is_admin": true, iss: "toptal.com"}}
        const jwtToken : string = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9." +
            "eyJpc3MiOiJ0b3B0YWwuY29tIiwiZXhwIjoxNDI2NDIwODAwLCJodHRwOi8vdG9wdGFsLmNvbS9qd3RfY2xhaW1zL2lzX2FkbWluIjp0cnVlLCJjb21wYW55IjoiVG9wdGFsIiwiYXdlc29tZSI6dHJ1ZX0." +
            "yRQYnWzskCZUxPwaQupWkiUzKELZ49eM7oWxAQK_ZXw";
        const myUser = getUserFromToken(jwtToken);
        const myJwtAction = createAction(JWT_RECEIVED, {token:jwtToken ,user:myUser})
        const authentication : Authentication = authenticationReducer(null, myJwtAction)
        expect(authentication).toEqual(expectedState);
    })
    it("should return the initial application state for LOGOUT action type", () => {
        StorageMock.isSupported.mockReturnValue(true);
        const initialApplicationState: Authentication = { user: null, token: null };
        const myLogoutAction = logoutActions.logout();
        const authentication : Authentication = authenticationReducer(null, myLogoutAction)
        expect(authentication).toEqual(initialApplicationState);
        expect(StorageMock.removeItem).toHaveBeenCalledTimes(1)
    })
    it("should not call webstorage.removeItem if not supported ", () => {
        StorageMock.isSupported.mockReturnValue(false);
        const initialApplicationState: Authentication = { user: null, token: null };
        const myLogoutAction = logoutActions.logout();
        const authentication : Authentication = authenticationReducer(null, myLogoutAction)
        expect(authentication).toEqual(initialApplicationState);
        expect(StorageMock.removeItem).toHaveBeenCalledTimes(0)
    })
});