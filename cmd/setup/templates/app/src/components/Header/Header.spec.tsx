import * as React from 'react';
import { shallow } from 'enzyme';
import { mountWithState } from '../../../test-setup/redux';
import { Button, ButtonPriority } from '@pyramidlabs/react-ui';
import Header, { HeaderComponent, mapStateToProps } from './Header.component';
import { Route } from '../../routes';
import { IUser } from '@pyramid-systems/domain';

jest.mock('../../routes', () => Array<Route>());

describe('header (unit/shallow)', () => {
  // if you've exported the class by itself, you can do quick unit tests like the following test
  // shallow() will not render child components

  it('should show a login link when not logged in', () => {
    const hdr = shallow(<HeaderComponent
        logout={() => ({
          type: 'LOGOUT'
        })}
    />);
    expect(
        hdr.contains(
            <a href="http://localhost:3000/api/auth/login">
              <Button text={'Login'} priority={ButtonPriority.Primary} />
            </a>
        )
    ).toBe(true);
  });

  it('should show user information while logged in', () => {
    const props = {
      userName: 'testUser',
      isAuthenticated: true
    };
    const hdr = shallow(<HeaderComponent
        {...props}
        logout={() => ({
          type: 'LOGOUT'
        })}/>);
    expect(hdr.contains(<span className="username">testUser</span>)).toBe(true);
  });

  it('should map state appropriately', () => {
    const fakeUser: IUser = {
      name: 'testUser',
      groups: [],
      sub: ''
    };
    const inputState = {
      user: fakeUser
    };
    const expectedProps = {
      userName: 'testUser',
      isAuthenticated: true
    };
    expect(mapStateToProps(inputState)).toEqual(expectedProps);
  });
});

describe('header (integration/mount)', () => {
  // following test uses mount, instead of shallow. it renders the whole component, which allows one
  // to test the redux+hot-wrapped default-exportcomponent like the following. Need to pass in a store as shown.

  it('should show a login link when not logged in', () => {
    const hdr = mountWithState(<Header />, {});
    expect(
      hdr.contains(
        <a href="http://localhost:3000/api/auth/login">
          <Button text={'Login'} priority={ButtonPriority.Primary} />
        </a>
      )
    ).toBe(true);
  });

  it('should show user information while logged in', () => {
    const defaultState = {
      user: {
        name: 'testUser'
      }
    };
    const hdr = mountWithState(<Header />, defaultState);
    expect(hdr.contains(<span className="username">testUser</span>)).toBe(true);
  });
});
