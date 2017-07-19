package main

import (
	"fmt"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"github.com/jnewmano/mqtt-nsq/mqtt"
	"github.com/jnewmano/mqtt-nsq/nsqexporter"
	"github.com/nsqio/go-nsq"
)

type nsqProducer struct {
	producer *nsq.Producer

	topic string
	wrap  bool
}

func newNSQProducer(addr string, topic string, wrap bool) (*nsqProducer, error) {
	config := nsq.NewConfig()

	np, err := nsq.NewProducer(addr, config)
	if err != nil {
		return nil, err
	}

	// TODO: intercept log messages
	np.SetLogger(nil, nsq.LogLevelDebug)

	p := nsqProducer{
		producer: np,
		topic:    topic,
		wrap:     wrap,
	}

	return &p, nil
}

func (n *nsqProducer) Handle(p *mqtt.Publish) error {

	var body []byte
	if n.wrap {
		t, err := ptypes.TimestampProto(time.Now())
		if err != nil {
			return err
		}

		m := nsqexporter.MQTTMessage{
			Timestamp:     t,
			Topic:         p.Topic,
			Payload:       p.Payload,
			SourceAddress: "", // we don't know where the message originated
			PacketID:      uint32(p.PacketID),
		}

		body, err = proto.Marshal(&m)
		if err != nil {
			return err
		}

	} else {
		body = p.Payload
	}

	err := n.producer.Publish(n.topic, body)
	if err != nil {
		fmt.Printf("unable to publish message to NSQ [%s]\n", err)
		return err
	}

	fmt.Printf(".")

	return nil
}
