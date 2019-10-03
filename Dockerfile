###
# Builder to compile our golang code
###
FROM golang:alpine AS builder

WORKDIR /build
COPY . .

RUN go build
RUN go install

RUN which mightyena


###
# Now generate our smaller image
###
FROM alpine

COPY --from=builder /go/bin/mightyena /go/bin/mightyena

WORKDIR /app
COPY mcping.py .

RUN \
    apk add --no-cache py-pip python && \
    apk add --no-cache py-pip && \
    pip install mcstatus && \
    apk del py-pip

CMD ["/go/bin/mightyena"]
