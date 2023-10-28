run:
	go run ./cmd/web

sqlc:
	docker run --rm -v $(CURDIR):/src -w /src sqlc/sqlc generate

test:
	go test -v -cover ./internal/db/sqlc

.PHONY: run sqlc