import * as React from 'react';
import { hot } from 'react-hot-loader/root';
import { Link } from 'react-router-dom';
import routes, { Route } from '@app/routes';
import './NavigationBar.scss';

export class NavigationBarComponent extends React.Component {
  render = () => {
    return (
      <div className="navigation-bar-component">
        {routes.map((route, key) => this.renderRoute(route, key))})}
      </div>
    );
  };

  renderRoute = (route: Route, key: number) => {
    if (route.displayName) {
      let classNames = 'route-button';
      if (route.path === location.pathname) {
        classNames += ' active-route';
      }
      return (
        <Link className={classNames} to={route.path} key={key}>
          <div className="route-button-inner">
            <span>{route.displayName}</span>
          </div>
        </Link>
      );
    }
  };
}

export default hot(NavigationBarComponent);
