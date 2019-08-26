import { config } from './config';
import {
  analyzeAllPages,
  buildWebDriver,
  combineReports,
  determineEnvironment,
  login,
  verifyEnvironmentIsUp
} from './steps';

const main = async () => {
  let driver = await buildWebDriver();
  try {
    const environment = determineEnvironment();
    await verifyEnvironmentIsUp(environment.frontEndUrl);
    await login(environment, driver);
    await analyzeAllPages(environment, driver);
    await combineReports();
  } catch(err) {
    console.error(err);
    await driver.quit();
    const exitCode = 4;
    process.exit(exitCode);
  } finally {
    await driver.quit();
  }
};

main().catch(console.error);
