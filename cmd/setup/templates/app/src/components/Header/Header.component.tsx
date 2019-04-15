import * as React from 'react';
import { hot } from 'react-hot-loader/root';
import { connect } from 'react-redux';
import { Button, ButtonPriority } from '@pyramidlabs/react-ui';
import { UrlConfig } from '../../config';
import './header.scss';
import { ApplicationState } from '@app/redux/Reducers';

export const mapStateToProps = (
  state: ApplicationState
): HeaderComponentState => ({
  userName: state.user ? state.user.name : null,
  isAuthenticated: Boolean(state.user)
});
export const mapDispatchToProps = () => ({});

type HeaderComponentState = { userName?: string; isAuthenticated?: boolean };
type ReduxStateToProps = ReturnType<typeof mapStateToProps>;
type ReduxDispatchToProps = ReturnType<typeof mapDispatchToProps>;
type HeaderProps = HeaderComponentState &
  ReduxDispatchToProps &
  ReduxStateToProps;

export class HeaderComponent extends React.Component<HeaderProps> {
  render = () => {
    return (
      <header className="header-component">
        <div className="section left">
          <div className="image" />
          <span className="application-title">Application Title</span>
        </div>
        <div className="section right">{this.renderLogin()}</div>
      </header>
    );
  };

  renderLogin = () => {
    if (this.props.isAuthenticated) {
      return (
        <div className="logged-in-user-section">
          <div className="user-menu-button">
            <span className="fas fa-user user-icon" />
            <span className="username">{this.props.userName}</span>
            <span className="fas fa-caret-down" />
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
  };
}

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(hot(HeaderComponent));
