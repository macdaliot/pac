import {
  decorate,
  Container,
  inject as Inject,
  injectable as Injectable
} from 'inversify';
import { Controller } from 'tsoa';
import "reflect-metadata";
import { Connectors, DynamoRepository } from './database-connectors';

decorate(Injectable(), Controller);

const iocContainer: Container = new Container();
iocContainer.bind(Connectors.DynamoDB).to(DynamoRepository);

export {
  iocContainer,
  Inject,
  Injectable
};