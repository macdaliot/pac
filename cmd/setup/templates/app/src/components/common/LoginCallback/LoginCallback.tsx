import * as React from 'react';
import { hot } from 'react-hot-loader/root';
import { Redirect } from 'react-router-dom';
import { connect } from 'react-redux';
import { Dispatch } from 'redux';
import { JwtReceivedAction } from '../../../types';

interface LoginCallbackProps {
  setToken(token: string): void;
  location: any;
}

export class LoginCallbackComponent extends React.Component<LoginCallbackProps> {
  componentDidMount = () => {
    let token = this.props.location.search.slice(1);
    this.props.setToken(token);
  }

  render = () => {
    return <Redirect to="/" />;
  }
}

const mapState = () => ({});
const mapDispatch = (dispatch: Dispatch<JwtReceivedAction>) => ({
  setToken: (token: string) => dispatch({ type: 'JWT_RECEIVED', token: token }) 
});

export default connect(mapState, mapDispatch)(hot(LoginCallbackComponent))
