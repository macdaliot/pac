import {
  appendToFile,
  createBlankFile,
  deleteFile,
  fileExists,
  getCurrentWorkingDirectory,
  readFile,
  trimTopLine
} from '../util';
import { config } from '../config';

export const combineReports = async (): Promise<void> => {
  const currentWorkingDirectory: string = getCurrentWorkingDirectory();
  const reportsDirectory: string = `${currentWorkingDirectory}/../..`;
  const combinedReportPath: string = `${reportsDirectory}/${config.report.name}.${config.report.type}`;
  let index: number = 0;
  for (;;) {
    const pageReportPath: string = `${reportsDirectory}/${config.report.name}-${index}.${config.report.type}`;
    if (!fileExists(pageReportPath)) {
      break;
    }
    let contents = await readFile(pageReportPath);
    if (index === 0) {
      await createBlankFile(combinedReportPath);
    } else {
      contents = trimTopLine(contents);
    }
    await appendToFile(combinedReportPath, contents);
    await deleteFile(pageReportPath);
    index++;
  }
};
