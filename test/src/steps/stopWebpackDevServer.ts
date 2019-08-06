import { config } from '../config';
import { killProcessOnPort } from '../util';

export const stopWebpackDevServer = (): void => {
  const programName: string = "node";
  killProcessOnPort(config.devServerPort, programName);
}
