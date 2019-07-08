import {Actions, JWT_RECEIVED, LOGOUT} from "@app/redux/Actions/Authentication";
import { IUser } from "@pyramid-systems/domain";

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
            return initialApplicationState;
        default:
            return { ...state };
    }
};