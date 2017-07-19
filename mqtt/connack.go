package mqtt

import (
	"fmt"
	"io"
)

/*
Connection request responses.
*/
const (
	ConnectionAccepted              = byte(0x00)
	ConnectionRefusedProtocol       = byte(0x01)
	ConnectionRefusedIdentifier     = byte(0x02)
	ConnectionRefusedUnavailable    = byte(0x03)
	ConnectionRefusedBadCredentials = byte(0x04)
	ConnectionRefusedUnauthorized   = byte(0x05)
)

/*
ConnAck is the connection request response
*/
type ConnAck struct {
	ConnectStatus byte
	OK            bool
}

/*
Send sends the connaction acknowledgment
*/
func (c *ConnAck) Send(w io.Writer) error {

	var err error

	// write out the header
	_, err = w.Write([]byte{ControlConnAck})
	if err != nil {
		return err
	}

	err = writeVarint(w, 2)
	if err != nil {
		return err
	}

	// TODO: implement session storage?
	n, err := w.Write([]byte{0, c.ConnectStatus})
	if err != nil {
		return err
	}
	if n != 2 {
		return fmt.Errorf("unable to write all ConnAck bytes")
	}

	return nil

}

/*
ReadConnack reads a ConnackAck message from the reader
*/
func ReadConnack(r io.Reader, length int) (ConnAck, error) {

	if length != 2 {
		return ConnAck{}, fmt.Errorf("unexpected length for ConnAck [%d]", length)
	}
	b := make([]byte, length)

	n, err := io.ReadFull(r, b)
	if err != nil {
		return ConnAck{}, err
	}
	if n != len(b) {
		return ConnAck{}, fmt.Errorf("unable to read full connack response")
	}

	// ignore b[0]&0x01 session present flag
	c := ConnAck{
		ConnectStatus: b[1],
		OK:            b[1] == ConnectionAccepted,
	}

	return c, nil
}
