// Package packet provides tools to work with packets
package packet

import (
	"bytes"
	"fmt"
)

// Packet is an alias of a `bytes.Buffer`
type Packet struct {
	buf bytes.Buffer
	fmt.Stringer
}

// Size returns the current size of the buffer
func (p Packet) Size() int {
	return int(len(p.buf.Bytes()))
}

// Bytes returns a slice copy of the current unread portion of the byte buffer.
// The slice is most current until a read or write operation is performed on the buffer.
func (p Packet) Bytes() []byte {
	return p.buf.Bytes()
}

// String implements the `fmt.Stringer` interface for the packet.
func (p Packet) String() string {
	return fmt.Sprintf("% X", p.buf.String())
}
