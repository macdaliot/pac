import NotFound from '../components/pages/NotFound/NotFound';
import LoginCallback from '../components/LoginCallback/LoginCallback';

export interface Route {
  path: string;
  component: any;
  exact?: boolean;
  restricted?: boolean;
  pageTitle?: string;
  color?: string;
}
var routes = new Array<Route>();

routes.push({
  path: '/login',
  exact: true,
  component: LoginCallback
});
routes.push({
  path: '*',
  restricted: false,
  component: NotFound
});

export default routes;
