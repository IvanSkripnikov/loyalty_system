FROM golang:1.20.10-alpine3.17 AS builder

RUN apk add --update  && \
    apk add --no-cache alpine-conf tzdata git

ADD ./src /go/src/loyalty_system
ADD ./src/config /go/config

RUN cd /go/src/loyalty_system && \
    go install loyalty_system

FROM alpine:3.18.4 AS app

COPY --from=builder /go/bin/* /go/bin/loyalty_system
COPY --from=builder /go/config /go/config

ENV CONTAINER_NAME=loyalty-system
ENV HAS_WRITE_LOG_TO_FILE=false
ENV LOG_LEVEL=5

EXPOSE 8080

WORKDIR "/go"
ENTRYPOINT ["/go/bin/loyalty_system"]