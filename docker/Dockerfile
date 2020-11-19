FROM golang:latest
RUN mkdir /go/src/music_app_api

WORKDIR /go/src/music_app_api
ADD . /go/src/music_app_api

RUN go get github.com/gin-gonic/gin
RUN go get github.com/go-sql-driver/mysql
RUN go get github.com/jinzhu/gorm
RUN go get -u bitbucket.org/liamstask/goose/cmd/goose
RUN go get -u github.com/gin-contrib/cors

