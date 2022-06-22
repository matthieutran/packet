package packet

import "encoding/binary"

// WriteByte appends a byte to the packet
func (p *Packet) WriteByte(d byte) (err error) {
	return binary.Write(&p.buf, binary.LittleEndian, d)
}

// WriteBytes appends an array of bytes to the packet
func (p *Packet) WriteBytes(d []byte) (err error) {
	return binary.Write(&p.buf, binary.LittleEndian, d)

}

// WriteShort appends a short to the packet
func (p *Packet) WriteShort(w uint16) (err error) {
	return binary.Write(&p.buf, binary.LittleEndian, w)

}

// Writeint appends an int to the packet
func (p *Packet) WriteInt(dw uint32) (err error) {
	return binary.Write(&p.buf, binary.LittleEndian, dw)
}

// WriteLong appends a long to the packet
func (p *Packet) WriteLong(qw uint64) (err error) {
	return binary.Write(&p.buf, binary.LittleEndian, qw)
}

// WriteString appends a string to the packet
//
// Strings are encoded on the packet with a short containing the length of the string and then the string in bytes.
func (p *Packet) WriteString(str string) (err error) {
	err = p.WriteShort(uint16(len(str))) // Length of the string
	if err != nil {
		return
	}
	err = p.WriteBytes([]byte(str)) // The string in bytes

	return
}

func (p *Packet) WriteSignedByte(b int8) (err error)   { return p.WriteByte(byte(b)) }
func (p *Packet) WriteSignedShort(b int16) (err error) { return p.WriteShort(uint16(b)) }
func (p *Packet) WriteSignedInt(b int32) (err error)   { return p.WriteInt(uint32(b)) }
func (p *Packet) WriteSignedLong(b int64) (err error)  { return p.WriteLong(uint64(b)) }
