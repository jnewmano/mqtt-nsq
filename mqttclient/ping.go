package mqttclient

import (
	"context"
	"fmt"
	"time"

	"github.com/jnewmano/mqtt-nsq/mqtt"
)

// send a ping once every interval
func (m *MQTTClient) keepAliveLoop(ctx context.Context) {

	t := time.NewTicker(m.keepAlive)
	defer t.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-t.C:
			err := m.Ping(ctx)
			if err != nil {
				return
			}
		}
	}

	return
}

func (m *MQTTClient) Ping(ctx context.Context) error {

	p := mqtt.PingReq{}

	select {
	case <-ctx.Done():
		return fmt.Errorf("context done")
	case m.sendChannel <- &p:
		m.lastPing = time.Now()
	}

	return nil
}
