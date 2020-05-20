FROM alpine:latest
MAINTAINER "lhlyu"
WORKDIR /app
RUN mkdir conf && mkdir log
COPY main ./main
COPY conf/config.yaml conf/config.yaml
COPY static static
RUN chmod 777 -R /app
CMD ["./main"]

# 创建容器运行
# docker build -t nrand .
# docker run -itd -p 9003:8080 -v /logs/nrand-log:/app/log --name gen-name nrand
