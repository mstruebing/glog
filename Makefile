test:
	@go test -race -coverprofile=coverage.txt -covermode=atomic ./...
	@go tool vet .
	@test -z $(shell gofmt -s -l . | tee /dev/stderr) || (echo "[ERROR] Fix formatting issues with 'gofmt'" && exit 1)
