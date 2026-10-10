import { DatabaseSync } from "node:sqlite";
import { ItemDbResponse } from "../models/Item.js";

export class ItemsService {
  private _db: DatabaseSync;

  constructor(db: DatabaseSync) {
    this._db = db;
  }

  async getAll(): Promise<ItemDbResponse[]> {
    const query = this._db.prepare(`
      SELECT id, name, description
      FROM items
    `);

    return query.all() as ItemDbResponse[];
  }

  async getOne(id: string): Promise<ItemDbResponse | undefined> {
    const query = this._db.prepare(`
      SELECT id, name, description
      FROM items
      WHERE id = ?
    `);

    return query.get(id) as ItemDbResponse | undefined;
  }
}
