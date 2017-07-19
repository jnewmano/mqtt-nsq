package mqtt

import (
	"bytes"
	"testing"
)

func TestSendPublish(t *testing.T) {
	expectedBytes := []byte{
		0x3D, 0x18, 0x00, 0x0A, 0x74, 0x65, 0x73, 0x74, 0x20, 0x74, 0x6F, 0x70, 0x69, 0x63, 0x00, 0x02, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0A,
	}

	p := Publish{
		Topic:   "test topic",
		Payload: []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		QOS:     2,
		Retain:  true,
		DUP:     true,
	}

	b := bytes.NewBuffer(nil)
	err := p.Send(b)
	if err != nil {
		t.Fatalf("unexpected err: %s\n", err)
	}

	if bytes.Compare(expectedBytes, b.Bytes()) != 0 {
		t.Fatalf("bytes are different [% 02X]", b.Bytes())
	}
}

func TestReceivePublish(t *testing.T) {
	b := []byte{
		0x3D, 0x18, 0x00, 0x0A, 0x74, 0x65, 0x73, 0x74, 0x20, 0x74, 0x6F, 0x70, 0x69, 0x63, 0x00, 0x02, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0A,
	}

	r := bytes.NewBuffer(b[2:])

	p, err := ReceivePublish(r, b[0], int(b[1]))
	if err != nil {
		t.Fatalf("unexpected err: %s\n", err)
	}

	if p.Topic != "test topic" {
		t.Fatalf("unexpected topic: %s", p.Topic)
	}
	if p.DUP != true {
		t.Fatalf("expected to be DUP")
	}
	if p.Retain != true {
		t.Fatalf("expected to be Retain")
	}
	if p.QOS != 2 {
		t.Fatalf("expected to be QOS 2 [%d]", p.QOS)
	}

	if bytes.Compare(p.Payload, b[16:]) != 0 {
		t.Fatalf("unexpected payload bytes")
	}
}
