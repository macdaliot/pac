import * as React from 'react';
import { shallow } from 'enzyme';
import { HomeComponent, mapStateToProps } from './Home.page';
import { ApplicationState } from '@app/redux/Reducers';

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
    } as ApplicationState;
    const expectedProps = {
      isAuthenticated: true
    };
    expect(mapStateToProps(inputState)).toEqual(expectedProps);
  });
  it('should map state appropriately', () => {
    const inputState = {
      user: undefined,
      token: undefined
    } as ApplicationState;
    const expectedProps = {
      isAuthenticated: false
    };
    expect(mapStateToProps(inputState)).toEqual(expectedProps);
  });
});
