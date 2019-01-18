package service

import (
	"github.com/PyramidSystemsInc/go/files"
)

// CreateTestFile creates a default test file based on passed in config
func CreateDefaultServiceSpecTs(filePath string, config map[string]string) {
	const template = `import 'mocha';
    import { expect } from 'chai';
    import { DefaultService } from '../src/services/defaultService';
    import { MockDynamoDB, mockDBResponse, mockDBobject } from './mockDynamoDb';
    
    describe('{{.serviceName}}: DefaultService', function () {
        let service = new DefaultService(new MockDynamoDB());
        before(done => {
            /* put any prerequisite here */
            done();
        })
    
        describe('get ', () => {
            it('should return DB object', async () => {
                let dbQueryResult = await service.get({ "mock": 1 });
                expect(dbQueryResult).to.equal(mockDBobject);
            });
        });
    
        describe('getById ', () => {
            it('should return DB object', async () => {
                let dbQueryResult = await service.getById("mock");
                expect(dbQueryResult).to.equal(mockDBobject);
            });
        });
    
        describe('post', () => {
            it('should return DB object', async () => {
                let dbQueryResult = await service.post({ "mock": 1 });
                expect(dbQueryResult).to.be.an('object');
            });
        });
    
        after(done => {
            /* Put any cleanup here */
            done();
        })
    });          
`
	files.CreateFromTemplate(filePath, template, config)
}
