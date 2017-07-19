package mqtt

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
)

const (
	publishDUPMask    = 0x08
	publishQOSMask    = 0x06
	publishRetainMask = 0x01
)

/*
Publish is the publish message.
*/
type Publish struct {
	DUP    bool
	Retain bool
	QOS    byte

	PacketID uint16
	Topic    string
	Payload  []byte
}

/*
Send writes the publish message to a writer.
*/
func (p *Publish) Send(w io.Writer) error {

	b := bytes.NewBuffer(nil)

	err := writeString(b, p.Topic)
	if err != nil {
		return fmt.Errorf("unable to write publish packet topic: %s", err)
	}

	// if QOS == 0 don't attach a packet ID
	if p.QOS > 0 {
		p.PacketID = 2 // TODO: do something better about generating packet ids
		err = binary.Write(b, binary.BigEndian, p.PacketID)
		if err != nil {
			return err
		}
	}

	n, err := b.Write(p.Payload)
	if err != nil {
		return fmt.Errorf("unable to write publish payload: %s", err)
	}
	if n != len(p.Payload) {
		return fmt.Errorf("unable to write complete publish payload %d/%d", n, len(p.Payload))
	}

	control := ControlPublish
	if p.DUP {
		control |= publishDUPMask
	}
	if p.QOS > 0 {
		control |= (p.QOS << 1) & publishQOSMask
	}
	if p.Retain {
		control |= publishRetainMask
	}

	err = send(w, control, b)
	if err != nil {
		return fmt.Errorf("unable to send publish packet: %s", err)
	}

	return nil

}

/*
ReceivePublish reads a publish message from a reader
*/
func ReceivePublish(r io.Reader, reserved byte, length int) (Publish, error) {

	dup := reserved&publishDUPMask > 0
	qos := (reserved & publishQOSMask) >> 1

	if qos > 2 {
		return Publish{}, fmt.Errorf("invalid QOS")
	}

	retain := reserved&publishRetainMask > 0

	var topic string
	err := readString(r, &length, &topic)
	if err != nil {
		return Publish{}, err
	}

	packetID := make([]byte, 2)
	if qos > 0 {
		if length < 2 {
			return Publish{}, fmt.Errorf("not enough bytes remaining for packetID")
		}
		// packet ID is only present if QOS > 0
		_, err = io.ReadFull(r, packetID)
		if err != nil {
			return Publish{}, err
		}
		length -= 2
	}

	if length < 0 {
		return Publish{}, fmt.Errorf("bad length")
	}

	payload := make([]byte, length)
	_, err = io.ReadFull(r, payload)
	if err != nil {
		return Publish{}, err
	}

	p := Publish{
		DUP:    dup,
		Retain: retain,
		QOS:    qos,

		Topic:    topic,
		PacketID: uint16(packetID[0])<<8 + uint16(packetID[1]),
		Payload:  payload,
	}

	return p, nil

}
