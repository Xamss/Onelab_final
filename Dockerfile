FROM golang:alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o main cmd/main.go

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/main .
COPY config.yaml .
EXPOSE 8080
CMD ["/app/main"]