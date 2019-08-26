import { config } from '../config';
import { Environment } from '../environment';

export const determineEnvironment = (): Environment => {
  return {
    authUrl: config.auth.environmentUrl.integration,
    frontEndUrl: config.ui.environmentUrl.integration,
    seleniumUrl: config.selenium.cloud
  }
};
