# Stage 1: Build the Go application
FROM golang:1.20.3-alpine AS builder

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o consumer

# Stage 2: Create a lightweight production image
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/consumer .

# Expose Prometheus metrics port
EXPOSE 9010

# Run the Go application
CMD ["./consumer"]