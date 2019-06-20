import { createStore } from 'redux';
import { rootReducer } from './Reducers';

export const ApplicationStore = createStore(rootReducer);