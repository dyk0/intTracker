FROM golang:1.13.6 AS build
MAINTAINER ddyke@dyk0.tech

ADD . /usr/src/
RUN cd /usr/src && make build-linux && make test

FROM alpine
WORKDIR /usr/src/
COPY --from=build /usr/src/bin/intTracker /usr/src/

ENTRYPOINT ["/usr/src/intTracker"]
