import { config } from '../config';
import { assert, logger, printDirectoryContents, resolveDirectory } from '../util';
import { join } from 'path';

export const verifyOnDiskContent = (projectName: string): string => {
  const itemsCreatedOutsideOfTemplates: number = 4; // (1) .git folder, (2) node_modules/ folder, (3) .pac.json, and (3) package-lock.json
  const assertion: string = `The new project directory has ${itemsCreatedOutsideOfTemplates} more items in its root folder than the PAC setup template directory`;
  const pacRootTemplatesDirectory: string = `${config.goPath}/src/github.com/PyramidSystemsInc/pac/cmd/setup/templates`;
  const projectDirectory: string = join(resolveDirectory(config.tempProjectDir), projectName);
  const pacTemplatesFileList: Array<string> = printDirectoryContents(pacRootTemplatesDirectory);
  const newProjectFileList: Array<string> = printDirectoryContents(projectDirectory);
  assert(assertion, pacTemplatesFileList.length === newProjectFileList.length - itemsCreatedOutsideOfTemplates);
  logger.info("The number of directories that were expected to be created were created");
  return projectDirectory;
}
