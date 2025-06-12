# Stage 1: Build
FROM golang:1.23.4-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/collector/main.go

# Stage 2: Run
FROM alpine:latest AS runner

WORKDIR /root/

COPY --from=builder /app/main .

CMD ["./main"]