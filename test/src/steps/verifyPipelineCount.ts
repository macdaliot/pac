import { config } from '../config';
import { assert, JenkinsPipelineStatus, logger, queryJenkinsBuildStatus } from '../util';

export const verifyPipelineCount = async (projectName: string): Promise<void> => {
  const expectedPipelineCount: number = config.jenkins.expectedPipelineCount;
  const assertion: string = `'pac automate jenkins' will create exactly ${expectedPipelineCount} pipelines`;
  const result: Array<JenkinsPipelineStatus> = await queryJenkinsBuildStatus(projectName);
  const actualPipelineCount: number = result.length;
  assert(assertion, actualPipelineCount === expectedPipelineCount);
  logger.info("All Jenkins pipelines were successfully created using 'pac automate jenkins'");
}
