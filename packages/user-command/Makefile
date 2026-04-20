DB_URL=postgres://postgres:postgress@127.0.0.1:5433/shopify_users_service?sslmode=disable
sqlc:
	sqlc generate

migration:
	migrate create -ext sql -dir db/migrations -seq $(name)

migrate-up:
	migrate -path db/migrations -database "$(DB_URL)" up

migrate-drop:
	migrate -path db/migrations -database "$(DB_URL)" drop

proto:
	@echo "--- Generating Proto  ---"
	buf generate
	@echo "--- Done! Check your /contact folder ---"

run:
	air

dev: migrate-up run