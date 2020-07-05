package inventory

import "unsafe"

const defaultBufferSize = 8094
const minumumGrowSize = 32

// BinaryStream encode bytes
type BinaryStream struct {
	buf []byte
	pos int
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
		buf: make([]byte, defaultBufferSize),
	}
}

// NewBinaryStreamWithSize new stream with buffer size
func NewBinaryStreamWithSize(n int) *BinaryStream {
	return &BinaryStream{
		pos: 0,
		buf: make([]byte, n),
	}
}

// NewBinaryStreamWithByteArray internal buffer use b
func NewBinaryStreamWithByteArray(b []byte) *BinaryStream {
	return &BinaryStream{
		pos: 0,
		buf: b,
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

func (b *BinaryStream) checkGrowN(n int) {
	if b.pos+n > len(b.buf) {
		b.growN(len(b.buf) - b.pos - 1 + n)
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

// EncodeCString encode string + NULL
// "hello" => 'h', 'e', 'l', 'l', 'o', '\0'
func (b *BinaryStream) EncodeCString(v string) {
	length := len(v)
	b.checkGrowN(length + 1) // size + string length

	copy(b.buf[b.pos:], *(*[]byte)(unsafe.Pointer(&v)))
	b.pos += length + 1
}

// EncodeUInt16LengthString encode length + string
// "hello" => int16(5) + 'h', 'e', 'l', 'l', 'o'
func (b *BinaryStream) EncodeUInt16LengthString(v string) {
	length := len(v)
	b.checkGrowN(length + 2) // size + string length

	b.buf[b.pos] = byte(length)
	b.buf[b.pos+1] = byte(length >> 8)

	copy(b.buf[b.pos+2:], *(*[]byte)(unsafe.Pointer(&v)))

	b.pos += length + 2
}

// GetBytes get buffer
func (b *BinaryStream) GetBytes() []byte {
	return b.buf[0:b.pos]
}

// DecodeByte decode value
func (b *BinaryStream) DecodeByte() byte {
	v := b.buf[b.pos]
	b.pos++
	return v
}

// DecodeUInt8 decode value
func (b *BinaryStream) DecodeUInt8() uint8 {
	return uint8(b.DecodeByte())
}

// DecodeInt8 decode value
func (b *BinaryStream) DecodeInt8() int8 {
	return int8(b.DecodeByte())
}

// DecodeUInt16 decode value
func (b *BinaryStream) DecodeUInt16() uint16 {
	v := *(*uint16)(unsafe.Pointer(&b.buf[b.pos]))
	b.pos += 2
	return v
}

// DecodeInt16 decode value
func (b *BinaryStream) DecodeInt16() int16 {
	return int16(b.DecodeUInt16())
}

// DecodeUInt32 decode value
func (b *BinaryStream) DecodeUInt32() uint32 {
	v := *(*uint32)(unsafe.Pointer(&b.buf[b.pos]))
	b.pos += 4
	return v
}

// DecodeInt32 decode value
func (b *BinaryStream) DecodeInt32() int32 {
	return int32(b.DecodeUInt32())
}

// DecodeUInt64 decode value
func (b *BinaryStream) DecodeUInt64() uint64 {
	v := *(*uint64)(unsafe.Pointer(&b.buf[b.pos]))
	b.pos += 8
	return v
}

// DecodeInt64 decode value
func (b *BinaryStream) DecodeInt64() int64 {
	return int64(b.DecodeUInt64())
}

// DecodeCString decode string + NULL
// 'h', 'e', 'l', 'l', 'o', '\0' => "hello"
func (b *BinaryStream) DecodeCString() string {
	r := make([]byte, 0, 32)
	for {
		v := b.DecodeByte()
		if v == 0 {
			break
		}
		r = append(r, v)
	}
	return *(*string)(unsafe.Pointer(&r))
}

// DecodeUInt16LengthString encode length + string
// int16(5) + 'h', 'e', 'l', 'l', 'o' => "hello"
func (b *BinaryStream) DecodeUInt16LengthString() string {
	length := int(b.DecodeUInt16())
	v := b.buf[b.pos : b.pos+length]
	b.pos += length
	return *(*string)(unsafe.Pointer(&v))
}
