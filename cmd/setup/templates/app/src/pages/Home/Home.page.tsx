import * as React from 'react';
import { hot } from 'react-hot-loader/root';
import { connect } from 'react-redux';
import { Alert, AlertType } from '@pyramidlabs/react-ui';
import './home.scss';
import { ApplicationState } from '@app/redux/Reducers/Reducer';
import { Dispatch } from 'redux';

export const mapStateToProps = (state: ApplicationState) => {
  return {
    isAuthenticated: Boolean(state.authentication && state.authentication.user)
  };
};
export const mapDispatchToProps = (dispatch: Dispatch) => ({});

type ReduxProps = ReturnType<typeof mapStateToProps>;
type ReduxActions = ReturnType<typeof mapDispatchToProps>;
type HomeProps = ReduxProps & ReduxActions;

export class HomeComponent extends React.Component<HomeProps> {
  render = () => {
    return !this.props.isAuthenticated ? (
      this.renderNotLoggedIn()
    ) : (
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
  };

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
  };
}

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(hot(HomeComponent));
