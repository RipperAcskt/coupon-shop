FROM golang:alpine AS builder

WORKDIR /app

RUN apk --no-cache add ca-certificates

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main ./cmd/server

FROM scratch

WORKDIR /app

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

COPY --from=builder /app/main .
COPY --from=builder /app/config ./config
COPY --from=builder /app/.env.dist ./.env.dist

EXPOSE 3000

# Command to run the application
CMD ["./main"]
