FROM golang:1.12-alpine3.10

ADD app /app

WORKDIR /app

RUN go build -o /usr/local/bin/main .

CMD ["/usr/local/bin/main"]
