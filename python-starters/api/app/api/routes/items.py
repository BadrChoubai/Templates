import uuid

from fastapi import APIRouter, HTTPException, status

from app.api import repository
from app.api.models import ItemCreate, ItemPublic
from app.core.db import SessionDep

router = APIRouter(prefix="/items", tags=["items"])


@router.get("", response_model=list[ItemPublic])
def list_items(session: SessionDep):
    return repository.get_all(session)


@router.get("/{item_id}", response_model=ItemPublic)
def get_item(session: SessionDep, item_id: uuid.UUID):
    item = repository.get_by_id(session, item_id)
    if item is None:
        raise HTTPException(
            status_code=status.HTTP_404_NOT_FOUND, detail="Item not found"
        )
    return item


@router.post("", response_model=ItemPublic, status_code=status.HTTP_201_CREATED)
def create_item(session: SessionDep, item: ItemCreate):
    return repository.create(session, item)
