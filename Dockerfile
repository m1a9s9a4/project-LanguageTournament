FROM golang:alpine

ENV GOPATH=''

WORKDIR /app

# EXPOSE 8000

WORKDIR /go
RUN apk add git vim openssh

RUN go get -u github.com/labstack/echo
RUN go get -u github.com/dgrijalva/jwt-go
RUN go get -u github.com/jinzhu/gorm
RUN go get -u github.com/go-sql-driver/mysql
RUN go get golang.org/x/tools/cmd/goimports
RUN go get gopkg.in/urfave/cli.v2
RUN go get github.com/oxequa/realize

RUN mkdir backend/