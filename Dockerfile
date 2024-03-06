# FROM golang:1.22

# WORKDIR /app

# COPY go.mod go.sum ./

# RUN go mod download && go mod verify

# COPY . .

# RUN CGO_ENABLED=0 go build -o broker ./


# EXPOSE 8000

# CMD ["./broker"]

# Use a lightweight base image
FROM golang:1.22-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files
COPY go.mod go.sum ./

# Download and verify Go module dependencies
RUN go mod download && go mod verify

# Copy the rest of the application source code
COPY . .

# Build the Go binary with CGO disabled
RUN CGO_ENABLED=0 go build -o /app/broker

# Start a new stage with a smaller base image
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the binary from the builder stage to the final stage
COPY --from=builder /app/broker .

# Expose port 8000
EXPOSE 8000

# Command to run the binary
CMD ["./broker"]
