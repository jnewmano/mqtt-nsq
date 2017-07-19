package mqtt

import (
	"encoding/binary"
	"fmt"
	"io"
)

/*
SubscribeAck is the subscribe message response.
*/
type SubscribeAck struct {
	PacketID  uint16
	TopicQOSs []byte
}

type subscribeAckHeader struct {
	PacketID uint16
}

/*
ReceiveSubscribeAck reads a SubscribeAck from a reader.
*/
func ReceiveSubscribeAck(r io.Reader, length int) (SubscribeAck, error) {

	h := subscribeAckHeader{}
	err := binary.Read(r, binary.BigEndian, &h)
	if err != nil {
		return SubscribeAck{}, err
	}
	length -= 2

	if length <= 0 {
		return SubscribeAck{}, fmt.Errorf("malformed subscribe ack, invalid remaining length: %d", length)
	}

	qs := make([]byte, length)

	_, err = io.ReadFull(r, qs)
	if err != nil {
		return SubscribeAck{}, err
	}

	s := SubscribeAck{
		PacketID:  h.PacketID,
		TopicQOSs: qs,
	}

	return s, nil
}

/*
Send write a subscribe ack to a writer.
*/
func (s *SubscribeAck) Send(w io.Writer) error {
	return fmt.Errorf("not implemented")
}
