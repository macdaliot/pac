import { DefaultController } from '../src/controllers/defaultController';
import { MockDefaultService } from './mockDefaultService';
import { serviceName, projectName } from '../src/config/appInfo.config';

/* TODO: setup way to mock express requests */
describe(`${projectName}-${serviceName} DefaultController`, () => {
    let controller = new DefaultController(new MockDefaultService());
    beforeEach(done => {
        /* put any prerequisite here */
        done();
    })

    /* Stubs for each of the default functions in DefaultController.ts
    */
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

    afterEach(done => {
        /* Put any cleanup here */
        done();
    })
});

