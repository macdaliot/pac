import * as React from 'react';
import { shallow } from 'enzyme';
import { NotFoundComponent } from './NotFound';

describe('Not Found Page', () => {
  it('should have an error alert', () => {
    const component = shallow(<NotFoundComponent />);
    expect(component.contains(`
      <div className="not-found-page-component">
        <Alert 
          foeTitle={'page-not-found'}
          foeId={'error-alert'}
          heading="Page Not Found" 
          type={AlertType.Error} 
          message="Sorry, the page you requested could not be found!" />
      </div>
    `));
  });
});
