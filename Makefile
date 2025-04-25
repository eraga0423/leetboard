upalem:
	docker-compose --profile tools run --rm migrate
	docker-compose up --build

up:
	docker compose up --build
	docker compose --profile tools run --rm migrate
downalem:
	docker-compose down
down:
	docker compose down
deleteimage:
	docker compose down
	docker volume prune
	docker compose down --rmi all
