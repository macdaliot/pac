import React from 'react';
import { hot } from 'react-hot-loader';
import Button from './parts/Button/Button';
import * as routeData from '../../routes/routes.json';
import './sidebar.scss';

class Sidebar extends React.Component {
  constructor(props) {
    super(props);
    this.sidebar = React.createRef();
    this.createClassSpec = this.createClassSpec.bind(this);
    this.collapse = this.collapse.bind(this);
    this.expand = this.expand.bind(this);
    this.linkRef = this.linkRef.bind(this);
    this.toggle = this.toggle.bind(this);
    this.state = {
      expanded: true
    };
  }

  componentDidMount() {
    this.linkRef(this);
  }

  componentWillUnmount() {
    this.linkRef(undefined);
  }

  render() {
    var classSpec = this.createClassSpec();
    return (
      <div className={classSpec.sidebar} ref={this.sidebar}>
        {
          routeData.default.map(function(route, key) {
            return (
              <div className="sidebar-element" key={key}>
                <Button route={route} />
              </div>
            )
          }.bind(this))
        }
      </div>
    );
  }

  createClassSpec() {
    var sidebarClasses = "sidebar-component";
    sidebarClasses += this.state.expanded ? ' expanded' : ' collapsed';
    return {
      sidebar: sidebarClasses
    };
  }

  linkRef(ref) {
    this.props.sidebar && this.props.sidebar(ref);
  }

  collapse() {
    this.setState({
      expanded: false
    });
  }

  expand() {
    this.setState({
      expanded: true
    });
  }

  toggle() {
    this.setState({
      expanded: !this.state.expanded
    });
  }
}

export default hot(module)(Sidebar);
