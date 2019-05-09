import * as express from 'express';
const mockPassport = {
    use: jest.fn(),
    authenticate: jest.fn(() => (req: Request, res: Response, next: express.NextFunction) => {}),
    initialize: jest.fn(() => (req: Request, res: Response, next: express.NextFunction) => {})
};
jest.doMock('passport', () => (mockPassport));
import app from '../src/server'

jest.mock('../docs/swagger.json', () => ({}));
jest.mock('../src/generated/routes', () => ({RegisterRoutes: jest.fn()}));

describe('the auth application', () => {
    it('should call passport appropriately', () => {
        expect(app).not.toBeNull();
        expect(mockPassport.use).toBeCalled()
        expect(mockPassport.initialize).toBeCalled()
    })
})