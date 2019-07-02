import { LoginCallbackActions } from '@app/components/LoginCallback/LoginCallback.component';
import { IUser } from '@pyramid-systems/domain';
import { JWT_RECEIVED, LOGOUT } from './Actions';
import { AuthenticationActions } from "@app/components/Header/Header.component";
type Authentication = {
  user?: IUser;
  token?: string;
};

export type ApplicationState = Authentication;

const initialApplicationState: ApplicationState = { user: null, token: null };

export const rootReducer = (
    state: ApplicationState = initialApplicationState,
    action: LoginCallbackActions | AuthenticationActions
): ApplicationState => {
  switch (action.type) {
    case JWT_RECEIVED:
      return {
        token: action.payload.token,
        user: action.payload.user
      };
    case LOGOUT:
      return initialApplicationState;
    default:
      return { ...state };
  }
};
