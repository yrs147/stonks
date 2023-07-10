# Start from the official Golang image
FROM golang:1.20.3-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the source code into the container
COPY . .

# Download the Go modules
RUN go mod download

# Build the Go application
RUN go build -o consumer

# Expose Prometheus metrics port
EXPOSE 9000

# Run the Go application
CMD ["./consumer"]
