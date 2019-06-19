import * as React from 'react';
import { hot } from 'react-hot-loader/root';
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
import { webStorage } from './config';
import { getUserFromToken } from './core/token.helper';
import { JWT_RECEIVED } from './redux/Actions';
import { createAction } from './core/action';

export class ApplicationComponent extends React.Component<{}> {
  constructor(props: any) {
    super(props);
    const tokenName = 'pac-testdsa-token';
    if (webStorage.isSupported() && webStorage.hasItem(tokenName)) {
      const token = webStorage.getItem(tokenName);
      const user = getUserFromToken(token);
      appStore.dispatch(createAction(JWT_RECEIVED, { token, user }));
    }
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
  };
}

export default hot(ApplicationComponent);
