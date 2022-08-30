# 基础镜像
# 43.140.196.111
FROM golang:latest
# 配置模块代理
ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn,direct
# 创建目录
RUN mkdir "/app"

WORKDIR /app

ADD . /app
EXPOSE 8000

RUN go build -o main ./server.go

CMD /app/main