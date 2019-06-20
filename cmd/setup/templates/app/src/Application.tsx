import * as React from 'react';
import { hot } from 'react-hot-loader/root';
import { Provider } from 'react-redux';
import { renderRoutes } from 'react-router-config';
import { BrowserRouter as Router, Switch } from 'react-router-dom';
import { ApplicationStore } from './redux/Store';
import { GovWebsiteBanner } from './components/GovWebsiteBanner';
import { Header } from './components/Header';
import { NavigationBar } from './components/NavigationBar';
import routes from './routes';
import '@pyramidlabs/react-ui/styles.css';
import './scss/main.scss';
import { WebStorage, tokenName } from './config';
import { getUserFromToken } from './core/token.helper';
import { JWT_RECEIVED } from './redux/Actions';
import { createAction } from './core/action';
export class ApplicationComponent extends React.Component<{}> {
  constructor(props: any) {
    super(props);
    if (WebStorage.isSupported() && WebStorage.hasItem(tokenName)) {
      const token = WebStorage.getItem(tokenName);
      const user = getUserFromToken(token);
      ApplicationStore.dispatch(createAction(JWT_RECEIVED, { token, user }));
    }
  }

  render = () => {
    return (
        <Provider store={ApplicationStore}>
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
