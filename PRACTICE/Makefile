include .env
export

COMPOSE := docker-compose

env-up:
	$(COMPOSE) up -d go-practice-postgres

env-down:
	$(COMPOSE) down

env-stop:
	$(COMPOSE) stop go-practice-postgres

env-logs:
	$(COMPOSE) logs -f go-practice-postgres

env-ps:
	$(COMPOSE) ps

debug:
	which docker
	docker --version
	which docker-compose
	docker-compose version