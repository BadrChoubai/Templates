import { defineCollection } from "astro:content";
import { glob } from "astro/loaders";
import { z } from "astro/zod";

// Example collection rendered by src/sections/Items.astro; one .md file per item
const items = defineCollection({
  loader: glob({ pattern: "**/*.md", base: "./src/content/items" }),
  schema: z.object({
    title: z.string(),
    description: z.string(),
    tags: z.array(z.string()).default([]),
    link: z.url().optional(),
    // Lower numbers show first
    order: z.number().default(100),
  }),
});

export const collections = { items };
