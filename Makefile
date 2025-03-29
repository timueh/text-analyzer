test:
	@cd zipfback && go test ./... -v --failfast

run:
	@cd zipfback && go run ./cmd/api

up:
	@docker compose up --build