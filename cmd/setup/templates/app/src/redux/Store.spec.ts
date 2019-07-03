import { ApplicationStore } from './Store';
import { JWT_RECEIVED } from './Actions/Authentication';
import { IUser } from '@pyramid-systems/domain';

const tokenFromJwtIo =
    'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c';

const sampleUser: IUser = {
  exp: 1234567890,
  groups: [],
  iat: 1516239022,
  iss: "urn:pacAuth",
  name: "John Doe",
  sub: "sample@pyramidsystems.com"
};
describe('Redux store', () => {
  it('should return state with token and user', () => {
    ApplicationStore.dispatch({ type: JWT_RECEIVED, payload: { token: tokenFromJwtIo, user: sampleUser } });
    const state = ApplicationStore.getState();
    expect(state).toBeTruthy();
    expect(state.authentication.user).toBeTruthy();
    expect(state.authentication.user.name).toEqual('John Doe');
    expect(state.authentication.token).toEqual(tokenFromJwtIo);
  });
});
