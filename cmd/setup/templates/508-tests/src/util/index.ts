import { analyzePage } from './accessibility';
import { getCurrentWorkingDirectory } from './directories';
import {
  appendToFile,
  createBlankFile,
  deleteFile,
  fileExists,
  readFile,
  trimTopLine
} from './files';
import { wasFlagProvided } from './flags';
import { httpGet } from './http';
import {
  buildChromeDriver,
  findElementByClassName,
  findElementById,
  findElementByName,
  loadPage
} from './selenium';
import { sleep, startTimer, stopTimer } from './timer';

export {
  analyzePage,
  appendToFile,
  buildChromeDriver,
  createBlankFile,
  deleteFile,
  fileExists,
  findElementByClassName,
  findElementById,
  findElementByName,
  getCurrentWorkingDirectory,
  httpGet,
  loadPage,
  readFile,
  sleep,
  startTimer,
  stopTimer,
  trimTopLine,
  wasFlagProvided
};
