import * as React from 'react';
import axios from 'axios';
import { shallow } from 'enzyme';
import { ApplicationComponent } from './Application';
import { appStore } from './redux/Store';
import { User } from './types';
 
const sampleUser: User = {
  exp: 1556083479,
  groups: [],
  iat: 1556047479,
  iss: "urn:pacAuth",
  name: "Sample",
  sub: "sample@pyramidsystems.com"
};

describe('Application component (unit/shallow)', () => {
  it('should render', () => {
    const component = shallow(<ApplicationComponent />);
    expect(component.exists()).toBe(true);
  });

  it('should have a default "loggedIn" state of false', () => {
    const component = shallow(<ApplicationComponent />);
    expect(component.state('loggedIn')).toBe(false);
  });
});

describe('Application component handles login correctly', () => {
  it('should keep the loggedIn state variable at false if the user is null', () => {
    appStore.getState().user = undefined;
    const component = shallow(<ApplicationComponent />);
    const instance = component.instance() as ApplicationComponent;
    instance.handleLogin();
    expect(component.state('loggedIn')).toEqual(false);
  });

  it('should set the token as an axios default header if the user is not null', () => {
    appStore.getState().user = sampleUser;
    const component = shallow(<ApplicationComponent />);
    const instance = component.instance() as ApplicationComponent;
    instance.handleLogin();
    expect(axios.defaults.headers.common['Authorization']).not.toBe(undefined);
  });

  it('should set the loggedIn state variable to true if the user is not null', () => {
    appStore.getState().user = sampleUser;
    const component = shallow(<ApplicationComponent />);
    const instance = component.instance() as ApplicationComponent;
    instance.handleLogin();
    expect(component.state('loggedIn')).toEqual(true);
  });
});
