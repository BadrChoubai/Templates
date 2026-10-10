# TypeScript API Starter

This is a TypeScript API starter using Express.js and SQLite (`node:sqlite`). It provides a development and production setup with
TypeScript compilation and linting.

## Prerequisites

Ensure you have the following installed:

- [Node.js](https://nodejs.org/) (Latest LTS recommended)
- [npm](https://www.npmjs.com/) (Comes with Node.js)

#### Development Mode (with Hot Reloading)

```sh
npm run dev
```

This will start the TypeScript compiler in watch mode and run the server using `nodemon`.

#### Production Mode

First, build the project:

```sh
npm run build
```

Then, start the server:

```sh
npm run serve
```

### Linting

To run ESLint and automatically fix issues:

```sh
npm run lint
```

### Database

The example `items` table lives in a SQLite database at `.db/api.db`. Create and
seed it before starting the server:

```sh
mkdir -p .db
npm run database:seed
```

## Using This Template

1. Set `name` in `package.json` and the image names in `package.json` and `compose.yaml`.
2. Replace the example `items` resource (`src/models`, `src/service`, `src/router`,
   `src/database/seedDb.ts`) with your own.
