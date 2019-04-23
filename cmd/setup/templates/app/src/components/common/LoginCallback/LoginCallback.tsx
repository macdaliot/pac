import * as React from 'react';
import { hot } from 'react-hot-loader/root';
import { Redirect } from 'react-router-dom';
import { connect } from 'react-redux';
import { Dispatch } from 'redux';
import { webStorage } from '../../../config';
import { JwtReceivedAction } from '../../../types';

interface LoginCallbackProps {
  setToken(token: string): void;
  location: any;
}

export class LoginCallbackComponent extends React.Component<LoginCallbackProps> {
  componentDidMount = () => {
    let token = this.props.location.search.slice(1);
    if (webStorage.isSupported()) {
      const tokenName = "pac-{{.projectName}}-token";
      webStorage.setItem(tokenName, token.toString());
    }
    this.props.setToken(token);
  }

  render = () => {
    return <Redirect to="/" />;
  }
}

const mapState = () => ({});
export const mapDispatch = (dispatch: Dispatch<JwtReceivedAction>) => ({
  setToken: (token: string) => dispatch({ type: 'JWT_RECEIVED', token: token }) 
});

export default connect(mapState, mapDispatch)(hot(LoginCallbackComponent))
