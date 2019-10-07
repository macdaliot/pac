import * as express from 'express';
import { Route, Get, Post, Controller, Request } from 'tsoa';
import { Injectable, createJwt, ILogger } from '@pyramid-systems/core';

@Injectable()
@Route('auth')
export class AuthenticationController extends Controller {
  constructor(private logger: ILogger) {
    super();
  }

  /**
   * Gets a login redirect. If the user is logged in, this will return "OK"
   * Usually called when the user is *not* logged in, and redirects to SAML auth.
   * @param request 
   */
  @Get('login')
  getLogin(@Request() request: express.Request) {
    request.res.status(200).send('OK');
  }

  @Post('callback')
  processCallback(@Request() request: express.Request) {
    const response = request.res;
    createJwt(request, (err, encoded) => {
      if (err) {
        this.logger.error('Error creating JWT ', err);
        response.status(500).send(err);
      } else {
        response.redirect(process.env.APP_ROOT + '/login?' + encoded);
      }
    });
  }
}
