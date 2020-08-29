FROM golang:1.14

WORKDIR /app

COPY . .
COPY atoxicer-project-firebase-adminsdk.json /firebase/atoxicer-project-firebase-adminsdk.json

RUN go get -v ./...

RUN go get github.com/cespare/reflex

CMD ["reflex", "-c", "/app/reflex.conf"]
