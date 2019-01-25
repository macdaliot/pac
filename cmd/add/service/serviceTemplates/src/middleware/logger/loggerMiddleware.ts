/* TODO: use not default logger like https://github.com/winstonjs/winston 
    need to create predefined logger level groups
*/
import { Logger } from './logger';
import * as e from 'express';
import * as _ from 'lodash';

let _logger = new Logger();
export const routeLogger = (req: e.Request, res: e.Response, next) => {
    _logger.info(`${req.method}: ${req.path}`);
    next();
}
export const timeLogger = (req: e.Request, res: e.Response, next) => {
    _logger.info(Date());
    next();
}

export const bodyLogger = (req: e.Request, res: e.Response, next) => {
    if (!_.isEmpty(req.body)) {
        _logger.info(`body: ${JSON.stringify(req.body, null, 2)}`);
    }
    next();
}

export const _loggers = [
    timeLogger,
    routeLogger,
    bodyLogger
]

/* TODO: need user logger */