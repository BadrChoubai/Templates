<script setup lang="ts">
import { RouterLink, RouterView } from "vue-router";
</script>

<template>
  <a href="#content" class="btn sm sr-only skip-link">Skip to main content</a>
  <header class="site-header container">
    <RouterLink to="/" class="logo">Vue Starter</RouterLink>
    <nav aria-label="Site navigation">
      <ul>
        <li><RouterLink to="/">items</RouterLink></li>
        <li><RouterLink to="/items/new">new item</RouterLink></li>
      </ul>
    </nav>
  </header>
  <main id="content" class="container">
    <RouterView />
  </main>
  <footer class="site-footer">
    <div class="container">
      <small>Built with Vue 3 and Vite</small>
    </div>
  </footer>
</template>

<style scoped>
.site-header {
  display: flex;
  flex-wrap: wrap;
  align-items: center;
  justify-content: space-between;
  gap: var(--app-spacing);
  padding-block: 1rem;
}

.logo {
  color: var(--app-color);
  font-family: var(--app-font-family-headings);
  font-size: 1.25rem;
  font-weight: 700;
  text-decoration: none;
}

/* RouterLink marks the link for the current page */
nav a.router-link-exact-active {
  color: var(--app-accent-color);
}

main {
  padding-block: var(--app-section-spacing);
}

.site-footer {
  padding-block: 1rem;
  border-top: var(--app-rule-width) solid var(--app-rule-color);
  color: var(--app-muted-color);
}
</style>
