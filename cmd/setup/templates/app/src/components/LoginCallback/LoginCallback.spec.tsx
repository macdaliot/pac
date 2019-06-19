import * as React from 'react';
import { shallow } from 'enzyme';
import {
    LoginCallbackComponent,
    mapDispatchToProps
} from './LoginCallback.component';

describe('Login Callback', () => {
    it('should render', () => {
        const location = {
            search: '/something/something-else'
        };
        shallow(
            <LoginCallbackComponent
                // tslint:disable-next-line: jsx-no-lambda
                setToken={() => ({
                    type: 'JWT_RECEIVED',
                    payload: { token: '8954gjsdfsfs', user: { name: 'someone', sub: 'sub', groups: [] } }
                })}
                location={location}
            />
        );
    });
    it('should map dispatch', () => {
        expect.assertions(2);
        const token: string = 'abcde.12345.abcde';
        const user = { name: 'someone', sub: 'sub', groups: [] };
        const result = mapDispatchToProps(action => {
            return expect(action.payload).toEqual(token);
        });
        expect(result.setToken).toBeTruthy();
        result.setToken(token, user);
    });
});
