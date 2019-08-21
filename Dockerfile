FROM golang:1.12-alpine3.10

ADD app /app

WORKDIR /app

RUN \
  apk add --no-cache bash git openssh && \
  go get -v github.com/shomali11/slacker && \
  go get -v github.com/google/uuid && \
  go get -v github.com/jackc/pgx

RUN go build -o /usr/local/bin/main .

CMD ["/usr/local/bin/main"]
