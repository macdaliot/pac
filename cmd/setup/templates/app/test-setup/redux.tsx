import * as React from 'react';
import configureStore from 'redux-mock-store';
import { Provider } from 'react-redux';
import { mount, shallow } from 'enzyme';

export const mountWithState = (element: JSX.Element, defaultState: any) => {
    let store = configureStore([])(defaultState);
    return mount(<Provider store={store}>{element}</Provider>);
}