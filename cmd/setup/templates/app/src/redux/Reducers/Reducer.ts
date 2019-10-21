import { combineReducers } from "redux";
import * as Auth from './Authentication';

export type ApplicationState = ReturnType<typeof rootReducer>;
export type GetState = () => ApplicationState

export const rootReducer = combineReducers({
  authentication: Auth.Reducer
});

