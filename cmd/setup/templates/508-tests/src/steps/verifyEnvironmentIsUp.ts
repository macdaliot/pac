import { httpGet } from '../util';

export const verifyEnvironmentIsUp = async (url: string): Promise<boolean> => {
  if (await httpGet(url)) {
    return true;
  } else {
    console.trace("The environment designated to test is not active yet. Failing gracefully. A blank report will be published");
    process.exit(0);
  }
}
