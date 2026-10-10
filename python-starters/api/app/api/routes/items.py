import uuid

from fastapi import APIRouter, HTTPException, status

from app.api import repository
from app.api.models import Item, ItemCreate

router = APIRouter(prefix="/items", tags=["items"])


@router.get("/", response_model=list[Item])
async def list_items() -> list[Item]:
    return repository.get_all()


@router.get("/{item_id}", response_model=Item)
async def get_item(item_id: uuid.UUID) -> Item:
    item = repository.get_by_id(item_id)
    if item is None:
        raise HTTPException(
            status_code=status.HTTP_404_NOT_FOUND, detail="Item not found"
        )
    return item


@router.post("/", response_model=Item, status_code=status.HTTP_201_CREATED)
async def create_item(item: ItemCreate) -> Item:
    return repository.create(item)
