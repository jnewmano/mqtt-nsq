package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"os"
	"time"

	"github.com/jnewmano/mqtt-nsq/mqttclient"
	"github.com/namsral/flag"
)

// generate a certificate CSR
// openssl req -newkey rsa:2048 -nodes -keyout client.key -out client.csr

func main() {

	var s Settings
	flag.String(flag.DefaultConfigFlagname, "", "path to config file")

	flag.StringVar(&s.MQTT.Address, "mqtt-address", "", "MQTT server address:port")
	flag.StringVar(&s.MQTT.Username, "mqtt-username", "", "MQTT username")
	flag.StringVar(&s.MQTT.Password, "mqtt-password", "", "MQTT password")
	flag.StringVar(&s.MQTT.ClientID, "mqtt-client-id", "", "MQTT client id")
	flag.DurationVar(&s.MQTT.KeepAlive, "mqtt-keep-alive", 0, "MQTT keep alive")
	flag.Var(&s.MQTT.Topics, "mqtt-topics", "MQTT topics, allows repeated")

	flag.StringVar(&s.MQTT.ClientCertificate, "mqtt-client-crt", "", "MQTT client certificate")
	flag.StringVar(&s.MQTT.ClientKey, "mqtt-client-key", "", "MQTT client certificate key")

	flag.StringVar(&s.NSQ.Address, "nsqd-address", "", "NSQd TCP address:port")
	flag.StringVar(&s.NSQ.Topic, "nsqd-topic", "", "NSQd publish topic")
	flag.BoolVar(&s.NSQ.WrapPayload, "nsqd-wrap-payload", true, "Wrap payloads in procol buffer")

	flag.Parse()

	ctx := context.Background()

	fmt.Printf("Connecting to MQTT [%s]\n", s.MQTT.Address)
	c, err := mqttclient.New(s.MQTT.Address, s.MQTT.Username, []byte(s.MQTT.Password))
	if err != nil {
		exit(err)
	}

	c.SetClientID(s.MQTT.ClientID)
	c.SetKeepAlive(time.Duration(s.MQTT.KeepAlive))

	fmt.Printf("Subscribing to topics %v\n", s.MQTT.Topics)
	c.Topics = s.MQTT.Topics

	fmt.Printf("Connecting to NSQd [%s]\n", s.NSQ.Address)
	p, err := newNSQProducer(s.NSQ.Address, s.NSQ.Topic, s.NSQ.WrapPayload)
	if err != nil {
		exit(err)
	}

	c.SetPublishHandler(p)
	c.SkipTLSVerify(true)

	if s.MQTT.ClientCertificate != "" && s.MQTT.ClientKey != "" {
		fmt.Println("loading client certificate and key")
		clientCertificate, err := tls.LoadX509KeyPair(s.MQTT.ClientCertificate, s.MQTT.ClientKey)
		if err != nil {
			fmt.Printf("invalid client certificate, not going to use [%s]\n", err)
		}
		c.SetClientTLSCertificate(clientCertificate)
	}

	err = c.Connect(ctx)
	if err != nil {
		exit(err)
	}

	fmt.Println("waiting for messages")
	select {}
}

func exit(err error) {
	fmt.Println(err)
	time.Sleep(time.Second * 2)
	os.Exit(1)
}
