import { createStore } from 'redux';
import { rootReducer } from './Reducers';

export const appStore = createStore(rootReducer);
