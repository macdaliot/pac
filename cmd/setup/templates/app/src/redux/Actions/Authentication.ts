import {ActionsUnion, createAction} from "@app/core/action";
import {IUser} from "@pyramid-systems/domain";

export const JWT_RECEIVED = 'JWT_RECEIVED'
export const LOGOUT = 'LOGOUT'

export const logoutActions = {
    logout: () => createAction(LOGOUT)
};

export const loginCallbackActions = {
    setToken: (token: string, user: IUser) =>
        createAction(JWT_RECEIVED, { token, user })
};

export type Actions = ActionsUnion<typeof logoutActions> | ActionsUnion<typeof loginCallbackActions>