import * as React from 'react';
import { shallow } from 'enzyme';
import Header, { HeaderComponent, mapState } from './Header';
import { mountWithState } from '../../../testSupport/redux';

describe('header (unit/shallow)', () => {
    // if you've exported the class by itself, you can do quick unit tests like the following test
    // shallow() will not render child components
    it('should be a header', () => {
        const hdr = shallow(<HeaderComponent />)
        expect(hdr.contains(<span className="application-title">Acme Employee Register</span>)).toBe(true);
    });

    it('should show a login link when not logged in', () => {
        const hdr = shallow(<HeaderComponent />)
        expect(hdr.contains(<span className="user-name"><a href="http://localhost:3000/api/auth/login">Login</a></span>)).toBe(true);
    });
    it('should show user information while logged in', () => {
        const props = {
            username: 'testUser',
            isAuthenticated: true
        };
        const hdr = shallow(<HeaderComponent {...props} />);
        expect(hdr.contains(<span className="user-name">testUser</span>)).toBe(true);
    });
    it('should map state appropriately', () => {
        const inputState = {
            user: {
                name: 'testUser'
            }};
        const expectedProps = {
            username: 'testUser',
            isAuthenticated: true
        };
        expect(mapState(inputState)).toEqual(expectedProps);
        });
  });

describe('header (integration/mount)', () => {
    // following test uses mount, instead of shallow. it renders the whole component, which allows one
    // to test the redux+hot-wrapped default-exportcomponent like the following. Need to pass in a store as shown.
    it('should be a header (mount)', () => {
        const hdr = mountWithState(<Header />, {});
        expect(hdr.contains(<span className="application-title">Acme Employee Register</span>)).toBe(true);
    });
    it('should show a login link when not logged in', () => {
        const hdr = mountWithState(<Header />, {});
        expect(hdr.contains(<span className="user-name"><a href="http://localhost:3000/api/auth/login">Login</a></span>)).toBe(true);
    });
    it('should show user information while logged in', () => {
        const defaultState = {
            user: {
                name: 'testUser'
            }};
        const hdr = mountWithState(<Header />, defaultState);
        expect(hdr.contains(<span className="user-name">testUser</span>)).toBe(true);
    });
  });