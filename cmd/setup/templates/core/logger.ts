import * as pino from "pino";
import { Injectable } from "@pyramidlabs/core";

const logger = pino();
export abstract class ILogger {
    abstract info(msg: string, ...args: any[]): void;
    abstract warn(msg: string, ...args: any[]): void;
    abstract error(msg: string, ...args: any[]): void;
    abstract trace(msg: string, ...args: any[]): void;
}

@Injectable()
export class Logger implements ILogger {
    private serviceLogger: pino.Logger;
    constructor(serviceName: string) {
        this.serviceLogger = logger.child({ serviceName: serviceName });
    }

    info(msg: string, ...args: any[]) {
        this.serviceLogger.info(msg, args);
    }

    warn(msg: string, ...args: any[]) {
        this.serviceLogger.warn(msg, args);
    }

    error(msg: string, ...args: any[]) {
        this.serviceLogger.error(msg, args);
    }

    trace(msg: string, ...args: any[]) {
        this.serviceLogger.trace(msg, args);
    }
}

export { logger };
