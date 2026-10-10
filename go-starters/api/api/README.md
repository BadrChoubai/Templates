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

This route returns a single item, or `404` if no item matches `id`.

```json
{
  "id": "c04ab9a0-e398-471d-8cfe-95317ccf0a20",
  "name": "First item",
  "description": "An example item"
}
```

## POST `/items`

This route creates an item and responds `201 Created` with the item and a
`Location` header.

```json
{
  "name": "New item",
  "description": "An optional description"
}
```

`name` is required and at most 255 characters; `description` is optional and at
most 1024 characters.

## Errors

Errors are [RFC 9457 Problem Details][rfc9457] (`application/problem+json`).
Validation errors list messages per request field under `errors`:

```json
{
  "type": "about:blank",
  "title": "One or more validation errors occurred.",
  "status": 400,
  "errors": {
    "name": ["name is required"]
  }
}
```

Other errors, such as a missing item, carry a `detail` instead:

```json
{
  "type": "about:blank",
  "title": "Not Found",
  "status": 404,
  "detail": "item not found"
}
```

[rfc9457]: https://www.rfc-editor.org/rfc/rfc9457
