server:
	docker compose up server

up:
	docker compose up

stop:
	docker compose stop

down:
	docker compose down --remove-orphans -v

logs:
	docker compose logs -f

tests:
	docker compose up tests

migrate:
	docker compose up migrations
	@make stop

migration:
	@read -p "Enter migration name: " migration_name; \
	file_name=$(shell date +%s)_$$migration_name; \
	touch internal/migrations/$$file_name.up.sql; \
	touch internal/migrations/$$file_name.down.sql

bash:
	docker compose exec server sh

