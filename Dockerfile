FROM golang:1.19-alpine AS builder

RUN apk update && apk add --no-cache git && apk add make

COPY . /src

RUN export GOPATH= && \
    cd /src && \
    go mod tidy && \
    go build -o /dist/main cmd/main.go

ADD script /dist/script

FROM alpine:latest

RUN apk update && apk add --no-cache
RUN apk add postgresql-client

COPY --from=builder /dist /dist
COPY ./bin ./dist

RUN chmod +x ./dist/wait-for-it

WORKDIR /dist

ENV PORT=3000

CMD until pg_isready --host=composepostgres; do sleep 1; done \
    && "/dist/main"

EXPOSE 3000