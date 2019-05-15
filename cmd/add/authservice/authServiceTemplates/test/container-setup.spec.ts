import { setupContainer } from '../src/container-setup';
import * as express from 'express';
const mockPassport = {
    use: jest.fn(),
    authenticate: jest.fn(() => (req: Request, res: Response, next: express.NextFunction) => {}),
    initialize: jest.fn(() => (req: Request, res: Response, next: express.NextFunction) => {})
};
jest.doMock('passport', () => (mockPassport));
var app = express();

describe('Container Setup', () => {
    it('Should call the container setup on a test app', () => {
        const iocContainer = setupContainer(app);
        expect(iocContainer).not.toBeNull();
        expect(iocContainer.isBound);
        //expect(mockPassport.use).toBeCalled()
        //expect(mockPassport.initialize).toBeCalled()
    });
});
