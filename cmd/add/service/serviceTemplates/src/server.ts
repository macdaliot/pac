import * as dotenv from 'dotenv';
dotenv.load(); // must remain here before other imports

import * as express from 'express';
import * as cors from 'cors';
import * as loggerMiddleware from './middleware/logger/loggerMiddleware';
import { Strategy as JwtStrategy, StrategyOptions as JwtStrategyOptions, ExtractJwt } from 'passport-jwt';
import * as passport from 'passport'
import { generateRandomString } from './utility/functions'

const app = express();
/* TODO: update error handling */
/* need configMap here */

import apiRouter from './routes/routes'

let opts: JwtStrategyOptions = {
  jwtFromRequest: ExtractJwt.fromAuthHeaderAsBearerToken(),
  secretOrKey: process.env.JWT_SECRET || generateRandomString(),
  issuer: process.env.JWT_ISSUER || generateRandomString()
}
let jwtStrategy = new JwtStrategy(opts, function (jwt_payload, done) {
  done(null, jwt_payload);
});

passport.use(jwtStrategy);

app
  .use(cors())
  .use(passport.initialize())
  
  /* parse middleware */
  /* https://expressjs.com/en/api.html#express.json */
  .use(express.json())
  .use(express.urlencoded({ extended: true }))

  /* logging middleware, order matters */
  .use('', loggerMiddleware._loggers)

  /* routes */
  .use('/api', apiRouter)

export default app;
