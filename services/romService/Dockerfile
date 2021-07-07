FROM golang:1.16.5-alpine3.14 AS builder
WORKDIR /app
COPY . .
RUN go build -o romService ./services/romService/romService.go

FROM alpine:3.13.5
WORKDIR /app
COPY --from=builder /app/romService .

EXPOSE 9092
CMD [ "/app/romService" ]