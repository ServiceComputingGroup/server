
FROM registry.docker-cn.com/library/golang:latest

MAINTAINER ServiceComputingGroup "1485578188@qq.com"
RUN mkdir -p  /go/src/github.com
WORKDIR  /go/src/github.com
RUN go get github.com/ServiceComputingGroup/simpleWebServer
WORKDIR  /go/src/github.com/ServiceComputingGroup/simpleWebServer
RUN go get
RUN go install
#暴露端口
EXPOSE 9090
#最终运行docker的命令
CMD ["simpleWebServer"]
