FROM golang:alpine3.19  AS builder
WORKDIR /app

COPY . .

RUN go build -o main main.go

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/main .
COPY --from=builder /app/app.env .
RUN apk --no-cache add ca-certificates


CMD ["/app/main"]