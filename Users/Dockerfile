FROM golang:1.19.2-alpine
ADD . /go/src/users
WORKDIR /go/src/users
RUN go get users
RUN go install
ENTRYPOINT [ "/go/bin/users" ]