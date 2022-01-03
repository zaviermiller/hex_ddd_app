FROM golang:1.17.5-alpine3.15

COPY . /app

WORKDIR /app

RUN go mod tidy

CMD ["go", "run", "cmd/rest/main.go"]