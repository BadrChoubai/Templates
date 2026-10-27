/**
 * Site-wide settings: the name in the header, page titles and footer, and
 * the links in the header menu. Set `site` in astro.config.mjs to match.
 */
export const SITE = {
  title: "Astro Starter",
  author: "Your Name",
  description: "A starting point for an Astro site.",
  // Browser UI color; match --app-background-color in src/styles/colors.css
  themeColor: { light: "#ffffff", dark: "#121417" },
};

export const MENU: { label: string; href: string }[] = [
  { label: "home", href: "/" },
  { label: "items", href: "/#items" },
];
