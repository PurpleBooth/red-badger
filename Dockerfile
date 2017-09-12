FROM golang:1

RUN mkdir -p /go/src/github.com/purplebooth/red-badger
WORKDIR /go/src/github.com/purplebooth/red-badger
COPY . .

RUN go get github.com/golang/dep/cmd/dep \
           github.com/onsi/ginkgo/ginkgo \
           github.com/onsi/gomega
RUN dep ensure
WORKDIR /go/src/github.com/purplebooth/red-badger
RUN go test
RUN go install
WORKDIR /go/src/github.com/purplebooth/red-badger/red-badger-robot
RUN go test
RUN go install

CMD ["go-wrapper", "run"]
