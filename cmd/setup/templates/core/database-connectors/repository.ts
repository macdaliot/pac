import { Newable } from './newable';

export interface Repository<TModel> {
  get(id: string, model: Newable<TModel>): Promise<TModel>

  // GET /api/fire?column=this&othercol=that
  // query || search (returns list of data rows that match on an arbitrary column)

  scan(value: Newable<TModel>): Promise<Array<TModel>>

  put(id: string, value: TModel, model: Newable<TModel>): Promise<TModel>

  post(value: TModel, model: Newable<TModel>): Promise<TModel>

  delete(id: string, model: Newable<TModel>): Promise<TModel>
}