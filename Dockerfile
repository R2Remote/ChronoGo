# Build Stage
FROM golang:1.25-alpine AS builder

# Set environment variables
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Set working directory
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN go build -o chronogo-master ./cmd/master/main.go

# Run Stage
FROM alpine:latest

# Install necessary packages
RUN apk --no-cache add ca-certificates tzdata

# Set working directory
WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/chronogo-master .

# Copy configuration files (if available in root or specific path)
# COPY --from=builder /app/config.yaml .

# Expose the application port
EXPOSE 8080

# Command to run the application
CMD ["./chronogo-master"]
