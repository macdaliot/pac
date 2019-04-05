import NotFound from '../components/pages/NotFound/NotFound';
import * as routeData from './routes.json';

var routes = [];
routeData.default.forEach(function(route) {
  route['component'] = eval(route.pageTitle).default;
  routes.push(route);
}.bind(this));
routes.push({
  path: '/',
  exact: true,
  component: NotFound
},
{
  path: '*',
  restricted: false,
  component: NotFound
});

export default routes;
