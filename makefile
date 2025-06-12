.PHONY: new_migration run run_app run_collector
new_migration:
	goose -s create -dir migrations $(MIGRATION_NAME) sql

run:
	docker compose --env-file .env up --build

run_app:
	./run_app.bash

run_collector:
	./run_collector.bash