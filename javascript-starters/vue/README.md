# Vue Starter

A [Vue 3](https://vuejs.org/) single-page app that lists, shows and creates
`items` through the API from either the [Go](../../go-starters/api/README.md)
or the [C#](../../csharp-dotnet-starters/api/README.md) starter. It shows the
patterns most API clients need: a typed fetch wrapper, loading and error
states, Problem Details errors shown on the page, and validation messages
matched to form fields.

## Prerequisites

- [Node.js](https://nodejs.org/) 22.18 or later (`.nvmrc` pins 24)
- One of the API starters, running locally
- [Docker](https://www.docker.com/) (for the container image)

## Running the Application

### Start an API

Run the [Go](../../go-starters/api/README.md) or the
[C#](../../csharp-dotnet-starters/api/README.md) starter. The app expects the
Go starter on `http://localhost:8084`; to use another address, set it in `.env`:

```sh
cp .env.example .env    # then edit API_PROXY_TARGET
```

### Run Locally

```sh
npm install
npm run dev
```

The app is then available at `http://localhost:5173`. The dev server forwards
requests for `/api/*` to the API (`/api/items` becomes `/items`), so the API
needs no CORS setup.

### Build and Run the Docker Image

The image serves the built app with nginx and forwards `/api/*` to `API_URL`,
just like the dev server:

```sh
docker build -t vue-starter .
docker run --rm -p 8080:8080 \
  -e API_URL=http://host.docker.internal:8084 \
  --add-host=host.docker.internal:host-gateway \
  vue-starter
```

The app is then available at `http://localhost:8080`. If the API runs in a
container on the same Docker network, use its name instead, e.g.
`API_URL=http://api:8084`.

## Commands

| Command                | Action                                          |
| :--------------------- | :---------------------------------------------- |
| `npm run dev`          | Start the dev server at `localhost:5173`        |
| `npm run build`        | Type-check and build the app to `./dist/`       |
| `npm run preview`      | Serve the build locally                         |
| `npm test`             | Run the unit tests once (`test:watch` to watch) |
| `npm run type-check`   | Type-check `.vue` and `.ts` files               |
| `npm run lint`         | Lint with oxlint and ESLint (`lint:fix` to fix) |
| `npm run format`       | Format every file with Prettier                 |
| `npm run format:check` | Check formatting without changing files         |

## Project Structure

```
.
├── public/                 # Copied to the build as-is (favicon)
├── src/
│   ├── api/
│   │   ├── client.ts       # fetch wrapper, ApiError and Problem Details types
│   │   ├── items.ts        # Item types and the items endpoints
│   │   └── __tests__/      # Tests for the client against both APIs' responses
│   ├── components/         # Shared components (ProblemAlert)
│   ├── router/             # Routes: list, new item, item detail, not found
│   ├── styles/             # Global CSS: base, colors, typography, document
│   ├── views/              # One component per route, with tests in __tests__/
│   ├── App.vue             # Header, footer and the <RouterView>
│   └── main.ts             # Creates the app and loads the global styles
├── .env.example            # Dev server settings (copy to .env)
├── Dockerfile              # Builds the app, serves it with nginx
├── nginx.conf.template     # nginx config: app routes and the /api proxy
├── index.html
└── vite.config.ts          # Vite config, including the dev /api proxy
```

## Talking to the API

Requests go through `src/api/client.ts`, which:

- prefixes every path with `/api`, or with `VITE_API_BASE_URL` if set at build
  time (to call the API directly; it must then allow this site's origin)
- throws an `ApiError` for any error response, carrying the
  [Problem Details](https://www.rfc-editor.org/rfc/rfc9457) body; views show it
  with `ProblemAlert`, and `error.fieldErrors("name")` returns the validation
  messages for one field

The two APIs differ in one place: `GET /items` returns `{ items, count }` from
Go and a plain array from C#. `listItems` in `src/api/items.ts` accepts both;
once you've picked an API, you can drop the other branch.

## Styling

The styles in `src/styles/` are the same as the
[Astro starter's](../astro/README.md#styling): system fonts, one accent colour,
and light and dark modes that follow the visitor's system setting, all set
through CSS custom properties. Components keep their own `<style scoped>` blocks
and use only those tokens.

## Using This Template

1. Set `name` in `package.json`, and the title and description in
   `index.html` and `src/App.vue`.
2. Point `API_PROXY_TARGET` (`.env`) and `API_URL` (Docker) at your API.
3. Replace `items` (`src/api/items.ts`, `src/views/`, `src/router/index.ts`)
   with your own resources.
4. Replace `public/favicon.svg`, and set the colours in
   `src/styles/colors.css`.
