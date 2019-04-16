import * as pino from "pino";
import { Injectable, isNullOrUndefined } from "@pyramid-systems/core";

const Logger: pino.Logger = pino();
if (!isNullOrUndefined(process.env.LOG_LEVEL)) {
    Logger.level = process.env.LOG_LEVEL;
}
export abstract class ILogger {
    abstract info(msg: any, ...args: any[]): void;
    abstract warn(msg: any, ...args: any[]): void;
    abstract error(msg: any, ...args: any[]): void;
    abstract trace(msg: any, ...args: any[]): void;
    abstract fatal(msg: any, ...args: any[]): void;
    abstract debug(msg: any, ...args: any[]): void;
    abstract getPinoLogger(): pino.Logger;
}


const normalizeLogMessage = (msg: any) => {
    return typeof msg === "string" ? '' : msg;
}
@Injectable()
export class ServiceLogger implements ILogger {
    private serviceLogger: pino.Logger;

    constructor(serviceName: any) {
        this.serviceLogger = Logger.child({ serviceName: serviceName });
    }

    info(msg: any, ...args: any[]): void {
        this.serviceLogger.info(msg, normalizeLogMessage(msg), args);
    }

    warn(msg: any, ...args: any[]): void {
        this.serviceLogger.warn(msg, normalizeLogMessage(msg), args);
    }

    error(msg: any, ...args: any[]): void {
        this.serviceLogger.error(msg, normalizeLogMessage(msg), args);
    }

    trace(msg: any, ...args: any[]): void {
        this.serviceLogger.trace(msg, normalizeLogMessage(msg), args);
    }

    debug(msg: any, ...args: any[]): void {
        this.serviceLogger.debug(msg, normalizeLogMessage(msg), args);
    }

    fatal(msg: any, ...args: any[]): void {
        this.serviceLogger.fatal(msg, normalizeLogMessage(msg), args);
    }

    getPinoLogger(): pino.Logger {
        return this.serviceLogger;
    }
}