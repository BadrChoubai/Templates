####################################################################################################
# Database Tasks
####################################################################################################

db/info: # @HELP prints the database URL used by migrate
db/info:
	@echo ${MIGRATE_DATABASE_URL}

db/connect: # @HELP connect to locally running database
db/connect:
	sqlite3 ${DB_CONNECTION_STRING}

db/migrations/up: # @HELP run migrations to create database tables
db/migrations/up:
	migrate -path ./migrations/migrations -database ${MIGRATE_DATABASE_URL} up


db/migrations/down: # @HELP run migrations to teardown database tables
db/migrations/down:
	migrate -path ./migrations/migrations -database ${MIGRATE_DATABASE_URL} down


db/migrations/refresh: # @HELP refresh database tables
db/migrations/refresh: db/migrations/down db/migrations/up

ENGINE := sqlite3
DB_CONNECTION_STRING ?= .db/api.db
MIGRATE_DATABASE_URL := ${ENGINE}://${DB_CONNECTION_STRING}

