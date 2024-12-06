# Use the official Go image as a base
FROM golang:1.23 AS builder

# Set environment variables
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Create a working directory inside the container
WORKDIR /app

# Copy go mod and sum files first to leverage caching
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the project files
COPY . .

# Build the binaries for your executables
RUN go build -o /bin/api-server ./cmd/api-server/main.go
RUN go build -o /bin/migration ./cmd/migration/main.go

# Use a minimal base image for running the application
FROM alpine:latest
RUN apk --no-cache add ca-certificates

# Copy binaries from builder
COPY --from=builder /bin/api-server /usr/local/bin/api-server
COPY --from=builder /bin/migration /usr/local/bin/migration

# Expose application ports 
EXPOSE 8080

# Set the default executable
CMD ["api-server"]