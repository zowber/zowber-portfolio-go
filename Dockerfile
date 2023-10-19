FROM golang:1.21-alpine

WORKDIR /app

COPY go.mod . 
COPY cmd/ .
COPY internal/ .
COPY templates/ .

RUN go mod download

RUN go build -o bin/main ./cmd/main/main.go

ENTRYPOINT ["./bin/main"]