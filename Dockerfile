FROM alpine
ADD kuu-locale /bin/
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories && \
    apk update && \
    apk -Uuv add --no-cache ca-certificates tini tzdata && \
    ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
ENTRYPOINT ["/sbin/tini","--", "kuu-locale"]