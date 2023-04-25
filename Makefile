all: lint compile

compile:
	go build CGO_ENABLED=0 GOOS=linux GOARCH=amd64 ./cmd/main.go

RUN_LOCAL_ENV = \
	PORT=localhost:3041 \
	MYSQL_DATABASE_NAME=testdb \
	MYSQL_USERNAME=root \
	MYSQL_PASSWORD=password \
	MYSQL_HOST=localhost \
	MYSQL_PORT=3040

run_local:
	env $(RUN_LOCAL_ENV) go run ./cmd/main.go run

init_mysql_db:
	env $(RUN_LOCAL_ENV) go run ./cmd/main.go init

lint:
	goimports -w .
	golangci-lint run -E goimports -E gocyclo -E gosec -E revive --timeout 5m

lint-deps:
	go install github.com/securego/gosec/v2/cmd/gosec@latest
	go get -d github.com/fzipp/gocyclo/cmd/gocyclo
	go install github.com/fzipp/gocyclo/cmd/gocyclo
	go install golang.org/x/tools/cmd/goimports
	go get -d github.com/golangci/golangci-lint/cmd/golangci-lint
	go install github.com/golangci/golangci-lint/cmd/golangci-lint
	go mod tidy