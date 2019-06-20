import { NextFunction, Response, Request } from 'express';
import { HttpException } from '../exceptions/http.exception';

export function errorMiddleware(
    error: HttpException,
    request: Request,
    response: Response,
    next: NextFunction
) {
  const status = error.status || 500;
  const message = error.message || 'Something went wrong';
  request.log.error(message);
  response.status(status).send({
    status,
    message
  });
}
