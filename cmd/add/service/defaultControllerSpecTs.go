package service

import (
	"github.com/PyramidSystemsInc/go/files"
)

// CreateTestFile creates a default test file based on passed in config
func CreateDefaultControllerSpecTs(filePath string, config map[string]string) {
	const template = `import 'mocha';
    import { expect } from 'chai';
    import { DefaultController } from '../src/controllers/defaultController';
    import { MockDefaultService } from './mockDefaultService';
    
    /* need to mock req */
    describe('{{.serviceName}}: DefaultController', function () {
        let controller = new DefaultController(new MockDefaultService());
        before(done => {
            /* put any prerequisite here */
            done();
        })
    
        describe('.get ', () => {
            it('should send response', async () => {
            });
        });
    
        describe('.getById ', () => {
            it('should send response', async () => {
            });
        });
    
        describe('.post', () => {
            it('should send response', async () => {
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
