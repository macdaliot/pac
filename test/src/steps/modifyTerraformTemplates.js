// TODO: Replace with a module/function that has TypeScript definitions
//   See: https://stackoverflow.com/questions/14177087/replace-a-string-in-a-file-with-nodejs

import * as replace from 'replace-in-file';
import { logger } from '../util';

export const modifyTerraformTemplates = async (projectDirectory) => {
  const options = {
    files: [ `${projectDirectory}/terraform/**/lambda*.tf` ],
    from: [ / source_code_hash/g, ],
    to: [ '#source_code_hash' ]
  };
  try {
    await replace(options, (error, files) => {
      if (error) {
        return logger.error(`An error occurred while modifying Terraform templates prior to running 'pac clean-cloud': ${error}`);
      }
    });
  } catch (error) {
    logger.error(`An error occurred while modifying Terraform templates prior to running 'pac clean-cloud': ${error}`);
  }
}
