import { Home } from './components/pages/Home';
import { NotFound } from './components/pages/NotFound';
import { LoginCallback } from './components/common/LoginCallback';

export interface Route {
  path: string;
  displayName?: string;
  exact?: boolean;
  restricted?: boolean;
  component: any;
}

const routes: Array<Route> = [
  {
    path: '/',
    displayName: 'Home',
    exact: true,
    component: Home
  },
  {
    path: '/login',
    exact: true,
    component: LoginCallback
  },
  {
    path: '*',
    restricted: false,
    component: NotFound
  }
]

export default routes;