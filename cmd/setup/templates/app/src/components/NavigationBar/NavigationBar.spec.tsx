import * as React from 'react';
import { shallow } from 'enzyme';
import { Link } from 'react-router-dom';
import { NavigationBarComponent } from './NavigationBar.component';
import { Route } from '../../routes';

jest.mock('../../routes', () => [
  {
    path: '/',
    displayName: 'Home',
    exact: true,
    component: true
  },
  {
    path: '/no-display-name',
    exact: true,
    component: true
  }
] as Array<Route>);

describe('header (unit/shallow)', () => {
  it('should always have at least one link', () => {
    const mockRoutesWithDisplayNamesCount = 1;
    const component = shallow<NavigationBarComponent>(<NavigationBarComponent />);
    const links = component.find(Link);
    expect(links.length).toEqual(mockRoutesWithDisplayNamesCount);
  });

  it('should call renderRoute() once for each route', () => {
    const mockRoutesCount = 2;
    const component = shallow<NavigationBarComponent>(<NavigationBarComponent />);
    const instance = component.instance();
    instance.renderRoute = jest.fn();
    instance.render();
    expect(instance.renderRoute).toBeCalledTimes(mockRoutesCount);
  });
});
