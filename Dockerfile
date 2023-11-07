FROM golang:1-19-alpine AS builder

RUN apk update && apk add --no-cache git && apk add amke

WORKDIR /app

COPY . .

RUN make build

FROM alpine:latest

RUN apk update && apk add --no-cache

WORKDIR /app

COPY --from=builder /app/bin/ .

ENV PORT=3000

CMD ["/app/bin"]

EXPOSE 3000