import { DefaultController } from '../controllers/defaultController';
import { errorHandler } from '../utility';
import { serviceName } from '../config';
import * as express from 'express';
let apiRouter = express.Router();
let defaultController = new DefaultController();
apiRouter
    // to require authentication for a particular route,
    // 	.get('/widgets', passport.authenticate('jwt', { session: false }), defaultController.get)
    .get(`/${serviceName}`, defaultController.get)
    .get(`/${serviceName}/:id`, defaultController.getById)
    .post(`/${serviceName}`, defaultController.post)
    // .put('/test/id', defaultController.update)
    // .delete('/test/id', defaultController.delete)

    .use('*', errorHandler)
export default apiRouter;