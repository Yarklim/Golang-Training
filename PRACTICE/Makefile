include .env
export

COMPOSE := docker-compose

env-up:
	$(COMPOSE) up -d go-practice-postgres

env-stop:
	$(COMPOSE) stop go-practice-postgres

env-down:
	$(COMPOSE) down

env-cleanup:
	@read -p "Clear all volume environment files? [y/N]: " ans; \
	if [ "$$ans" = "y" ]; then \
		$(COMPOSE) down -v; \
	else \
		echo "Cleanup cancelled"; \
	fi

env-logs:
	$(COMPOSE) logs -f go-practice-postgres

env-ps:
	$(COMPOSE) ps -a

debug:
	which docker
	docker --version
	which docker-compose
	docker-compose version