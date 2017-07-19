package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"net"
	"os"
	"time"

	"github.com/jnewmano/mqtt-nsq/mqttclient"
)

// generate a certificate CSR
// openssl req -newkey rsa:2048 -nodes -keyout client.key -out client.csr

func main() {

	// MQTT address
	// MQTT username
	// MQTT password
	// MQTT topic

	mqttAddress := net.JoinHostPort("test.mosquitto.org", mqttclient.DefaultClientTLSPort)
	mqttUsername := ""
	mqttPassword := []byte("")
	mqttClientID := "someone"
	mqttKeepAlive := time.Second * 25
	mqttSubscribeTopic := "#"

	clientCertificate, err := tls.LoadX509KeyPair("client.crt", "client.key")
	if err != nil {
		fmt.Printf("invalid client certificate, not going to use [%s]\n", err)
	}

	// NSQd address
	// NSQd topic
	// NSQd wrap payload
	nsqdAddress := "localhost:4150"
	nsqdTopic := "mqtt#ephemeral"
	nsqdWrapPayload := true

	ctx := context.Background()

	c, err := mqttclient.New(mqttAddress, mqttUsername, mqttPassword)
	if err != nil {
		exit(err)
	}

	c.SetClientID(mqttClientID)
	c.SetKeepAlive(mqttKeepAlive)
	c.Topics = []string{mqttSubscribeTopic}

	p, err := newNSQProducer(nsqdAddress, nsqdTopic, nsqdWrapPayload)
	if err != nil {
		exit(err)
	}

	c.SetPublishHandler(p)
	c.SkipTLSVerify(true)
	c.SetClientTLSCertificate(clientCertificate)

	err = c.Connect(ctx)
	if err != nil {
		exit(err)
	}

	ignoreError()

	fmt.Println("waiting for messages")
	select {}
}

func exit(err error) {
	fmt.Println(err)
	time.Sleep(time.Second * 2)
	os.Exit(1)
}

func ignoreError() error {
	return fmt.Errorf("error")
}
