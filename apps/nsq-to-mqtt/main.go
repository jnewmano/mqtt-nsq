package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"os"
	"os/signal"
	"syscall"
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
	flag.StringVar(&s.MQTT.Topic, "mqtt-topic", "", "MQTT publish topic")

	flag.StringVar(&s.MQTT.ClientCertificate, "mqtt-client-crt", "", "MQTT client certificate")
	flag.StringVar(&s.MQTT.ClientKey, "mqtt-client-key", "", "MQTT client certificate key")

	flag.Var(&s.NSQ.LookupdAddresses, "nsq-lookupd-address", "NSQ lookupd HTTP address:port (only supply NSQd or Lookupd addresses)")
	flag.Var(&s.NSQ.NSQdAddresses, "nsqd-address", "NSQ nsqd TCP address:port")

	flag.StringVar(&s.NSQ.Topic, "nsq-topic", "", "NSQ publish topic")
	flag.StringVar(&s.NSQ.Channel, "nsq-channel", "", "NSQ consumer topic")
	flag.BoolVar(&s.NSQ.WrapPayload, "nsq-wrap-payload", true, "Wrap payloads in procol buffer")

	flag.Parse()

	ctx := context.Background()

	fmt.Printf("Connecting to MQTT [%s]\n", s.MQTT.Address)
	c, err := mqttclient.New(s.MQTT.Address, s.MQTT.Username, []byte(s.MQTT.Password))
	if err != nil {
		exit(err)
	}

	c.SetClientID(s.MQTT.ClientID)
	c.SetKeepAlive(time.Duration(s.MQTT.KeepAlive))

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

	fmt.Printf("Connecting to NSQ\n")
	p, err := newNSQConsumer(s.NSQ.LookupdAddresses, s.NSQ.NSQdAddresses, s.NSQ.Topic, s.NSQ.Channel, s.NSQ.WrapPayload, c.Publish, s.MQTT.Topic)
	if err != nil {
		exit(err)
	}

	wait()

	p.Stop()
}

func wait() {

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	select {
	case <-c:
	}

}

func exit(err error) {
	fmt.Println(err)
	time.Sleep(time.Second * 2)
	os.Exit(1)
}
