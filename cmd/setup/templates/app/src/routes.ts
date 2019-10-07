import { Home } from './pages/Home';
import { NotFound } from './pages/NotFound';
import { LoginCallback } from './components/LoginCallback';

export type Route = {
  path: string;
  displayName?: string;
  exact?: boolean;
  restricted?: boolean;
  component: any;
};

const routes: Route[] = [
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
];

export default routes;
