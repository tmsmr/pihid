FROM alpine:3.15

RUN apk update && apk add --no-cache imagemagick python3 py3-pip
RUN pip install Wand

ADD gen_icon_src.py /gen_icon_src.py

ENTRYPOINT ["/usr/bin/python3"]
