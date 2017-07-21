package main

import (
	"context"
	"fmt"

	"github.com/golang/protobuf/proto"
	"github.com/jnewmano/mqtt-nsq/nsqexporter"
	"github.com/nsqio/go-nsq"
)

type nsqConsumer struct {
	consumer  *nsq.Consumer
	publisher func(context.Context, string, []byte) error
	topic     string

	wrap bool

	Stop func()
}

func newNSQConsumer(lookupdAddrs []string, nsqdAddrs []string, topic string, channel string, wrap bool, publish func(context.Context, string, []byte) error, mqttTopic string) (*nsqConsumer, error) {

	if len(lookupdAddrs) > 0 && len(nsqdAddrs) > 0 {
		return nil, fmt.Errorf("only supply nsqd or lookupd addresses, not both")
	}

	config := nsq.NewConfig()
	config.MaxInFlight = 5

	c, err := nsq.NewConsumer(topic, channel, config)
	if err != nil {
		return nil, err
	}

	n := nsqConsumer{
		consumer:  c,
		wrap:      wrap,
		publisher: publish,
	}

	c.AddConcurrentHandlers(&n, 5)

	if len(lookupdAddrs) > 0 {
		err = c.ConnectToNSQLookupds(lookupdAddrs)
	} else {
		err = c.ConnectToNSQDs(nsqdAddrs)
	}
	if err != nil {
		return nil, err
	}

	// TODO: intercept NSQ log messages
	// np.SetLogger(nil, nsq.LogLevelDebug)

	return &n, nil
}

func (n *nsqConsumer) HandleMessage(msg *nsq.Message) error {

	ctx := context.Background()
	var body []byte

	if n.wrap {

		m := nsqexporter.MQTTMessage{}

		err := proto.Unmarshal(msg.Body, &m)
		if err != nil {
			return err
		}

		body = m.Payload

	} else {
		body = msg.Body
	}

	err := n.publisher(ctx, n.topic, body)
	if err != nil {
		fmt.Printf("unable to publish message to MQTT [%s]\n", err)
		return err
	}

	fmt.Printf(".")

	return nil
}
