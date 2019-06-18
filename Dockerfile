FROM golang:alpine

COPY mcping.py .

RUN \
    apk add --no-cache git py-pip python && \
    go get github.com/lordralex/mightyena && \
    go install github.com/lordralex/mightyena && \
    rm -rf /go/src && \
    apk add --no-cache py-pip && \
    pip install mcstatus && \
    apk del git

CMD ["mightyena"]
