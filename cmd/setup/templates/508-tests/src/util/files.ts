import { appendFileSync, closeSync, existsSync, openSync, readFileSync, unlink } from 'fs';

export const appendToFile = async (filePath: string, newContent: string): Promise<void> => {
  try {
    newContent = replaceLineEndings(newContent);
    appendFileSync(filePath, newContent);
  } catch (err) {
    console.error(`ERROR: Appending content to ${filePath} failed`);
  }
}

export const deleteFile = async (filePath: string): Promise<void> => {
  await unlink(filePath, (err) => {
    if (err) {
      console.error(`ERROR: Attempting to delete ${filePath} caused an error`);
    };
  });
};

export const fileExists = (filePath: string): boolean => {
  return existsSync(filePath);
}

export const createBlankFile = async (filePath: string): Promise<void> => {
  // `a` write mode avoids overwriting
  const writeMode: string = 'a';
  const descriptor = openSync(filePath, writeMode);
  closeSync(descriptor);
}

export const readFile = async (filePath: string): Promise<string> => {
  const contents = await readFileSync(filePath).toString();
  return replaceLineEndings(contents);
}

export const trimTopLine = (fileContents: string): string => {
  const lineBreakCharacter: string = '\n';
  let fileLines: Array<string> = fileContents.split(lineBreakCharacter);
  fileLines.shift();
  return fileLines.join(lineBreakCharacter);
}

const replaceLineEndings = (fileContents: string): string => {
  return fileContents.replace('\r\n', '\n').replace('\r', '\n');
}
