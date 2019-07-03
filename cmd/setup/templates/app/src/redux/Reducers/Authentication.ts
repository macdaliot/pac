import {Authentication} from "@app/redux/Reducers/Reducer";
import {Actions, JWT_RECEIVED, LOGOUT} from "@app/redux/Actions/Authentication";


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