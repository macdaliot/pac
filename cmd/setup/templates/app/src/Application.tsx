import * as React from 'react';
import { hot } from 'react-hot-loader/root';
import axios from 'axios';
import { Provider } from 'react-redux';
import { renderRoutes } from 'react-router-config';
import { BrowserRouter as Router, Switch } from 'react-router-dom';
import { appStore } from './redux/Store';
import { GovWebsiteBanner } from './components/common/GovWebsiteBanner';
import { Header } from './components/common/Header';
import { NavigationBar } from './components/common/NavigationBar';
import routes from './routes';
import '@pyramidlabs/react-ui/styles.css';
import './scss/main.scss';

interface State {
  loggedIn: boolean;
}

class ApplicationComponent extends React.Component<{}, State> {
  constructor(props: any) {
    super(props);
    appStore.subscribe(this.handleLogin);
    this.state = {
      loggedIn: false
    };
  }

  render = () => {
    return (
      <Provider store={appStore}>
        <Router>
          <div className="app">
            <GovWebsiteBanner />
            <Header />
            <NavigationBar />
            <main className="main">
              <Switch>{renderRoutes(routes)}</Switch>
            </main>
          </div>
        </Router>
      </Provider>
    );
  }

  handleLogin = () => {
    if (appStore.getState().user != null) {
      axios.defaults.headers.common['Authorization'] = `Bearer ${appStore.getState().token}`;
      return this.setState({
        loggedIn: true
      });
    }
    this.setState({
      loggedIn: false
    });  
  }
}

export default hot(ApplicationComponent);