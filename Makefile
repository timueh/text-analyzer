test:
	@cd backend && go test ./... -v --failfast

run:
	@cd backend && go run ./cmd/api

up:
	@docker compose up --build