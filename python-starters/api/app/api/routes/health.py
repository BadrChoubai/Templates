from fastapi import APIRouter

router = APIRouter()


@router.get("/health/", name="Health Check")
def health() -> bool:
    return True
