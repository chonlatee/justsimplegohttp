FROM golang:1.19

WORKDIR /app

COPY go.mod ./

COPY *.go ./

RUN go build -o /go-simple-server

EXPOSE 8000

CMD ["/go-simple-server"]