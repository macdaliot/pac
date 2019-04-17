import * as React from 'react';
import { shallow } from 'enzyme';
import { GovWebsiteBannerComponent } from './GovWebsiteBanner';

describe('Government Website Banner', () => {
  it('should have an image and standard text', () => {
    const component = shallow(<GovWebsiteBannerComponent />);
    expect(component.contains(`
      <div className="gov-website-banner-component">
        <div className="flag-image" />
        <span>An official website of the United States government</span>
      </div>
    `));
  });
});
