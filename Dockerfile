# Stage 1: Build environment
FROM golang:latest AS builder

WORKDIR /app

COPY go.mod ./
RUN go mod download && go mod verify

COPY . .

RUN go build ./main.go  # Replace ./cmd/main.go with your main Go file path

WORKDIR /app

EXPOSE 8000

CMD ["./main"]