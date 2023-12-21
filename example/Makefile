COMPOSE := docker compose -f compose/docker-compose.yaml

.PHONY: up
up:
	@$(COMPOSE) up -d --wait --force-recreate --remove-orphans
	@echo "Visit http://localhost:3000/dashboards"

.PHONY: down
down:
	@$(COMPOSE) down -v --remove-orphans
