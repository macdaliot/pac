import { NextFunction, Response } from 'express';
import { HttpException } from '../exceptions/http.exception';
import { Request } from '@pyramidlabs/core';

export function errorMiddleware(
    error: HttpException,
    request: Request,
    response: Response,
    next: NextFunction
) {
  request.log.error(JSON.stringify(error));
  const status = error.status || 500;
  const message = error.message || 'Something went wrong';
  response.status(status).send({
    status,a
    message
  });
}
