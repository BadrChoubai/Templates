import { ItemsService } from "../../service/items.js";
import { ItemDbResponse, ItemResponse } from "../../models/Item.js";
import { transformResponse } from "../../lib/transformResponse.js";

export const getItems = (itemsService: ItemsService) => {
  return async function (req, res) {
    try {
      const items = await itemsService.getAll();

      const itemsListResponse: ItemResponse[] = items.map(
        (item: ItemDbResponse): ItemResponse =>
          transformResponse<ItemDbResponse, ItemResponse>(item, toItemResponse),
      );

      const data = {
        items: itemsListResponse || [],
      };

      res.status(200).send(data);
    } catch (error) {
      console.error("Error fetching items data:", error);
      res.status(500).send({ error: "Internal server error" });
    }
  };
};

function toItemResponse(data: ItemDbResponse): ItemResponse {
  return {
    id: data.id,
    name: data.name,
    description: data.description,
  };
}
