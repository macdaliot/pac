import * as React from 'react';
import { hot } from 'react-hot-loader/root';
import { Alert, AlertType } from '@pyramidlabs/react-ui';
import './home.scss';

export class HomeComponent extends React.Component {
  render = () => {
    return (
      <div className="home-page-component">
        <Alert 
          foeId="sample-alert"
          foeTitle="sample-alert"
          heading="Congratulations!" 
          type={AlertType.Success} 
          message="You are on the home page!" />
      </div>
    );
  }
}

export default hot(HomeComponent);
