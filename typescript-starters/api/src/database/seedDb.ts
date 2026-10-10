import database from "./sqlite.js";
import { Item } from "../models/Item.js";

database.exec(`
  DROP TABLE IF EXISTS items
`);

database.exec(`
  CREATE TABLE IF NOT EXISTS items
  (
    id          TEXT PRIMARY KEY,
    name        TEXT NOT NULL,
    description TEXT NOT NULL DEFAULT ''
  )
`);

const items: Item[] = [
  {
    id: "c04ab9a0-e398-471d-8cfe-95317ccf0a20",
    name: "First item",
    description: "An example item",
  },
  {
    id: "1e95757d-066d-41a6-b5d2-5e34b71d6051",
    name: "Second item",
    description: "Another example item",
  },
];

const insertItem = database.prepare(`
  INSERT INTO items (id, name, description)
  VALUES (?, ?, ?)
`);

items.forEach((item) => {
  insertItem.run(item.id, item.name, item.description);
});
