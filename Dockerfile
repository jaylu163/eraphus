FROM alpine:latest

# 设置时区为上海
RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
RUN echo 'Asia/Shanghai' >/etc/timezone
#设置编码
#ENV LANG C.UTF-8
# 安装curl
#RUN apk add curl
# 挂载容器目录
VOLUME ["/app/data"]
#在容器根目录 创建一个 apps 目录

ENV env_conf_suffix=prod
ENV GIN_MODE=release

WORKDIR /app

# 打包平台先把项目编译，然后在执行docker build
COPY  eraphus.server  /app/eraphus.server
COPY  env/config.yaml /app/env/config.yaml.prod
EXPOSE 8888

#workdir目录的路径，登录docker shell里面自动到此目录下
ENTRYPOINT ["/app/eraphus.server"]