# workspace (GOPATH) configured at /go
FROM golang:1.13.1 as builder


#
RUN mkdir -p $GOPATH/src/bitbucket.org/alien_soft/courier_service
WORKDIR $GOPATH/src/bitbucket.org/alien_soft/courier_service

# Copy the local package files to the container's workspace.
COPY . ./

# installing depends and build
RUN export CGO_ENABLED=0 && \
    export GOOS=linux && \
    make build && \
    mv ./bin/courier_service /



FROM alpine
COPY --from=builder courier_service .
RUN apk add --no-cache tzdata
ENV TZ Asia/Tashkent
ENTRYPOINT ["/courier_service"]



