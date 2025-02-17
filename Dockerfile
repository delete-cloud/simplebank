# Build stage
FROM golang:1.23.4-alpine AS builder 
WORKDIR /app
COPY . .
RUN go build -o main main.go

# Run stage
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/main .
COPY app.env .
COPY start.sh .
COPY wait-for.sh .
RUN chmod +x /app/start.sh /app/wait-for.sh
COPY db/migration ./db/migration

EXPOSE 8080
# 当 CMD 和 ENTRYPOINT 同时存在时，CMD 会作为 ENTRYPOINT 的参数
CMD [ "/app/main" ]
ENTRYPOINT [ "/app/start.sh" ]