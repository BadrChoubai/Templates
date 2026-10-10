# Python API Starter

This is a Python API starter using FastAPI, SQLModel and Alembic, backed by
SQLite. It exposes an example `items` resource.

## Prerequisites

- [uv](https://docs.astral.sh/uv/) for Python package and environment management
- [Docker](https://www.docker.com/) (For containerized build and development)

### Installing Dependencies and Attaching to `venv` Locally

```shell
# Sync dependencies
uv sync

# After syncing, activate the virtual environment
source .venv/bin/activate
```

## Running the Application

### Configure the Environment

```shell
cp .env.example .env
```

Settings are read from `.env` by `app/core/config.py`:

| Variable       | Default                | Description                 |
| :------------- | :--------------------- | :-------------------------- |
| `PROJECT_NAME` | `API`                  | Title shown in the API docs |
| `DOCS_URL`     | `/docs`                | Path of the Swagger UI      |
| `DATABASE_URL` | `sqlite:///.db/api.db` | SQLAlchemy database URL     |

### Local Development

```shell
make migrate   # apply database migrations (creates and seeds the items table)
make dev       # run with auto-reload on http://localhost:8000
```

### Docker

```shell
make       # build the image
make run   # start it with docker compose
```

The container applies migrations on startup and stores its database in `./.db`.

## Database Migrations

Migrations live in `app/alembic/versions` and are managed with Alembic:

```shell
alembic upgrade head                            # apply migrations
alembic revision --autogenerate -m "describe"   # generate a migration from model changes
alembic downgrade -1                            # roll back the last migration
```

## Testing the Application

No tests have been added yet.

```shell
make test
```

## Linting

Linting and formatting is achieved and enforced by [`ruff`](https://docs.astral.sh/ruff/),
with type checking by [`mypy`](https://mypy-lang.org/), on files in the `app` and
`scripts` directories.

- `make format` - runs `format.sh`
- `make lint` - runs `lint.sh`

## Using This Template

1. Set `name` in `pyproject.toml` (then run `uv lock`) and the values in `.env.example`.
2. Replace the example `items` resource (`app/api/models.py`, `app/api/repository.py`,
   `app/api/routes/items.py`, `app/alembic/versions/0001_create_items.py`) with your own.
