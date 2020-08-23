FROM golang:1.14

WORKDIR /app

COPY . .

RUN go get -v ./...

RUN go get github.com/cespare/reflex

CMD ["reflex", "-c", "/app/reflex.conf"]
