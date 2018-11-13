FROM golang:1.9-alpine

RUN apk add --no-cache --update ca-certificates \
		netcat-openbsd \
        dpkg \
        gcc \
        git \
        musl-dev \
        bash

WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH

COPY ./entrypoint.sh /entrypoint.sh 

RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" \
    && chmod -R 777 "$GOPATH" \
    && chmod +x /entrypoint.sh

RUN go get github.com/tockins/realize
WORKDIR $GOPATH

ENTRYPOINT ["/entrypoint.sh"]
CMD ["realize", "start"]