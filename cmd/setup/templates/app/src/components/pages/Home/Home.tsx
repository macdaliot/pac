import * as React from 'react';
import { hot } from 'react-hot-loader/root';
import { connect } from 'react-redux';
import { AuthState } from '../../../types/AuthState';
import { Alert, AlertType } from '@pyramidlabs/react-ui';
import './home.scss';

interface HomeProps {
  isAuthenticated: boolean
}

export class HomeComponent extends React.Component<HomeProps> {
  render = () => {
    return !this.props.isAuthenticated ? this.renderNotLoggedIn() : (
      <div className="home-page-component">
        <Alert
          foeId="sample-alert"
          foeTitle="sample-alert"
          heading="Congratulations!"
          message="You are on the home page and logged in!"
          type={AlertType.Success}
        />
      </div>
    );
  }

  renderNotLoggedIn = () => {
    return (
      <div className="home-page-component">
        <Alert
          foeId="sample-alert"
          foeTitle="sample-alert"
          heading="You Are Unauthenticated"
          message="You are on the home page, but you are not logged in"
          type={AlertType.Information}
        />
      </div>
    );
  }
}

export const mapState = (state: AuthState): HomeProps => {
  return {
    isAuthenticated: Boolean(state.user)
  };
}
const mapDispatch = () => ({});
export default connect(mapState, mapDispatch)(hot(HomeComponent));
