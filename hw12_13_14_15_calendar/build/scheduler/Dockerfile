FROM golang:1.22-alpine AS builder

WORKDIR /usr/local/src

# Dependencies
COPY ["go.mod", "go.sum", "./"]
RUN go mod download

# Build
COPY . ./
RUN go build -o ./bin/scheduler cmd/scheduler/main.go
#
FROM alpine:3.9 AS runner
#
COPY --from=builder /usr/local/src/bin/scheduler /

COPY configs/scheduler_config.yaml /scheduler_config.yaml

CMD ["/scheduler"]
