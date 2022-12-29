FROM golang:1.19.4

WORKDIR /app

RUN go install github.com/cosmtrek/air@latest

COPY go.mod ./

RUN go clean -modcache

RUN go mod download

CMD ["air", "-c", ".air.toml"]