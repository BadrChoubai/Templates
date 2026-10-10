from pydantic_settings import BaseSettings, SettingsConfigDict


class Settings(BaseSettings):
    model_config = SettingsConfigDict(env_file=".env", extra="ignore")

    PROJECT_NAME: str = "API"
    DOCS_URL: str = "/docs"
    DATABASE_URL: str = "sqlite:///.db/api.db"


settings = Settings()
