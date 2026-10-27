<script setup lang="ts">
import { ApiError } from "@/api/client";

/** Shows an error from the API, or any other thrown value, as an alert */
const { error } = defineProps<{ error: unknown }>();
</script>

<template>
  <div class="problem" role="alert">
    <template v-if="error instanceof ApiError">
      <strong>{{ error.problem.title ?? "Something went wrong" }}</strong>
      <p v-if="error.problem.detail">{{ error.problem.detail }}</p>
    </template>
    <strong v-else>Something went wrong</strong>
    <slot />
  </div>
</template>

<style scoped>
.problem {
  padding: var(--app-spacing);
  border: var(--app-border-width) solid var(--app-danger-color);
  border-radius: var(--app-border-radius);
  background-color: var(--app-danger-background);
}

p {
  margin: 0.25rem 0 0;
}
</style>
