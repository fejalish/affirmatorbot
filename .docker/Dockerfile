FROM 1.12-alpine3.10

ADD app /app

WORKDIR /app

RUN go build -o main .

CMD ["/app/main"]
