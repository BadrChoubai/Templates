import type { Request, Response, NextFunction } from "express";

const OK: string = "OK";

export const heartbeat = (req: Request, res: Response, next: NextFunction) => {
  if (req.method === "GET") {
    res.setHeader("Content-Type", "text/plain");
    res.status(200).send(OK);
  } else if (req.method === "HEAD") {
    res.setHeader("Content-Type", "text/plain");
    res.status(200);
  }

  next();
};
