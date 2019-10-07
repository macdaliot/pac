import { Controller } from 'tsoa';
import {
    decorate,
    Container,
    injectable as Injectable,
    inject as Inject
} from 'inversify';
import "reflect-metadata";
import { config, DynamoDB } from 'aws-sdk';
import { awsConfig } from './database-connectors/amazon-dynamodb/config';
import { DataMapper } from '@aws/dynamodb-data-mapper';

decorate(Injectable(), Controller);
let env = process.env.ENVIRONMENT || 'cloud';
config.update(awsConfig[env]);
const iocContainer = new Container();
iocContainer.bind(DataMapper).toDynamicValue(() => {
    return new DataMapper({
        client: new DynamoDB()
    })
})
export {
    iocContainer,
    Inject,
    Injectable
};
