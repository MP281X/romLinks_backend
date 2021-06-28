FROM golang:1.16.5-alpine3.14 AS builder
WORKDIR /app
COPY . .
RUN go build -o deviceService ./services/deviceService/deviceService.go

FROM alpine:3.13.5
WORKDIR /app
COPY --from=builder /app/deviceService .

ENV jwtKey=""
ENV mongoUri=""
ENV logFile="false"

EXPOSE 9090
CMD [ "/app/deviceService" ]