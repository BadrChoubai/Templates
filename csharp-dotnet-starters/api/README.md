# C# API Starter

This is a C# API starter using ASP.NET Core minimal APIs (.NET 10) and Entity
Framework Core with SQLite. It exposes an example `items` resource.

## Prerequisites

Ensure you have the following installed:

- [.NET 10 SDK](https://dotnet.microsoft.com/en-us/download)
- [SQLite](https://www.sqlite.org/download.html) (optional, for inspecting the database)

## Run the application locally

Unless noted otherwise, commands below are run from the `Api/` project directory.

```sh
cd Api
```

Pending Entity Framework migrations are applied on startup, creating the SQLite
database at `.db/api.db` (set `ConnectionStrings:DbConnectionString` to change it).

`Program.cs` looks for a `seed` argument and, when it is passed, adds example
items to the database and exits.

```sh
dotnet run seed
dotnet run
```

The API is then available at `http://localhost:8086`. In Development, the
OpenAPI document is served at `/openapi/v1.json` and Swagger UI at `/`. See
`Api.http` for example requests.

### Database Migrations

Migrations live in `Api/Migrations`. After changing the model, add a migration
with the Entity Framework CLI:

```sh
dotnet tool install --global dotnet-ef
dotnet ef migrations add <Name>
```

### Docker

```sh
docker compose up --build
docker compose run --rm api seed   # optional: add example items
```

The API is then available at `http://localhost:8086`.

### Rate Limiting

`/items` is limited to 100 requests per minute per user, or per client IP for
anonymous callers. Override the limits in the `RateLimiting` configuration
section (for example `RateLimiting__TokenLimit=10`).

### Testing

Run from the repository root (where `Api.slnx` lives):

```sh
dotnet test
```

`Api.Tests` hosts the API in memory with `WebApplicationFactory` and swaps the
database for an in-memory SQLite connection (see `Api.Tests/ApiFactory.cs`).
Tests use xUnit v3 on Microsoft.Testing.Platform, enabled in `global.json`.

### Linting

```sh
dotnet format
```

## Using This Template

1. Rename the `Api` project and namespace if needed, and set the image name in `compose.yaml`.
2. Replace the example `items` resource (`Api/Model`, `Api/Repository`,
   `Api/Service`, `Api/Endpoints.cs`, `Api/DataSeeder.cs`) with your own, then
   delete `Api/Migrations` and run `dotnet ef migrations add Initial`.
