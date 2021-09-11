FROM golang:1.17-alpine3.14 AS builder
WORKDIR /app
COPY ./packages ./packages
COPY ./services/romService ./services/romService
COPY ./go.mod ./go.mod
COPY ./go.sum ./go.sum
RUN go build -ldflags="-s -w" -o romService ./services/romService/romService.go
RUN apk add upx
RUN upx --best --lzma /app/romService

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/romService .

CMD  [ "/app/romService" ]