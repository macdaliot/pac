import { Buffer } from 'buffer';
import { LoginCallbackActions } from '@app/components/LoginCallback/LoginCallback.component';
import { IUser } from '@pyramid-systems/domain';
import { JWT_RECEIVED } from './Actions';

type Authentication = {
  user?: IUser;
  token?: string;
};

export type ApplicationState = Authentication;

const initialApplicationState: ApplicationState = { user: null, token: null };

export const rootReducer = (
  state = initialApplicationState,
  action: LoginCallbackActions
) => {
  switch (action.type) {
    case JWT_RECEIVED:
      try {
        const token = action.payload as string;
        const splitToken = token.split('.')[1]; // lop off header and signature
        const json = Buffer.from(splitToken, 'base64').toString('ascii');
        const jwt = JSON.parse(json);
        return {
          user: jwt,
          token
        };
      } catch (err) {
        return state;
      }
    default:
      return { ...state };
  }
};
