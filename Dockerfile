FROM golang:1.21.0-bookworm as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Use a lightweight image (alpine) for the final image
FROM alpine:latest
WORKDIR /root/

# Copy the binary from the builder image
COPY --from=builder /app/main .
COPY --from=builder /app/config.json .

CMD ["./main"]
