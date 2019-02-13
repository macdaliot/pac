import { DefaultService } from '../src/services/defaultService';
import { MockDynamoDB, mockDBobject } from './mockDynamoDb';
import { serviceName, projectName } from '../src/config/appInfo.config';

describe(`${projectName}-${serviceName}: DefaultService`, function () {
    let service = new DefaultService(new MockDynamoDB());
    beforeEach(done => {
        /* put any prerequisite here */
        done();
    })

    /* Trivial tests and stubs for each of the default functions in DefaultService.ts
    */
    describe('get ', () => {
        it('should return DB object', async () => {
            let dbQueryResult = await service.get({ "mock": 1 });
            expect(dbQueryResult).toEqual(mockDBobject);
        });
    });

    describe('getById ', () => {
        it('should return DB object', async () => {
            let dbQueryResult = await service.getById("mock");
            expect(dbQueryResult).toEqual(mockDBobject);
        });
    });

    describe('post', () => {
        it('should return DB object', async () => {
            let dbQueryResult = await service.post({ "mock": 1 });
            expect(dbQueryResult).toBeInstanceOf(Object);
        });
    });

    afterEach(done => {
        /* Put any cleanup here */
        done();
    })
});

