FROM golang:1.22.5 AS builder
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o myapp ./cmd/main.go

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/myapp .

# Install CA certificates (essential for Kafka SSL connection)
RUN apk update && apk add ca-certificates

COPY .env .

# Make sure the CA certificates are in the trusted store
ENV SSL_CERT_FILE=/etc/ssl/certs/ca-certificates.crt

EXPOSE 8082
CMD ["./myapp"]