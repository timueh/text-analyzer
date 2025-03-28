FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o app ./cmd/api

FROM alpine:latest
RUN adduser -D appuser
WORKDIR /app
COPY --from=builder /app/app .
RUN chown appuser:appuser /app/app
USER appuser
EXPOSE 8080
CMD [ "./app" ]