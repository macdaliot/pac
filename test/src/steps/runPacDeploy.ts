import { CommandResult, logger, logAndQuit, runSync } from '../util';

export const runPacDeploy = (projectDirectory: string): void => {
  const result: CommandResult = runSync('pac deploy', projectDirectory);
  logAndQuit(result.stderr);
  logger.info("The PAC project infrastructure was provisioned without an error");
}
