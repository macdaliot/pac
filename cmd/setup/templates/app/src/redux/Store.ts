import { JWT_RECEIVED } from './Actions'
import { createStore } from 'redux';
import { Buffer } from 'buffer';
import { AuthState, JwtReceivedAction, User } from '../types';

const reducer = (state = { user: null, token: null } as AuthState, action: JwtReceivedAction) => {
  switch (action.type) {
    case JWT_RECEIVED: // this action occurs when a user logs in (posted back from auth0)
      try {
        let payload = action.token.split('.')[1]; // lop off header and signature
        let json = Buffer.from(payload, 'base64').toString('ascii')
        let jwt = JSON.parse(json);
        debugger;
        return {
          user: jwt,
          token: action.token
        };
      }
      catch (err) {
        return state;
      }
    default:
      return state;
  }
}

export const appStore = createStore(reducer);
