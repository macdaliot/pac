import { Controller } from 'tsoa';
import {
    decorate,
    Container,
    injectable as Injectable
} from 'inversify';
import "reflect-metadata";
import { config } from 'aws-sdk';
import { awsConfig } from './database-connectors/config';

decorate(Injectable(), Controller);
let env = process.env.ENVIRONMENT || 'cloud';
config.update(awsConfig[env]);
const iocContainer = new Container();

export {
    iocContainer,
    Injectable
};
