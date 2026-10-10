"""create items

Revision ID: 0001
Revises:
Create Date: 2026-10-09 00:00:00.000000

"""

import uuid
from collections.abc import Sequence

import sqlalchemy as sa
from alembic import op

# revision identifiers, used by Alembic.
revision: str = "0001"
down_revision: str | None = None
branch_labels: str | Sequence[str] | None = None
depends_on: str | Sequence[str] | None = None


def upgrade() -> None:
    """Upgrade schema."""
    items = op.create_table(
        "items",
        sa.Column("id", sa.Uuid(), nullable=False),
        sa.Column("name", sa.String(length=255), nullable=False),
        sa.Column("description", sa.String(length=1024), nullable=False),
        sa.PrimaryKeyConstraint("id"),
    )

    # Example data for the starter; remove when replacing the items resource.
    op.bulk_insert(
        items,
        [
            {
                "id": uuid.UUID("c04ab9a0-e398-471d-8cfe-95317ccf0a20"),
                "name": "First item",
                "description": "An example item",
            },
            {
                "id": uuid.UUID("1e95757d-066d-41a6-b5d2-5e34b71d6051"),
                "name": "Second item",
                "description": "Another example item",
            },
        ],
    )


def downgrade() -> None:
    """Downgrade schema."""
    op.drop_table("items")
