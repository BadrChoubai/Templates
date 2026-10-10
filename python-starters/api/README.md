# Python API Starter

This is a Python API starter using FastAPI. It exposes an example `items` resource
backed by an in-memory store in `app/api/repository.py`.

## Prerequisites

- [Docker](https://www.docker.com/) (For containerized build and development)
- [uv](https://docs.astral.sh/uv/) for Python package and environment management

### Installing Dependencies and Attaching to `venv` Locally

```shell
# Sync dependencies
uv sync

# After syncing, activate the virtual environment
source .venv/bin/activate
```

## Running the Application

Setting up and running the application may be done using a task
contained in `Makefile` or with `docker compose watch`

### Local Development

To run the application locally, use the `fastapi` CLI which can run our application locally,
you may also use docker compose after building the container image. 

```shell
# Run the application locally
fastapi run

# Or, run the built container image
make run
```

## Testing the Application

No tests have been added yet.

```shell
make test
```

## Linting

Linting and formatting is achieved and enforced by [`ruff`](https://docs.astral.sh/ruff/) and 
is run on all files located in the `app` and `scripts` directories.

- `make format` - runs `format.sh`
- `make lint` - runs `lint.sh`

## Using This Template

1. Set `name` in `pyproject.toml` (then run `uv lock`) and the values in `.env.example`.
2. Replace the example `items` resource (`app/api/models.py`, `app/api/repository.py`,
   `app/api/routes/items.py`) with your own.
