import { JenkinsPipelineConfig } from './util';

const createAuthService: boolean = true;
const devServerPort: number = 8080;
const goPath: string = process.env.GOPATH || 'Optionally, manually set $GOPATH here';
const jenkinsCredentialsToVerify: Array<string> = ['gitcredentials', 'SONARLOGIN', 'slack-notification-token', 'TEST_EMAIL', 'TEST_PASSWORD'];
const jenkinsExpectedPipelineCount: number = 5;
const jenkinsPipelinesBuiltByWebhook: Array<string> = ['release-through-staging'];
const jenkinsPipelinesBuiltManually: Array<JenkinsPipelineConfig> = [
  {
    name: 'release-to-production',
    parameters: {
      BUILD_NUMBER: '1'
    }
  }
];
const jenkinsPollingMaxAttempts: number = 50;
const jenkinsPollingSecondDelayBetweenAttempts: number = 30;
const pyramidGitHubAuth: string = 'Basic amRpZWRlcmlrc0Bwc2ktaXQuY29tOkRpZWRyZV4yMDE4';
const servicesToCreate: number = 2;
const stressTestBuildCount: number = 50;
const tempProjectDir: string = '~/tmpPac';

interface Config {
  createAuthService: boolean,
  devServerPort: number,
  goPath: string,
  jenkins: {
    credentials: {
      username: string,
      password: string
    },
    credentialsToVerify: Array<string>,
    expectedPipelineCount: number,
    pipelines: {
      builtByWebhook: Array<string>,
      builtManually: Array<JenkinsPipelineConfig>
    },
    polling: {
      maxAttempts: number,
      secondDelayBetweenAttempts: number
    }
  },
  pyramidGitHubAuth: string,
  servicesToCreate: number,
  stressTestBuildCount: number,
  tempProjectDir: string
}

export const config: Config = {
  createAuthService: createAuthService,
  devServerPort: devServerPort,
  goPath: goPath,
  jenkins: {
    credentials: {
      username: 'pyramid',
      password: 'systems'
    },
    credentialsToVerify: jenkinsCredentialsToVerify,
    expectedPipelineCount: jenkinsExpectedPipelineCount,
    pipelines: {
      builtByWebhook: jenkinsPipelinesBuiltByWebhook,
      builtManually: jenkinsPipelinesBuiltManually
    },
    polling: {
      maxAttempts: jenkinsPollingMaxAttempts,
      secondDelayBetweenAttempts: jenkinsPollingSecondDelayBetweenAttempts
    }
  },
  pyramidGitHubAuth: pyramidGitHubAuth,
  servicesToCreate: servicesToCreate,
  stressTestBuildCount: stressTestBuildCount,
  tempProjectDir: tempProjectDir,
};
