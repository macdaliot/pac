import { Request } from "types";
import { ILogger } from "@pyramid-systems/core";
export interface Request extends Request {
  log: ILogger;
}
