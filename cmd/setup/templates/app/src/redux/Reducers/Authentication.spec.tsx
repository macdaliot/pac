import * as A from "./Authentication";
import { createAction } from "@app/core/action";
import { JWT_RECEIVED, Actions } from "@app/redux/Actions/Authentication";
import { IWebStorage, WebStorage } from "@app/config";

jest.mock('@app/config');

let StorageMock: jest.Mocked<IWebStorage> = WebStorage as any;

describe('authenticationReducer', () => {

    beforeEach(() => {
        jest.resetAllMocks();
    });

    it("Should set the right values in state", () => {
        const expectedState: A.State = {
            token: "TestToken",
            user: { name: "Test", sub: "Testsub", groups: []}
        }
        const myJwtAction = createAction(JWT_RECEIVED, { token: expectedState.token, user: expectedState.user })
        const authentication = A.Reducer(null, myJwtAction)
        expect(authentication).toEqual(expectedState);
    })
    it("should return the initial application state for LOGOUT action type", () => {
        StorageMock.isSupported.mockReturnValue(true);
        const initialState = { user: null, token: null };
        const myLogoutAction = Actions.logout();
        const authentication = A.Reducer(null, myLogoutAction)
        expect(authentication).toEqual(initialState);
        expect(StorageMock.removeItem).toHaveBeenCalledTimes(1)
    })
    it("should not call webstorage.removeItem if not supported ", () => {
        StorageMock.isSupported.mockReturnValue(false);
        const initialState = { user: null, token: null };
        const myLogoutAction = Actions.logout();
        const authentication = A.Reducer(null, myLogoutAction)
        expect(authentication).toEqual(initialState);
        expect(StorageMock.removeItem).toHaveBeenCalledTimes(0)
    })
});