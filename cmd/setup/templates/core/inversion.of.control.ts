import { Controller } from 'tsoa';
import {
    decorate,
    Container,
    injectable as Injectable,
    inject as Inject
} from 'inversify';
import "reflect-metadata";
import { config, DynamoDB, Service } from 'aws-sdk';
import { awsConfig } from './database-connectors/amazon-dynamodb/config';

decorate(Injectable(), Controller);
let env = process.env.ENVIRONMENT || 'cloud';
config.update(awsConfig[env]);
const iocContainer = new Container();
iocContainer.bind(DynamoDB).toDynamicValue(() => {
    return new DynamoDB()
});
export {
    iocContainer,
    Inject,
    Injectable
};
