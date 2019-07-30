import * as selenium from 'selenium-webdriver';
import {
  findElementByClassName,
  findElementById,
  findElementByName,
  loadPage,
  sleep
} from '../util';
import { config } from '../config';
import { Environment } from '../environment';

export const login = async (environment: Environment, driver: selenium.WebDriver): Promise<void> => {
  if (config.auth.deployed) {
    await loadPage(environment.authUrl, driver);
    let emailField = await findElementById('1-email', driver);
    await emailField.sendKeys(process.env.TEST_EMAIL);
    let passwordField = await findElementByName('password', driver);
    await passwordField.sendKeys(process.env.TEST_PASSWORD);
    let button = await findElementByName('submit', driver);
    await button.click();
    // Sleeping here is needed because the next page will load before the
    //   redirects are followed and the token gets stored
    await sleep(3);
  }
};
