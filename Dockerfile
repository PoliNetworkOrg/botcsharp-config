# Dockerfile
FROM golang:alpine AS builder

WORKDIR /app
COPY main.go .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/main.go

FROM alpine:3.16  
WORKDIR /app
COPY --from=builder /app/main .
CMD ["./main"]
