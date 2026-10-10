from fastapi import APIRouter
from fastapi.responses import PlainTextResponse

router = APIRouter()


@router.api_route("/health", methods=["GET", "HEAD"], name="Health Check")
def health() -> PlainTextResponse:
    return PlainTextResponse("OK")
