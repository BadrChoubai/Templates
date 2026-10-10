# C# API Starter

This is a C# API starter using ASP.NET Core minimal APIs (.NET 9.0) and Entity
Framework Core with SQLite. It exposes an example `items` resource.

## Prerequisites

Ensure you have the following installed:

- [.NET SDK](https://dotnet.microsoft.com/en-us/download)
- [SQLite](https://www.sqlite.org/download.html)

## Run the application locally

All commands below are run from the `Api/` project directory.

```sh
cd Api
```

Before we can run the application locally, you should ensure that the Entity
Framework CLI is installed.

```sh
dotnet tool install --global dotnet-ef
```

Once that's been done we can create and apply our initial database migration.

```sh
mkdir -p .db
dotnet ef migrations add Initial
dotnet ef database update
```

Initially there will be no data in the database so we'll want to seed it.

`Program.cs` looks for a `seed` argument and, when it is passed, adds example
items to the database.

```sh
dotnet run seed
dotnet run
```

The API is then available at `http://localhost:8086`; see `Api.http` for example
requests.

### Linting

```sh
dotnet format
```

## Using This Template

1. Rename the `Api` project and namespace if needed, and set the image name in `compose.yaml`.
2. Replace the example `items` resource (`Api/Model`, `Api/Repository`,
   `Api/Service`, `Api/Endpoints.cs`, `Api/DataSeeder.cs`) with your own.
