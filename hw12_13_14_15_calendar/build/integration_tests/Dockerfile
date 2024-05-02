FROM golang:1.22-alpine

WORKDIR /usr/local/src

# Dependencies
COPY ["go.mod", "go.sum", "./"]
RUN go mod download

# Build
COPY . ./

CMD go test ./tests/...

