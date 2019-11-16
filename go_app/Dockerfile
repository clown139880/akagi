# Multi-stage builds require Docker 17.05 or higher on the daemon and client.
# see: https://docs.docker.com/engine/userguide/eng-image/multistage-build/

# build the go app binary
FROM golang:latest
WORKDIR /root/
COPY . /root/
ENV SHELL /bin/bash
RUN perl -pi -e "s/tcp\(.*?:/tcp\(db:/; s/host=\S+? /host=db /" models/db.go
ENV GO111MODULE on
ENV GOPROXY https://goproxy.io
ENV GIT_SSL_NO_VERIFY 1
RUN go mod download
RUN go get github.com/oxequa/realize