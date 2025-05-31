# Stage 1: Build binary
FROM golang:1.23-alpine AS builder

WORKDIR /app

# Copy only go.mod/go.sum to dependency cache
COPY go.mod go.sum ./
RUN go mod download

# Copy all code
COPY . .

# Compile static and optimized binary
RUN CGO_ENABLED=0 GOOS=linux go build -trimpath -o auth-service ./cmd/auth

# Stage 2: Minimal runtime# CA bundle for HTTPS
FROM scratch

# CA bundle for HTTPS
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

# Compiled binary
COPY --from=builder /app/auth-service /usr/local/bin/auth-service

# Listening port (default 9000, but you can override via ENV)
EXPOSE 9000

ENTRYPOINT ["/usr/local/bin/auth-service"]
