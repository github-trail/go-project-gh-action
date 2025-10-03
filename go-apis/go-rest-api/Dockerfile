FROM golang:1.24.3 AS builder

# The first WORKDIR /app prepares the /app folder and copies your project there.
WORKDIR /app

# Copy the whole project (including go.mod/go.sum)
COPY . .

# The second WORKDIR /app/src moves into the source code directory to run Go build commands.
WORKDIR /app/src

RUN go mod tidy
RUN go build -o /app/main .

FROM alpine:latest

# Install ca-certificates for HTTPS requests and basic tools
RUN apk --no-cache add ca-certificates tzdata curl

WORKDIR /root/

COPY --from=builder /app/main /
COPY --from=builder /app/.env* /

# Create log directory
RUN mkdir -p /var/log

# Default exposed port for the Go REST API is 8080
CMD ["/main"]