# Build stage
FROM golang:1.23.4-alpine AS builder 
WORKDIR /app
COPY . .
RUN go build -o main main.go
RUN apk add --no-cache curl
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.18.1/migrate.linux-amd64.tar.gz | tar xvz

# Run stage
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/main .
COPY --from=builder /app/migrate ./migrate
COPY app.env .
COPY start.sh .
RUN chmod +x /app/start.sh
COPY wait-for.sh .
COPY db/migration ./migration

EXPOSE 8080
# 当 CMD 和 ENTRYPOINT 同时存在时，CMD 会作为 ENTRYPOINT 的参数
CMD [ "/app/main" ]
ENTRYPOINT [ "/app/start.sh" ]