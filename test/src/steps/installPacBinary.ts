import { config } from '../config';
import { CommandResult, logger, logAndQuit, runSync, tryRunSync } from '../util';

export const installPacBinary = (): void => {
  removePacBinaryIfExists();
  installPac();
}

const removePacBinaryIfExists = (): void => {
  tryRunSync(`rm ${config.goPath}/bin/pac`);
}

const installPac = (): void => {
  const result: CommandResult = runSync('packr install github.com/PyramidSystemsInc/pac');
  logAndQuit(result.stderr);
  logger.info("The pac binary was installed");
}
