import { existsSync, mkdirSync, readdirSync } from 'fs';
import { homedir } from "os";

export const createDirectory = (directory: string): void => {
  directory = resolveDirectory(directory);
  if (!existsSync(directory)) {
    mkdirSync(directory);
  }
}

export const printDirectoryContents = (directory: string): Array<string> => {
  let fileList: Array<string> = [];
  readdirSync(directory).forEach((file: string) => {
    fileList.push(file);
  });
  return fileList;
}

export const resolveDirectory = (directory: string): string => {
  const currentDirShortcut: string = ".";
  const homeDirShortcut: string = "~";
  if (directory && directory.indexOf(homeDirShortcut) === 0) {
    return directory.replace(homeDirShortcut, homedir());
  } else if (directory && directory.indexOf(currentDirShortcut) === 0) {
    return directory.replace(currentDirShortcut, process.cwd());
  }
  return directory;
}
