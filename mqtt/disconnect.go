package mqtt

import (
	"fmt"
	"io"
)

/*
Disconnect is the disconnect message.
*/
type Disconnect struct{}

/*
Send sends the disconnect message.
*/
func (p *Disconnect) Send(w io.Writer) error {

	// write out the header
	err := send(w, ControlDisconnect, nil)
	if err != nil {
		return err
	}

	return nil
}

/*
ReceiveDisconnect reads a disconnect message from the reader.
*/
func ReceiveDisconnect(r io.Reader, reserved byte, l int) error {
	if l != 0 {
		return fmt.Errorf("disconnect request length must be 0")
	}

	return nil
}
