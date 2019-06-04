import { Request } from "express";
import { ILogger } from "@pyramid-systems/core";
export interface Request extends Request {
  log: ILogger;
}
