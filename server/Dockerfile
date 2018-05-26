FROM golang:alpine

RUN apk update && \
    apk add make git && \
    rm -rf /var/cache/apk/*

COPY ./ /go/src/github.com/okashoi/go-websocket-test
WORKDIR /go/src/github.com/okashoi/go-websocket-test

RUN go get -u github.com/golang/dep/cmd/dep

RUN ${GOPATH}/bin/dep ensure -vendor-only

RUN go build -o bin/server src/main.go

CMD ["./bin/server"]
