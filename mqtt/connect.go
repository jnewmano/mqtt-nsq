package mqtt

import (
	"bytes"
	"encoding/binary"
	"io"
)

var (
	connectProtocol  = [6]byte{0x00, 0x04, 'M', 'Q', 'T', 'T'}
	defaultKeepAlive = 60 * 5 // default to a five minute keep alive
)

const (
	connectProtocolLevel = 0x04

	connectFlagCleanSession = 0x02 // not supported
	connectFlagWill         = 0x04
	connectFlagWillQOSLo    = 0x08
	connectFlagWillQOSHi    = 0x10
	connectFlagWillRetain   = 0x20
	connectFlagPassword     = 0x40
	connectFlagUsername     = 0x80
)

/*
Connect is the connection request message
*/
type Connect struct {
	ClientID  string // ^[0-9a-zA-Z]{1-23}$
	KeepAlive uint16

	CleanSession bool

	WillTopic   string
	WillMessage []byte

	Username string
	Password []byte
}

type connectHeader struct {
	ConnectProtocol [6]byte
	Level           byte
	Flags           byte
	KeepAlive       uint16
}

func (c *Connect) flags() byte {

	var f byte

	if c.FlagWill() {
		f += connectFlagWill
	}

	if c.FlagUsername() {
		f += connectFlagUsername
	}

	if c.FlagPassword() {
		f += connectFlagPassword
	}

	if c.FlagCleanSession() {
		f += connectFlagCleanSession
	}

	return f
}

/*
FlagWill returns true if the client supplied a Will
*/
func (c *Connect) FlagWill() bool {
	// an empty will message is okay
	return c.WillTopic != ""
}

/*
FlagUsername returns true if the client supplied a username
*/
func (c *Connect) FlagUsername() bool {
	return c.Username != ""
}

/*
FlagPassword returns true if the client supplied a password
*/
func (c *Connect) FlagPassword() bool {
	return len(c.Password) > 0
}

/*
FlagCleanSession returns true if the client requests a clean session
*/
func (c *Connect) FlagCleanSession() bool {
	return c.CleanSession
}

/*
Send transmits the connect packet
 http://docs.oasis-open.org/mqtt/mqtt/v3.1.1/os/mqtt-v3.1.1-os.html#_Toc385349240
*/
func (c *Connect) Send(w io.Writer) error {

	b := bytes.NewBuffer(nil)

	h := connectHeader{
		ConnectProtocol: connectProtocol,
		Level:           connectProtocolLevel,
		Flags:           c.flags(),
		KeepAlive:       uint16(c.KeepAlive),
	}

	err := binary.Write(b, binary.BigEndian, h)
	if err != nil {
		return err
	}

	// client identifier
	err = writeString(b, c.ClientID) // only 0-9a-zA-Z MUST be accepted, 1-23 characters in length MUST be accepted
	if err != nil {
		return err
	}

	if c.FlagWill() {
		err = writeString(b, c.WillTopic)
		if err != nil {
			return err
		}
		err = writeSlice(b, c.WillMessage)
		if err != nil {
			return err
		}
	}

	if c.FlagUsername() {
		err = writeString(b, c.Username)
		if err != nil {
			return err
		}
	}

	if c.FlagPassword() {
		err = writeSlice(b, c.Password)
		if err != nil {
			return err
		}
	}

	err = send(w, ControlConnect, b)
	if err != nil {
		return err
	}

	return nil

}

/*
ReadConnect reads a connection message from the reader
*/
func ReadConnect(r io.Reader, length int) (Connect, error) {

	var h connectHeader
	err := binary.Read(r, binary.BigEndian, &h)
	if err != nil {
		return Connect{}, err
	}

	length -= 10
	// read out any information specified in the flags

	var clientID string
	err = readString(r, &length, &clientID)
	if err != nil {
		return Connect{}, err
	}

	var willTopic string
	var willMessage []byte
	if h.Flags&connectFlagWill > 0 {
		err = readString(r, &length, &willTopic)
		if err != nil {
			return Connect{}, err
		}
		err = readSlice(r, &length, &willMessage)
		if err != nil {
			return Connect{}, err
		}
	}

	var username string
	if h.Flags&connectFlagUsername > 0 {
		err = readString(r, &length, &username)
		if err != nil {
			return Connect{}, err
		}
	}

	var password []byte
	if h.Flags&connectFlagPassword > 0 {
		err = readSlice(r, &length, &password)
		if err != nil {
			return Connect{}, err
		}
	}

	if h.KeepAlive == 0 {
		h.KeepAlive = uint16(defaultKeepAlive)
	}

	connect := Connect{
		KeepAlive: h.KeepAlive,

		WillTopic:   willTopic,
		WillMessage: willMessage,

		Username: username,
		Password: password,
	}

	return connect, nil
}
