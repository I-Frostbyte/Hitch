-include .env

install-tools:
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $$(shell go env GOPATH)/bin v1.64.8
	go install github.com/sqlc-dev/sqlc/cmd/sqlc@1.26.0
	go install go.uber.org/mock/mockgen@v0.6.0