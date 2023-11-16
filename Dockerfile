# Pull Geth into a second stage deploy alpine container
FROM alpine:latest

WORKDIR /app

RUN apk add --no-cache ca-certificates

ADD application .
COPY bin/docker-start.sh .

ENV GIN_MODE release
EXPOSE 80

CMD ["sh", "./docker-start.sh"]