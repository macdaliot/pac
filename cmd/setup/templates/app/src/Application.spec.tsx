import * as React from 'react';
import { shallow } from 'enzyme';
import { ApplicationComponent } from './Application';
import { appStore } from './redux/Store';
import { IUser } from '@pyramid-systems/domain';

const sampleUser: IUser = {
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

describe('Application component dispatches an action if a token is in local storage', () => {
  it('should keep the loggedIn state variable at false if the user is null', () => {
    appStore.getState().user = undefined;
    const component = shallow(<ApplicationComponent />);
    const instance = component.instance() as ApplicationComponent;

    expect(component.state('loggedIn')).toEqual(false);
  });
});
