# Database Makefile Commands

> Migration tasks must be run at the root of the project.

This repository contains a `Makefile` with useful commands for managing a SQLite
database, including connecting to the database, running migrations, and
refreshing tables.

## Prerequisites

Ensure you have the following installed on your system:

- `sqlite3`
- `migrate` (Database migration tool)
- DataGrip or similar Visual Database Management tool

### Installing `migrate`

Installation of the tool can be done using Go toolchain, the value of `-tags` is
used
to install the correct database engine library with the tool.

> ```shell
>  go install -tags 'sqlite3' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
> ```
>
> [Documentation][1]

## Usage

Run the following `make` commands as needed:

### Database Information

```sh
make db/info
```

Prints the database URL passed to `migrate`.

### Connect to the Database

```sh
make db/connect
```

Opens a connection to the locally running SQLite database.

### Run Migrations (Up)

```sh
make db/migrations/up
```

Runs migrations to create database tables.

### Rollback Migrations (Down)

```sh
make db/migrations/down
```

Rolls back migrations, tearing down database tables.

### Refresh Database Migrations

```sh
make db/migrations/refresh
```

Drops all tables and reruns the migrations from scratch.

## Environment Variables

The database file is set by `DB_CONNECTION_STRING` (default `.db/api.db`), the
same variable the application reads. Set it in `.env` or pass it to `make`:

```sh
make DB_CONNECTION_STRING=".db/other.db" db/migrations/up
```

---

[1]: https://github.com/golang-migrate/migrate/tree/master/cmd/migrate
