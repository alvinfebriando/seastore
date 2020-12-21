FROM golang:alpine

WORKDIR /app/

COPY go.mod .
COPY go.sum .

RUN go mod download

RUN go get github.com/githubnemo/CompileDaemon

COPY . .

ENTRYPOINT CompileDaemon --build="go build -o bin/ ./cmd/seastore.go" --command=./bin/seastore