import React from 'react';
import ReactDOM from 'react-dom';
import { renderRoutes } from 'react-router-config';
import { BrowserRouter as Router, Switch } from 'react-router-dom';
import { hot } from 'react-hot-loader';
import routes from './routes/routes';
import Header from './components/Header/Header';
import Sidebar from './components/Sidebar/Sidebar';
import './scss/main.scss';

class Application extends React.Component {
  constructor(props) {
    super(props);
    this.setSidebarRef = this.setSidebarRef.bind(this);
    this.state = {
      sidebar: undefined
    }
  }

  render() {
    return (
      <Router>
        <div className="app">
          <Header sidebar={this.state.sidebar} />
          <main className="main">
            <Sidebar sidebar={this.setSidebarRef} />
            <div className="content">
              <Switch>{renderRoutes(routes)}</Switch>
            </div>
          </main>
        </div>
      </Router>
    );
  }

  setSidebarRef(ref) {
    this.setState({
      sidebar: ref
    });
  }
}

ReactDOM.render(<Application />, document.getElementById('container'));
export default hot(module)(Application);
