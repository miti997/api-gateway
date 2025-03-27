CMD = docker compose

up:
	$(CMD) up -d

down:
	$(CMD) down -v --remove-orphans

build:
	$(CMD) build
	docker image prune -f

recreate: down build up

test:
	go test ./...

test-v:
	go test ./... -v
