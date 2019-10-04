import { MongoRepository } from '../../core/database-connectors/amazon-documentdb/repository';
import { {{.serviceNamePascal}} } from './{{.serviceName}}';
import { Injectable } from '@pyramid-systems/core';

@Injectable()
export class {{.serviceNamePascal}}Repository extends MongoRepository<{{.serviceNamePascal}}> {
    //add additional queries by accessing this._collection

}
