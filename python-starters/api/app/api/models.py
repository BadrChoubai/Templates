import uuid

from sqlmodel import Field, SQLModel


class ItemBase(SQLModel):
    name: str = Field(nullable=False, max_length=255, description="Item name")
    description: str = Field(
        default="", max_length=1024, description="Item description"
    )


class ItemCreate(ItemBase):
    pass


class Item(ItemBase):
    id: uuid.UUID = Field(primary_key=True)
