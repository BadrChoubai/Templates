import express from "express";

export async function createApp(): Promise<express.Application> {
  const app = express();

  const { default: loadApp } = await import("./loaders/index.js");
  if (loadApp) await loadApp({ app });

  return app;
}
