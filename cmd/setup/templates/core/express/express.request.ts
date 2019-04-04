import { Request } from "express";
import { ILogger } from "@pyramidlabs/core";
export interface Request extends Request {
  log: ILogger;
}
