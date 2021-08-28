FROM golang:1.17-alpine3.14 AS builder
WORKDIR /app
COPY . .
RUN go build -ldflags="-s -w" -o romService ./services/romService/romService.go
RUN apk add upx
RUN upx /app/romService

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/romService .

CMD [ "/app/romService" ]