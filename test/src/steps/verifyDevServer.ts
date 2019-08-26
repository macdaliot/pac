import * as axios from 'axios';
import { config } from '../config';
import { assert, httpGet, logger, sleep } from '../util';

export const verifyDevServer = async (): Promise<void> => {
  const assertion: string = "An HTTP Get request against the Webpack development server gives back a 200/OK response";
  await waitForDevServerToStart();
  const url: string = 'http://localhost:' + config.devServerPort;
  const response: axios.AxiosResponse<any> = await httpGet(url);
  assert(assertion, response.status === 200);
  logger.info(assertion);
}

const waitForDevServerToStart = async (): Promise<void> => {
  const delay: number = 7;
  await sleep(delay);
}
