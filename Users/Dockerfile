FROM golang:1.19.2-alpine
WORKDIR /app
COPY . .
RUN go get -d github.com/gorilla/mux
RUN go build -v .
CMD ["go", "run", "main.go"]