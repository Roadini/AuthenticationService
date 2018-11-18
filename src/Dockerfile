# realize.dockerfile
FROM golang:1.11beta2-alpine3.8

RUN apk --update add git bash build-base netcat-openbsd && \
    rm -rf /var/lib/apt/lists/* && \
    rm /var/cache/apk/*

WORKDIR /app
VOLUME /app

RUN go get github.com/oxequa/realize

COPY . .

RUN go get -v -d ./...

ENTRYPOINT ["/app/test_db_connection.sh"]

CMD ["realize", "start", "--install", "--run", "--nc"]