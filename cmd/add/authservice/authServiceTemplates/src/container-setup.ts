import { iocContainer, ILogger, ServiceLogger } from "@pyramid-systems/core";
import { AuthenticationController } from "./authentication-controller";
import { interfaces, Container } from "inversify";
import * as pino from "express-pino-logger";
import { Express } from "express";

export const setupContainer = (app: Express): Container => {
  const serviceLogger = new ServiceLogger("AuthenticationService");
  serviceLogger.info("Setting up container");
  iocContainer.bind(AuthenticationController).to(AuthenticationController);
  iocContainer.bind(ILogger).toDynamicValue((context: interfaces.Context) => {
    return serviceLogger;
  });
  app.use(
      pino({
        logger: serviceLogger.getPinoLogger(),
        level: process.env.LOG_LEVEL
      })
  );
  serviceLogger.debug('Container Value:', iocContainer);
  app.set("logger", serviceLogger);
  return iocContainer;
};

export const serviceContainer = iocContainer;