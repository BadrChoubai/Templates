import type { Request, Response, NextFunction } from "express";

const OK: string = "OK";

/**
 * Responds 200 OK to GET and HEAD requests and passes any other method on.
 */
export const heartbeat = (req: Request, res: Response, next: NextFunction) => {
  if (req.method !== "GET" && req.method !== "HEAD") {
    next();
    return;
  }

  res.status(200).type("text/plain").send(OK);
};
