import * as express from 'express';
import { Route, Get, Post, Controller, Request, Query, Path, Put } from 'tsoa';
import { Injectable } from '@pyramid-systems/core';

@Injectable()
@Route('{{.serviceName}}')
export class {{.serviceNamePascal}}Controller extends Controller {

    constructor(private {{.serviceName}}Repository: {{.serviceNamePascal}}Repository) {
        super();
    }

    @Get('{id}')
    getById(@Path('id') id: string) {
        return this.{{.serviceName}}Repository.getById(id);
    }

    @Get()
    getAll() {
        return this.{{.serviceName}}Repository.getAll();
    }

    @Post()
    post(@Body() newItem: {{.serviceNamePascal}}) {
        return this.{{.serviceName}}Repository.add(newItem);
    }

    @Put('{id}')
    put(@Path('id') idToUpdate: string, @Body() itemWithUpdatedValues: {{.serviceNamePascal}}) {
        return this.{{.serviceName}}Repository.update(idToUpdate, itemWithUpdatedValues);
    }

    @Delete('{id}')
    deleteById(@Path('id') idToDelete: string) {
        return this.{{.serviceName}}Repository.deleteById(idToDelete);
    }
}
