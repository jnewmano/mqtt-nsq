package mqttclient

import (
	"context"
	"fmt"

	"github.com/jnewmano/mqtt-nsq/mqtt"
)

type PublishHandler interface {
	Handle(*mqtt.Publish) error
}

// Publish sends a QOS 0 payload to the topic
func (m *MQTTClient) Publish(ctx context.Context, topic string, payload []byte) error {

	p := mqtt.Publish{
		Topic:   topic,
		Payload: payload,
		QOS:     0x01,
	}

	select {
	case <-ctx.Done():
		return fmt.Errorf("context expired")
	case m.sendChannel <- &p:

	}

	return nil

}

func (m *MQTTClient) SetPublishHandler(f PublishHandler) {
	m.publishHandler = f
}

func (m *MQTTClient) handlePublish(p mqtt.Publish) error {
	err := m.publishHandler.Handle(&p)
	if err != nil {
		return err
	}

	return nil

}
