import * as express from 'express';
import { Route, Get, Post, Controller, Request } from 'tsoa';
import { Injectable, createJwt, ILogger } from '@pyramidlabs/core';

@Injectable()
@Route('auth')
export class AuthenticationController extends Controller {
  constructor(private logger: ILogger) {
    super();
  }

  @Get('login')
  getLogin(@Request() request: express.Request) {
    // Leo: I'm leaving this here. Mike implemented it this way, but I don't think this will ever be hit unless there's a client that uses saml authentication.
    request.res.status(200).send('OK');
  }

  @Post('callback')
  login(@Request() request: express.Request) {
    const response = request.res;
    createJwt(request, (err, encoded) => {
      if (err) {
        this.logger.error('Error creating JWT ', err);
        response.status(500).send(err);
      } else {
        response.redirect(process.env.APP_ROOT + '/api/auth/login?' + encoded);
      }
    });
  }
}
