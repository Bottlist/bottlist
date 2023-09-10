FROM golang:1.19.1-alpine

WORKDIR /app

COPY app .
RUN  go mod download


CMD ["go", "run", "cmd/main.go"]