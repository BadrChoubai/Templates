# Astro Starter

A single [Astro](https://astro.build) site with neutral light and dark styles,
a responsive header menu, a back-to-top button, and an example `items`
content collection. Builds to static files and deploys to
[Cloudflare Workers](https://developers.cloudflare.com/workers/static-assets/).

## Prerequisites

- [Node.js](https://nodejs.org/) 22.12 or later (`.nvmrc` pins 24)
- A [Cloudflare](https://dash.cloudflare.com/) account, to deploy

## Commands

All commands are run from the root of the project:

| Command                  | Action                                             |
| :----------------------- | :------------------------------------------------- |
| `npm install`            | Install dependencies                               |
| `npm run dev`            | Start the dev server at `localhost:4321`           |
| `npm run build`          | Build the site to `./dist/`                        |
| `npm run preview`        | Preview the build locally                          |
| `npm run check`          | Type-check `.astro` and `.ts` files                |
| `npm run lint`           | Check formatting with Prettier                     |
| `npm run format`         | Format every file with Prettier                    |
| `npm run deploy:preview` | Build and serve through Wrangler (`wrangler dev`)  |
| `npm run deploy`         | Build and deploy to Cloudflare (`wrangler deploy`) |

## Project Structure

```
.
├── public/                 # Copied to the build as-is (favicon, robots.txt, manifest)
├── src/
│   ├── components/         # Header, Footer, BackToTop, Head (meta tags), Section
│   ├── content/items/      # Example collection: one Markdown file per item
│   ├── layouts/            # Page shell: head, skip link, header, footer
│   ├── pages/              # File-based routes
│   ├── sections/           # Full-width page sections (Hero, Items)
│   ├── styles/             # Global CSS: base, colors, typography, document
│   ├── content.config.ts   # Content collection schemas
│   └── site.ts             # Site name, author, description, menu links
├── astro.config.mjs
├── wrangler.jsonc          # Cloudflare Workers deployment config
└── package.json
```

## Styling

The starting styles are deliberately plain: system fonts, one accent colour,
and light and dark modes that follow the visitor's system setting. Everything
is driven by CSS custom properties, so a new look is mostly token edits:

- `src/styles/colors.css`: surface, text, border and accent colours, each with
  a light and a dark value (`light-dark()`)
- `src/styles/typography.css`: font stacks and the heading scale
- `src/styles/base.css`: spacing, radius, container width, links, buttons
  (`.btn`, `.btn.ghost`, `.btn.sm`)
- `src/styles/document.css`: box-model reset and root font settings

Components keep their own scoped styles in a `<style>` block and use only these
tokens. To add a web font, install it (for example
`npm i @fontsource-variable/inter`), import it in `src/layouts/Layout.astro`, and
put its name first in `--app-font-family`.

## Using This Template

1. Set the name, author, description and menu in `src/site.ts`, and `name` in
   `package.json` and `public/site.webmanifest`.
2. Set `site` in `astro.config.mjs` and the `Sitemap:` line in
   `public/robots.txt` to your production URL.
3. Replace the example `items` collection (`src/content.config.ts`,
   `src/content/items/`, `src/sections/Items.astro`) with your own content.
4. Replace `public/favicon.svg`, and set the colours in `src/styles/colors.css`
   and `themeColor` in `src/site.ts`.
5. Set `name` in `wrangler.jsonc`, and add a custom domain under `routes` when
   you have one.

If a page needs server rendering, add the Cloudflare adapter with
`npx astro add cloudflare`.
