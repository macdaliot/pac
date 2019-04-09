import * as React from 'react';
import { shallow } from 'enzyme';
import { HomeComponent, mapState } from './Home';
import { AuthState } from '@app/types';

describe('Home Page', () => {
  it('does not call renderNotLoggedIn() if authenticated', () => {
    const component = shallow<HomeComponent>(
        <HomeComponent isAuthenticated={true} />
    );
    const instance = component.instance();
    instance.renderNotLoggedIn = jest.fn();
    instance.render();
    expect(instance.renderNotLoggedIn).toBeCalledTimes(0);
  });
  it('calls renderNotLoggedIn() if not authenticated', () => {
    const component = shallow<HomeComponent>(
        <HomeComponent isAuthenticated={false} />
    );
    const instance = component.instance();
    instance.renderNotLoggedIn = jest.fn();
    instance.render();
    expect(instance.renderNotLoggedIn).toBeCalledTimes(1);
  });
  it('should map state appropriately', () => {
    const inputState = {
      user: {
        name: 'testUser'
      },
      token: 'abcde.12345.abcde'
    } as AuthState;
    const expectedProps = {
      isAuthenticated: true
    };
    expect(mapState(inputState)).toEqual(expectedProps);
  });
  it('should map state appropriately', () => {
    const inputState = {
      user: undefined,
      token: undefined
    } as AuthState;
    const expectedProps = {
      isAuthenticated: false
    };
    expect(mapState(inputState)).toEqual(expectedProps);
  });
});
