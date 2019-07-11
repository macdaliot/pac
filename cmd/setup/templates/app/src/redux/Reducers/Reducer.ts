// copy commented sections and replace every instane of "Entity" with the name of 
// yours and correct the imports as necessary.

import {authenticationReducer} from './Authentication';
//import { reducer } from './Entity';
import {combineReducers} from "redux";

export type ApplicationState = ReturnType<typeof rootReducer>;
export type GetState = () => ApplicationState

export const rootReducer = combineReducers({
  authentication: authenticationReducer,
//  entity: reducer
  })

