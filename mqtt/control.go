package mqtt

import (
	"encoding/binary"
	"fmt"
	"io"
)

/*
MQTT control bytes.
*/
const (
	ControlConnect         = byte(0x10)
	ControlConnAck         = byte(0x20)
	ControlPublish         = byte(0x30)
	ControlPublishAck      = byte(0x40)
	ControlPublishReceived = byte(0x50)
	ControlPublishRelease  = byte(0x60)
	ControlPublishComplete = byte(0x70)
	ControlSubscribe       = byte(0x82)
	ControlSubscribeAck    = byte(0x90)
	ControlUnsubscribe     = byte(0xA0)
	ControlUnsubsribeAck   = byte(0xB0)
	ControlPingReq         = byte(0xC0)
	ControlPingResp        = byte(0xD0)
	ControlDisconnect      = byte(0xE0)
)

func readSlice(r io.Reader, l *int, b *[]byte) error {

	// first two bytes are the length
	if *l < 2 {
		return fmt.Errorf("Read* requires at least 2 bytes")
	}
	var d = make([]byte, 2)
	_, err := io.ReadFull(r, d)
	if err != nil {
		return err
	}
	*l = *l - 2

	sl := int(d[0])<<8 + int(d[1])

	if *l < sl {
		return fmt.Errorf("require %d bytes, only have %d", sl, *l)
	}

	s := make([]byte, sl)
	_, err = io.ReadFull(r, s)
	if err != nil {
		return err
	}
	*l = *l - sl

	*b = s

	return nil

}

func writeSlice(w io.Writer, d []byte) error {

	l := len(d)

	n, err := w.Write([]byte{byte(l >> 8), byte(l)})
	if err != nil {
		return err
	}
	if n != 2 {
		return fmt.Errorf("write slice length: unable to write full message")
	}

	n, err = w.Write(d)
	if err != nil {
		return err
	}
	if n != len(d) {
		return fmt.Errorf("write slice bytes: unable to write full message")
	}

	return nil
}

func readString(r io.Reader, l *int, s *string) error {

	var b []byte
	err := readSlice(r, l, &b)
	if err != nil {
		return err
	}

	*s = string(b)
	return nil
}

func writeString(w io.Writer, s string) error {

	err := writeSlice(w, []byte(s))
	if err != nil {
		return err
	}

	return nil
}

func writeVarint(w io.Writer, l int) error {
	b := make([]byte, 4)

	n := binary.PutUvarint(b, uint64(l))

	nn, err := w.Write(b[:n])
	if err != nil {
		return err
	}
	if nn != n {
		fmt.Errorf("unable to write all bytes")
	}

	return nil
}

/*
ReadVarint reads a varint from the reader.
*/
func ReadVarint(r io.Reader) (int, error) {

	l := make([]byte, 0, 4)
	for {

		b := make([]byte, 1)
		n, err := io.ReadFull(r, b)
		if err != nil {
			return 0, err
		}
		if n != 1 {
			return 0, fmt.Errorf("varint: did not read byte %d", n)
		}

		l = append(l, b[0])

		if b[0]&0x80 == 0 {
			break
		}
	}

	// decode the varint
	length, _ := binary.Uvarint(l)

	return int(length), nil
}

/*
ReadCommand reads an MQTT command from the reader
*/
func ReadCommand(r io.Reader) (byte, byte, int, error) {

	// first step is to authenticate the connection
	command := make([]byte, 1)
	n, err := io.ReadFull(r, command)
	if err != nil {
		return 0, 0, 0, err
	}
	if n != 1 {
		return 0, 0, 0, fmt.Errorf("read command, didn't read byte: %d", n)
	}

	l, err := ReadVarint(r)
	if err != nil {
		return 0, 0, 0, err
	}

	cmd := command[0] & 0xF0
	reserved := command[0] & 0x0F

	return cmd, reserved, int(l), nil

}
