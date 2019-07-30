import * as selenium from 'selenium-webdriver';
import * as chrome from 'selenium-webdriver/chrome';
import { sleep, startTimer, stopTimer } from './timer';
import { config } from '../config';

export const buildChromeDriver = async (): Promise<selenium.WebDriver> => {
  const browser: string = 'chrome';
  return await new selenium.Builder()
    .forBrowser(browser)
    .withCapabilities(selenium.Capabilities.chrome())
    .setChromeOptions(new chrome.Options())
    .usingServer(config.selenium.cloud)
    .build();
}

export const findElementByClassName = async (className: string, driver: selenium.WebDriver): Promise<selenium.WebElement> => {
  process.stdout.write("\rSearching for element by class name: " + className + " ...");
  const locator: selenium.By = selenium.By.className(className);
  return await findElement(locator, driver);
}

export const findElementById = async (id: string, driver: selenium.WebDriver): Promise<selenium.WebElement> => {
  process.stdout.write("\rSearching for element by id: " + id + " ...");
  const locator: selenium.By = selenium.By.id(id);
  return await findElement(locator, driver);
}

export const findElementByName = async (name: string, driver: selenium.WebDriver): Promise<selenium.WebElement> => {
  process.stdout.write("\rSearching for element by name: " + name + " ...");
  const locator: selenium.By = selenium.By.name(name);
  return await findElement(locator, driver);
}

export const findElement = async (locator: selenium.By, driver: selenium.WebDriver): Promise<selenium.WebElement> => {
  let timer = await startTimer(driver);
  let element = await driver.wait(selenium.until.elementLocated(locator));
  console.log(" Located element");
  await stopTimer(timer);
  return element;
}

export const loadPage = async (url: string, driver: selenium.WebDriver): Promise<void> => {
  try {
    process.stdout.write(`\rLoading page: ${url} ...`);
    await driver.get(url);
    await sleep(3);
    console.log(" DONE");
  } catch(err) {
    console.error(err);
  }
}
