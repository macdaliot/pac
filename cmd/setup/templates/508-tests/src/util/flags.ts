export const wasFlagProvided = (flag: string): boolean => {
  for (let arg of process.argv) {
    if (arg === flag) {
      return true;
    }
  }
  return false;
}
