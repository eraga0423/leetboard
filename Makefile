.PHONY: up down clear logs restart

up:
	docker-compose --profile tools run --rm migrate
	docker-compose up --build

down:
	docker-compose down

clear:
	docker-compose down --rmi all -v
	docker volume prune -f

logs:
	docker-compose logs -f

restart:
	make down && make up
