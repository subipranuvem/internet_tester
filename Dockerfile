FROM golang:1.23.6 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o ping_monitor

# Final image
FROM scratch

WORKDIR /app/

COPY --from=builder /app/ping_monitor ping_monitor
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

CMD ["/app/ping_monitor"]
