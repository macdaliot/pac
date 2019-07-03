import { createStore, applyMiddleware } from 'redux';
import { rootReducer } from './Reducers/Reducer';
import thunk from 'redux-thunk';

export const ApplicationStore = createStore(rootReducer, applyMiddleware(thunk));