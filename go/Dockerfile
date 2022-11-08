# syntax=docker/dockerfile:1

FROM golang:1.19.1-alpine
WORKDIR /golang
COPY . .
RUN go mod download
RUN go build main.go
CMD ["go", "run", "main.go"]