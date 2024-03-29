import 'module-alias/register';
import * as dotenv from 'dotenv';
dotenv.load(); // must remain here before other imports

import * as express from "express";
import * as passport from "passport";
import { samlStrategy, errorMiddleware } from "@pyramid-systems/core";
import * as swaggerUi from "swagger-ui-express";
import * as swaggerDocument from "../docs/swagger.json";
import { RegisterRoutes } from "./generated/routes";
import { setupContainer } from "./container-setup";

passport.use(samlStrategy);
const app = express();
const container = setupContainer(app);
app
    .use(passport.initialize())
    .use(express.json())
    .use(express.urlencoded({ extended: false }))
    .use("/api/auth/docs", swaggerUi.serve, swaggerUi.setup(swaggerDocument))
    .use(passport.authenticate("saml", { session: false }));

RegisterRoutes(app, container);

app.use(errorMiddleware);

export default app;
