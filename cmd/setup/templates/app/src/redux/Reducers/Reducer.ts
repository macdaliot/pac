import { IUser } from '@pyramid-systems/domain';
import {authenticationReducer} from './Authentication';
import {combineReducers} from "redux";

export type Authentication = {
  user?: IUser;
  token?: string;
};

export type ApplicationState = ReturnType<typeof rootReducer>;


export const rootReducer = combineReducers({
  authentication: authenticationReducer
  })

