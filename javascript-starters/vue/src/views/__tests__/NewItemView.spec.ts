import { flushPromises, mount } from "@vue/test-utils";
import { afterEach, describe, expect, it, vi } from "vitest";
import { createMemoryHistory, createRouter } from "vue-router";
import { ApiError } from "@/api/client";
import * as items from "@/api/items";
import NewItemView from "../NewItemView.vue";

function mountView() {
  const router = createRouter({
    history: createMemoryHistory(),
    routes: [
      { path: "/", component: { template: "<div />" } },
      { path: "/items/new", component: NewItemView },
      { path: "/items/:id", name: "item", component: { template: "<div />" } },
    ],
  });
  return {
    router,
    wrapper: mount(NewItemView, { global: { plugins: [router] } }),
  };
}

afterEach(() => {
  vi.restoreAllMocks();
});

describe("NewItemView", () => {
  it("opens the new item once it's created", async () => {
    const created = { id: "abc", name: "Lamp", description: "" };
    const createItem = vi.spyOn(items, "createItem").mockResolvedValue(created);
    const { router, wrapper } = mountView();
    const push = vi.spyOn(router, "push");

    await wrapper.find("#name").setValue("Lamp");
    await wrapper.find("form").trigger("submit");
    await flushPromises();

    expect(createItem).toHaveBeenCalledWith({ name: "Lamp", description: "" });
    expect(push).toHaveBeenCalledWith({ name: "item", params: { id: "abc" } });
  });

  it("shows the API's validation messages under their fields", async () => {
    vi.spyOn(items, "createItem").mockRejectedValue(
      new ApiError(400, {
        title: "One or more validation errors occurred.",
        status: 400,
        errors: { name: ["name is required"] },
      }),
    );
    const { wrapper } = mountView();

    await wrapper.find("form").trigger("submit");
    await flushPromises();

    expect(wrapper.find("#name-errors").text()).toBe("name is required");
    expect(wrapper.find("#name").attributes("aria-invalid")).toBe("true");
    expect(wrapper.find('[role="alert"]').text()).toContain(
      "One or more validation errors occurred.",
    );
  });
});
