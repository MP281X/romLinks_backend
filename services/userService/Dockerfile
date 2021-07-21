FROM golang:1.16.5-alpine3.14 AS builder
WORKDIR /app
COPY . .
RUN go build -ldflags="-s -w" -o userService ./services/userService/userService.go
RUN apk add upx
RUN upx /app/userService

FROM alpine:3.13.5
WORKDIR /app
COPY --from=builder /app/userService .

EXPOSE 9093
CMD [ "/app/userService" ]