FROM golang:1.23 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . ./
RUN go build -o main ./cmd

FROM debian:bookworm-slim
WORKDIR /app
COPY .env .env
COPY --from=builder /app/main ./main
EXPOSE 9111
CMD ["./main"]