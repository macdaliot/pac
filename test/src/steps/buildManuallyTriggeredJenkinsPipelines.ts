import { config } from '../config';
import { buildJenkinsPipeline, logger } from '../util';

export const buildManuallyTriggeredJenkinsPipelines = async (projectName: string): Promise<void> => {
  for (const pipelineConfig of config.jenkins.pipelines.builtManually) {
    await buildJenkinsPipeline(pipelineConfig, projectName);
    logger.info(`Successfully triggered the build of the ${pipelineConfig.name} pipeline`);
  }
}
