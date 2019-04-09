import * as express from 'express';
import { Route, Get, Post, Controller, Request, Query, Path, Put } from 'tsoa';
import { Injectable } from '@pyramid-systems/core';

@Injectable()
@Route('{{.serviceName}}')
export class {{.serviceNamePascal}}Controller extends Controller {
    @Get('{id}')
    getById(@Path('id') id: string) {
        throw new Error('Not yet implemented');
    }

    @Get()
    getAll() {
        throw new Error('Not yet implemented');
    }

    @Post()
    post(@Request() request: express.Request) {
        throw new Error('Not yet implemented');
    }

    @Put()
    put() {
        throw new Error('Not yet implemented');
    }
}
