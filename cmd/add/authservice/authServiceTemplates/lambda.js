'use strict'
const awsServerlessExpress = require('aws-serverless-types');
const app = require('./dist/server').default;
const server = awsServerlessExpress.createServer(app);
 
exports.handler = (event, context) => awsServerlessExpress.proxy(server, event, context);
