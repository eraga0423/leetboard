up:
	docker-compose --profile tools run --rm migrate
	docker-compose up --build
down:
	docker-compose down
clear:
	docker-compose down
	docker volume prune
	docker-compose down --rmi all
