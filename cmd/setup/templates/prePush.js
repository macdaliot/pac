const { spawnSync } = require('child_process')

const runSync = (cmd = "") => {
    const parts = cmd.split(" ");
    if (parts.length > 0);
    const proc = spawnSync(parts[0], parts.slice(1));
    console.log(proc.stdout.toString());
    console.log(proc.stderr.toString());
    return proc.status;
}
const body = () => {
    let returnCode = runSync('git stash --include-untracked');
    if (returnCode) return returnCode;

    returnCode = runSync('npx lerna run build');
    if (returnCode) return returnCode;

    returnCode = runSync('npx lerna run test');

    runSync('git stash pop');
    return returnCode;
}
process.exit(body());