FROM golang:1.17-alpine3.14 AS builder
WORKDIR /app
COPY ./packages ./packages
COPY ./services/deviceService ./services/deviceService
COPY ./go.mod ./go.mod
COPY ./go.sum ./go.sum
RUN go build -ldflags="-s -w" -o deviceService ./services/deviceService/deviceService.go
RUN apk add upx
RUN upx /app/deviceService

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/deviceService .

CMD [ "/app/deviceService" ]