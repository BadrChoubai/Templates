import uuid

from sqlmodel import Session, select

from app.api.models import Item, ItemCreate


def get_all(session: Session) -> list[Item]:
    return list(session.exec(select(Item)).all())


def get_by_id(session: Session, item_id: uuid.UUID) -> Item | None:
    return session.get(Item, item_id)


def create(session: Session, item_in: ItemCreate) -> Item:
    item = Item.model_validate(item_in)
    session.add(item)
    session.commit()
    session.refresh(item)
    return item
