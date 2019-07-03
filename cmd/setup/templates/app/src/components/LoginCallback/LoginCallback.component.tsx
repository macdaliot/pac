import * as React from 'react';
import { hot } from 'react-hot-loader/root';
import { Redirect } from 'react-router-dom';
import { connect } from 'react-redux';
import { Dispatch, bindActionCreators } from 'redux';
import { WebStorage, tokenName } from '@app/config';
import { loginCallbackActions } from '@app/redux/Actions/Authentication';
import { getUserFromToken } from '@app/core/token.helper';

export const mapStateToProps = () => ({});
export const mapDispatchToProps = (dispatch: Dispatch) =>
    bindActionCreators(loginCallbackActions, dispatch);

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
