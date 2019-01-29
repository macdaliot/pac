import * as dotenv from 'dotenv';
dotenv.load(); // must remain here before other imports

import * as express from 'express';
import bodyParser = require('body-parser');
import * as passport from 'passport'
import { Strategy as SamlStrategy} from 'passport-saml';
import { authRouter } from './authRouter';
import { generateRandomString } from './functions'

const port: number = 3000;
const serviceName: string = 'auth';

const signingCert = process.env.SAML_SIGNER ?
  Buffer.from(process.env.SAML_SIGNER, 'base64').toString('ascii') :
  generateRandomString();

const app = express();
app.use(bodyParser.json());
app.use(bodyParser.urlencoded({extended: false}));

var samlStrategy = new SamlStrategy(
  {
    callbackUrl: process.env.SAML_CALLBACK || `http://localhost/${generateRandomString()}`,
    entryPoint: 
      (process.env.AUTH0_CLIENT_ID && process.env.AUTH0_DOMAIN) ? 
        `https://${process.env.AUTH0_DOMAIN}/samlp/${process.env.AUTH0_CLIENT_ID}` :
        `https://localhost/samlp/${generateRandomString()}`,
    issuer: 
      (process.env.AUTH0_DOMAIN) ? 
        'urn:' + process.env.AUTH0_DOMAIN + '-saml' :
        `urn:${generateRandomString()}`,
    cert: signingCert,
    acceptedClockSkewMs: 60000
  },
  function (user, done) {
      return done(null, user, null);
  });


app.use(passport.initialize());
passport.use(samlStrategy);

app.use('/api/auth', authRouter);

export default app;