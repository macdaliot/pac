import * as React from 'react';
import { hot } from 'react-hot-loader/root';
import { Redirect } from 'react-router-dom';
import { connect } from 'react-redux';
import { Dispatch } from 'redux';
import { WebStorage, tokenName } from '@app/config';

import { getUserFromToken } from '@app/core/token.helper';
import { IUser } from '@pyramid-systems/domain';
import { Actions } from '@app/redux/Actions/Authentication';

export const mapStateToProps = () => ({});
export const mapDispatchToProps = (dispatch: Dispatch) =>
  ({ setToken: (token: string, user: IUser) => dispatch(Actions.setToken(token, user)) });

type Location = {
  location: { search: string };
};
type ReduxStateToProps = ReturnType<typeof mapStateToProps>;
type ReduxDispatchToProps = ReturnType<typeof mapDispatchToProps>;
type LoginCallbackProps = Location & ReduxDispatchToProps & ReduxStateToProps;

export class LoginCallbackComponent extends React.Component<LoginCallbackProps> {
  componentDidMount = () => {
    const token = this.props.location.search.slice(1);
    if (WebStorage.isSupported()) {
      WebStorage.setItem(tokenName, token.toString());
    }

    this.props.setToken(token, getUserFromToken(token));
  };

  render = () => {
    return <Redirect to="/" />;
  };
}

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(hot(LoginCallbackComponent));
