# API

An HTTP API exposing an example `items` resource. Rename or replace `items`
with your own resources.

## GET `/items`

This route returns a list of all items

```json
{
  "items": [
    {
      "id": "c04ab9a0-e398-471d-8cfe-95317ccf0a20",
      "name": "First item",
      "description": "An example item"
    }
  ],
  "count": 1
}
```

## GET `/items/{id}`

This route returns a single item, `400` if `id` is not a UUID, or `404` if no
item matches.

```json
{
  "id": "c04ab9a0-e398-471d-8cfe-95317ccf0a20",
  "name": "First item",
  "description": "An example item"
}
```
