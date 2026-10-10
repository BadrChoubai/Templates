import { DatabaseSync } from "node:sqlite";

const database: DatabaseSync = new DatabaseSync(".db/api.db");

export default database;
