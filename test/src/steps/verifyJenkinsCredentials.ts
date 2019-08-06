import { config } from '../config';
import { doesJenkinsCredentialExist, logger, poll } from '../util';

export const verifyJenkinsCredentials = async (projectName: string): Promise<void> => {
  const credentialsToVerify: Array<string> = config.jenkins.credentialsToVerify;
  const maxAttempts: number = config.jenkins.polling.maxAttempts;
  const secondDelayBetweenAttempts: number = config.jenkins.polling.secondDelayBetweenAttempts;
  for (const credentialId of credentialsToVerify) {
    await poll(() => doesJenkinsCredentialExist(credentialId, projectName), secondDelayBetweenAttempts, maxAttempts);
    logger.info(`The ${credentialId} credential exists in Jenkins`);
  }
}
