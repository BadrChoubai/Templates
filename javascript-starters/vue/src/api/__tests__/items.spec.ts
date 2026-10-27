import { afterEach, describe, expect, it, vi } from "vitest";
import { ApiError } from "../client";
import { createItem, getItem, listItems } from "../items";

const item = {
  id: "c04ab9a0-e398-471d-8cfe-95317ccf0a20",
  name: "A",
  description: "",
};

function respond(status: number, body: unknown) {
  const fetch = vi.fn<typeof globalThis.fetch>().mockResolvedValue(
    new Response(JSON.stringify(body), {
      status,
      headers: { "Content-Type": "application/json" },
    }),
  );
  vi.stubGlobal("fetch", fetch);
  return fetch;
}

afterEach(() => {
  vi.unstubAllGlobals();
});

describe("listItems", () => {
  it("reads the Go starter's { items, count } response", async () => {
    const fetch = respond(200, { items: [item], count: 1 });

    await expect(listItems()).resolves.toEqual([item]);
    expect(fetch).toHaveBeenCalledWith("/api/items", expect.anything());
  });

  it("reads the C# starter's bare array response", async () => {
    respond(200, [item]);

    await expect(listItems()).resolves.toEqual([item]);
  });
});

describe("getItem", () => {
  it("throws an ApiError carrying the Problem Details body", async () => {
    respond(404, { title: "Not Found", status: 404, detail: "item not found" });

    const error = await getItem(item.id).catch((e: unknown) => e);

    expect(error).toBeInstanceOf(ApiError);
    expect(error).toMatchObject({ status: 404, message: "item not found" });
  });
});

describe("createItem", () => {
  it("posts the item as JSON", async () => {
    const fetch = respond(201, item);

    await createItem({ name: "A", description: "" });

    const [url, init] = fetch.mock.calls[0]!;
    expect(url).toBe("/api/items");
    expect(init?.method).toBe("POST");
    expect(JSON.parse(init?.body as string)).toEqual({
      name: "A",
      description: "",
    });
    expect(new Headers(init?.headers).get("Content-Type")).toBe(
      "application/json",
    );
  });

  it("exposes validation errors by field", async () => {
    respond(400, {
      title: "One or more validation errors occurred.",
      status: 400,
      errors: { name: ["name is required"] },
    });

    const error = (await createItem({ name: "", description: "" }).catch(
      (e: unknown) => e,
    )) as ApiError;

    expect(error.fieldErrors("name")).toEqual(["name is required"]);
    expect(error.fieldErrors("description")).toEqual([]);
  });
});

it("reports an unreachable API as a network error", async () => {
  vi.stubGlobal(
    "fetch",
    vi.fn<typeof globalThis.fetch>().mockRejectedValue(new TypeError("failed")),
  );

  await expect(listItems()).rejects.toMatchObject({
    status: 0,
    problem: { title: "Network error" },
  });
});
