import { DefaultController } from '../controllers/defaultController';
import { errorHandler } from '../utility';
import * as express from 'express';
import { serviceName } from '../config'
import { passportAuthenticator } from '../middleware/jwtAuth';
import { DynamoDB } from '../database/dynamo.db';

let apiRouter = express.Router();
let defaultController = new DefaultController(new DynamoDB());

apiRouter
  // Uncomment the following line to enable passport JWT authentication for all routes
  // Ensure you have correct settings in .env (for local use) or Lambda
  // environment variables to support authentication and authorization
  // If you are receiving 401 errors, check logs.
    // .use(passportAuthenticator)
    .get(`/${serviceName}`, defaultController.get)
    .get(`/${serviceName}/:id`, defaultController.getById)
    .post(`/${serviceName}`, defaultController.post)
    .put(`/${serviceName}`, defaultController.update)
    .put(`/${serviceName}/:id`, defaultController.update)
    .delete(`/${serviceName}/:id`, defaultController.delete)
    .delete(`/${serviceName}`, defaultController.delete)
    .use('*', errorHandler)

export default apiRouter;
