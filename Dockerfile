FROM golang:1.21-alpine

WORKDIR /usr/local/go/src/zowber-portfolio-go

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY cmd/ cmd/
COPY internal/ internal/
COPY pkg/ pkg/
COPY templates/ templates/

RUN go build -o bin/main ./cmd/main/main.go

ENTRYPOINT ["./bin/main"]