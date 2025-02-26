DATABASE_URL="postgresql://dbadmin:superadminsecret@localhost:5432/identity-provider-local-db?sslmode=disable"
MIGRATION_PATH="internal/database/migration"

.PHONY: checkdeps help
help: ## Displays the help for each command.
	@grep -E '^[a-zA-Z_-]+:.*## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

dbu: ## Migrates the database to the latest version.
	migrate -database="$(DATABASE_URL)" -path=$(MIGRATION_PATH) -lock-timeout=20  -verbose up

dbd: ## Migrates the database down by first version.
	migrate -database="$(DATABASE_URL)" -path=$(MIGRATION_PATH) -lock-timeout=20 -verbose down

dbdrop: ## Drops the database.
	migrate -database="$(DATABASE_URL)" -path=$(MIGRATION_PATH) -lock-timeout=20 -verbose drop -f

dbseq: ## Creates a new migration file with the given name. usage: name=create_users_table
	migrate create -ext sql -dir $(MIGRATION_PATH) $(name)

dbver: ## Prints the current migration version.
	migrate -database="$(DATABASE_URL)" -path=$(MIGRATION_PATH) -lock-timeout=20 version

dbf: ## Force the migrates the database to the given version. usage: version=1
	migrate -database="$(DATABASE_URL)" -path=$(MIGRATION_PATH) -lock-timeout=20 -verbose force $(version)

dbt: ## Goto migrates the database to the given version. usage: version=1
	migrate -database="$(DATABASE_URL)" -path=$(MIGRATION_PATH) -lock-timeout=20 -verbose goto $(version)

dbreset: dbdrop dbu ## Reset database to the latest version.
