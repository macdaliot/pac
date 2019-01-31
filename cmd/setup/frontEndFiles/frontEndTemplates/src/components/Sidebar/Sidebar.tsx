import * as React from 'react';
import { hot } from 'react-hot-loader/root';
import Button from './parts/Button/Button';
import * as routeData from '../../routes/routes.json';
import './sidebar.scss';

interface Props {
  sidebar: any;
}
interface State {
  expanded: boolean;
}
export class Sidebar extends React.Component<Props, State> {
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
  sidebar: React.Ref<HTMLDivElement>;

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
          routeData.routes.map(function(route, key) {
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

export default hot(Sidebar);
