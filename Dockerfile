FROM golang:alpine

COPY . /go/src/github.com/lordralex/mightyena

RUN \
    apk add --no-cache git py-pip && \
    go get github.com/lordralex/mightyena && \
    go install github.com/lordralex/mightyena && \
    rm -rf /go/src && \
    apk add --no-cache py-pip && \
    pip install mcstatus && \
    mcstatus -h && \
    apk del git

CMD ["mightyena"]
