import { DefaultController } from '../src/controllers/defaultController';
import { serviceName, projectName } from '../src/config/appInfo.config';
import { MockDynamoDB, BrokenDynamoDB } from './mockDynamoDb';
import express = require('express');

/* TODO: setup way to mock express requests */
describe(`${projectName}-${serviceName} DefaultController`, () => {
    let result: any = null;
    let capturer = (output: any) => {result = output};
    let controller: DefaultController;
    beforeEach(done => {
        /* put any prerequisite here */
        controller = new DefaultController(new MockDynamoDB());
        done();
    });

    /* Stubs for each of the default functions in DefaultController.ts
    */
    describe('.get', () => {
        it('should send response', async () => {
            await controller.get({ body: {}} as any as express.Request, { send: capturer } as any as express.Response, capturer);
            expect(result).not.toBeNull();
            expect(result).not.toBeInstanceOf(Error);
        });
        it('should fail response', async () => {
            controller = new DefaultController(new BrokenDynamoDB());
            await controller.get({ body: {}} as any as express.Request, { send: capturer } as any as express.Response, capturer);
            expect(result).toBeInstanceOf(Error);
        });
    });

    describe('.getById', () => {
        it('should send response', async () => {
            await controller.getById({ body: {}, params: {id: 1}} as any as express.Request, { send: capturer } as any as express.Response, capturer);
            expect(result).not.toBeNull();
            expect(result).not.toBeInstanceOf(Error);
        });
        it('should fail response', async () => {
            controller = new DefaultController(new BrokenDynamoDB());
            await controller.getById({ body: {}} as any as express.Request, { send: capturer } as any as express.Response, capturer);
            expect(result).toBeInstanceOf(Error);
        });
    });

    describe('.post', () => {
        it('should send response', async () => {
            await controller.post({ body: {} } as any as express.Request, { send: capturer } as any as express.Response, capturer);
            expect(result).not.toBeNull();
            expect(result).not.toBeInstanceOf(Error);
        });
        it('should fail response', async () => {
            controller = new DefaultController(new BrokenDynamoDB());
            await controller.post({ body: {} } as any as express.Request, { send: capturer } as any as express.Response, capturer);
            expect(result).toBeInstanceOf(Error);
        });
    });

    describe('.delete', () => {
        it('should send response', async () => {
            await controller.delete({ body: {} } as any as express.Request, { send: capturer } as any as express.Response, capturer);
            expect(result).not.toBeNull();
            expect(result).not.toBeInstanceOf(Error);
        });
        it('should send response with param', async () => {
            await controller.delete({ body: {}, params: {id: 1} } as any as express.Request, { send: capturer } as any as express.Response, capturer);
            expect(result).not.toBeNull();
            expect(result).not.toBeInstanceOf(Error);
        });
        it('should fail response', async () => {
            controller = new DefaultController(new BrokenDynamoDB());
            await controller.delete({ body: {} } as any as express.Request, { send: capturer } as any as express.Response, capturer);
            expect(result).toBeInstanceOf(Error);
        });
    });
    describe('.update', () => {
        it('should send response', async () => {
            await controller.update({ body: {} } as any as express.Request, { send: capturer } as any as express.Response, capturer);
            expect(result).not.toBeNull();
            expect(result).not.toBeInstanceOf(Error);
        });
        it('should send response with param', async () => {
            await controller.update({ body: {}, params: {id: 1} } as any as express.Request, { send: capturer } as any as express.Response, capturer);
            expect(result).not.toBeNull();
            expect(result).not.toBeInstanceOf(Error);
        });
        it('should fail response', async () => {
            controller = new DefaultController(new BrokenDynamoDB());
            await controller.update({ body: {} } as any as express.Request, { send: capturer } as any as express.Response, capturer);
            expect(result).toBeInstanceOf(Error);
        });
    });
    afterEach(done => {
        /* Put any cleanup here */
        result = null;
        done();
    })
});
