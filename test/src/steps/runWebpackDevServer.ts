import { logger, runDetached, runSync } from '../util';
import { join } from 'path';

export const runWebpackDevServer = async (projectDirectory: string): Promise<void> => {
  const appDirectory: string = join(projectDirectory, "app");
  runSync('npm i', appDirectory);
  logger.info("The front-end dependencies were successfully installed by NPM");
  runDetached('npm run dev-server', appDirectory);
  logger.info("The front-end development server was started");
}
