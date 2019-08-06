interface LogFunction {
  (message: string): void
}

const log = (message: string, messageType: string, logFunction: LogFunction, newLine?: boolean): void => {
  message = `[ ${timestamp()} ${messageType} ]: ${message}`;
  if (newLine || newLine === undefined) {
    logFunction(message);
  } else {
    process.stdout.write(message);
  }
}

const logError = (message: string, newLine?: boolean): void => {
  log(message, "  ERROR", console.error, newLine);
}

const logInformation = (message: string, newLine?: boolean): void => {
  log(message, "   INFO", console.log, newLine);
}

const logWarning = (message: string, newLine?: boolean): void => {
  log(message, "WARNING", console.warn, newLine);
}

const timestamp = (): string => {
  return new Date().toISOString();
}

export const logger = {
  error: logError,
  info: logInformation,
  warn: logWarning,
}
