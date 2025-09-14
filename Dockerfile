# Multi-stage Dockerfile for Go Chat Server
# Stage 1: Build stage
FROM golang:1.21-alpine AS builder

# Install git and ca-certificates (needed for fetching dependencies and HTTPS)
RUN apk update && apk add --no-cache git ca-certificates tzdata && update-ca-certificates

# Create appuser for security
RUN adduser -D -g '' appuser

# Set working directory
WORKDIR /build

# Copy go mod files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download
RUN go mod verify

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
    -ldflags='-w -s -extldflags "-static"' \
    -a -installsuffix cgo \
    -o app ./cmd/server

# Stage 2: Final stage
FROM scratch

# Import from builder
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /etc/passwd /etc/passwd

# Copy the binary
COPY --from=builder /build/app /app

# Copy web assets
COPY --from=builder /build/web /web

# Use unprivileged user
USER appuser

# Expose port
EXPOSE 8080

# Health check
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
    CMD ["/app", "--health-check"] || exit 1

# Run the application
ENTRYPOINT ["/app"]