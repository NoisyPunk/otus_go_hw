BIN := "./bin"
DOCKER_IMG="calendar:develop"

GIT_HASH := $(shell git log --format="%h" -n 1)
LDFLAGS := -X main.release="develop" -X main.buildDate=$(shell date -u +%Y-%m-%dT%H:%M:%S) -X main.gitHash=$(GIT_HASH)

build-calendar:
	go build -v -o $(BIN) -ldflags "$(LDFLAGS)" ./cmd/calendar
build-scheduler:
	go build -v -o $(BIN) -ldflags "$(LDFLAGS)" ./cmd/scheduler
build-sender:
	go build -v -o $(BIN) -ldflags "$(LDFLAGS)" ./cmd/sender

run-calendar: build-calendar
	$(BIN)/calendar -config ./configs/calendar_config.yaml
run-scheduler: build-scheduler
	$(BIN)/scheduler -config ./configs/scheduler_config.yaml
run-sender: build-sender
	$(BIN)/sender -config ./configs/sender_config.yaml

build-img:
	docker build \
		--build-arg=LDFLAGS="$(LDFLAGS)" \
		-t $(DOCKER_IMG) \
		-f build/Dockerfile .

run-img: build-img
	docker run $(DOCKER_IMG)

version: build
	$(BIN) version

generate:
	go generate ./...

test:
	docker-compose up -d \
	&& go install 'github.com/pressly/goose/cmd/goose@latest' \
	&& goose -dir migrations postgres "host=127.0.0.1 port=5436 user=postgres password=postgres dbname=calendar sslmode=disable" up \
	&& go test -race ./internal/... \
	&& docker-compose down -v

install-lint-deps:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

lint: install-lint-deps
	golangci-lint run ./...

.PHONY: build-calendar build-scheduler build-sender run-calendar run-scheduler run-sender build-img run-img version test lint generate
