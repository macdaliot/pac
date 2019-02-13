import { DynamoDB } from '../src/database/dynamo.db';
import * as AWSMock from 'aws-sdk-mock';
import { QueryOutput } from 'aws-sdk/clients/dynamodb';
import { mockDbResponses, mockDBResponse } from './mockDynamoDb';
import { serviceName, projectName } from '../src/config/appInfo.config';

describe(`${projectName}-${serviceName}: DynamoDB`, function () {
    let database: DynamoDB;

    beforeEach(done => {
        /* put any prerequisite here */
        mockDbResponses();
        database = new DynamoDB();
        done();
    })

    /* Trivial tests and stubs for each of the default functions in DynamoDB.ts
    */
    describe('Should implement DB interface', () => {
        describe('query ', async () => {
            it('should return DB Query', async () => {
                let dbQueryResult = await database.query({ TableName: '' });
                expect(dbQueryResult).toEqual(mockDBResponse);
            });
        });

        describe('create ', () => {
            it('should return DB put', async () => {
                let dbQueryResult = await database.create({ test: "test" });
                expect(dbQueryResult).toEqual(mockDBResponse);
            });
        });

        describe('update', () => {
            it('should return DB Update', async () => {
                let dbQueryResult = await database.update({ TableName: '' }, {});
                expect(dbQueryResult).toEqual(mockDBResponse);
            });
        });

        describe('delete', () => {
            it('should not error', async () => {
                let dbQueryResult = await database.delete({ TableName: '' });
            });
        });
    })

    afterEach(done => {
        /* Put any cleanup here */
        AWSMock.restore('DynamoDB.DocumentClient');
        done();
    });
});

