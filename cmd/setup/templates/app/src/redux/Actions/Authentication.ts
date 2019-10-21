import { ActionsUnion, createAction } from "@app/core/action";
import { IUser } from "@pyramid-systems/domain";

export const JWT_RECEIVED = 'JWT_RECEIVED'
export const LOGOUT = 'LOGOUT'

export const Actions = {
    logout: () => createAction(LOGOUT),
    setToken: (token: string, user: IUser) => createAction(JWT_RECEIVED, { token, user })
};

export type Actions = ActionsUnion<typeof Actions>;