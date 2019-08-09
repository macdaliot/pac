import { Route, Get, Post, Controller, Path, Put, Body, Delete, Security, SuccessResponse } from 'tsoa';
import { Connectors, Inject, Injectable, Repository } from '@pyramid-systems/core';
import { {{.serviceNamePascal}} } from '@pyramid-systems/domain';

@Injectable()
@Route('{{.serviceName}}')
// Uncomment the line below and add your own scope names from Auth0 (do not change `'groups'`) in order to require a valid JWT to use this service
// @Security('groups', ['<one-or-more-scope-names>'])
export class {{.serviceNamePascal}}Controller extends Controller {
  private repository: Repository<{{.serviceNamePascal}}>;

  constructor(@Inject(Connectors.DynamoDB) private _repository: Repository<{{.serviceNamePascal}}>) {
    super();
    this.repository = _repository;
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
    return this.repository.get(id);
  }

  /**
   * @summary Gets all the {{.serviceNamePascal}}
   */
  @Get()
  @SuccessResponse(200, "A list of all existing {{.serviceNamePascal}}s")
  getAll() {
    return this.repository.scan(Fire);
  }

  /**
   * Post a new {{.serviceNamePascal}} to the backing store
   *
   * @summary Supply a {{.serviceNamePascal}} object and have it be stored
   * 
   * @param newItem the new {{.serviceNamePascal}} object to persist
   */
  @Post()
  @SuccessResponse(200, "The inserted {{.serviceNamePascal}}")
  add(@Body() newItem: {{.serviceNamePascal}}) {
    return this.repository.post(newItem, Fire);
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
  update(@Path('id') idToUpdate: string, @Body() itemWithUpdatedValues: {{.serviceNamePascal}}) {
    return this.repository.put(idToUpdate, itemWithUpdatedValues, Fire);
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
    return this.repository.delete(idToDelete, Fire);
  }
}
