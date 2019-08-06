import { CommandResult, logAndQuit, logger, runSync } from '../util';

export const pushToGitHub = (projectDirectory: string): void => {
  let exitCodeSum: number = 0;
  exitCodeSum += runGitAdd(projectDirectory);
  exitCodeSum += runGitCommit(projectDirectory);
  exitCodeSum += runGitPush(projectDirectory);
  if (exitCodeSum > 0) {
    logAndQuit("ERROR: The Git push to GitHub failed");
  }
}

const runGitAdd = (directory: string): number => {
  const command: string = 'git add --all';
  return runGitCommand(command, directory);
}

const runGitCommand = (command: string, directory: string): number => {
  // const errorToIgnore: string = "error: src refspec master does not match any."
  const result: CommandResult = runSync(command, directory);
  if (result.stderr) {
    // if (result.stderr !== errorToIgnore) {
    //  return 0;
    // }
    logger.error(result.stderr);
  }
  return result.exitCode;
}

const runGitCommit = (directory: string): number => {
  const command: string = 'git commit -m \"Automated-Push-Via-NodeJS\"';
  return runGitCommand(command, directory);
}

const runGitPush = (directory: string): number => {
  const command: string = 'git push origin master';
  return runGitCommand(command, directory);
}
