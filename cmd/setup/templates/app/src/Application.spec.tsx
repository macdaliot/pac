import * as React from 'react';
import { shallow } from 'enzyme';
import { appStore } from './redux/Store';
import axios from 'axios';
import { ApplicationComponent } from './Application';

// TODO: Create a mock for appStore and axios
const sampleUser: string = 'thisPerson';

describe('Application component (unit/shallow)', () => {
  it('should render', () => {
    const component = shallow(<ApplicationComponent />)
    expect(component.exists()).toBe(true);
  });

  it('should have a default "loggedIn" state of false', () => {
    const component = shallow(<ApplicationComponent />)
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

  it('should fire onSubmit() when onNext() is called', () => {
    const mockFunction = () => {
      console.log("Test");
    }
    const component = shallow(<CreateReportFormComponent onSubmit={mockFunction} />
  });
});
