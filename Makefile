export GOTOOLCHAIN := $(filter go1.%,$(file <go.mod))

lint:
	go run github.com/golangci/golangci-lint/cmd/golangci-lint@v1.64.8 run --timeout 5m --verbose -c .golangci.yaml

test:
	go test -v -coverpkg ./... -covermode=atomic -coverprofile cover.out -timeout 5m ./...
	go tool cover -func cover.out | grep total && unlink cover.out
