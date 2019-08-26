import { CommandResult, logger, logAndQuit, runSync } from '../util';

export const runPacCleanCloud = (projectDirectory: string): void => {
  const result: CommandResult = runSync(`pac clean-cloud`, projectDirectory);
  logAndQuit(result.stderr);
  logger.info(`The PAC project was successfully cleaned up`);
}
