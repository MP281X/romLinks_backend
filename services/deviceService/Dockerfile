FROM golang:1.17-alpine3.14 AS builder
WORKDIR /app
COPY . .
RUN go build -ldflags="-s -w" -o deviceService ./services/deviceService/deviceService.go
RUN apk add upx
RUN upx /app/deviceService

FROM alpine:3.13.5
WORKDIR /app
COPY --from=builder /app/deviceService .

EXPOSE 9090
CMD [ "/app/deviceService" ]