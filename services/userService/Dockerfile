FROM golang:1.17-alpine3.14 AS builder
WORKDIR /app
COPY ./packages ./packages
COPY ./services/userService ./services/userService
COPY ./go.mod ./go.mod
COPY ./go.sum ./go.sum
RUN go build -ldflags="-s -w" -o userService ./services/userService/userService.go
RUN apk add upx
RUN upx --best --lzma /app/userService

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/userService .

CMD [ "/app/userService" ]