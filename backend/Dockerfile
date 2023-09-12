FROM golang:1.19.1-alpine

WORKDIR /app

COPY app .
RUN  go mod download
RUN go install -tags mysql github.com/golang-migrate/migrate/v4/cmd/migrate@latest


CMD ["go", "run", "cmd/main.go"]