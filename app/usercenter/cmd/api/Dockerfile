# 表示依赖 alpine 最新版
FROM alpine:latest
ENV VERSION 1.0

# 在容器根目录 创建一个 apps 目录
WORKDIR /apps

# 挂载容器目录
#VOLUME ["/apps/conf"]

# 拷贝当前目录下 go_docker_demo1 可以执行文件
COPY test-zero /apps/test-zero

# 拷贝配置文件到容器中
COPY etc/usercenter.yaml /apps/etc/usercenter.yaml

# 设置时区为上海
RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
RUN echo 'Asia/Shanghai' >/etc/timezone

# 设置编码
ENV LANG C.UTF-8

# 暴露端口
EXPOSE 8089

# 运行golang程序的命令
CMD ["/apps/test-zero", "-f", "/apps/etc/usercenter.yaml"]