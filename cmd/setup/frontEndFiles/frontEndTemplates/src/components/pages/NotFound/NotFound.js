import React from 'react';
import { hot } from 'react-hot-loader';
import { Alert } from '@pyramidlabs/react-ui';
import './not-found.scss';

class NotFound extends React.Component {
  render() {
    return (
      <div className="not-found-page-component">
        <Alert heading="Page Not Found" type="error" message="Sorry, the page you requested could not be found" />
      </div>
    );
  }
}

export default hot(module)(NotFound);
