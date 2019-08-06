import { config } from '../config';
import { isPortInUse, logger, logAndQuit } from '../util';

export const validateTestStartConditions = async (): Promise<void> => {
  const devServerPort: number = config.devServerPort;
  if (await devServerPortIsInUse(devServerPort)) {
    logAndQuit(`ERROR: Port ${devServerPort} must be free before running pac-test in order to properly test the local Webpack development server`);
  }
  logger.info("Config values located in config.ts are valid");
}

const devServerPortIsInUse = async (devServerPort: number): Promise<boolean> => {
  return await isPortInUse(devServerPort);
}
