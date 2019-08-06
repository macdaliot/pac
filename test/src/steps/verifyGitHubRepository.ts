import * as axios from 'axios';
import { config } from '../config';
import { assert, httpGetWithAuth, logger } from '../util';

export const verifyGitHubRepository = async (projectName: string): Promise<void> => {
  const assertion: string = "GitHub repository for the " + projectName + " project was created";
  const gitHubRepositoryStatusUrl: string = 'https://api.github.com/repos/PyramidSystemsInc/' + projectName;
  const result: axios.AxiosResponse<any> = await httpGetWithAuth(gitHubRepositoryStatusUrl, config.pyramidGitHubAuth);
  assert(assertion, result.status === 200);
  logger.info(assertion);
}
