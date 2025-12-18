# --- Stage 1: Builder ---
# We use the specific version matching your local environment
FROM golang:1.25-alpine AS builder

WORKDIR /app

# Copy dependency files first (better caching)
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build the binary
# CGO_ENABLED=0 creates a static binary (no external C libraries needed)
RUN CGO_ENABLED=0 GOOS=linux go build -o main cmd/server/main.go
# --- Stage 2: Runner ---
# We use a tiny, secure Alpine Linux image for production
FROM alpine:latest

# Install certificates so your app can make HTTPS calls if needed
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy only the compiled binary from Stage 1
COPY --from=builder /app/main .

# Expose the port (informative only, doesn't actually open it)
EXPOSE 8080

# Run the app
CMD ["./main"]
