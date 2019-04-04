export abstract class Database {
    abstract query(params: any);
    abstract create(object: any);
    abstract update(params: any, object: any);
    abstract delete(params: any);
}
