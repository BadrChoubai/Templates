import { Router } from "express";
import database from "../database/sqlite.js";

import { getItems } from "./handlers/getItems.js";
import { getItem } from "./handlers/getItem.js";

import { ItemsService } from "../service/items.js";

const items = Router();
const itemsService = new ItemsService(database);

// Entire API surface mapped here
items.get("/", getItems(itemsService));
items.get("/:id", getItem(itemsService));

export default items;
