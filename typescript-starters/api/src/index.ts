import { createApp } from "./app.js";
import type { Application } from "express";
import * as http from "node:http";

async function main() {
  const port = process.env.PORT || 8082;
  let server: http.Server;

  try {
    const app: Application = await createApp();

    server = app.listen(port, () => {
      console.log(`Server is running on http://localhost:${port}`);
    });

    // Handle graceful shutdown on SIGTERM and SIGINT signals
    process.on("SIGTERM", () => shutdown(server, "SIGTERM"));
    process.on("SIGINT", () => shutdown(server, "SIGINT"));
  } catch (error) {
    console.error("Failed to start application:", error);
    process.exit(1);
  }
}

function shutdown(server: http.Server, signal: string) {
  console.debug(`${signal} signal received: closing HTTP server`);

  server.close((err?: Error) => {
    if (err) {
      console.error("Error while closing server:", err);
      process.exit(1);
    }
    console.log("HTTP server closed");
    process.exit(0); // Graceful shutdown
  });
}

await main();
