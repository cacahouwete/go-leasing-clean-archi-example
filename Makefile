# Docker commands
DOCKER = docker
DOCKER_COMPOSE = docker-compose
APP_RUN = $(DOCKER_COMPOSE) run --rm app
APP_EXEC = $(DOCKER_COMPOSE) exec app
GO_EXEC = $(APP_EXEC) go
WFI = $(APP_EXEC) wait-for-it

##
## Project
## -------

init: ## Create override files if not exist
init: docker-compose.override.yml

pull: ## Pull the images (this will refresh them if needed)
pull: init
	$(DOCKER_COMPOSE) pull --ignore-pull-failures

build: ## Pull and build the images
build: pull
	COMPOSE_DOCKER_CLI_BUILD=1 DOCKER_BUILDKIT=1 $(DOCKER_COMPOSE) build --pull

install: ## Pull and build the images
install: build
	$(MAKE) start
	$(MAKE) db

run: ## Run the containers... (no detach --> all logs are visible in the console)
run: init
	$(DOCKER_COMPOSE) up --remove-orphans --force-recreate

start: ## Start the containers in the background (detach the processes)
start: init
	$(DOCKER_COMPOSE) up -d --remove-orphans --force-recreate

stop: ## Stop the containers
stop: init
	$(DOCKER_COMPOSE) stop

kill: ## Force to stop and remove the containers and volumes
kill: init
	$(DOCKER_COMPOSE) kill || true
	$(DOCKER_COMPOSE) down --volumes --remove-orphans

reset: ## Kill and install the containers
reset: kill
	$(MAKE) install

restart: ## Stop and start the containers
restart: kill
	$(MAKE) --no-print-directory start

db-wait: ## Wait for db up
	@$(WFI) db:5432 --timeout=20

db: ## Run migrations
db: db-wait
	$(GO_EXEC) run . db init
	$(GO_EXEC) run . db migrate
	$(MAKE) fixtures

fixtures: ## Run fixture
	$(GO_EXEC) run . db fixtures

.PHONY: init pull build install run start stop kill reset restart db db-wait db-reset fixtures

##
## Utils
## -----

swag: ## Run go swag to generate all api docs
	$(APP_EXEC) swag init

generate: ## Run go generate to generate all mocks
	$(GO_EXEC) generate ./...

.PHONY: swag generate

##
## Tests
## -----

test: ## Launch the tests suite (godog, unit)
	$(MAKE) godog
	$(MAKE) unit

godog: ## Launch the godog test suite
	$(DOCKER_COMPOSE) up -d db app
	$(MAKE) db-wait
	$(DOCKER_COMPOSE) kill app
	$(APP_RUN) sh -c "cd tests && go test"
	$(DOCKER_COMPOSE) up -d app

unit: ## Launch all unit test suite
unit: generate
	$(GO_EXEC) test -race $$(go list ./... | grep -v /tests) -v -coverprofile=coverage.out -coverpkg=./...
	$(GO_EXEC) tool cover -func=coverage.out

.PHONY: godog unit

##
## Quality assurance
## -----------------

qa: ## Launch all QA
qa: format tidy lint

format: ## Launch goimports for formatting analysis
	$(APP_EXEC) goimports -w .

tidy: ## Launch tidy for update go mod
	$(GO_EXEC) mod tidy

lint: generate ## Launch lint form formatting and other errors
	$(DOCKER_COMPOSE) run --rm golangci golangci-lint run --fix

.PHONY: format tidy lint

docker-compose.override.yml: docker-compose.override.yml.dist
	cp -n docker-compose.override.yml.dist docker-compose.override.yml

.DEFAULT_GOAL := help
help: init
	@printf " \033[33mCommands\033[0m:\n"
	@grep -E '(^[a-zA-Z0-9_-]+:.*?##.*$$)|(^##)' Makefile | awk 'BEGIN {FS = ":.*?##s*?"}; {printf "  \033[32m%-20s\033[0m %s\n", $$1, $$2}' | sed -e 's/\[32m##/[33m/'

.PHONY: help
