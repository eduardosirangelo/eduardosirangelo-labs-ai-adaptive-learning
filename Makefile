# Makefile — Root do monorepo

# Nome do projeto, usado pelo Compose
# Você pode sobrepor via CLI: make docker PROJECT_NAME=custom
PROJECT_NAME ?= ai_adaptive_learning

# Caminho para o docker-compose
COMPOSE_FILE := docker-compose.yml

# Opções padrão para docker-compose
DC := docker compose -p $(PROJECT_NAME) -f $(COMPOSE_FILE)

# Tempo de espera para healthchecks
HEALTH_TIMEOUT ?= 30s

.PHONY: help docker up down logs ps restart build clean

## help: exibe este menu
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) \
		| sort \
		| awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}'

## docker: sobe o ambiente em background (docker-compose up -d)
docker: ## exporta COMPOSE_PROJECT_NAME e sobe containers
	@echo "📦  Starting services with project name '$(PROJECT_NAME)'..."
	@$(DC) up -d

## up: alias para docker
up: docker

## down: derruba o ambiente (docker-compose down)
down: ## para e remove containers, redes e volumes anônimos
	@echo "🛑  Stopping services..."
	@$(DC) down --remove-orphans

## logs: acompanha os logs de todos os serviços
logs: ## tail -f de logs de containers
	@echo "📝  Streaming logs..."
	@$(DC) logs -f --tail=100

## ps: lista containers do projeto
ps: ## lista containers gerenciados pelo Compose
	@echo "🔍  Project containers:"
	@$(DC) ps

## restart: reinicia os containers
restart: ## quick restart (down + up)
	@$(MAKE) down
	@$(MAKE) up

## build: força rebuild das imagens
build: ## faz build de todas as imagens sem cache
	@echo "🔨  Rebuilding images..."
	@$(DC) build --no-cache

## clean: remove imagens locais geradas pelo Compose
clean: down ## executa down e depois prune
	@echo "🧹  Cleaning up unused images and volumes..."
	@docker system prune -f
	@docker volume prune -f
