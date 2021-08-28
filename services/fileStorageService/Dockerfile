FROM golang:1.17-alpine3.14 AS builder
WORKDIR /app
COPY . .
RUN go build -ldflags="-s -w" -o fileStorageService ./services/fileStorageService/fileStorageService.go
RUN apk add upx
RUN upx /app/fileStorageService

FROM alpine:latest
RUN apk add optipng
RUN apk add jpegoptim
WORKDIR /app
COPY --from=builder /app/fileStorageService .

CMD [ "/app/fileStorageService" ]