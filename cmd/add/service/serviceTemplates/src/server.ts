import * as express from 'express';
import apiRouter from './routes/routes'
import * as cors from 'cors';
import * as loggerMiddleware from './middleware/logger/loggerMiddleware';
import { Strategy as JwtStrategy, StrategyOptions as JwtStrategyOptions, ExtractJwt } from 'passport-jwt';
import * as passport from 'passport'
import * as dotenv from 'dotenv';

const app = express();
const port = 3000;
/* TODO: update error handling */
/* need configMap here */
dotenv.load();

let opts: JwtStrategyOptions = {
  jwtFromRequest: ExtractJwt.fromAuthHeaderAsBearerToken(),
  secretOrKey: process.env.JWT_SECRET,
  issuer: process.env.JWT_ISSUER,
}
//opts.audience = 'yoursite.net';
passport.use(new JwtStrategy(opts, function (jwt_payload, done) {
  done(jwt_payload);
}));

passport.initialize();
app
  .use(cors())

  /* parse middleware */
  /* https://expressjs.com/en/api.html#express.json */
  .use(express.json())
  .use(express.urlencoded({ extended: true }))

  /* logging middleware, order matters */
  .use('', loggerMiddleware._loggers)

  /* routes */
  .use('/api', apiRouter)

const isInLambda = !!process.env.LAMBDA_TASK_ROOT;
if (isInLambda) {
  module.exports = app;
} else {
  app.listen(port, () => console.log(` + "`" + `{{.serviceName}} is running on port ${port}!` + "`" + `))
}  
