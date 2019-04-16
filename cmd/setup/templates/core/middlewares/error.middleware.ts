import { NextFunction, Response } from 'express';
import { HttpException } from '../exceptions/http.exception';
import { Request } from '@pyramid-systems/core';

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
