export const logAndQuit = (stderr: string): void => {
  if (stderr === '' || stderr === undefined || stderr === null) {
    return;
  } else {
    const exitCode = 2;
    console.trace(stderr);
    process.exit(exitCode);
  }
}
