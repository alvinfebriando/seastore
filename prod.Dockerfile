FROM golang:alpine AS builder

WORKDIR /app/

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o /app/bin/seastore /app/cmd/seastore

EXPOSE 8080

FROM alpine
COPY --from=builder /app/bin/seastore /app/bin/seastore

CMD [ "app/bin/seastore" ]