import { JWT_RECEIVED } from './Actions'
import { createStore } from 'redux';
import { Buffer } from 'buffer';
import { webStorage } from '../config';
import { AuthState, JwtReceivedAction, User } from '../types';

let initialState: AuthState = {
  token: null,
  user: null,
};

const getUser = (token: string): User => {
  let payload = token.split('.')[1]; // lop off header and signature
  let json = Buffer.from(payload, 'base64').toString('ascii');
  return JSON.parse(json);
}

const tokenName = "pac-{{.projectName}}-token";
if (webStorage.isSupported() && webStorage.hasItem(tokenName)) {
  const token = webStorage.getItem(tokenName);
  initialState.token = token;
  initialState.user = getUser(token);
}

const reducer = (state = initialState, action: JwtReceivedAction) => {
  if (action.type === JWT_RECEIVED) { // this action occurs when a user logs in (posted back from auth0)
    try {
      return {
        token: action.token,
        user: getUser(action.token)
      };
    }
    catch (err) {
      return state;
    }
  } else {
    return state;
  }
}

export const appStore = createStore(reducer);
