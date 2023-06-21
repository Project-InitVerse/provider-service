# Pull Geth into a second stage deploy alpine container
FROM alpine:latest

RUN apk add --no-cache ca-certificates

COPY application /usr/local/bin/

ENTRYPOINT ["application"]