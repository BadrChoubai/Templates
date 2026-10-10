import winston from "winston";

const transports = [
  new winston.transports.Console({
    level: "error",
  }),
];

const loggerInstance = winston.createLogger({
  level: "error",
  levels: winston.config.npm.levels,
  format: winston.format.combine(
    winston.format.timestamp({
      format: "YYYY-MM-DD HH:mm:ss",
    }),
    winston.format.errors({ stack: true }),
    winston.format.splat(),
    winston.format.json(),
  ),
  transports: transports,
});

export default loggerInstance;
