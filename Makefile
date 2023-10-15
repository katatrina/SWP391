run:
	go run ./cmd/web

sqlc:
	docker run --rm -v $(CURDIR):/src -w /src sqlc/sqlc generate

test:
	go test -v -cover ./...

.PHONY: run sqlc