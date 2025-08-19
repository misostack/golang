include .env.local

.PHONY: hello
hello:
	@echo "Hello, World!"

.PHONY: watch_dev
watch_dev:
	@air

.PHONY: db_migrate
db_migrate:
	@echo "Migrating database..."


# @goose create $@ sql
.PHONY: db_migrate
db_migrate:
	@goose up