lint: golint

golint:
	golangci-lint run -E goimports,gocritic -v ./...

tools:
	go get github.com/couchbase/gocb/v2@latest

run: 
	go run cmd/app/main.go

up-down:
	docker compose down -v && docker compose up -d