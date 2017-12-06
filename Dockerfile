FROM golang:1.9.2-alpine3.6
MAINTAINER An Nguyen <an.nguyenhoai@tiki.vn>

RUN apk add --update bash && apk add --no-cache git mercurial && rm -rf /var/cache/apk/*

ENV APP_HOME $GOPATH/src/talaria-recentlyviewed-go
RUN mkdir -p $APP_HOME
WORKDIR $APP_HOME
RUN mkdir -p $APP_HOME/vendor


RUN go get golang.org/x/tools/cmd/cover
RUN go get github.com/mattn/goveralls
RUN go get -u github.com/kardianos/govendor
RUN go get github.com/gorilla/mux
RUN go get github.com/BurntSushi/toml
RUN go get gopkg.in/mgo.v2
RUN go get gopkg.in/mgo.v2/bson
RUN go get github.com/joho/godotenv
RUN go get github.com/newrelic/go-agent
RUN go get -u github.com/kardianos/govendor

ADD . $APP_HOME
RUN govendor sync

RUN cd app/ && go build  
RUN chmod a+x app 
CMD ["./app"]