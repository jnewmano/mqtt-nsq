package mqttclient

import (
	"bytes"
	"context"
	"crypto/tls"
	"fmt"
	"io"
	"net"
	"time"

	"github.com/jnewmano/mqtt-nsq/mqtt"
)

/*
Common MQTT ports
*/
const (
	DefaultTCPPort       = "1883"
	DefaultTLSPort       = "8883"
	DefaultClientTLSPort = "8884"
	DefaultWSPort        = "8080"
	DefaultWSSPort       = "8081"
)

// Connect establishes a connection with the MQTT broker
// making a best guess on the connection protocol based on
// the MQTT broker port
func (m *MQTTClient) Connect(ctx context.Context) error {

	_, port, err := net.SplitHostPort(m.address)
	if err != nil {
		return fmt.Errorf("unable to split host port: %s", err)
	}

	switch port {
	case DefaultTCPPort:
		err = m.ConnectTCP(ctx)
	case DefaultTLSPort, DefaultClientTLSPort:
		err = m.ConnectTLS(ctx)
	case DefaultWSPort:
		err = m.ConnectWS(ctx)
	case DefaultWSSPort:
		err = m.ConnectWSS(ctx)
	default:
		return fmt.Errorf("unrecognized port [%s], connect using ConnectSSS")
	}
	if err != nil {
		return err
	}

	return nil
}

func (m *MQTTClient) ConnectTCP(ctx context.Context) error {

	conn, err := net.Dial("tcp", m.address)
	if err != nil {
		return err
	}

	return m.connect(ctx, conn)
}

func (m *MQTTClient) ConnectTLS(ctx context.Context) error {

	cfg := &tls.Config{
		Certificates:       m.tlsClientCertificates,
		InsecureSkipVerify: m.skipTLSVerify,
	}

	conn, err := tls.Dial("tcp", m.address, cfg)
	if err != nil {
		return err
	}

	return m.connect(ctx, conn)
}

func (m *MQTTClient) ConnectWS(ctx context.Context) error {
	return fmt.Errorf("not implemented")
}

func (m *MQTTClient) ConnectWSS(ctx context.Context) error {
	return fmt.Errorf("not implemented")
}

func (m *MQTTClient) connect(ctx context.Context, conn net.Conn) error {

	m.connLock.Lock()
	defer m.connLock.Unlock()

	if m.connected {
		return fmt.Errorf("already connected")
	}

	if m.conn != nil {
		m.conn.Close()
	}

	var err error

	r := io.Reader(conn)
	w := io.Writer(conn)

	debug := false
	if debug {
		rb := bytes.NewBuffer(nil)
		r = io.TeeReader(conn, rb)

		defer func() {
			fmt.Printf("[% 02X]\n", rb.Bytes())
		}()

		wb := bytes.NewBuffer(nil)
		w = io.MultiWriter(conn, wb)
	}
	m.conn = conn
	m.connected = true

	if m.clientID == "" {
		return fmt.Errorf("client id is required")
	}

	// send connect request
	connect := mqtt.Connect{
		CleanSession: true,
		ClientID:     m.clientID,
		Username:     m.username,
		Password:     m.password,
		KeepAlive:    uint16(m.keepAlive / time.Second),
	}

	fmt.Println("Sending connect command")
	err = connect.Send(w)
	if err != nil {
		return fmt.Errorf("unable to send connect packet: %s", err)
	}

	// wait for connack response

	fmt.Println("Reading response command")

	conn.SetReadDeadline(time.Now().Add(time.Second * 5))
	cmd, _, length, err := mqtt.ReadCommand(r)
	if err != nil {
		return fmt.Errorf("connect response: unable to read command: %s", err)
	}

	if cmd != mqtt.ControlConnAck {
		return fmt.Errorf("unexpected command, expected conn ack: %v", cmd)
	}

	conn.SetReadDeadline(time.Now().Add(time.Second * 5))
	ack, err := mqtt.ReadConnack(r, length)
	if err != nil {
		return err
	}

	if ack.OK == false {
		return fmt.Errorf("connack error: [%d]", ack.ConnectStatus)
	}

	// send subscription requests
	subscribe := mqtt.Subscribe{}

	for _, v := range m.Topics {
		err := subscribe.AddTopic(v, 1)
		if err != nil {
			fmt.Errorf("unable to generate subscribe: %s", err)
		}
	}

	_, err = subscribe.Send(w)
	if err != nil {
		return fmt.Errorf("unable to subscribe: %s", err)
	}

	if m.keepAlive > 0 {
		go m.keepAliveLoop(ctx)
	}

	go func() {
		err := m.readLoop(ctx, conn)
		if err != nil {
			fmt.Println(err)
		}
		_ = conn.Close()
	}()

	go func() {
		err := m.writeLoop(ctx, conn) // block on write loop, will return when done
		if err != nil {
			fmt.Println(err)
		}
		_ = conn.Close()
	}()

	return nil
}

func (m *MQTTClient) readLoop(ctx context.Context, c net.Conn) error {

	for {

		c.SetReadDeadline(time.Now().Add(time.Minute * 5))
		cmd, reserved, length, err := mqtt.ReadCommand(c)
		if err != nil {
			return fmt.Errorf("unable to read command: %s", err)
		}

		c.SetReadDeadline(time.Now().Add(time.Second * 5))

		switch cmd {
		case mqtt.ControlPublish:
			p, err := mqtt.ReceivePublish(c, reserved, length)
			if err != nil {
				return err
			}

			m.handlePublish(p)

			if p.QOS == 0x01 {
				ack := mqtt.PublishAck{
					PacketID: p.PacketID,
				}

				select {
				case <-ctx.Done():
				case m.sendChannel <- &ack:
				}
			}

		case mqtt.ControlPublishAck:
			_, err := mqtt.ReceivePublishAck(c, reserved, length)
			if err != nil {
				return fmt.Errorf("unable to read command: %s", err)
			}

			// TODO: handle ack

		case mqtt.ControlSubscribeAck:
			_, err := mqtt.ReceiveSubscribeAck(c, length)
			if err != nil {
				return fmt.Errorf("unable to read command: %s", err)
			}

			// TODO: handle ack

		case mqtt.ControlPingResp:
			err := mqtt.ReceivePingResp(c, reserved, length)
			if err != nil {
				return err
			}

			// TODO: handle ack

		default:
			_, err = mqtt.ReceiveSubscribeAck(c, length)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (m *MQTTClient) writeLoop(ctx context.Context, c net.Conn) error {

	for {
		select {
		case <-ctx.Done():

			d := mqtt.Disconnect{}
			_ = d.Send(c)
			c.Close()

			return nil

		case d := <-m.sendChannel:
			err := d.Send(c)
			if err != nil {
				return fmt.Errorf("unable to send: %s", err)
			}
		}
	}
}
