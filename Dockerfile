FROM golang

ADD . /go/src/github.com/fajardm/inventories

RUN go install github.com/fajardm/inventories

ENTRYPOINT /go/bin/inventories

EXPOSE 3000