import items from "../router/items.js";
import cors from "cors";
import express, { Application } from "express";
import { heartbeat } from "../middleware/index.js";

/**
 * Configures an express application with necessary middleware, mounts a router, and
 * adds global error handling.
 *
 */
export default async (app: Application): Promise<void> => {
  app.use(cors()); // add cors middleware
  app.use(express.json()); // JSON Body Parser

  app.use("/health", heartbeat);

  app.use("/items", items);
};
