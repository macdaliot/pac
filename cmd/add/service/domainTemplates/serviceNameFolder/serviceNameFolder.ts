import {
    autoGeneratedHashKey,
    hashKey,
    rangeKey,
    table,
    attribute
} from '@aws/dynamodb-data-mapper-annotations';

@table('{{.serviceName}}')
export class {{.serviceNamePascal}} {
    @autoGeneratedHashKey()
    id: string;
}