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
    state: ApplicationState = initialApplicationState,
    action: LoginCallbackActions
): ApplicationState => {
  switch (action.type) {
    case JWT_RECEIVED:
      return {
        token: action.payload.token,
        user: action.payload.user
      };
    default:
      return { ...state };
  }
};
