import * as React from 'react';
import { hot } from 'react-hot-loader/root';
import { connect } from 'react-redux';
import routes, { Route } from '../../routes/routes';
import './header.scss';
import { UrlConfig } from '../../config';

interface HeaderProps {
  isAuthenticated?: boolean;
  username?: string;
}
export class HeaderComponent extends React.Component<HeaderProps> {
  constructor(props: HeaderProps){
    super(props);
  }
  render() {
    var matchingRoute = this.findMatchingRoute();
    var pageTitle = this.getPageTitle(matchingRoute);
    var styleSpec = this.createStyleSpec(matchingRoute);
    const loginUrl = `${UrlConfig.apiUrl}auth/login`;
    let login = <a href={loginUrl}>Login</a>
    return (
      <header className="header-component" style={styleSpec.header}>
        <div className="section left">
          <span className="page-title">{pageTitle}</span>
        </div>
        <div className="section center">
        <span className="application-title">{{.projectName}}</span>
        </div>
        <div className="section right">
          <span className="profile-button">
            <span className="fas fa-user-circle"></span>
            <span className="user-name">{this.props.isAuthenticated ? this.props.username : login }</span>
          </span>
        </div>
      </header>
    );
  }

  createStyleSpec(matchingRoute: Route) {
    if (matchingRoute && matchingRoute.color) {
      return {
        header: {
          backgroundColor: matchingRoute.color
        }
      };
    } else {
      const notFoundHeaderColor = '#616161';
      return {
        header: {
          backgroundColor: notFoundHeaderColor
        }
      };
    }
  }

  findMatchingRoute() {
    const pathName = location.pathname;
    let matchingRoute: Route | null;
    routes && routes.forEach((route) => {
      if (route.path === pathName) {
        matchingRoute = route;
      }
    });
    return matchingRoute;
  }

  getPageTitle(matchingRoute: Route) {
    return (matchingRoute && matchingRoute.pageTitle) ? matchingRoute.pageTitle : 'Not Found';
  }
}

export const mapState = state => {
  return {
    username: state.user ? state.user.name : null, 
    isAuthenticated: state.user ? true : false
  };
}
const mapDispatch = dispatch => ({});
export default connect(mapState, mapDispatch)(hot(HeaderComponent));
