FROM golang:1.16.5-alpine3.14 AS builder
WORKDIR /app
COPY . .
RUN go build -o fileStorageService ./services/fileStorageService/fileStorageService.go

FROM alpine:3.13.5
WORKDIR /app
COPY --from=builder /app/fileStorageService .

EXPOSE 9091
CMD [ "/app/fileStorageService" ]