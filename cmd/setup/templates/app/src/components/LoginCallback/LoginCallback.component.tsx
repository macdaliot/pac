import * as React from 'react';
import { hot } from 'react-hot-loader/root';
import { Redirect } from 'react-router-dom';
import { connect } from 'react-redux';
import { Dispatch, bindActionCreators } from 'redux';
import { createAction, ActionsUnion } from '@app/core/action';
import { webStorage } from '@app/config';
import { JWT_RECEIVED } from '@app/redux/Actions';

const actions = {
  setToken: (token: string) => createAction(JWT_RECEIVED, token)
};
export type LoginCallbackActions = ActionsUnion<typeof actions>;

export const mapStateToProps = () => ({});
export const mapDispatchToProps = (dispatch: Dispatch) =>
    bindActionCreators(actions, dispatch);

type Location = {
  location: { search: string };
};
type ReduxStateToProps = ReturnType<typeof mapStateToProps>;
type ReduxDispatchToProps = ReturnType<typeof mapDispatchToProps>;
type LoginCallbackProps = Location & ReduxDispatchToProps & ReduxStateToProps;

export class LoginCallbackComponent extends React.Component<
    LoginCallbackProps
    > {
  componentDidMount = () => {
    const token = this.props.location.search.slice(1);
    if (webStorage.isSupported()) {
      const tokenName = "pac-wwedone-token";
      webStorage.setItem(tokenName, token.toString());
    }
    this.props.setToken(token);
  };

  render = () => {
    return <Redirect to="/" />;
  };
}

export default connect(
    mapStateToProps,
    mapDispatchToProps
)(hot(LoginCallbackComponent));
