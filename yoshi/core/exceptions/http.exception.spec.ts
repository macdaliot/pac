import { HttpException } from './http.exception';

describe('HTTP Exception', () => {
    it('should create the constructor and set the status and number', () => {
        const testMessage = "Hello";
        const testStatus = 500;
        const instance = new HttpException(testStatus, testMessage);

        expect(instance.message).toEqual(testMessage);
        expect(instance.status).toEqual(testStatus);
    })    
})