import { ElasticIndex } from "../database-connectors/elasticSearch/elasticIndex";
import { MongoRepository } from "../database-connectors/amazon-documentdb/repository";
import { IWritable } from "../database-connectors/IWritable";
import { ObjectId } from "bson";
import { Client } from "@elastic/elasticsearch";
import { Db } from "mongodb";

export class IndexingService<T> implements IWritable<T> {
    elasticIndex: ElasticIndex<T>;
    mongoRepository: MongoRepository<T>;

    constructor(mongodb: Db, indexName: string) {
        const client:Client = new Client({ node:process.env.ES_ENDPOINT , ssl:{rejectUnauthorized:false} });
        this.mongoRepository = new class extends MongoRepository<T>{}(mongodb,indexName);
        this.elasticIndex = new class extends ElasticIndex<T> {}(client, indexName);
    }

    /**
     * Deletes an item of type <T> based on ID from both a Primary Db as well as an index db.
     * @param id of item to be deleted
     */
    async deleteById(id: string): Promise<boolean> {
        let self = this;
        const item: T = await this.mongoRepository.findById(id);
        return IndexingService.handleTransaction(
            function () {
                return self.mongoRepository.deleteById(id)
            },
            function () {
                return self.elasticIndex.deleteById(id)
            },
            function () {
                return self.mongoRepository.insertOne(item)
            });
    }

    /**
     * Deletes an array of items of type <T> based on ID from both a Primary Db as well as an index db.
     * @param ids to be deleted
     */
    async deleteManyById(ids: string[]): Promise<boolean> {
        let self = this;
        let documentIds: ObjectId[] = ids.map(function (id) {
            return new ObjectId(id);
        });
        const items: T[] = await this.mongoRepository.findByFilter({_id: {$in: documentIds}});
        return IndexingService.handleTransaction(
            function () {
                return self.mongoRepository.deleteManyById(ids)
            },
            function () {
                return self.elasticIndex.deleteManyById(ids)
            },
            function () {
                return self.mongoRepository.insertMany(items)
            });
    }

    /**
     * Inserts an array of items of type <T> into both a Primary Db as well as an index db.
     * @param items to be inserted/indexed
     */
    async insertMany(items: T[]): Promise<string[]> {
        let self = this;
        return IndexingService.handleTransaction(
            function () {
                return self.mongoRepository.insertMany(items)
            },
            function () {
                return self.elasticIndex.insertMany(items)
            },
            function (ids) {
                return self.mongoRepository.deleteManyById(ids)
            });
    }

    /**
     * Inserts an item of type <T> into both a Primary Db as well as an index db.
     * @param item to be inserted/indexed
     */
    async insertOne(item: T): Promise<string> {
        let self = this;
        return IndexingService.handleTransaction(
            function () {
                return self.mongoRepository.insertOne(item)
            },
            function () {
                return self.elasticIndex.insertOne(item)
            },
            function (id) {
                return self.mongoRepository.deleteById(id)
            });
    }

    async update(id: string, item: any): Promise<boolean> {
        let self = this;
        const originalItem: T = await this.mongoRepository.findById(id);
        return IndexingService.handleTransaction(
            function () {
                return self.mongoRepository.update(id, item)
            },
            function () {
                return self.elasticIndex.update(id, item)
            },
            function (id) {
                return self.mongoRepository.update(id, originalItem)
            });
    }

    /**
     * This method is in charge of maintaining a consistent state between the primary db and the index db.
     * It does this by executing the same insert/delete/update transaction to both primary/index dbs. Should writting
     * to the index db fail, a rollback function is executed against the primary db.
     * @param dbfunc - function which interacts with the primary db, insert/update/delete
     * @param indexFunc - function which interacts with the indexing db. this should mimic the behavior of the dbfunc
     * @param rollbackFunc - function which rollsback transaction to primary db should indexing fail.
     * @return returns whatever data type the primary dbfunc returns (inserts return id/s, deletes return boolean)
     */
    private static async handleTransaction(dbfunc: Function,
                                           indexFunc: Function,
                                           rollbackFunc: Function) {
        let mongoResponse = await dbfunc();
        if (mongoResponse) {
            let elasticResponse = await indexFunc();
            if (elasticResponse) {
                return mongoResponse;
            } else {
                let rollbackResponse = rollbackFunc(mongoResponse);
                if (!rollbackResponse) {
                    const errorMsg: string = `Unable to rollback. Verify logs.`;
                    this.logAndThrowError(errorMsg);
                } else {
                    //TODO better error handling, this error message was pretty useless while testing
                    const errorMsg: string = `Unable to index items. Verify logs.`;
                    this.logAndThrowError(errorMsg);
                }
            }
        } else {
            const errorMsg: string = `Unable to complete action to DocumentDb. Verify logs.`;
            this.logAndThrowError(errorMsg);
        }
    }

    private static logAndThrowError(errorMsg: string) {
        console.error(errorMsg);
        throw new Error(errorMsg)
    }

}
