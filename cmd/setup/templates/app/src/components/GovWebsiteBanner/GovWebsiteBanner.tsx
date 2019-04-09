import * as React from 'react';
import { hot } from 'react-hot-loader/root';
import './gov-website-banner.scss';

export class GovWebsiteBannerComponent extends React.Component {
  render = () => {
    return (
      <div className="gov-website-banner-component">
        <div className="flag-image" />
        <span>An official website of the United States government</span>
      </div>
    );
  }
}

export default hot(GovWebsiteBannerComponent);
