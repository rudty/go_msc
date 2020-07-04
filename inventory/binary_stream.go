package inventory

import "unsafe"

const minumumGrowSize = 32

// BinaryStream encode bytes
type BinaryStream struct {
	pos int
	buf []byte
}

func maxValue(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// NewBinaryStream new stream
func NewBinaryStream() *BinaryStream {
	return &BinaryStream{
		pos: 0,
		buf: make([]byte, 8094),
	}
}

// NewBinaryStreamWithSize new stream with buffer size
func NewBinaryStreamWithSize(n int) *BinaryStream {
	return &BinaryStream{
		pos: 0,
		buf: make([]byte, n),
	}
}

func (b *BinaryStream) growN(n int) {
	n = maxValue(minumumGrowSize, n)
	var newBuf = make([]byte, n)
	copy(newBuf, b.buf[0:b.pos])
	b.buf = newBuf
}

func (b *BinaryStream) checkGrow(n int) {
	if b.pos+n > len(b.buf) {
		b.growN(len(b.buf) * 2)
	}
}

// EncodeByte encode value
func (b *BinaryStream) EncodeByte(v byte) {
	b.checkGrow(1)
	b.buf[b.pos] = v
	b.pos++
}

// EncodeUInt8 encode value
func (b *BinaryStream) EncodeUInt8(v uint8) {
	b.EncodeByte(v)
}

// EncodeInt8 encode value
func (b *BinaryStream) EncodeInt8(v int8) {
	b.EncodeByte(byte(v))
}

// EncodeUInt16 encode value
func (b *BinaryStream) EncodeUInt16(v uint16) {
	b.checkGrow(2)
	a := b.buf[b.pos:]
	a[0] = byte(v)
	a[1] = byte(v >> 8)
	b.pos += 2
}

// EncodeInt16 encode value
func (b *BinaryStream) EncodeInt16(v int16) {
	b.EncodeUInt16(uint16(v))
}

// EncodeUInt32 encode value
func (b *BinaryStream) EncodeUInt32(v uint32) {
	b.checkGrow(4)
	a := b.buf[b.pos:]
	a[0] = byte(v)
	a[1] = byte(v >> 8)
	a[2] = byte(v >> 16)
	a[3] = byte(v >> 24)
	b.pos += 4
}

// EncodeInt32 encode value
func (b *BinaryStream) EncodeInt32(v int32) {
	b.EncodeUInt32(uint32(v))
}

// EncodeUInt64 encode value
func (b *BinaryStream) EncodeUInt64(v uint64) {
	b.checkGrow(8)
	a := b.buf[b.pos:]
	a[0] = byte(v)
	a[1] = byte(v >> 8)
	a[2] = byte(v >> 16)
	a[3] = byte(v >> 24)
	a[4] = byte(v >> 32)
	a[5] = byte(v >> 40)
	a[6] = byte(v >> 48)
	a[7] = byte(v >> 56)
	b.pos += 8
}

// EncodeInt64 encode value
func (b *BinaryStream) EncodeInt64(v int64) {
	b.EncodeUInt64(uint64(v))
}

// EncodeByteArray encode value
func (b *BinaryStream) EncodeByteArray(v []byte) {
	dist := b.pos + len(v) - len(b.buf)
	if dist > 0 {
		b.growN(len(b.buf) + dist)
	}
	copy(b.buf[b.pos:], v)
	b.pos += len(v)
}

// EncodeString encode value
func (b *BinaryStream) EncodeString(v string) {
	remainSize := len(b.buf) - b.pos
	requireSize := len(v) + 1 - remainSize // string + NULL
	if requireSize > 0 {
		b.growN(len(b.buf) + requireSize)
	}
	copy(b.buf[b.pos:], *(*[]byte)(unsafe.Pointer(&v)))
	b.pos += len(v)
}

// GetBytes get buffer
func (b *BinaryStream) GetBytes() []byte {
	return b.buf
}
