import type { Request, Response } from "express";
import Logger from "../../logger/winston.js";
import { ItemsService } from "../../service/items.js";
import { ItemDbResponse, ItemResponse } from "../../models/Item.js";
import { transformResponse } from "../../lib/transformResponse.js";

export const getItem = (itemsService: ItemsService) => {
  return async function (req: Request<{ id: string }>, res: Response) {
    try {
      const { id } = req.params;
      const item = await itemsService.getOne(id);

      if (!item) {
        return res.status(404).send({
          error: `Item with given id ${id} not found`,
        });
      }

      const data = transformResponse<ItemDbResponse, ItemResponse>(
        item,
        toItemResponse,
      );

      res.status(200).send(data);
    } catch (error) {
      Logger.error("Error fetching item data", error);
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
