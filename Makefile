MIGRATE_VERSION=v4.19.0
BIN_DIR=./bin
MIGRATE=$(BIN_DIR)/migrate

DB_URL=postgresql://postgre:postgre@localhost:5432/dianping?sslmode=disable
MIGRATIONS_PATH=./migrations

.PHONY: migrate-install 
migrate-install:
	@mkdir -p $(BIN_DIR)
	@if [ ! -f $(MIGRATE) ]; then \
		echo "Installing migrate $(MIGRATE_VERSION)..."; \
		GOBIN=$(PWD)/bin go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@$(MIGRATE_VERSION); \
	fi

.PHONY: migrate-up
migrate-up: migrate-install
	$(MIGRATE) -path $(MIGRATIONS_PATH) -database "$(DB_URL)" up

.PHONY: migrate-down
migrate-down: migrate-install
	$(MIGRATE) -path $(MIGRATIONS_PATH) -database "$(DB_URL)" down 1

.PHONY: migrate-force
migrate-force: migrate-install
	$(MIGRATE) -path $(MIGRATIONS_PATH) -database "$(DB_URL)" force $(version)

.PHONY: migrate-create
migrate-create: migrate-install
	$(MIGRATE) create -ext sql -dir $(MIGRATIONS_PATH) -seq $(name)
