# Build stage
FROM golang:1.22 AS builder

WORKDIR /app/auth-service

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN CGO_ENABLED=0 go build -o /app/authservice ./cmd/api

# Final stage
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/authservice .

EXPOSE 8082

CMD ["./authservice"]
