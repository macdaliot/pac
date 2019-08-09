import { iocContainer, ILogger, ServiceLogger } from "@pyramid-systems/core";
import { Express } from "express";
import * as pino from "express-pino-logger";
import { interfaces, Container } from "inversify";
import { {{.serviceNamePascal}}Controller } from './{{.serviceName}}.controller';

export const setupContainer = (app: Express): Container => {
  const serviceLogger = new ServiceLogger("{{.serviceNamePascal}}");
  serviceLogger.info("Setting up container");
  iocContainer.bind({{.serviceNamePascal}}Controller).to({{.serviceNamePascal}}Controller);
  iocContainer.bind(ILogger).toDynamicValue((context: interfaces.Context) => {
    return serviceLogger;
  });
  app.use(
    pino({
      logger: serviceLogger.getPinoLogger(),
      useLevel: 'trace'
    })
  );
  serviceLogger.debug('Container Value:', iocContainer);
  app.set("logger", serviceLogger);
  return iocContainer;
};

export const serviceContainer = iocContainer;