# Templates

This repository contains code boilerplate and is used by [buddy.config](https://www.github.com/BadrChoubai/buddy.config)
to hydrate `XDG_TEMPLATES_DIR`.

## Templates

There are API starter kits for Go and C#/.NET. Each exposes `/health` and an example `items`
resource, wired through the layers the starter uses so the patterns are visible.
Replace `items` with your own resources when starting a project.

There is also a website starter for Astro, with an example `items` content collection
rendered on its home page, and a Vue 3 app that lists and creates `items` through either API.

| Template                    | Type    | Language   | Stack                                |
| :-------------------------- | :------ | :--------- | :----------------------------------- |
| [Go][go]                    | API     | Go         | `net/http`, SQLite (Go 1.26)         |
| [C#][csharp]                | API     | C#         | ASP.NET Core 10, EF Core             |
| [Astro][astro]              | Website | TypeScript | Astro 6, Cloudflare Workers          |
| [Vue][vue]                  | Web app | TypeScript | Vue 3, Vite, Vue Router, Vitest      |

Each template is self-contained and documents how to build, run, and adapt it in its own README.

[go]: ./go-starters/api/README.md
[csharp]: ./csharp-dotnet-starters/api/README.md
[astro]: ./javascript-starters/astro/README.md
[vue]: ./javascript-starters/vue/README.md
