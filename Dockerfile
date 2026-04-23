# Build stage: full Go toolchain to compile the binary
FROM golang:1.26-alpine AS builder

WORKDIR /app

# Copy module files first to cache dependencies as a separate layer
COPY go.mod ./
RUN go mod download

# Copy source and build a static binary
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o server ./cmd/server


# Runtime stage: alpine instead of scratch to include TLS certs and tzdata
FROM alpine:latest
# ca-certificates: enables outbound HTTPS calls (absent in alpine by default)
# tzdata: enables timezone lookups via time.LoadLocation (absent in alpine by default)
# --no-cache: skips storing the apk index on disk, keeps the image smaller
RUN apk --no-cache add ca-certificates tzdata
WORKDIR /root/
# Copy the compiled binary from the builder stage
COPY --from=builder /app/server .
# Expose the default port
EXPOSE 8080
# Start the server
CMD [ "./server" ]

# To build: docker build -t status-monitor .
# To run:   docker run -p 8080:8080 status-monitor