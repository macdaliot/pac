import { ILogger } from "@pyramid-systems/core";

declare global {
  namespace Express {
    export interface Request {
      log: ILogger;
    }
  }
}