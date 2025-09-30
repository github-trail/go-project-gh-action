FROM golang:1.24.3 AS builder

# The first WORKDIR /app prepares the /app folder and copies your project there.
WORKDIR /app

# Copy the whole project (including go.mod/go.sum)
COPY . .

# The second WORKDIR /app/src moves into the source code directory to run Go build commands.
WORKDIR /app/src

RUN go mod tidy
RUN go build -o /app/main .

FROM gcr.io/distroless/base

COPY --from=builder /app/main /
COPY --from=builder /app/.env* /

CMD ["/main"]