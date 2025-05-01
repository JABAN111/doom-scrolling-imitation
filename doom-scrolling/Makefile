lint: golint

golint:
	golangci-lint run -E gocritic -v ./...

tools:
	go install golang.org/x/tools/cmd/goimports@latest
	go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.64.2
	go mod tidy