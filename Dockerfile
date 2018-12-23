
FROM registry.docker-cn.com/library/golang:latest

MAINTAINER ServiceComputingGroup "1485578188@qq.com"

RUN mkdir -p  /go/src/github.com/ServiceComputingGroup/simpleWebServer
WORKDIR  /go/src/github.com/ServiceComputingGroup/simpleWebServer
COPY . .
RUN go get
RUN go install

CMD ["simpleWebServer"]
