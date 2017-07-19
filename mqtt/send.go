package mqtt

import (
	"bytes"
	"fmt"
	"io"
)

func send(w io.Writer, control byte, b *bytes.Buffer) error {

	// write out the header
	n, err := w.Write([]byte{control})
	if err != nil {
		return fmt.Errorf("unable to send control byte: %s", err)
	}
	if n != 1 {
		return fmt.Errorf("unable to write control byte")
	}

	l := 0
	if b != nil {
		l = b.Len()
	}

	err = writeVarint(w, l)
	if err != nil {
		return fmt.Errorf("unable to send control length: %s", err)
	}

	if l == 0 {
		return nil
	}

	ln, err := io.Copy(w, b)
	if err != nil {
		return fmt.Errorf("unable to copy to writer: %s", err)
	}
	if int(ln) != l {
		return fmt.Errorf("unable to write full packet data")
	}

	return nil
}
