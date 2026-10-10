# Templates

This repository contains code boilerplate and is used by [buddy.config](https://www.github.com/BadrChoubai/buddy.config)
to hydrate `XDG_TEMPLATES_DIR`.

## Templates

Each language has an API starter kit exposing `/health` and an example `items`
resource, wired through the layers the starter uses so the patterns are visible.
Replace `items` with your own resources when starting a project.

| Template                    | Language   | Stack                          |
| :-------------------------- | :--------- | :----------------------------- |
| [Go][go]                    | Go         | `net/http`, SQLite (Go 1.26)   |
| [TypeScript][typescript]    | TypeScript | Express 5, `node:sqlite`       |
| [C#][csharp]                | C#         | ASP.NET Core 10, EF Core       |
| [Python][python]            | Python     | FastAPI, SQLModel, Alembic     |

Each template is self-contained and documents how to build, run, and adapt it in its own README.

[go]: ./go-starters/api/README.md
[typescript]: ./typescript-starters/api/README.md
[csharp]: ./csharp-dotnet-starters/api/README.md
[python]: ./python-starters/api/README.md
