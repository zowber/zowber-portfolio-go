FROM golang:1.21-alpine

WORKDIR /app

COPY go.mod . 
COPY cmd/ cmd/
COPY internal/ cmd/
COPY pkg/ pkg/
COPY templates/ templates/

RUN go mod download

RUN go build -o bin/main ./cmd/main/main.go

ENTRYPOINT ["./bin/main"]