import { DynamoDB } from 'aws-sdk';
import { Repository } from '../../core/database-connectors/amazon-dynamodb/repository';
import { {{.serviceNamePascal}} } from './{{.serviceName}}';

export class {{.serviceNamePascal}}Repository extends Repository<{{.serviceNamePascal}}> {
    constructor(protected dynamoDb: DynamoDB) {
        super(dynamoDb);
    }

    async getById(id: string) {
        return await this.get(id, {{.serviceNamePascal}});
    }

    async getAll(): Promise<{{.serviceNamePascal}}[]> {
        const partners: Partner[] = [];
        const iterator = await this.scan(Partner);
        for await (const record of iterator) {
            partners.push(record);
        }
        return partners;
    }

    async update(partnerUpdate: {{.serviceNamePascal}}) {
        return await this.put(partnerUpdate);
    }

    async deleteById(idToDelete: string) {
        return await this.delete(idToDelete, {{.serviceNamePascal}});
    }
}
