FROM golang:1.20-alpine AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o auth-service cmd/api/main.go

FROM alpine:latest
RUN apk add --no-cache ca-certificates

WORKDIR /root/
COPY --from=builder /app/auth-service .

EXPOSE 8000

ENTRYPOINT ["./auth-service"]
