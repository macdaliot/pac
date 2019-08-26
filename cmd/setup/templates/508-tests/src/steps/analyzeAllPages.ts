import * as selenium from 'selenium-webdriver';
import { analyzePage, loadPage } from '../util';
import { config } from '../config';
import { Environment } from '../environment';

export const analyzeAllPages = async (environment: Environment, driver: selenium.WebDriver): Promise<void> => {
  let index: number = 0;
  for await (const route of config.ui.routes) {
    const fullUrl: string = `${environment.frontEndUrl}${route}`;
    await loadPage(fullUrl, driver);
    await analyzePage(fullUrl, index, driver);
    index++;
  }
};
