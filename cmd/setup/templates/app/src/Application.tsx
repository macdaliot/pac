import * as React from 'react';
import { hot } from 'react-hot-loader/root';
import axios from 'axios';
import { Provider } from 'react-redux';
import { renderRoutes } from 'react-router-config';
import { BrowserRouter as Router, Switch } from 'react-router-dom';
import { appStore } from './redux/Store';
import { GovWebsiteBanner } from './components/GovWebsiteBanner';
import { Header } from './components/Header';
import { NavigationBar } from './components/NavigationBar';
import routes from './routes';
import '@pyramidlabs/react-ui/styles.css';
import './scss/main.scss';

interface State {
  loggedIn: boolean;
}

export class ApplicationComponent extends React.Component<{}, State> {
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
    const loggedIn = appStore.getState().user != null;
    if (loggedIn) {
      axios.defaults.headers.common['Authorization'] = `Bearer ${appStore.getState().token}`;
    }
    this.setState({
      loggedIn: loggedIn
    });
  }
}

export default hot(ApplicationComponent);
