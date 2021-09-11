FROM golang:1.17-alpine3.14 AS builder
WORKDIR /app
COPY ./packages ./packages
COPY ./services/fileStorageService ./services/fileStorageService
COPY ./go.mod ./go.mod
COPY ./go.sum ./go.sum
RUN go build -ldflags="-s -w" -o fileStorageService ./services/fileStorageService/fileStorageService.go
RUN apk add upx
RUN upx --best --lzma /app/fileStorageService

FROM alpine:latest
RUN apk add optipng
RUN apk add jpegoptim
WORKDIR /app
COPY --from=builder /app/fileStorageService .

CMD [ "/app/fileStorageService" ]