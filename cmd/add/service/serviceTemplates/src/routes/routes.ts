import { DefaultController } from '../controllers/defaultController';
import { errorHandler } from '../utility';
import * as express from 'express';
import { passportAuthenticator } from '../middleware/jwtAuth';

let apiRouter = express.Router();
let defaultController = new DefaultController();

apiRouter
	// Uncomment the following line to enable passport JWT authentication for all routes
	// Ensure you have correct settings in .env (for local use) or Lambda
	// environment variables to support authentication and authorization
	// If you are receiving 401 errors, check logs.
    .use(passportAuthenticator)
    .get(`/${serviceName}`, defaultController.get)
    .get(`/${serviceName}/:id`, defaultController.getById)
    .post(`/${serviceName}`, defaultController.post)
    // .put('/test/id', defaultController.update)
    // .delete('/test/id', defaultController.delete)

    .use('*', errorHandler)
export default apiRouter;