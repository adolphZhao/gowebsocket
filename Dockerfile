FROM alpine:3.5

MAINTAINER zhaomx zhaomingxing@medatc.com

RUN echo "build ws image "

RUN mkdir -p /usr/local/bin/

ADD ws /usr/local/bin/

ADD config.ini /usr/local/bin/
ADD server.crt /usr/local/bin/
ADD server.key /usr/local/bin/

EXPOSE 8080

VOLUME /usr/local/bin/

WORKDIR /usr/local/bin/

CMD ["ws"]

