##STEP1
FROM golang:alpine as builder

RUN mkdir -p /app

WORKDIR /app/

RUN apk update && apk add --no-cache git && apk add --update gcc go git mercurial

ENV GO111MODULE=auto

ADD . /app

RUN go get "github.com/gorilla/mux" && go get "github.com/mattn/go-sqlite3" && go get "github.com/sirupsen/logrus"

RUN go build main.go

EXPOSE 8080

CMD ["./backend" ]

##STEP2
FROM alpine as backend

WORKDIR /app

COPY --from=builder ./app/main .

RUN adduser -S -D -H -h /app appuser && chown -R appuser: /app && apk add curl 

USER appuser

CMD ["./main" ]
