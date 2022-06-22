package packet

import (
	"encoding/binary"
)

// ReadBytes reads a fixed number of bytes (`size`) from the packet
func (p *Packet) ReadBytes(size int) []byte {
	buf := make([]byte, size)
	binary.Read(&p.buf, binary.LittleEndian, &buf)

	return buf
}

// ReadShort reads a short from the packet
func (p *Packet) ReadShort() (w uint16) {
	binary.Read(&p.buf, binary.LittleEndian, &w)

	return
}

// ReadInt reads an int from the packet
func (p *Packet) ReadInt() (dw uint32) {
	binary.Read(&p.buf, binary.LittleEndian, &dw)

	return
}

// ReadLong reads a long from the packet
func (p *Packet) ReadLong() (qw uint64) {
	binary.Read(&p.buf, binary.LittleEndian, &qw)

	return
}

// ReadString reads a string from the packet
func (p *Packet) ReadString() (length int, str string) {
	length = int(p.ReadShort())
	str = string(p.ReadBytes(length))

	return
}
