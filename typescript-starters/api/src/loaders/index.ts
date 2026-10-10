import express from "./express.js";
import type { Application } from "express";
import Logger from "../logger/winston.js";

/**
 * Configure and initialize any external service (infrastructure, side effects, etc.)
 * used by our Express application asynchronously
 *
 */
export default async ({ app: expressApp }: { app: Application }) => {
  Logger.info("Created App");
  await express(expressApp);
};
