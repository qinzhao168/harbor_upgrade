FROM alpine:3.6
RUN echo "https://mirror.tuna.tsinghua.edu.cn/alpine/v3.6/main/" > /etc/apk/repositories && \
    apk update && apk add bash && \
    apk add supervisor openssh  curl  && \
    apk add su-exec  && \
    rm -rf /var/cache/apk/*
COPY harbor_upgrade /harbor_upgrade
#copy config.yaml.template /config.yaml
copy config.yaml /config.yaml
COPY run.sh /run.sh

WORKDIR /
ESPOSE 8090
ENTRYPOINT ["./run.sh"]