import { setupContainer } from './container-setup';
import express = require('express');

var app = express();
describe('Container Setup', () => {
    it('Should call the container setup on a test app and expect Container to be bound', () => {
        const iocContainer = setupContainer(app);
        expect(iocContainer.isBound);
    });
});