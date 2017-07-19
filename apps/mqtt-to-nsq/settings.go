package main

import (
	"strings"
	"time"
)

/*
Settings is used for importing settings
 from a settings configuration file.
*/
type Settings struct {
	MQTT MQTTSettings
	NSQ  NSQdSettings
}

/*
MQTTSettings contains MQTT client related settings.
*/
type MQTTSettings struct {
	Address   string
	Username  string
	Password  string
	ClientID  string
	KeepAlive time.Duration
	Topics    repeatedString

	ClientCertificate string
	ClientKey         string
}

/*
NSQdSettings contains NSQd producer settings.
*/
type NSQdSettings struct {
	Address     string
	Topic       string
	WrapPayload bool
}

type repeatedString []string

func (s *repeatedString) String() string {
	return strings.Join([]string(*s), ",")
}

func (s *repeatedString) Set(v string) error {
	split := strings.Split(v, ":")
	*s = append([]string(*s), split...)
	return nil
}
