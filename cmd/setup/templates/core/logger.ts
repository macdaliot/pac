import * as pino from "pino";
import { Injectable, isNullOrUndefined } from "@pyramidlabs/core";

const Logger: pino.Logger = pino();
if (!isNullOrUndefined(process.env.LOG_LEVEL)) {
  Logger.level = process.env.LOG_LEVEL;
}
export abstract class ILogger {
  abstract info(msg: string, ...args: any[]): void;
  abstract warn(msg: string, ...args: any[]): void;
  abstract error(msg: string, ...args: any[]): void;
  abstract trace(msg: string, ...args: any[]): void;
  abstract fatal(msg: string, ...args: any[]): void;
  abstract debug(msg: string, ...args: any[]): void;
  abstract getPinoLogger(): pino.Logger;
}

@Injectable()
export class ServiceLogger implements ILogger {
  private serviceLogger: pino.Logger;

  constructor(serviceName: string) {
    this.serviceLogger = Logger.child({ serviceName: serviceName });
  }

  info(msg: string, ...args: any[]): void {
    this.serviceLogger.info(msg, args);
  }

  warn(msg: string, ...args: any[]): void {
    this.serviceLogger.warn(msg, args);
  }

  error(msg: string, ...args: any[]): void {
    this.serviceLogger.error(msg, args);
  }

  trace(msg: string, ...args: any[]): void {
    this.serviceLogger.trace(msg, args);
  }

  debug(msg: string, ...args: any[]): void {
    this.serviceLogger.debug(msg, args);
  }

  fatal(msg: string, ...args: any[]): void {
    this.serviceLogger.fatal(msg, args);
  }

  getPinoLogger(): pino.Logger {
    return this.serviceLogger;
  }
}