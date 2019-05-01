import { ILogger } from "../../../core/logger";

export const logMock: ILogger = {
    debug: jest.fn(),
    info: jest.fn(),
    warn: jest.fn(),
    error: jest.fn(),
    trace: jest.fn(),
    fatal: jest.fn(),
    getPinoLogger: jest.fn()
  };