import * as React from 'react';
import { renderRoutes } from 'react-router-config';
import { BrowserRouter as Router, Switch } from 'react-router-dom';
import { Provider } from 'react-redux';
import { hot } from 'react-hot-loader/root';
import routes from './routes/routes';
import { appStore } from './redux/Store';
import '@pyramidlabs/react-ui/styles.css';
import './scss/main.scss';
import axios from 'axios';
import Header from './components/Header/Header';

interface State {
  loggedIn: boolean;
}
class ApplicationComponent extends React.Component<{}, State> {
  constructor(props) {
    super(props);
    appStore.subscribe(this.handleLogin);
    this.state = { loggedIn: false};
  }
  handleLogin = () => {
    if (appStore.getState().user != null){
      axios.defaults.headers.common['Authorization'] = `Bearer ${appStore.getState().token}`
      return this.setState({loggedIn: true});
    }
    this.setState({ loggedIn: false});  
  }

  render() {
    return (
      <Provider store={appStore}>
        <Router>
          <div className="app">
            <Header />
            <main className="main">
              <div className="content">
                <Switch>{renderRoutes(routes)}</Switch>
              </div>
            </main>
          </div>
        </Router>
      </Provider>
    );
  }
}

export default hot(ApplicationComponent);
