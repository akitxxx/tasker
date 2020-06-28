FROM golang:latest as builder
ENV SRC_DIR=/go/src/github.com/lelouch99v/tasker
ENV GOBIN=/go/bin
WORKDIR $GOBIN
COPY . $SRC_DIR
RUN cd /go/src && go get github.com/labstack/echo/... && \
    go install github.com/lelouch99v/tasker/ && \
    apt-get update && apt-get install -y vim

# runtime image
FROM alpine:latest
COPY --from=builder /go/bin/tasker /tasker
EXPOSE 5010
ENTRYPOINT ["/tasker"]
