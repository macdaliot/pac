import { CommandResult, logger, logAndQuit, runSync, sleep } from '../util';

export const runPacAutomateJenkins = async (projectDirectory: string): Promise<void> => {
  const result: CommandResult = runSync('pac automate jenkins', projectDirectory);
  logAndQuit(result.stderr);
  logger.info("Jenkins was automated successfully");
}
