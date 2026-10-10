import { DatabaseSync } from "node:sqlite";

export class ItemsService {
  private _db: DatabaseSync;

  constructor(db: DatabaseSync) {
    this._db = db;
  }

  async getAll(): Promise<unknown[]> {
    const query = this._db.prepare(`
      SELECT id, name, description
      FROM items
    `);

    return query.all();
  }

  async getOne(id: string): Promise<unknown> {
    const query = this._db.prepare(`
      SELECT id, name, description
      FROM items
      WHERE id = ?
    `);

    return query.get(id);
  }
}
