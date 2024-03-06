
FROM golang:1.22 AS builder

WORKDIR /app/user-microservice

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN CGO_ENABLED=0 go build -o /app/user ./cmd/api

# Final stage
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/user .

EXPOSE 8081

CMD ["./user"]
