import * as selenium from 'selenium-webdriver';
import { buildChromeDriver } from '../util';
import { config } from '../config';

export const buildWebDriver = async (): Promise<selenium.WebDriver> => {
  switch(config.browser) {
    case 'chrome':
      return await buildChromeDriver();
    default:
      console.error("ERROR: Unable to build the web driver. An unsupported browser was specified");
  }
};
