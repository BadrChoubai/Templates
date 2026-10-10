# Go API Starter

This is a Go-based API project with a Makefile-driven build and deployment process. The project includes support for
building, linting, containerizing, and running the application.

## Prerequisites

Ensure you have the following installed:

- [Go](https://go.dev/) (Latest stable version recommended)
- [Docker](https://www.docker.com/) (For containerized builds)
- [golangci-lint](https://golangci-lint.run/) (For linting)
- [Make](https://www.gnu.org/software/make/) (For build automation)

## Running the Application

### Build the Project

Build our application's docker image

```sh
make build
```

### Run the Application

To run the application locally:

```sh
make run
```

### Clean Build Artifacts

To remove build artifacts:

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

This will generate a Dockerfile and build the corresponding image.

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
└── .env              # Environment variables (not committed)
```

## Using This Template

1. Rename the module in `go.mod` (`github.com/badrchoubai/api`) and update imports.
2. Set `DOCKER_IMAGE_NAME` in `Makefile` and `CONNECTION_STRING` in `migrations/Makefile`.
3. Replace the example `items` resource in `internal/` and `migrations/` with your own.
