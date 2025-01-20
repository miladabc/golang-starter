up:
	docker compose up

start:
	docker compose start

stop:
	docker compose stop

down:
	docker compose down --remove-orphans -v

logs:
	docker compose logs -f

bash:
	docker compose exec server sh

migrate:
	docker compose up migrations

migration:
	@read -p "Enter migration name: " migration_name; \
	file_name=$(shell date +%s)_$$migration_name; \
	touch internal/migrations/$$file_name.up.sql; \
	touch internal/migrations/$$file_name.down.sql

