# Use the official Golang image as the base image
FROM golang:1.20.3-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the source code into the container
COPY . .

# Download the Go modules
RUN go mod download

# Build the Go application
RUN go build -o producer1

# Set the entry point for the container
CMD ["./producer1"]
