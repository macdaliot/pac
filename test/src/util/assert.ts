import { logAndQuit } from './errors';

export const assert = (assertion: string, condition: boolean): void => {
  if (!condition) {
    logAndQuit(`ERROR: The following assertion failed: ${assertion}`);
  }
}
