
import * as pino from 'pino';

interface MockPino {
    level?: string;
    child(options: any): any;
}
process.env.LOG_LEVEL = "fatal";

const pinoCastToAny = pino as any;
const mockPino = <jest.Mock<MockPino>>pinoCastToAny;
let childMock = jest.fn(() => ({
    warn: jest.fn().mockReturnThis(),
    trace: jest.fn(),
    error: jest.fn(),
    info: jest.fn(),
    debug: jest.fn(),
    fatal: jest.fn()
}));
let pinoMock = {
    child: childMock
}
mockPino.mockImplementation(() => pinoMock);
// Order of statement is important.
import { ServiceLogger, normalizeLogMessage } from '@pyramid-systems/core';

describe('Logger', () => {

    it('should call the pino logger child', () => {
        const instance = new ServiceLogger('someServiceName');

        expect(pinoMock.child).toBeCalledWith({ serviceName: 'someServiceName' });


        expect(instance.getPinoLogger()).toBeTruthy()
    })


    it('given one parameter only should call the pass the parameters to the logger child with empty string as second parameter', () => {
        const instance = new ServiceLogger('someServiceName');

        const message = 'Some message';
        instance.warn(message);
        instance.error(message);
        instance.info(message);
        instance.debug(message);
        instance.trace(message);
        instance.fatal(message);
        expect(instance.getPinoLogger().warn).toBeCalledWith(message, '', []);
        expect(instance.getPinoLogger().error).toBeCalledWith(message, '', []);
        expect(instance.getPinoLogger().info).toBeCalledWith(message, '', []);
        expect(instance.getPinoLogger().debug).toBeCalledWith(message, '', []);
        expect(instance.getPinoLogger().trace).toBeCalledWith(message, '', []);
        expect(instance.getPinoLogger().fatal).toBeCalledWith(message, '', []);
    })


    it('given one parameter that is object should call the pass the parameters to the logger child with undefined string as second parameter', () => {
        const instance = new ServiceLogger('someServiceName');

        const message = { message: 'Some message' };
        instance.warn(message);
        instance.error(message);
        instance.info(message);
        instance.debug(message);
        instance.trace(message);
        instance.fatal(message);
        expect(instance.getPinoLogger().warn).toBeCalledWith(message, message, []);
        expect(instance.getPinoLogger().error).toBeCalledWith(message, message, []);
        expect(instance.getPinoLogger().info).toBeCalledWith(message, message, []);
        expect(instance.getPinoLogger().debug).toBeCalledWith(message, message, []);
        expect(instance.getPinoLogger().trace).toBeCalledWith(message, message, []);
        expect(instance.getPinoLogger().fatal).toBeCalledWith(message, message, []);
    })
    it('normalizeLogMessage should return the empty string given type is string', () => {
        const value = normalizeLogMessage('message');
        expect(value).toBe('');
    })
    it('normalizeLogMessage should return the itself string given type is not string', () => {
        const message = { x: 'message' };
        const value = normalizeLogMessage(message);
        expect(value).toEqual(message);
    })
})