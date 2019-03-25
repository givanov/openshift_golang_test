FROM golang:1.11

USER nobody

RUN mkdir -p /go/src/github.com/givanov/openshift_golang_test
WORKDIR /go/src/github.com/givanov/openshift_golang_test

COPY . /go/src/github.com/givanov/openshift_golang_test
RUN go build -o openshift_golang_test cmd/main.go

CMD ["./openshift_golang_test"]