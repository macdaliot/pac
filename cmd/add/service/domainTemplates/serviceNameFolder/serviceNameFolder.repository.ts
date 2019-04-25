import { DataMapper } from '@aws/dynamodb-data-mapper';
import { Repository } from '../../core/database-connectors/amazon-dynamodb/repository';
import { {{.serviceNamePascal}} } from './{{.serviceName}}';
import { Injectable } from '@pyramid-systems/core';

@Injectable()
export class {{.serviceNamePascal}}Repository extends Repository<{{.serviceNamePascal}}> {
    constructor(dataMapper: DataMapper) {
        super(dataMapper);
    }

    async getById(id: string) {
        return await this.get(id, {{.serviceNamePascal}});
    }

    async getAll(): Promise<{{.serviceNamePascal}}[]> {
        const partners:  {{.serviceNamePascal}}[] = [];
        const iterator = await this.scan({{.serviceNamePascal}});
        for await (const record of iterator) {
            partners.push(record);
        }
        return partners;
    }

    async add({{.serviceName}}: {{.serviceNamePascal}}) {
        return await this.post({{.serviceName}}, {{.serviceNamePascal}});
    }

    async update(idToUpdate: string, {{.serviceName}}: {{.serviceNamePascal}}) {
        return await this.put(idToUpdate, {{.serviceName}}, {{.serviceNamePascal}});
    }

    async deleteById(idToDelete: string) {
        return await this.delete(idToDelete, {{.serviceNamePascal}});
    }
}
