package mqtt

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
)

/*
PublishAck is the publish acknowledgement message.
*/
type PublishAck struct {
	PacketID uint16
}

/*
Send writes the PublishAck to a writer.
*/
func (p *PublishAck) Send(w io.Writer) error {

	b := bytes.NewBuffer(nil)

	err := binary.Write(b, binary.BigEndian, p)
	if err != nil {
		return fmt.Errorf("unable to generate publish ack header: %s", err)
	}

	err = send(w, ControlPublishAck, b)
	if err != nil {
		return fmt.Errorf("unable to send publish ack: %s", err)
	}

	return nil
}

/*
ReceivePublishAck reads a PublishAck message from a reader.
*/
func ReceivePublishAck(r io.Reader, reserved byte, length int) (PublishAck, error) {

	if length != 2 {
		return PublishAck{}, fmt.Errorf("invalid publish ack length")
	}

	p := PublishAck{}

	err := binary.Read(r, binary.BigEndian, &p)
	if err != nil {
		return PublishAck{}, fmt.Errorf("unable to read publish ack: %s", err)
	}

	return p, nil
}
