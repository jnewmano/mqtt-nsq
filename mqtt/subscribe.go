package mqtt

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
)

/*
Topic is a topic and QOS pair
*/
type Topic struct {
	Topic string
	QOS   byte
}

/*
Subscribe is the subscribe message.
*/
type Subscribe struct {
	Topics []Topic
}

type subscribeHeader struct {
	packetID uint16
}

/*
ReadSubscribe reads a subscribe message from a reader.
*/
func ReadSubscribe(r io.Reader, reserved byte, length int) (Subscribe, error) {
	return Subscribe{}, fmt.Errorf("not implemented")
}

/*
AddTopic adds a topic and QOS pair to a subscribe message.
*/
func (s *Subscribe) AddTopic(t string, q byte) error {
	if q > 2 {
		return fmt.Errorf("invalid QOS")
	}

	topic := Topic{
		Topic: t,
		QOS:   q,
	}

	s.Topics = append(s.Topics, topic)

	return nil
}

/*
Send writes a subscribe message to a writer.
*/
func (s *Subscribe) Send(w io.Writer) (uint16, error) {

	b := bytes.NewBuffer(nil)

	h := subscribeHeader{
		packetID: 1,
	}

	err := binary.Write(b, binary.BigEndian, h)
	if err != nil {
		return 0, err
	}

	// write out the requested topic subscriptions
	for _, v := range s.Topics {
		fmt.Printf("subscribing to topic %s QOS %d\n", v.Topic, v.QOS)

		err = writeString(b, v.Topic)
		if err != nil {
			return 0, err
		}

		if v.QOS > 2 {
			return 0, fmt.Errorf("invalid subscribe QOS: %d", v.QOS)
		}

		n, err := b.Write([]byte{v.QOS})
		if err != nil {
			return 0, err
		}
		if n != 1 {
			return 0, fmt.Errorf("unable to write subscribe topic QOS")
		}
	}

	err = send(w, ControlSubscribe, b)
	if err != nil {
		return 0, fmt.Errorf("unable to send control subscribe: %s", err)
	}

	return h.packetID, nil

}
