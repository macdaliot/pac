export interface IReadable<T> {
    findByFilter(query: any, projection?: string[]): Promise<T[]>;
    findById(id: string): Promise<T>;
}
