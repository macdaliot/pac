import {
    Collection,
    Cursor,
    Db,
    DeleteWriteOpResultObject,
    InsertOneWriteOpResult,
    InsertWriteOpResult,
    ObjectId,
    UpdateQuery
} from 'mongodb';
import { IWritable } from "../IWritable";
import { IReadable } from "../IReadable";

// TODO need to remove duplicate try/catch/finally without compromising strong typing.
// TODO need to modify return types to be valuable instead of boolean.
export abstract class MongoRepository<T> implements IWritable<T>, IReadable<T> {
    // expose collection as an immutable property so that concrete classes may add additional functionality
    public readonly _collection: Collection;

    constructor(db: Db, collectionName: string) {
        this._collection = db.collection(collectionName);
    }

    /**
     * Inserts one document type <T>
     * @param item - single object
     * @returns Id of inserted document if action completed successfully
     */
     async insertOne(item: T): Promise<string> {
        let result: string;
        try {
            let response: InsertOneWriteOpResult = await this._collection.insertOne(item);
            result = response.insertedId.toHexString();
        } catch (error) {
            console.error(`Error while inserting to DocumentDB. Item: ${JSON.stringify(item)} Stack: ${error}`)
        } finally {
            return result;
        }
    }

    /**
     * Inserts many documents of type <T>
     * @param items - array of objects
     * @returns Array of IDs if action completed successfully
     */
    async insertMany(items: T[]): Promise<string[]> {
        let result: string[] = [];
        try {
            let response: InsertWriteOpResult = await this._collection.insertMany(items);
            let ids: { [p: number]: ObjectId } = response.insertedIds;
            for(let key in ids){
                result.push(ids[key].toHexString())
            }
        } catch (error) {
            console.error(`Error while bulk inserting to DocumentDB. Items: ${items.forEach(item => JSON.stringify(item))} Stack: ${error}`)
        } finally {
            return result;
        }
    }

    /**
     * Updates one item of type <T>
     * @param id - document objectId as a string. this should correspond to the _id field
     * @param newItem - updated document
     * @link https://docs.mongodb.com/manual/reference/operator/update
     * @returns boolean if action complete successfully
     */
    async update(id: string, newItem: T): Promise<boolean> {
        let result: boolean = false;
        try {
            let response = await this._collection.replaceOne({"_id": new ObjectId(id)}, newItem);
            result = Boolean(response.result.ok);
        } catch (error) {
            console.error(error)
        } finally {
            return result;
        }
    }

    /**
     * Deletes one document
     * @param id this should correspond to the _id field
     * @returns boolean if action complete successfully
     */
    async deleteById(id: string): Promise<boolean> {
        let result: boolean = false;
        try {
            let response: DeleteWriteOpResultObject = await this._collection.deleteOne({"_id": new ObjectId(id)});
            result = Boolean(response.result.ok);
        } catch (error) {
            console.error(`Error while deleting from DocumentDB. Id: ${id} Stack: ${error}`)
        } finally {
            return result;
        }
    }

    /**
     * Deletes many documents
     * @param ids - array of ids. this should correspond to the _id field
     * @returns boolean if action complete successfully
     */
    async deleteManyById(ids: string[]): Promise<boolean> {
        const objectIds: ObjectId[] = new Array<ObjectId>();
        //populate array of objectIds
        ids.map(id => {
            objectIds.push(new ObjectId(id));
        });
        let result: boolean = false;
        try {
            let response: DeleteWriteOpResultObject = await this._collection.deleteMany(
                {
                    _id:
                        {$in: objectIds}
                });
            result = Boolean(response.result.ok);
        } catch (error) {
            console.error(`Error while bulk deleting from DocumentDB. Ids: ${ids.toString()} Stack: ${error}`)
        } finally {
            return result;
        }
    }

    /**
     * Retrieves a list of documents that match the query criteria
     * @param query - can be a RootQuerySelector, QuerySelector or combination of both
     * @link https://docs.mongodb.com/manual/reference/operator/query/
     * @param projection (optional) - array of fields
     * @link http://mongodb.github.io/node-mongodb-native/3.1/api/Cursor.html
     * @example
     *      //find all documents with the name John
     *      const query: any = {name:{$eq:"Jeff"}};
     *      myRepo.findByFilter(query);
     * @example
     *      //find all documents with the name John but only return lastname
     *      //find all documents with the name John
     *      const query: any = {name:{$eq:"John"}};
     *      const projList: string[] = ["lastName"];
     *      myRepo.findByFilter(query, projList);
     * @returns array of objects of type <T>
     */
     async findByFilter(query: any, projection?: string[]): Promise<T[]> {
        let result: Cursor;
        if (projection) {
            let projectedVal: any = {};
            projection.map(field => {
                projectedVal[field] = 1;
            });

            result = await this._collection.find(query).project(projectedVal);
        } else {
            result = await this._collection.find(query)
        }
        return result.toArray();
    }

    /**
     * @description Finds one document with the given id
     * @param id
     * @returns document as object of type <T> if it exists; else and empty object
     */
     async findById(id: string): Promise<T> {
        let result;
        try {
            let response: Cursor = await this._collection.find({"_id": new ObjectId(id)});
            result = await response.hasNext() ? await response.next() : {};
        } catch (error) {
            console.error(error)
        } finally {
            return result;
        }
    }

    /**
     * @description The following aggregation operation selects documents with name
     * equal to "John", groups the matching documents by the name field and
     * calculates the total for each name field from the sum of the amount field,
     * and sorts the results by the total field in descending order
     * @example
     * let response: AggregationCursor<T> = await this._collection.aggregate().match({"name":"John"}).group({_id:"$name", total: {$sum:"$amount"}}).sort({total:-1});
     * return response.toArray();
     */
    //public abstract aggregate(): void //TODO to be removed, just used as a placeholder but aggregations should be implemented by concrete classes
}




