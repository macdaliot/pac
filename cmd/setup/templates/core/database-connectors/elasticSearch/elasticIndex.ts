import { IWritable } from "../IWritable";
import { Client } from "@elastic/elasticsearch";

export class ElasticIndex<T> implements IWritable<T> {

    constructor(public client: Client, private indexName: string){
        this.client = client;
        this.indexName = indexName
    }

    /**
     * Deletes one document
     * @param id this should correspond to the _id field
     * @returns boolean if action complete successfully
     */
    async deleteById(id: string): Promise<boolean> {
        let result: boolean = false;
        try {
            await this.client.delete({id, index:this.indexName});
            result =  true;
        } catch (error) {
            console.error(`Error while deleting from elasticsearch. ID: ${id} Stack: ${error}`)
        } finally {
            return result;
        }
    }

    /**
     * Deletes many documents
     * @param ids - array of ObjectId's. this should correspond to the _id field
     * @returns boolean if action complete successfully
     */
    async deleteManyById(ids: string[]): Promise<boolean> {
        let result: boolean = false;
        try {
            let bulkList = [];
            //create action-data pairs
            ids.map(id => {
                // @ts-ignore
                bulkList.push({ delete: {_index: this.indexName, _id: id}});
            });
            await this.client.bulk({body: bulkList});
            result =  true;
        } catch (error) {
            console.error(`Error while bulk deleting from elasticsearch. ID: ${ids.toString()} Stack: ${error}`)
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
        let result: string[];
        let ids: string[] = [];
        try {
            let bulkList = [];
            //create action-data pairs
            items.map(item => {
                // @ts-ignore
                let _id: string = item._id;
                ids.push(_id);
                bulkList.push({ create: {_index: this.indexName, _id}});
                // @ts-ignore
                delete item._id;
                bulkList.push(item);
            });

            await this.client.bulk({body: bulkList});
            result = ids;
        } catch (error) {
            console.error(`Error while bulk inserting to elasticsearch. ID: ${ids.toString()} Stack: ${error}`)
        } finally {
            return result;
        }
    }

    /**
     * Inserts one document type <T>
     * @param item - single object
     * @returns Id of inserted document if action completed successfully
     */
    async insertOne(item: T): Promise<string> {
        let result: string;
        // @ts-ignore
        const id: string = item._id;
        try {
            // @ts-ignore
            delete item._id;
            await this.client.create({id, index:this.indexName, body:item});
            result = id;
        } catch (error) {
            console.error(`Error while inserting to elasticsearch. ID: ${id} Stack: ${error}`)
        } finally {
            return result;
        }
    }

    /**
     * Updates a document of type <T>
     * @param item - single document
     * @returns Id of updated document if action completed successfully
     */
    async update(id: string, item: T): Promise<boolean> {
        let result: boolean;
        try {
            // @ts-ignore
            delete item._id;
            await this.client.update({id, index:this.indexName, body: { doc:item }});
            result = true;
        } catch (error) {
            console.error(`Error while updating elasticsearch. ID: ${id} Stack: ${error}`)
        } finally {
            return result;
        }
    }

}
