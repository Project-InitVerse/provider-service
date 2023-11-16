# Pull Geth into a second stage deploy alpine container
FROM alpine:latest

WORKDIR /app

RUN apk add --no-cache ca-certificates

ADD application .

ENTRYPOINT ["/app/application"]