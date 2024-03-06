# Build stage
FROM golang:1.22 AS builder

WORKDIR /app/task-service

COPY go.mod go.sum ./

RUN go mod download && go mod verify

COPY . .

RUN CGO_ENABLED=0 go build -o /app/task ./cmd/api


FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/task .

EXPOSE 8083

CMD ["./task"]
