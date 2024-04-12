# Собираем в гошке
FROM golang:1.22-alpine AS builder

WORKDIR /usr/local/src

RUN apk --no-cache add bash make

# Зависимости
COPY ["go.mod", "go.sum", "./"]
RUN go mod download

# Build
COPY . ./
RUN go build -o ./bin/calendar cmd/calendar/main.go
#
FROM alpine:3.9 AS runner
#
COPY --from=builder /usr/local/src/bin/calendar /

COPY configs/calendar_config.yaml /calendar_config.yaml

CMD ["/calendar"]

