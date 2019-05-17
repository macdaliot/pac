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
        setToken={() => ({ type: 'test', payload: '8954gjsdfsfs' })}
        location={location}
      />
    );
  });
  it('should map dispatch', () => {
    expect.assertions(2);
    const token: string = 'abcde.12345.abcde';
    const result = mapDispatchToProps(action => {
      return expect(action.payload).toEqual(token);
    });
    expect(result.setToken).toBeTruthy();
    result.setToken(token);
  });
});
