import * as net from 'net';
import { CommandResult, runSync } from './commands';
import { logAndQuit } from './errors';

export const isPortFree = async (portNumber: number): Promise<boolean> => {
  return !(await isPortInUse(portNumber));
}

export const isPortInUse = async (portNumber: number) => new Promise<boolean>((resolve, reject) => {
  const server = net.createServer();
  server.once('error', (error: Error) => {
    error.name === 'EADDRINUSE' ? resolve(true) : logAndQuit(error.message); reject(error);
  });
  server.once('listening', () => {
    server.close();
    resolve(false);
  });
  server.listen(portNumber);
});

// TODO: Clean up
export const killProcessOnPort = (port: number, programName: string): void => {
  const allProcesses: CommandResult = runSync("netstat -putlan");
  const processRegex: RegExp = new RegExp("((^tcp)|(^udp)).*127\\.0\\.(0|1)\\.(0|1):" + port + ".*", "gm");
  const matches: RegExpMatchArray = allProcesses.stdout.match(processRegex);
  if (matches === null) {
    const localhost: string = '127.0.0.1';
    logAndQuit(`ERROR: No process was found running at ${localhost}:${port}`);
  }
  const pidRegex: RegExp = /([0-9]+)\/.*/gm
  let matchingPidPrograms: Array<string> = [];
  matches.forEach((match: string) => {
    let pidProgram: RegExpMatchArray = match.match(pidRegex);
    if (pidProgram && pidProgram.length === 1) {
      matchingPidPrograms.push(pidProgram[0]);
    }
  });
  const pidProgram: string = matchingPidPrograms[0];
  if (pidProgram.match(programName)) {
    const numbersOnlyRegex: RegExp = /([0-9]+)/gm
    const pid: number = Number(pidProgram.match(numbersOnlyRegex)[0]);
    process.kill(pid);
  } else {
    logAndQuit(`ERROR: The program running on port ${port} does not match ${programName}`)
  }
}
