import { request } from "./client";

export interface Item {
  id: string;
  name: string;
  description: string;
}

export interface CreateItemRequest {
  name: string;
  description: string;
}

// The Go starter wraps the list as { items, count }; the C# starter returns
// a bare array. Accept both so either API works unchanged.
type ListResponse = Item[] | { items: Item[]; count: number };

export async function listItems(): Promise<Item[]> {
  const body = await request<ListResponse>("/items");
  return Array.isArray(body) ? body : body.items;
}

export function getItem(id: string): Promise<Item> {
  return request<Item>(`/items/${encodeURIComponent(id)}`);
}

export function createItem(item: CreateItemRequest): Promise<Item> {
  return request<Item>("/items", {
    method: "POST",
    body: JSON.stringify(item),
  });
}
