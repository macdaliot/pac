import { createStore, applyMiddleware } from 'redux';
import { rootReducer } from './Reducers';
import thunk from 'redux-thunk';

export const ApplicationStore = createStore(rootReducer, applyMiddleware(thunk));