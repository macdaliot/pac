import { appendFileSync, createWriteStream, unlink, existsSync, WriteStream } from 'fs';
import { get } from 'http';
import * as path from 'path';
import { resolveDirectory } from './directories';
import { logAndQuit } from './errors';

export const appendToFile = async (filePath: string, newConent: string): Promise<void> => {
  try {
    appendFileSync(filePath, newConent);
  } catch (err) {
    logAndQuit(`ERROR: Appending content to ${filePath} failed`);
  }
}

export const deleteFile = (filePath: string): void => {
  unlink(filePath, (err) => {
    if (err) {
      logAndQuit(`ERROR: Deleting ${filePath} caused an error`);
    };
  });
};

export const download = async (url: string, destination?: string) => new Promise((resolve, reject) => {
  if (!destination) {
    const fileName: string = url.substr(url.lastIndexOf('/') + 1, url.length - 1);
    destination = path.join(resolveDirectory('.'), fileName);
  }
  let file: WriteStream = createWriteStream(destination);
  get(url, (response) => {
    response.pipe(file);
    file.on('finish', () => {
      file.close();
      resolve();
    });
  }).on('error', (err: Error) => {
    logAndQuit(`ERROR: Downloading file from ${url} caused an error: ${err}`);
  });
});

export const fileExists = (filePath: string): boolean => {
  return existsSync(filePath);
}
