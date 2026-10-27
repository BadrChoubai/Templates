<script setup lang="ts">
import { reactive, ref } from "vue";
import { RouterLink, useRouter } from "vue-router";
import { ApiError } from "@/api/client";
import { createItem } from "@/api/items";
import ProblemAlert from "@/components/ProblemAlert.vue";

// Same limits the API enforces; maxlength stops typing past them. The form
// is novalidate, so a missing name reaches the API and its validation
// message shows under the field.
const maxNameLength = 255;
const maxDescriptionLength = 1024;

const router = useRouter();

const form = reactive({ name: "", description: "" });
const submitting = ref(false);
const error = ref<unknown>(null);

/** Field messages from the API's validation problem, if that's what failed */
function fieldErrors(field: string): string[] {
  return error.value instanceof ApiError ? error.value.fieldErrors(field) : [];
}

async function submit() {
  submitting.value = true;
  error.value = null;
  try {
    const item = await createItem({ ...form });
    await router.push({ name: "item", params: { id: item.id } });
  } catch (e) {
    error.value = e;
  } finally {
    submitting.value = false;
  }
}
</script>

<template>
  <p><RouterLink to="/">← All items</RouterLink></p>
  <h1>New item</h1>

  <form class="form" novalidate @submit.prevent="submit">
    <ProblemAlert v-if="error" :error="error" />

    <div class="field">
      <label for="name">Name</label>
      <input
        id="name"
        v-model="form.name"
        name="name"
        required
        :maxlength="maxNameLength"
        :aria-invalid="fieldErrors('name').length > 0"
        aria-describedby="name-errors"
      />
      <ul id="name-errors" class="field__errors">
        <li v-for="message in fieldErrors('name')" :key="message">
          {{ message }}
        </li>
      </ul>
    </div>

    <div class="field">
      <label for="description">Description <span>(optional)</span></label>
      <textarea
        id="description"
        v-model="form.description"
        name="description"
        rows="4"
        :maxlength="maxDescriptionLength"
        :aria-invalid="fieldErrors('description').length > 0"
        aria-describedby="description-errors"
      ></textarea>
      <ul id="description-errors" class="field__errors">
        <li v-for="message in fieldErrors('description')" :key="message">
          {{ message }}
        </li>
      </ul>
    </div>

    <div>
      <button type="submit" :disabled="submitting">
        {{ submitting ? "Creating…" : "Create item" }}
      </button>
    </div>
  </form>
</template>

<style scoped>
.form {
  display: flex;
  flex-direction: column;
  gap: 1.25rem;
  max-width: 36rem;
}

.field {
  display: flex;
  flex-direction: column;
  gap: 0.375rem;
}

label {
  font-weight: 600;
}

label span {
  color: var(--app-muted-color);
  font-weight: 400;
}

[aria-invalid="true"] {
  border-color: var(--app-danger-color);
}

.field__errors {
  margin: 0;
  padding: 0;
  list-style: none;
  color: var(--app-danger-color);
  font-size: 0.875rem;
}

.field__errors:empty {
  display: none;
}

button:disabled {
  cursor: progress;
  opacity: 0.7;
}
</style>
