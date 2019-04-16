import { iocContainer, ILogger, ServiceLogger } from "@pyramid-systems/core";
import { interfaces, Container } from "inversify";
import * as pino from "express-pino-logger";
import { Express } from "express";
import { {{.serviceNamePascal}}Repository } from '@pyramid-systems/domain';
import { {{.serviceNamePascal}}Controller } from './{{.serviceName}}.controller';


export const setupContainer = (app: Express): Container => {
    const serviceLogger = new ServiceLogger("{{.serviceNamePascal}}");
    serviceLogger.info("Setting up container");
    iocContainer.bind({{.serviceNamePascal}}Repository).to({{.serviceNamePascal}}Repository);
    iocContainer.bind({{.serviceNamePascal}}Controller).to({{.serviceNamePascal}}Controller);
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