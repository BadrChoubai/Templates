<script setup lang="ts">
import { ref, watch } from "vue";
import { RouterLink } from "vue-router";
import { ApiError } from "@/api/client";
import { getItem, type Item } from "@/api/items";
import ProblemAlert from "@/components/ProblemAlert.vue";

const { id } = defineProps<{ id: string }>();

const item = ref<Item | null>(null);
const loading = ref(true);
const error = ref<unknown>(null);

// Reload when navigating between items, since the view is reused
watch(
  () => id,
  async (id) => {
    loading.value = true;
    error.value = null;
    try {
      item.value = await getItem(id);
    } catch (e) {
      error.value = e;
    } finally {
      loading.value = false;
    }
  },
  { immediate: true },
);
</script>

<template>
  <p><RouterLink to="/">← All items</RouterLink></p>

  <p v-if="loading" aria-live="polite">Loading item…</p>
  <template v-else-if="error">
    <h1 v-if="error instanceof ApiError && error.status === 404">
      Item not found
    </h1>
    <ProblemAlert :error="error" />
  </template>
  <article v-else-if="item">
    <h1>{{ item.name }}</h1>
    <p v-if="item.description">{{ item.description }}</p>
    <p v-else class="muted">No description.</p>
    <dl>
      <dt>ID</dt>
      <dd>
        <code>{{ item.id }}</code>
      </dd>
    </dl>
  </article>
</template>

<style scoped>
.muted,
dt {
  color: var(--app-muted-color);
}

dl {
  margin-top: 1.5rem;
}

dt {
  font-size: 0.875rem;
}

dd {
  margin: 0.25rem 0 0;
}
</style>
