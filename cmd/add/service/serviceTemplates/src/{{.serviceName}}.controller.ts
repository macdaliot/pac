import { Route, Get, Post, Controller, Request, Query, Path, Put, Body, Delete, Security, SuccessResponse  } from 'tsoa';
import { Injectable } from '@pyramid-systems/core';
import { {{.serviceNamePascal}}, {{.serviceNamePascal}}Repository } from '@pyramid-systems/domain';
import { MongoClient } from "mongodb";
import { IndexingService } from "../../../core/services/indexingService";

@Injectable()
@Route('{{.serviceName}}')
// Uncomment the line below and add your own scope names from Auth0 (do not change `'groups'`) in order to require a valid JWT to use this service
// @Security('groups', ['<one-or-more-scope-names>'])
export class {{.serviceNamePascal}}Controller extends Controller {
    domainModel: string = `${process.env.PROJECTNAME}-{{.serviceName}}`;
    indexingService: IndexingService<{{.serviceNamePascal}}>;
    private {{.serviceName}}Repository: {{.serviceNamePascal}}Repository;

    constructor() {
        super();
    }

    /**
     * @summary Deletes DB collections from Elastic and DocDb
     * PLEASE REMOVE ONCE data model is mature and confident in persisted data
     */
    @Get('wipedb')
    @SuccessResponse(200, "DEV feature - wipe db collections")
    async wipeDbs() {
        await this.initConnection();
        await this.indexingService.wipeDbs(this.domainModel);
        return "Wiped DocumentDB and Elastic Search Collections";
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
    async getById(@Path('id') id: string) {
        await this.initConnection();
        return this.{{.serviceName}}Repository.findById(id);
    }
    /**
     * @summary Gets all the {{.serviceNamePascal}}
     */
    @Get()
    @SuccessResponse(200, "A list of all existing {{.serviceNamePascal}}s")
    async getAll() {
        await this.initConnection();
        return this.{{.serviceName}}Repository.findByFilter({});
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
    async post(@Body() newItems: {{.serviceNamePascal}}[]) {
        await this.initConnection();
        return this.indexingService.insertMany(newItems);
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
    async put(@Path('id') idToUpdate: string, @Body() itemWithUpdatedValues: {{.serviceNamePascal}}) {
        await this.initConnection();
        return this.indexingService.update(idToUpdate, itemWithUpdatedValues);
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
    async deleteById(@Path('id') idToDelete: string) {
        await this.initConnection();
        return this.indexingService.deleteById(idToDelete);
    }

    private async initConnection() {
        if (!this.{{.serviceName}}Repository) {
            //local mongo connString `mongodb://<user>:<password>@<host>:27017/<db>`
            const connection = await MongoClient.connect(process.env.MONGO_CONN_STRING, {
                useNewUrlParser: true,
                useUnifiedTopology: true
            });
            const db = connection.db(this.domainModel);
            this.{{.serviceName}}Repository = new {{.serviceNamePascal}}Repository(db, this.domainModel);
            this.indexingService = new IndexingService(db, this.domainModel)
        }
    }
}
