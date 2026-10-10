# TypeScript API Starter

This is a TypeScript API starter using Express.js and SQLite (`node:sqlite`). It
provides a development and production setup with TypeScript compilation and
linting, and exposes an example `items` resource.

## Prerequisites

Ensure you have the following installed:

- [Node.js](https://nodejs.org/) 24 or later (see `.nvmrc`)
- [npm](https://www.npmjs.com/) (Comes with Node.js)

Install dependencies:

```sh
npm ci
```

## Running the Application

### Database

The example `items` table lives in a SQLite database at `.db/api.db` (override
with `DATABASE_PATH`). Build the project, then create and seed the database:

```sh
npm run build
npm run database:seed
```

### Development Mode (with Hot Reloading)

```sh
npm run dev
```

This will start the TypeScript compiler in watch mode and run the server using
`nodemon`. The API is available at `http://localhost:8082` (override with `PORT`).

### Production Mode

```sh
npm run build
npm run serve
```

### Docker

```sh
npm run docker:dev    # build and start with docker compose
npm run docker:stop   # stop the containers
```

The image seeds its own database at build time.

## Configuration

| Variable        | Default      | Description                      |
| :-------------- | :----------- | :------------------------------- |
| `PORT`          | `8082`       | HTTP port                        |
| `DATABASE_PATH` | `.db/api.db` | SQLite database file             |
| `LOG_LEVEL`     | `info`       | Winston log level                |

## Linting

```sh
npm run lint        # ESLint with --fix
npm run typecheck   # tsc --noEmit
```

## Using This Template

1. Set `name` in `package.json` and the image names in `package.json` and `compose.yaml`.
2. Replace the example `items` resource (`src/models`, `src/service`, `src/router`,
   `src/database/seedDb.ts`) with your own.
