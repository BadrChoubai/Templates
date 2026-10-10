# Go API Starter

This is a Go-based API project with a Makefile-driven build and deployment process. The project includes support for
building, linting, containerizing, and running the application.

## Prerequisites

Ensure you have the following installed:

- [Go](https://go.dev/) 1.26 or later
- [Docker](https://www.docker.com/) (For containerized builds)
- [golangci-lint](https://golangci-lint.run/) (For linting)
- [Make](https://www.gnu.org/software/make/) (For build automation)

## Running the Application

### Configure the Environment

```sh
cp .env.example .env
```

### Create the Database

```sh
mkdir -p .db
make db/migrations/up
```

See [Database Migrations](#database-migrations) for installing `migrate`.

### Run Locally

```sh
make dev
```

The API is then available at `http://localhost:8084`; see `api/api.http` for
example requests.

### Build and Run the Docker Image

```sh
make       # builds the image
make run   # builds the image and starts it with docker compose
```

### Stop the Containers

To stop and remove the docker compose containers:

```sh
make clean
```

### Linting

To check code quality using `golangci-lint`:

```sh
make lint
```

If linting fails, you will be prompted to retry with the `--fix` flag.

## Database Migrations

Database migrations are driven by tasks in `/migrations/Makefile`.

[Documentation](./migrations/README.md)

## Project Structure

```
.
├── api/              # HTTP request examples and API documentation
├── build/scripts/    # Build and lint scripts
├── cmd/api/          # Application entry point and HTTP server setup
├── internal/
│   ├── data/         # Domain types
│   ├── handlers/     # HTTP handlers
│   ├── repository/   # Database access
│   └── routes/       # Route registration
├── migrations/       # SQL migrations and database Makefile tasks
├── pkg/              # Reusable packages (database, encoding, middleware, observability)
├── Dockerfile
├── docker-compose.yml
├── Makefile
├── .env.example      # Example environment variables
└── .env              # Environment variables (not committed)
```

## Using This Template

1. Rename the module in `go.mod` (`github.com/badrchoubai/api`) and update imports.
2. Set `DOCKER_IMAGE` and `DB_CONNECTION_STRING` in `.env.example`.
3. Replace the example `items` resource in `internal/` and `migrations/` with your own.
