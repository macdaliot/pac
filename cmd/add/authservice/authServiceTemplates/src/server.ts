import * as express from 'express';
import bodyParser = require('body-parser');
import * as passport from 'passport'
import { Strategy as SamlStrategy} from 'passport-saml';
import { authRouter } from './authRouter';
import * as dotenv from 'dotenv';

const port: number = 3000;
const serviceName: string = 'auth';

dotenv.load();
const signingCert = Buffer.from(process.env.SAML_SIGNER, 'base64').toString('ascii');
console.log(signingCert);
const app = express();
app.use(bodyParser.json());
app.use(bodyParser.urlencoded({extended: false}));

var samlStrategy = new SamlStrategy(
  {
    callbackUrl: process.env.SAML_CALLBACK || 'http://localhost:3000/api/auth/callback',
    entryPoint: 
      (process.env.AUTH0_CLIENT_ID && process.env.AUTH0_DOMAIN) ? 
        "https://" + process.env.AUTH0_CLIENT_ID + "/samlp/" + process.env.AUTH0_DOMAIN :
        "https://pyramidsystems.auth0.com/samlp/PJqs70Pr0VRH67aR2TnHf4Sn6DDldiNR",
    issuer: 
      (process.env.AUTH0_CLIENT_ID) ? 
        'urn:' + process.env.AUTH0_CLIENT_ID + '-saml' :
        'urm:pyramidsystems.auth0.com-saml',
    passReqToCallback: true,
    cert: signingCert,
    acceptedClockSkewMs: 60000
  },
  function (req, user, done) {
      //user.loginKey = 'arn:aws:iam::118104210923:saml-provider/auth0-provider';
      console.log("got user");
      //console.log(user);
      //user['token'] = req.body.SAMLResponse;
      return done(null, user, null);
  });


app.use(passport.initialize());
passport.use(samlStrategy);

// You can use this section to keep a smaller payload
passport.serializeUser(function (user, done) {
  done(null, user);
});

passport.deserializeUser(function (user, done) {
  done(null, user);
});

app.use('/api/auth', authRouter);

export default app;