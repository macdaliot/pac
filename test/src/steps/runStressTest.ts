import { config } from '../config';
import { pushToGitHub, verifyWebhookTriggeredPipelineSuccesses } from '.';
import { appendToFile, wasFlagProvided } from '../util';

export const runStressTest = async (projectDirectory: string, projectName: string): Promise<void> => {
  if (wasFlagProvided('--stress')) {
    for (let i = 0; i < config.stressTestBuildCount; i++) {
      changeProjectReadme(projectDirectory);
      pushToGitHub(projectDirectory);
      await verifyWebhookTriggeredPipelineSuccesses(projectName);
    }
  }
}

const changeProjectReadme = (projectDirectory: string): void => {
  const newContent: string = "This is a new line to simulate a change in the project\n";
  appendToFile(`${projectDirectory}/README.md`, newContent);
}
