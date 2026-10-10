import uuid

from app.api.models import Item, ItemCreate

# In-memory store for the example `items` resource. Replace with a database
# session (SQLModel is already a dependency) when adding persistence.
_items: dict[uuid.UUID, Item] = {
    item.id: item
    for item in [
        Item(
            id=uuid.UUID("c04ab9a0-e398-471d-8cfe-95317ccf0a20"),
            name="First item",
            description="An example item",
        ),
        Item(
            id=uuid.UUID("1e95757d-066d-41a6-b5d2-5e34b71d6051"),
            name="Second item",
            description="Another example item",
        ),
    ]
}


def get_all() -> list[Item]:
    return list(_items.values())


def get_by_id(item_id: uuid.UUID) -> Item | None:
    return _items.get(item_id)


def create(item_in: ItemCreate) -> Item:
    item = Item(id=uuid.uuid4(), **item_in.model_dump())
    _items[item.id] = item
    return item
