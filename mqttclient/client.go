package mqttclient

import (
	"crypto/tls"
	"io"
	"net"
	"sync"
	"time"
)

type MQTTClient struct {
	address string

	Topics []string

	username string
	password []byte
	clientID string // if unset a random id may be generated

	connected bool
	connLock  sync.Mutex
	conn      net.Conn

	keepAlive   time.Duration
	connectedAt time.Time
	lastPing    time.Time

	sendChannel chan Sendable

	publishHandler PublishHandler

	// TLS config options
	skipTLSVerify         bool
	tlsClientCertificates []tls.Certificate
}

type Sendable interface {
	Send(w io.Writer) error
}

func New(addr string, username string, password []byte) (*MQTTClient, error) {

	c := MQTTClient{
		address:  addr,
		username: username,
		password: password,

		sendChannel: make(chan Sendable),
	}

	return &c, nil
}

func (c *MQTTClient) SetClientID(id string) {
	c.clientID = id
}

func (c *MQTTClient) SetKeepAlive(d time.Duration) {
	c.keepAlive = d
}

func (c *MQTTClient) SkipTLSVerify(v bool) {
	c.skipTLSVerify = v
}

func (c *MQTTClient) SetClientTLSCertificate(v ...tls.Certificate) {
	c.tlsClientCertificates = v
}
