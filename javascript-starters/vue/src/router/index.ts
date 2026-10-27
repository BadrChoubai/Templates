import { createRouter, createWebHistory } from "vue-router";
import ItemsView from "@/views/ItemsView.vue";

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    { path: "/", name: "items", component: ItemsView },
    // Loaded on first visit, so they stay out of the initial bundle
    {
      path: "/items/new",
      name: "new-item",
      component: () => import("@/views/NewItemView.vue"),
    },
    {
      path: "/items/:id",
      name: "item",
      component: () => import("@/views/ItemView.vue"),
      props: true,
    },
    {
      path: "/:pathMatch(.*)*",
      name: "not-found",
      component: () => import("@/views/NotFoundView.vue"),
    },
  ],
});

export default router;
