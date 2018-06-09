FROM golang:alpine

WORKDIR /go/src/github.com/okashoi/go-websocket-test

RUN apk update && \
    apk add make git && \
    rm -rf /var/cache/apk/* && \
    go get -u github.com/golang/dep/cmd/dep

COPY ./ /go/src/github.com/okashoi/go-websocket-test
RUN ${GOPATH}/bin/dep ensure -vendor-only &&\
    go build -o bin/server src/main.go

CMD ["./bin/server"]
