import { logAndQuit } from './errors';

interface ActionFunction {
  (): Promise<boolean>
}

// NOTE: action is a function that either returns `false` to keep the polling
//   going or `true` to indicate a success and end the polling
export const poll = async (action: ActionFunction, secondDelayBetweenAttempts: number, maxAttempts: number): Promise<void> => {
  let currentAttemptIndex: number = 0;
  while (currentAttemptIndex < maxAttempts) {
    const result: boolean = await action();
    if (result) {
      return;
    }
    currentAttemptIndex++;
    if (currentAttemptIndex === maxAttempts) {
      logAndQuit("ERROR: Polling failed.")
    }
    await sleep(secondDelayBetweenAttempts);
  }
}

// NOTE: sleep() must be used with the await keyword: i.e. `await sleep(5)`
export const sleep = async (seconds: number) => new Promise((resolve, reject) => {
  const milliseconds: number = 1000;
  setTimeout(() => { resolve(); }, (seconds * milliseconds));
});
