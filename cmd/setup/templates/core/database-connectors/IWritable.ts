/**
 * This interface provides a consistent contract for documentdb to elasticsearch synchronization.
 */
export interface IWritable<T> {
    insertOne(item: T): Promise<string>;
    insertMany(items: T[]): Promise<string[]>;
    update(id: string, newItem: T): Promise<boolean>;
    deleteById(id: string): Promise<boolean>;
    deleteManyById(ids: string[]): Promise<boolean>;
}
