import 'module-alias/register';
import * as dotenv from 'dotenv';
dotenv.load(); // must remain here before other imports

import * as express from 'express';
import * as cors from 'cors';
import * as passport from 'passport';
import { jwtStrategy, errorMiddleware, } from '@pyramid-systems/core';
import * as swaggerUi from 'swagger-ui-express';
import * as swaggerDocument from '../docs/swagger.json';
import { RegisterRoutes } from './generated/routes';
import { setupContainer } from './container-setup';

passport.use(jwtStrategy);
const app = express();
const container = setupContainer(app);
app.use(cors())
    .use(passport.initialize())
    .use(express.json())
    .use(express.urlencoded({ extended: false }))
    .use('/api/{{.serviceName}}/docs', swaggerUi.serve, swaggerUi.setup(swaggerDocument));

RegisterRoutes(app, container);

app.use(errorMiddleware);

export default app;
