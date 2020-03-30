FROM golang:alpine
RUN apk add --no-cache bash
WORKDIR /build
ENTRYPOINT ["bash", "./build.sh"]
VOLUME /build

