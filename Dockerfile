FROM golang:1.16-alpine

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o bin/main ./cmd/main/main.go

EXPOSE 8080

CMD ["./bin/main"]
