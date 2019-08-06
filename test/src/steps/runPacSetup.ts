import { config } from '../config';
import {
  CommandResult,
  createDirectory,
  logAndQuit,
  logger,
  randomString,
  runSync
} from '../util';

export const runPacSetup = (): string => {
  const tempProjectDir: string = config.tempProjectDir;
  createDirectory(tempProjectDir);
  const projectNameLength: number = 7;
  const projectName: string = randomString(projectNameLength);
  const result: CommandResult = runSync(`pac setup --name ${projectName}`, tempProjectDir);
  logAndQuit(result.stderr);
  logger.info("A PAC project was setup without an error");
  return projectName;
}
