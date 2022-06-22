package packet

import "encoding/binary"

// WriteByte appends a byte to the packet
func (p *Packet) WriteByte(d byte) (err error) {
	binary.Write(&p.buf, binary.LittleEndian, d)

	return
}

// WriteBytes appends an array of bytes to the packet
func (p *Packet) WriteBytes(d []byte) {
	binary.Write(&p.buf, binary.LittleEndian, d)
}

// WriteShort appends a short to the packet
func (p *Packet) WriteShort(w uint16) {
	binary.Write(&p.buf, binary.LittleEndian, w)
}

// Writeint appends an int to the packet
func (p *Packet) WriteInt(dw uint32) {
	binary.Write(&p.buf, binary.LittleEndian, dw)
}

// WriteLong appends a long to the packet
func (p *Packet) WriteLong(qw uint64) {
	binary.Write(&p.buf, binary.LittleEndian, qw)
}

// WriteString appends a string to the packet
//
// Strings are encoded on the packet with a short containing the length of the string and then the string in bytes.
func (p *Packet) WriteString(str string) {
	p.WriteShort(uint16(len(str))) // Length of the string
	p.WriteBytes([]byte(str))      // The string in bytes
}

func (p *Packet) WriteSignedByte(b int8)   { p.WriteByte(byte(b)) }
func (p *Packet) WriteSignedShort(b int16) { p.WriteShort(uint16(b)) }
func (p *Packet) WriteSignedInt(b int32)   { p.WriteInt(uint32(b)) }
func (p *Packet) WriteSignedLong(b int64)  { p.WriteLong(uint64(b)) }
