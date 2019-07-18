import {Actions, JWT_RECEIVED, LOGOUT} from "@app/redux/Actions/Authentication";
import { IUser } from "@pyramid-systems/domain";
import { WebStorage, tokenName } from '@app/config';

export type Authentication = {
    user?: IUser;
    token?: string;
  };

const initialApplicationState: Authentication = { user: null, token: null };

export function authenticationReducer (
    state: Authentication = initialApplicationState,
    action: Actions
): Authentication{
    switch (action.type) {
        case JWT_RECEIVED:
            return {
                token: action.payload.token,
                user: action.payload.user
            };
        case LOGOUT:
            if (WebStorage.isSupported()) {
                WebStorage.removeItem(tokenName);
            }
            return initialApplicationState;
        default:
            return { ...state };
    }
}