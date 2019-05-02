import * as express from 'express';
import { Route, Get, Post, Controller, Request, Query, Path, Put, Body, Delete, Security, SuccessResponse  } from 'tsoa';
import { Injectable } from '@pyramid-systems/core';
import { {{.serviceNamePascal}}, {{.serviceNamePascal}}Repository } from '@pyramid-systems/domain';

@Injectable()
@Route('{{.serviceName}}')
export class {{.serviceNamePascal}}Controller extends Controller {

    constructor(private {{.serviceName}}Repository: {{.serviceNamePascal}}Repository) {
        super();
    }

    /**
     * Get a {{.serviceNamePascal}} by Id
     *
     * @summary Get a {{.serviceNamePascal}}, by Id
     * 
     * @param id id of the {{.serviceNamePascal}} to search
     */
    @Get('{id}')
    @SuccessResponse(200, "The {{.serviceNamePascal}} with the given Id")
    getById(@Path('id') id: string) {
        return this.{{.serviceName}}Repository.getById(id);
    }
    /**
     * @summary Gets all the {{.serviceNamePascal}}
     */
    @Get()
    @SuccessResponse(200, "A list of all existing {{.serviceNamePascal}}s")
    getAll() {
        return this.{{.serviceName}}Repository.getAll();
    }
    /**
     * Post a new {{.serviceNamePascal}} to the backing store
     *
     * @summary Supply a {{.serviceNamePascal}} object and have it be stored
     * 
     * @param newItem the new {{.serviceNamePascal}} object to persist
     */
    @Post()
    @SuccessResponse(200, "The inserted {{.serviceNamePascal}} (with the new Id?)")
    post(@Body() newItem: {{.serviceNamePascal}}) {
        return this.{{.serviceName}}Repository.add(newItem);
    }
    /**
     * Update an existing {{.serviceNamePascal}} in the backing store
     *
     * @summary Supply a {{.serviceNamePascal}} object and have it be updated
     * 
     * @param idToUpdate the id of the {{.serviceNamePascal}} object to update
     * @param itemWithUpdatedValues the actual {{.serviceNamePascal}} object that will be used to update the stored object
     */
    @Put('{id}')
    @SuccessResponse(200, "The updated {{.serviceNamePascal}} with the given Id")
    put(@Path('id') idToUpdate: string, @Body() itemWithUpdatedValues: {{.serviceNamePascal}}) {
        return this.{{.serviceName}}Repository.update(idToUpdate, itemWithUpdatedValues);
    }
    /**
     * Delete an existing {{.serviceNamePascal}} in the backing store
     *
     * @summary Remove a {{.serviceNamePascal}}
     * 
     * @param id the id of the {{.serviceNamePascal}} object to delete
     * @param itemWithUpdatedValues the actual {{.serviceNamePascal}} object that will be used to update the stored object
     */
    @Delete('{id}')
    @SuccessResponse(200, "The {{.serviceNamePascal}} with the given Id")
    deleteById(@Path('id') idToDelete: string) {
        return this.{{.serviceName}}Repository.deleteById(idToDelete);
    }
}
