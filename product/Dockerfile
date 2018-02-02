FROM golang:1.8.3-alpine

WORKDIR /go/src/github.com/wuriyanto48/ecommerce-grpc-microservice/product

# Add the source code
ENV SRC_DIR=/go/src/github.com/wuriyanto48/ecommerce-grpc-microservice/product/

ENV BUILD_PACKAGES="git curl"

ADD . $SRC_DIR

RUN apk update && apk add --no-cache $BUILD_PACKAGES \
  && curl https://glide.sh/get | sh \
  && glide install \
  && apk del $BUILD_PACKAGES

RUN cd $SRC_DIR; CGO_ENABLED=0 GOOS=linux go build -ldflags '-w -s' -a -o main

ENTRYPOINT ["./main"]
