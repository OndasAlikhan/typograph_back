FROM golang:1.21-bookworm as builder

WORKDIR /app

# RUN go install github.com/cosmtrek/air@1.49
# RUN go install github.com/swaggo/swag/cmd/swag@1.16

COPY . .
COPY .env .
RUN go mod download

# Build the binary.
RUN go build -v -o server

# CMD ["ls -a"]
# Use the official Debian slim image for a lean production container.
# https://hub.docker.com/_/debian
# https://docs.docker.com/develop/develop-images/multistage-build/#use-multi-stage-builds
FROM debian:bookworm-slim
RUN set -x && apt-get update && DEBIAN_FRONTEND=noninteractive apt-get install -y \
    ca-certificates && \
    rm -rf /var/lib/apt/lists/*

# Copy the binary to the production image from the builder stage.
COPY --from=builder /app/server /app/server
COPY --from=builder /app/.env /app/.env

# Run the web service on container startup.
CMD ["/app/server"]