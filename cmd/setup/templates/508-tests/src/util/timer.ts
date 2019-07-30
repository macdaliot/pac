import * as selenium from 'selenium-webdriver';
import { config } from '../config';

export const sleep = async (seconds: number): Promise<void> => {
  const millisecondsConversion: number = 1000;
  return new Promise(resolve => setTimeout(resolve, seconds * millisecondsConversion));
}

export const startTimer = async (driver: selenium.WebDriver): Promise<NodeJS.Timeout> => {
  const millisecondsConversion = 1000;
  return setTimeout(async () => {
    console.warn('Process exited: The application took more than a minute to load its elements');
    await driver.quit();
    process.exit(22);
  }, config.timeout * millisecondsConversion);
}

export const stopTimer = async (timer: NodeJS.Timeout): Promise<void> => {
  clearTimeout(timer);
}
