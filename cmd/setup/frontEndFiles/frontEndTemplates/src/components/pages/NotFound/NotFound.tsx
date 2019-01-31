import * as React from 'react';
import { hot } from 'react-hot-loader/root';
import { Alert } from '@pyramidlabs/react-ui';
import './not-found.scss';

class NotFound extends React.Component {
  constructor(props){
    super(props);
  }

  render() {
    return (
      <div className="not-found-page-component">
        <Alert heading="Page Not Found" type="error" message="Sorry, the page you requested could not be found!" />
      </div>
    );
  }
}

export default hot(NotFound);
