import { getUserFromToken } from "./token.helper";
import {IUser} from "@pyramid-systems/domain";

describe('getUser from token', () => {

    it("should get user from token", () => {
        const jwtToken : string = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9." +
            "eyJpc3MiOiJ0b3B0YWwuY29tIiwiZXhwIjoxNDI2NDIwODAwLCJodHRwOi8vdG9wdGFsLmNvbS9qd3RfY2xhaW1zL2lzX2FkbWluIjp0cnVlLCJjb21wYW55IjoiVG9wdGFsIiwiYXdlc29tZSI6dHJ1ZX0." +
            "yRQYnWzskCZUxPwaQupWkiUzKELZ49eM7oWxAQK_ZXw"
        const expectedUser = {
            iss: 'toptal.com',
            exp: 1426420800,
            'http://toptal.com/jwt_claims/is_admin': true,
            company: 'Toptal',
            awesome: true }
        const user: IUser = getUserFromToken(jwtToken);
        expect(user).toEqual(expectedUser);
    })
});