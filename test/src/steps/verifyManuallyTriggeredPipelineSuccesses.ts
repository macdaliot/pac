import * as axios from 'axios';
import { config } from '../config';
import {
  checkForSuccessfulPipelineBuild,
  httpGetWithBasicAuth,
  logger,
  logAndQuit,
  poll,
  sleep
} from '../util';

export const verifyManuallyTriggeredPipelineSuccesses = async (projectName: string): Promise<void> => {
  const maxAttempts: number = config.jenkins.polling.maxAttempts;
  const secondDelayBetweenAttempts: number = config.jenkins.polling.secondDelayBetweenAttempts;
  for (const pipelineConfig of config.jenkins.pipelines.builtManually) {
    const pipelineName: string = pipelineConfig.name;
    await poll(() => checkForSuccessfulPipelineBuild(pipelineName, projectName), secondDelayBetweenAttempts, maxAttempts);
    logger.info(`The ${pipelineName} pipeline was built successfully`);
  }
}
