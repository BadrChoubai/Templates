import { mkdirSync } from "node:fs";
import { dirname } from "node:path";
import { DatabaseSync } from "node:sqlite";

const databasePath = process.env.DATABASE_PATH || ".db/api.db";

mkdirSync(dirname(databasePath), { recursive: true });

const database: DatabaseSync = new DatabaseSync(databasePath);

export default database;
