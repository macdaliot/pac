import { spawn, spawnSync, ChildProcess, CommonOptions, SpawnSyncReturns } from 'child_process';
import { existsSync } from 'fs';
import { resolveDirectory } from './directories';
import { logAndQuit } from './errors';

export interface CommandResult {
  error: Error,
  exitCode: number,
  pid: number,
  stderr: string,
  stdout: string
}

interface CommandResultTruncated {
  stderr: string,
  stdout: string
}

export const runAsync = async (fullCommand: string, directory: string): Promise<CommandResultTruncated> => {
  const command: string = trimCommand(fullCommand);
  const args: Array<string> = getArgs(fullCommand);
  let child: ChildProcess;
  try {
    const detached: boolean = false;
    child = spawn(command, args, createOptions(directory, detached));
  } catch (err) {
    logAndQuit(err);
  }
  return {
    stderr: String(child.stderr),
    stdout: String(child.stdout)
  };
}

export const runDetached = (fullCommand: string, directory: string): void => {
  const command: string = trimCommand(fullCommand);
  const args: Array<string> = getArgs(fullCommand);
  const detached: boolean = true;
  const child: ChildProcess = spawn(command, args, createOptions(directory, detached));
  child.unref();
}

export const runSync = (fullCommand: string, directory?: string, detached?: boolean): CommandResult => {
  const command: string = trimCommand(fullCommand);
  const args: Array<string> = getArgs(fullCommand);
  if (!detached) {
    detached = false;
  }
  const output: SpawnSyncReturns<Buffer> = spawnSync(command, args, createOptions(directory, detached));
  return {
    error: output.error,
    exitCode: output.status,
    pid: output.pid,
    stderr: cleanOutput(output.stderr),
    stdout: cleanOutput(output.stdout)
  };
}

export const tryRunSync = (command: string, directory?: string): CommandResult => {
  try {
    return runSync(command, directory, false);
  } catch(e) {
    // Purposefully suppressing errors. This command is designed to "try" to
    //   run and fail silently if it fails
  }
}

const cleanOutput = (output: Buffer): string => {
  const searchPattern: string = '\n';
  let result: string = output.toString('utf-8');
  if (result.endsWith(searchPattern)) {
    result = result.replace(new RegExp(searchPattern + '$'), '');
  }
  return result;
}

const createOptions = (directory: string, detached: boolean): CommonOptions => {
  let cwd: string = null;
  if (directory) {
    directory = resolveDirectory(directory);
    if (existsSync(directory)) {
      cwd = directory;
    } else {
      logAndQuit(`ERROR: The following directory does not exist: ${directory}`);
    }
  }
  if (detached) {
    return {
      cwd: cwd,
      detached: true,
      stdio: ['ignore', 'ignore', 'ignore']
    } as CommonOptions
  } else {
    return {
      cwd: cwd
    }
  }
}

const getArgs = (fullCommand: string): Array<string> => {
  let args: Array<string> = [];
  const parts: Array<string> = fullCommand.split(' ');
  for (let i = 0; i < parts.length; i++) {
    if (i !== 0) {
      args.push(parts[i]);
    }
  }
  return args;
}

const trimCommand = (fullCommand: string): string => {
  return fullCommand.split(' ')[0];
}
