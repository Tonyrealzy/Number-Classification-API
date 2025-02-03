# Use the official Golang image to build the application
FROM golang:1.23.5-alpine as builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the Go Modules manifests to the container
COPY go.mod go.sum ./

# Download dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go app
RUN go build -o main .

# Start a new stage from scratch
FROM alpine:latest  

# Install necessary dependencies (in case your application needs any)
RUN apk --no-cache add ca-certificates

# Set the Current Working Directory inside the container
WORKDIR /root/

# Copy the binary from the previous stage
COPY --from=builder /app/main .

# Expose the port your Go application will run on
EXPOSE 8080

# Command to run the executable
CMD ["./main"]
