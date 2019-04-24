import { Buffer } from 'buffer';
import { LoginCallbackActions } from '@app/components/LoginCallback/LoginCallback.component';
import { IUser } from '@pyramid-systems/domain';
import { JWT_RECEIVED } from './Actions';
import { webStorage } from '../config';
type Authentication = {
  user?: IUser;
  token?: string;
};

export type ApplicationState = Authentication;

const initialApplicationState: ApplicationState = { user: null, token: null };

const getUser = (token: string): IUser => {
  let payload = token.split('.')[1]; // lop off header and signature
  let json = Buffer.from(payload, 'base64').toString('ascii');
  return JSON.parse(json);
}

const tokenName = "pac-{{.projectName}}-token";
if (webStorage.isSupported() && webStorage.hasItem(tokenName)) {
  const token = webStorage.getItem(tokenName);
  initialApplicationState.token = token;
  initialApplicationState.user = getUser(token);
}

export const rootReducer = (
  state = initialApplicationState,
  action: LoginCallbackActions
) => {
  switch (action.type) {
    case JWT_RECEIVED:
      try {
        return {
          token: action.payload,
          user: getUser(action.payload)
        };
      } catch (err) {
        return state;
      }
    default:
      return { ...state };
  }
};
