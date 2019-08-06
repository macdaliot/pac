import { config } from '../config';
import { CommandResult, logger, logAndQuit, randomString, runSync } from '../util';

export const createPacServices = async (projectDirectory: string): Promise<void> => {
  await createAuthService(projectDirectory);
  await createRegularServices(projectDirectory);
}

const createAuthService = async (projectDirectory: string): Promise<void> => {
  if (config.createAuthService) {
    const result: CommandResult = runSync(`pac add authService`, projectDirectory);
    logAndQuit(result.stderr);
    logger.info(`The authentication service was successfully created`);
  }
}

const createRegularServices = async (projectDirectory: string): Promise<void> => {
  for (let i = 0; i < config.servicesToCreate; i++) {
    const serviceNameLength: number = 7;
    const serviceName: string = randomString(serviceNameLength);
    const result: CommandResult = runSync(`pac add service --name ${serviceName}`, projectDirectory);
    logAndQuit(result.stderr);
    logger.info(`The ${serviceName} service was successfully created`);
  }
}
