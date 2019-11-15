import * as dotenv from 'dotenv';
dotenv.load();

const projectName: string = "[[.projectName]]";
const authDeployed: boolean = false;
const authEndpoint: string = "/api/auth/login";
const browser: string = "chrome";
const reportName: string = "test-results";
const reportType: string = "csv";
const seleniumCloud: string = `https://selenium.${projectName}.pac.pyramidchallenges.com/wd/hub`;
const timeout: number = 60;
const uiEnvironmentUrlIntegration: string = `https://integration.${projectName}.pac.pyramidchallenges.com`;
const uiRoutes: Array<string> = ['/'];

interface Config {
  auth: {
    deployed: boolean,
    environmentUrl: {
      integration: string
    }
  },
  browser: string,
  credentials: {
    email: string,
    password: string
  },
  report: {
    name: string,
    type: string
  },
  selenium: {
    cloud: string
  },
  timeout: number,
  ui: {
    environmentUrl: {
      integration: string
    },
    routes: Array<string>
  }
}

export const config: Config = {
  auth: {
    deployed: authDeployed,
    environmentUrl: {
      integration: `${uiEnvironmentUrlIntegration}${authEndpoint}`
    }
  },
  browser: browser,
  credentials: {
    email: process.env.TEST_EMAIL || "",
    password: process.env.TEST_PASSWORD || ""
  },
  report: {
    name: reportName,
    type: reportType
  },
  selenium: {
    cloud: seleniumCloud
  },
  timeout: timeout,
  ui: {
    environmentUrl: {
      integration: uiEnvironmentUrlIntegration
    },
    routes: uiRoutes
  }
};
