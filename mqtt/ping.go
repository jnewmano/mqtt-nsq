package mqtt

import (
	"fmt"
	"io"
)

/*
PingReq is a ping message.
*/
type PingReq struct{}

/*
Send writes the PingReq message to the writer.
*/
func (p *PingReq) Send(w io.Writer) error {

	// write out the header
	err := send(w, ControlPingReq, nil)
	if err != nil {
		return err
	}

	return nil
}

/*
ReceivePingPingReq reads a PingReq rom the reader.
*/
func ReceivePingPingReq(r io.Reader, reserved byte, l int) error {
	if l != 0 {
		return fmt.Errorf("ping request length must be 0")
	}

	return nil
}

/*
PingResp is the ping request response message.
*/
type PingResp struct{}

/*
Send writes the PingResp message to a writer.
*/
func (p *PingResp) Send(w io.Writer) error {

	// write out the header
	err := send(w, ControlPingResp, nil)
	if err != nil {
		return err
	}

	return nil
}

/*
ReceivePingResp reads a PingResp message from a reader.
*/
func ReceivePingResp(r io.Reader, reserved byte, l int) error {
	if l != 0 {
		return fmt.Errorf("ping request length must be 0")
	}

	return nil
}
