// @ts-check
import { defineConfig } from "astro/config";

import sitemap from "@astrojs/sitemap";

// https://astro.build/config
export default defineConfig({
  // Your production URL: used for canonical links, social cards and the sitemap
  site: "https://example.com",
  server: { port: 4321 },

  integrations: [sitemap()],
});
