import { Actions, JWT_RECEIVED, LOGOUT } from "@app/redux/Actions/Authentication";
import { IUser } from "@pyramid-systems/domain";
import { WebStorage, tokenName } from '@app/config';

export type State = {
    user?: IUser;
    token?: string;
  };

const initialState: State = { user: null, token: null };

export function Reducer (
    state: State = initialState,
    action: Actions
): State {
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
            return initialState;
        default:
            return { ...state };
    }
}