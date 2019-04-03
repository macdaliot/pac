jest.mock("../src/config", () => ({projectName: 'testProject', serviceName: 'testService', awsConfig: {cloud: '', local: ''}}))

import { DynamoDB } from '../src/database/dynamo.db';
import * as AWSMock from 'aws-sdk-mock';
import { QueryOutput } from 'aws-sdk/clients/dynamodb';
import { mockDbResponses, mockDBResponse, MockOptions } from './mockDynamoDb';

// mock config

describe('DynamoDB', function () {
    let database: DynamoDB;
    let needClean = false;
    const recreateDB = (options?: MockOptions) => {
        if (needClean)
            AWSMock.restore('DynamoDB.DocumentClient');
        mockDbResponses(options);
        database = new DynamoDB();
        needClean = true;
    }
    beforeEach(done => {
        /* put any prerequisite here */
        recreateDB();
        done();
    })

    describe('Should generate query parameters given various inputs', () => {
        it('should generate a key string for one', () => {
            let result = database.createKeyConditionExpression({ id: 12345 });
            expect(result).toEqual("id = :_id");
        })
        it('should generate a key string for multiple', () => {
            let result = database.createKeyConditionExpression({ id: 12345, id2: 12345 });
            expect(result).toEqual("id = :_id and id2 = :_id2");
        })
        it('should generate key values for one', () => {
            let result = database.createExpressionAttributeValues({ id: 12345});
            expect(result).toEqual({":_id": 12345});
        })
        it('should generate key values for multiple', () => {
            let result = database.createExpressionAttributeValues({ id: 12345, id2: 12346});
            expect(result).toEqual({":_id": 12345, ":_id2": 12346});
        })
        it('should not crash when passed an empty object', () => {
            let result = database.createExpressionAttributeValues({});
            expect(result).toEqual({});
        })
    });
    describe('Should build dynamo pieces appropriately', () => {
        it('should build get params header', () => {
            const blob = {id: 12345};
            let result = database.buildGetParams(blob);
            expect(result.TableName).toEqual("pac-testProject-i-testService");
            expect(result.Key).toEqual(blob);
        })
    })
    /* Trivial tests and stubs for each of the default functions in DynamoDB.ts
    */
    describe('Should implement DB interface', () => {
        describe('query ', async () => {
            it('should return DB Query', async () => {
                let dbQueryResult = await database.query({ TableName: '' });
                expect(dbQueryResult).toEqual(mockDBResponse);
            });
            it('should return DB Query with no params', async () => {
                let dbQueryResult = await database.query({});
                expect(dbQueryResult).toEqual(mockDBResponse);
            });
            it('should error on error', async () => {
                recreateDB({ queryResponse: new Error()})
                try {
                    let dbQueryResult = await database.query({ TableName: '' });
                    fail("Should have errored");
                }
                catch (err){

                }
            });
        });

        describe('create ', () => {
            it('should return DB put', async () => {
                let dbQueryResult = await database.create({ test: "test" });
                expect(dbQueryResult).toEqual({});
            });
            it('should error on error', async () => {
                recreateDB({ putResponse: new Error()})
                try {
                    let dbQueryResult = await database.create({ test: "test"  });
                    fail("Should have errored");
                }
                catch (err){  
                }
            });

        });

        describe('update', () => {
            it('should return DB Update', async () => {
                let dbQueryResult = await database.update({ TableName: '' }, {});
                expect(dbQueryResult).toEqual(mockDBResponse);
            });
            it('should error on error', async () => {
                recreateDB({ updateResponse: new Error()})
                try {
                    let dbQueryResult = await database.update({ test: "test"  }, { test: "test"  });
                    fail("Should have errored");
                }
                catch (err){  
                }
            });
        });

        describe('delete', () => {
            it('should not error', async () => {
                let dbQueryResult = await database.delete({ test: "test"  });
            });
            it('should error on error', async () => {
                recreateDB({ deleteResponse: new Error()})
                try {
                    let dbQueryResult = await database.delete({ test: "test"  });
                    fail("Should have errored");
                }
                catch (err){  
                }
            });
        });
    })

    afterEach(done => {
        /* Put any cleanup here */
        AWSMock.restore('DynamoDB.DocumentClient');
        needClean = false;
        done();
    });
});

