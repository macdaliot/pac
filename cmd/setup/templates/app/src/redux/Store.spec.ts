import { appStore } from './Store';
import { JWT_RECEIVED } from './Actions';

const tokenFromJwtIo =
  'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c';

describe('Redux store', () => {
  it('should process a broken jwt token without blowing up', () => {
    appStore.dispatch({ type: JWT_RECEIVED, payload: 'foo' });
    const state = appStore.getState();
    expect(state).toBeTruthy();
    expect(state.user).toBeNull();
  });
  it('should process a real jwt token', () => {
    appStore.dispatch({ type: JWT_RECEIVED, payload: tokenFromJwtIo });
    const state = appStore.getState();
    expect(state).toBeTruthy();
    expect(state.user).toBeTruthy();
    expect(state.user.name).toEqual('John Doe');
    expect(state.token).toEqual(tokenFromJwtIo);
  });
});
