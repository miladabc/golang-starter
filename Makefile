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

