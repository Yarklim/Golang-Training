include .env
export

COMPOSE := docker-compose

MIGRATE_DATABASE_URL := postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@go-practice-postgres:5432/$(POSTGRES_DB)?sslmode=disable

env-up:
	@$(COMPOSE) up -d go-practice-postgres

env-stop:
	@$(COMPOSE) stop go-practice-postgres

env-down:
	@$(COMPOSE) down

env-cleanup:
	@read -p "Clear all volume environment files? [y/N]: " ans; \
	if [ "$$ans" = "y" ]; then \
		$(COMPOSE) down -v; \
	else \
		echo "Cleanup cancelled"; \
	fi

env-logs:
	@$(COMPOSE) logs -f go-practice-postgres

env-ps:
	@$(COMPOSE) ps -a

migrate-create:
	@if [ -z "$(name)" ]; then \
		echo "A required parameter 'name' is missing. Usage: make migrate-create name=create_users"; \
		exit 1; \
	fi
	$(COMPOSE) run --rm go-practice-migrate \
		create -ext sql -dir . -seq $(name)

migrate-up:
	@make migrate-action action=up

migrate-down:
	@make migrate-action action=down

migrate-action:
	@if [ -z "$(action)" ]; then \
		echo "A required parameter 'action' is missing. Usage: make migrate-action action=up"; \
		exit 1; \
	fi
	$(COMPOSE) run --rm go-practice-migrate \
		-path . \
		-database "$(MIGRATE_DATABASE_URL)" \
		"${action}"