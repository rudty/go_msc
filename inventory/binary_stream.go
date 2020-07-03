package inventory

// BinaryStream encode bytes
type BinaryStream struct {
	pos int
	buf []byte
}

// NewBinaryStream new stream
func NewBinaryStream() *BinaryStream {
	return &BinaryStream{
		pos: 0,
		buf: make([]byte, 10),
	}
}

func (b *BinaryStream) growN(n int) {
	var newBuf = make([]byte, n)
	copy(newBuf, b.buf)
	b.buf = newBuf
}

func (b *BinaryStream) grow() {
	b.growN(len(b.buf) * 2)
}

// EncodeInt8 encode value
func (b *BinaryStream) EncodeInt8(v int8) {
	if b.pos == len(b.buf) {
		b.grow()
	}
	b.buf[b.pos] = byte(v)
	b.pos++
}

// // SerializeInt8 byte slice 에 int8 을 인코딩합니다.
// func SerializeInt8(b []byte, v int8) []byte {
// 	b = append(b, byte(v))
// 	return b
// }

// // SerializeInt16 byte slice 에 int16 을 인코딩합니다.
// func SerializeInt16(b []byte, v int16) []byte {
// 	b = append(b, byte(v))
// 	b = append(b, byte(v>>8))
// 	return b
// }

// // SerializeInt32 byte slice 에 int32 을 인코딩합니다.
// func SerializeInt32(b []byte, v int32) []byte {
// 	b = append(b, byte(v))
// 	b = append(b, byte(v>>8))
// 	b = append(b, byte(v>>16))
// 	b = append(b, byte(v>>24))
// 	return b
// }

// // SerializeInt64 byte slice 에 int64 을 인코딩합니다.
// func SerializeInt64(b []byte, v int64) []byte {
// 	b = append(b, byte(v))
// 	b = append(b, byte(v>>8))
// 	b = append(b, byte(v>>16))
// 	b = append(b, byte(v>>24))
// 	b = append(b, byte(v>>32))
// 	b = append(b, byte(v>>40))
// 	b = append(b, byte(v>>48))
// 	b = append(b, byte(v>>56))
// 	return b
// }
