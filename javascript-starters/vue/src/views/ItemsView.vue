<script setup lang="ts">
import { onMounted, ref } from "vue";
import { RouterLink } from "vue-router";
import { listItems, type Item } from "@/api/items";
import ProblemAlert from "@/components/ProblemAlert.vue";

const items = ref<Item[]>([]);
const loading = ref(true);
const error = ref<unknown>(null);

async function load() {
  loading.value = true;
  error.value = null;
  try {
    items.value = await listItems();
  } catch (e) {
    error.value = e;
  } finally {
    loading.value = false;
  }
}

onMounted(load);
</script>

<template>
  <div class="heading">
    <h1>Items</h1>
    <RouterLink to="/items/new" class="btn">New item</RouterLink>
  </div>

  <p v-if="loading" aria-live="polite">Loading items…</p>
  <ProblemAlert v-else-if="error" :error="error">
    <button type="button" class="btn sm retry" @click="load">Try again</button>
  </ProblemAlert>
  <p v-else-if="items.length === 0">No items yet.</p>
  <ul v-else class="items" role="list">
    <li v-for="item in items" :key="item.id" class="item">
      <h2 class="item__name">
        <RouterLink :to="{ name: 'item', params: { id: item.id } }">
          {{ item.name }}
        </RouterLink>
      </h2>
      <p v-if="item.description" class="item__desc">{{ item.description }}</p>
    </li>
  </ul>
</template>

<style scoped>
.heading {
  display: flex;
  flex-wrap: wrap;
  align-items: center;
  justify-content: space-between;
  gap: var(--app-spacing);
  margin-bottom: 1.5rem;
}

h1 {
  margin: 0;
}

.retry {
  margin-top: 0.75rem;
}

.items {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(16rem, 1fr));
  gap: var(--app-spacing);
  margin: 0;
  padding: 0;
  list-style: none;
}

.item {
  position: relative;
  padding: 1.25rem;
  border: var(--app-border-width) solid var(--app-border-color);
  border-radius: var(--app-border-radius);
  transition: border-color var(--app-transition);
}

.item:hover,
.item:focus-within {
  border-color: var(--app-accent-color);
}

.item__name {
  margin: 0;
  font-size: 1.125rem;
}

.item__name a {
  color: var(--app-color);
  text-decoration: none;
}

/* The whole card is the link's click target */
.item__name a::after {
  content: "";
  position: absolute;
  inset: 0;
}

.item__desc {
  margin: 0.5rem 0 0;
  color: var(--app-muted-color);
}
</style>
