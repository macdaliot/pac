import * as React from 'react';
import { hot } from 'react-hot-loader/root';
import { connect } from 'react-redux';
import { Button, ButtonPriority } from '@pyramidlabs/react-ui';
import { UrlConfig } from '../../config';
import { User } from '../../types';
import './header.scss';

interface AuthState {
  user: User;
  token?: string;
}

interface HeaderProps {
  isAuthenticated?: boolean;
  username?: string;
}

export class HeaderComponent extends React.Component<HeaderProps> {
  render = () => {
    return (
      <header className="header-component">
        <div className="section left">
          <div className="image" />
          <span className="application-title">Application Title</span>
        </div>
        <div className="section right">
          {this.renderLogin()}
        </div>
      </header>
    );
  }

  renderLogin = () => {
    if (this.props.isAuthenticated) {
      return (
        <div className="logged-in-user-section">
          <div className="user-menu-button">
            <span className="fas fa-user user-icon" />
            <span className="username">{this.props.username}</span>
            <span className="fas fa-caret-down"></span>
          </div>
          <Button text={'Logout'} priority={ButtonPriority.Primary} />
        </div>
      );
    } else {
      const loginUrl = `${UrlConfig.apiUrl}auth/login`;
      return (
        <a href={loginUrl}>
          <Button text={'Login'} priority={ButtonPriority.Primary} />
        </a>
      );
    }
  }
}

export const mapState = (state: AuthState) => {
  return {
    username: state.user ? state.user.name : null, 
    isAuthenticated: Boolean(state.user)
  };
}

const mapDispatch = () => ({});
export default connect(mapState, mapDispatch)(hot(HeaderComponent));
