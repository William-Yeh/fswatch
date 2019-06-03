# fswatch - Watch for changes in file system.
#

#======================================#
# builder env
#

FROM golang:1.12.5 AS builder
LABEL author="william.pjyeh@gmail.com"

ENV EXE_NAME  fswatch
WORKDIR       /opt

# unset GOPATH in favor of GO111MODULE
ENV GO111MODULE on
ENV GOPATH ""

# fetch imported Go lib...
COPY go.mod .
COPY go.sum .
COPY fswatch.go    .
COPY copy-files.sh .


# compile...
RUN  GOOS=windows GOARCH=amd64  CGO_ENABLED=0 \
     go build -o $EXE_NAME-x86_64.exe

RUN  GOOS=linux   GOARCH=amd64  CGO_ENABLED=0 \
     go build -o $EXE_NAME-linux-x86_64

RUN  GOOS=darwin  GOARCH=amd64  CGO_ENABLED=0 \
     go build -o $EXE_NAME-mac


#======================================#
# deployment env
#

FROM alpine:3.9.4
LABEL author="william.pjyeh@gmail.com"

WORKDIR /opt

# copy executable
COPY --from=builder /opt/copy-files.sh   /opt/

COPY --from=builder /opt/fswatch-*   /opt/
RUN ln -s /opt/fswatch-linux-x86_64  /usr/local/bin/fswatch
CMD ["fswatch"]
