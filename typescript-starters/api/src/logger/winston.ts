import winston from "winston";

const level = process.env.LOG_LEVEL || "info";

const loggerInstance = winston.createLogger({
  level,
  levels: winston.config.npm.levels,
  format: winston.format.combine(
    winston.format.timestamp({
      format: "YYYY-MM-DD HH:mm:ss",
    }),
    winston.format.errors({ stack: true }),
    winston.format.splat(),
    winston.format.json(),
  ),
  transports: [new winston.transports.Console()],
});

export default loggerInstance;
