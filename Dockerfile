FROM scratch

ADD bin/mqtt-to-nsq mqtt-to-nsq
ADD bin/nsq-to-mqtt nsq-to-mqtt

VOLUME /etc/ssl/certs
