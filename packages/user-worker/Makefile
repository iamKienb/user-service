.PHONY: es-migrate
migrate-up:
	@echo "🚀 Starting Elasticsearch Migrations..."
	go run cmd/migrate/main.go

run:
	air

dev: migrate-up run