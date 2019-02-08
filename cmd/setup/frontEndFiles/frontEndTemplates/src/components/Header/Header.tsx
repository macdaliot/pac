import * as React from 'react';
import { hot } from 'react-hot-loader/root';
import { connect } from 'react-redux';
import * as routeData from '../../routes/routes.json';
import './header.scss';
import { UrlConfig } from '../../config'

class Header extends React.Component<any> {
  render() {
    var matchingRoute = this.findMatchingRoute();
    var pageTitle = this.getPageTitle(matchingRoute);
    var styleSpec = this.createStyleSpec(matchingRoute);
    let login = <a href={UrlConfig.apiUrl + "/api/auth/login"}>Login</a>

    return (
      <header className="header-component" style={styleSpec.header}>
        <div className="section left">
          <span className="menu-button material-icons" onClick={this.onClickMenuButton}>menu</span>
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

  createStyleSpec(matchingRoute) {
    if (matchingRoute) {
      return {
        header: {
          backgroundColor: matchingRoute.color
        }
      };
    } else {
      var notFoundHeaderColor = '#616161';
      return {
        header: {
          backgroundColor: notFoundHeaderColor
        }
      };
    }
  }

  findMatchingRoute() {
    var pathName = location.pathname;
    if (pathName === '/') {
      return routeData.routes[0];
    }
    var matchingRoute;
    routeData.routes.forEach((route) => {
      if (route.path === pathName) {
        matchingRoute = route;
      }
    });
    return matchingRoute;
  }

  getPageTitle(matchingRoute) {
    return matchingRoute ? matchingRoute.pageTitle : 'Not Found';
  }

  onClickMenuButton() {
    if (this.props.sidebar) {
      this.props.sidebar.toggle();
    }
  }
}
const mapState = state => ({username: state.user ? state.user.name : null, isAuthenticated: state.user ? true : false});
const mapDispatch = dispatch => ({});
export default connect(mapState, mapDispatch)(hot(Header));
