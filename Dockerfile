FROM golang:1.20.1
ADD iceyePoker /go/src/iceyePoker
WORKDIR /go/src/iceyePoker
RUN go get -d -v iceyePoker
RUN go install iceyePoker
ENTRYPOINT ["/go/bin/iceyePoker"]
