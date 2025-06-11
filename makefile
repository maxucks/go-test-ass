.PHONY: new_migration run dev
new_migration:
	goose -s create -dir migrations $(MIGRATION_NAME) sql

run:
	docker compose --env-file .env up --build

dev:
	./run.bash