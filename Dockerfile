FROM scratch

ADD apps/mqtt-to-nsq/mqtt-to-nsq mqtt-to-nsq
ADD apps/nsq-to-mqtt/nsq-to-mqtt nsq-to-mqtt

VOLUME /etc/ssl/certs
