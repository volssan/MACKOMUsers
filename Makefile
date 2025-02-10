args=`arg="$(filter-out $@,$(MAKECMDGOALS))" && echo $${arg:-${1}}`

env=./.env
composefile=./docker/docker-compose.yml

## Запуск в docker-compose с ребилдом контейнера detached
dc-up: dc-down
	docker compose -f $(composefile) --env-file $(env) up --build -d

dc-down:
	docker compose -f $(composefile) --env-file $(env) down

# Выполнить миграции
dc-migrate:
	docker compose -f $(composefile) --env-file $(env) run --rm --entrypoint "/bin/sh" api -c "./migrations.sh $(call args)"