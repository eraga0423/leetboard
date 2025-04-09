```
docker-compose --profile tools run --rm migrate up
docker compose --profile tools run --rm migrate create -ext sql -dir ./migrations NAME_OF_MIGRATION_FILE
```