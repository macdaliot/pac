import { Logger } from '../middleware/logger/logger';
let _logger = new Logger();
export const getError = error => {
    /* put error mapping here */
    return 500;
}
export const errorHandler = (err, req, res, next) => {
    _logger.error(`error: ${err}`);
    res.status(getError(err)).send(err);
}
