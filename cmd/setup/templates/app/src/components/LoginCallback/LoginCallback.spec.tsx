import * as React from 'react';
import { shallow } from 'enzyme';
import { LoginCallbackComponent, mapDispatch } from './LoginCallback';

describe('Login Callback', () => {
  it('should render', () => {
    let location = {
      search: '/something/something-else'
    };
    shallow(<LoginCallbackComponent setToken={() => {}} location={location} />);
  });
  it('should map dispatch', () => {
    expect.assertions(2);
    let token: string = 'abcde.12345.abcde';
    let result = mapDispatch((action: any) => expect(action.token).toEqual(token));
    expect(result.setToken).toBeTruthy();
    result.setToken(token);
  })
});
