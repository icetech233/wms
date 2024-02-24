FROM golang:alpine

ENV GO111MODULE=on
ENV GOPROXY=https://repo.huaweicloud.com/repository/goproxy/,direct

WORKDIR /server
COPY . .

RUN go env
RUN go build -trimpath -ldflags "-s -w -buildid=" .

FROM alpine:latest
LABEL MAINTAINER="dsg@qq.com"

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories

RUN apk add --no-cache gettext tzdata && \
    cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
    echo "Asia/Shanghai" >  /etc/timezone && \
    date && \
    apk del tzdata

WORKDIR /server
COPY --from=0 /server/wms ./

ENTRYPOINT ["./wms"]