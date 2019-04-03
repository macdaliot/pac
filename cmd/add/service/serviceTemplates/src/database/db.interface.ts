export interface Database {
    query(params: any);
    create(object: any);
    update(params: any, object: any);
    delete(params: any);
}
