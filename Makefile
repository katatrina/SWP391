run:
	go run ./cmd/web

sqlc:
	docker run --rm -v $(CURDIR):/src -w /src sqlc/sqlc generate

.PHONY: run sqlc