package inventory

import (
	"fmt"
	"testing"
)

func checkbyteArrayEquals(a, b []byte) bool {
	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func Test_Serialize_Int8_Not_Grow(t *testing.T) {
	b := NewBinaryStreamWithSize(10)
	for i := 1; i < 11; i++ {
		b.EncodeInt8(int8(i))
	}

	if len(b.buf) != 10 {
		t.Error("must not grow")
	}
}

func Test_Serialize_Int8_Grow(t *testing.T) {
	b := NewBinaryStreamWithSize(10)
	for i := 1; i <= 11; i++ {
		b.EncodeInt8(int8(i))
	}
}

func Test_Serialize_Byte(t *testing.T) {
	b := NewBinaryStreamWithSize(10)

	for i := 1; i < 10; i++ {
		b.EncodeByte(byte(i))
	}
	for i := 1; i < 10; i++ {
		if b.buf[i-1] != byte(i) {
			t.Error("encode byte fail")
			break
		}
	}
}

func Test_Serialize_Int16(t *testing.T) {
	b := NewBinaryStream()
	b.EncodeInt16(-1)

	for i := 0; i < 2; i++ {
		if b.buf[i] != 255 {
			t.Error("encode int32 -1 encode [255, 255, 255, 255]")
		}
	}

	for i := 2; i < len(b.buf); i++ {
		if b.buf[i] != 0 {
			t.Error("must 0")
			break
		}
	}
}

func Test_Serialize_Int32(t *testing.T) {
	b := NewBinaryStream()
	b.EncodeInt32(-1)

	for i := 0; i < 4; i++ {
		if b.buf[i] != 255 {
			t.Error("encode int32 -1 encode [255, 255, 255, 255]")
		}
	}

	for i := 4; i < len(b.buf); i++ {
		if b.buf[i] != 0 {
			t.Error("must 0")
			break
		}
	}
}

func Test_Serialize_Int64(t *testing.T) {
	b := NewBinaryStream()
	b.EncodeInt64(-1)

	for i := 0; i < 8; i++ {
		if b.buf[i] != 255 {
			t.Error("encode int64 -1 encode [255, 255, 255, 255, 255, 255, 255, 255]")
		}
	}

	for i := 8; i < len(b.buf); i++ {
		if b.buf[i] != 0 {
			t.Error("must 0")
			break
		}
	}
}

func Test_Serialize_Int64_Grow(t *testing.T) {
	b := NewBinaryStreamWithSize(10)
	b.EncodeInt64(-1)
	b.EncodeInt64(-1)

	for i := 0; i < 16; i++ {
		if b.buf[i] != 255 {
			t.Error("encode int64 -1, -1 encode ([255, 255, 255, 255, 255, 255, 255, 255] * 2)")
		}
	}

	for i := 16; i < len(b.buf); i++ {
		if b.buf[i] != 0 {
			t.Error("must 0")
			break
		}
	}
}

func Test_Serialize_Int64_Size(t *testing.T) {
	b := NewBinaryStreamWithSize(16)
	b.EncodeInt64(-1)
	b.EncodeInt64(-1)

	for i := 0; i < len(b.buf); i++ {
		if b.buf[i] != 255 {
			t.Error("must 255")
		}
	}
}

func Test_Serialize_ByteArray_Size_1(t *testing.T) {
	b := NewBinaryStreamWithSize(1)
	b.EncodeByteArray([]byte{255, 1, 3, 4, 5})
	if len(b.buf) < 5 {
		t.Error("must grow")
	}
}

func Test_Serialize_ByteArray_Size_2(t *testing.T) {
	b := NewBinaryStreamWithSize(1)
	b.EncodeUInt8(1)
	b.EncodeByteArray([]byte{255, 1, 3, 4, 5})
	if len(b.buf) < 6 {
		t.Error("must grow")
	}
}

func Test_Serialize_ByteArray_Size_5(t *testing.T) {
	b := NewBinaryStreamWithSize(5)
	b.EncodeByteArray([]byte{255, 1, 3, 4, 5})
	if len(b.buf) != 5 {
		t.Error("must grow")
	}
}

func Test_Serialize_String(t *testing.T) {
	b := NewBinaryStreamWithSize(6)
	b.EncodeCString("Hello")

	if len(b.buf) != 6 {
		t.Error("'H','e','l','l','o',NULL = 6")
	}

	b = NewBinaryStreamWithSize(4)
	b.EncodeCString("Hello")

	if len(b.buf) < 6 {
		t.Error("grow 'H','e','l','l','o',NULL = 6")
	}

	const longString = "HelloWorld HelloWorld HelloWorld HelloWorld HelloWorld HelloWorld"
	b = NewBinaryStreamWithSize(1)
	b.EncodeCString(longString)

	if len(b.buf) <= len(longString) {
		t.Error(fmt.Sprint("grow verylong string size:", len(longString)))
	}
}

func Test_Serialize_String2(t *testing.T) {
	b := NewBinaryStreamWithSize(32)
	b.EncodeCString("Hello")
	b.EncodeCString("Hello")

	if len(b.GetBytes()) != 12 {
		t.Error("hello + \\0 * 2 = 12 ")
	}
}

func Test_Serialize_Length_String(t *testing.T) {
	b := NewBinaryStreamWithSize(10)
	b.EncodeUInt16LengthString("Hello")

	if !checkbyteArrayEquals(
		b.GetBytes(),
		[]byte{5, 0, 72, 101, 108, 108, 111}) {
		t.Error("encode 5Hello")
	}

	b = NewBinaryStreamWithSize(1)
	b.EncodeUInt16LengthString("Hello")

	if !checkbyteArrayEquals(
		b.GetBytes(),
		[]byte{5, 0, 72, 101, 108, 108, 111}) {
		t.Error("encode 5Hello")
	}

	const longString = "HelloHelloHelloHelloHelloHelloHelloHelloHelloHelloHelloHelloHelloHelloHelloHelloHello"
	b = NewBinaryStreamWithSize(1)
	b.EncodeUInt16LengthString(longString)
}

func Test_Decode_Byte(t *testing.T) {
	b := NewBinaryStream()
	b.EncodeUInt8(255)
	b.EncodeInt8(-123)

	d := NewBinaryStreamWithByteArray(b.GetBytes())
	if d.DecodeUInt8() != 255 {
		t.Error("encode 255")
	}
	if d.DecodeInt8() != -123 {
		t.Error("encode -123")
	}
}

func Test_Decode_int16(t *testing.T) {
	b := NewBinaryStream()
	b.EncodeUInt16(65535)
	b.EncodeInt16(32767)
	b.EncodeInt16(-32768)

	d := NewBinaryStreamWithByteArray(b.GetBytes())
	if d.DecodeUInt16() != 65535 {
		t.Error("decode fail 65535")
	}
	if d.DecodeInt16() != 32767 {
		t.Error("decode fail 32767")
	}
	if d.DecodeInt16() != -32768 {
		t.Error("decode fail -32768")
	}
}

func Test_Decode_int32(t *testing.T) {
	b := NewBinaryStream()
	b.EncodeUInt32(4294967295)
	b.EncodeInt32(2147483647)
	b.EncodeInt32(-2147483648)

	d := NewBinaryStreamWithByteArray(b.GetBytes())
	if d.DecodeUInt32() != 4294967295 {
		t.Error("decode fail 4294967295")
	}
	if d.DecodeInt32() != 2147483647 {
		t.Error("decode fail 2147483647")
	}
	if d.DecodeInt32() != -2147483648 {
		t.Error("decode fail -2147483648")
	}
}

func Test_Decode_int64(t *testing.T) {
	b := NewBinaryStream()
	b.EncodeUInt64(18446744073709551615)
	b.EncodeInt64(9223372036854775807)
	b.EncodeInt64(-9223372036854775808)

	d := NewBinaryStreamWithByteArray(b.GetBytes())
	if d.DecodeUInt64() != 18446744073709551615 {
		t.Error("decode fail 18446744073709551615")
	}
	if d.DecodeInt64() != 9223372036854775807 {
		t.Error("decode fail 9223372036854775807")
	}
	if d.DecodeInt64() != -9223372036854775808 {
		t.Error("decode fail -9223372036854775808")
	}
}

func Test_Decode_CString(t *testing.T) {
	b := NewBinaryStream()
	b.EncodeCString("HELLO")
	b.EncodeCString("WORLD")
	b.EncodeCString("the go programming language")

	d := NewBinaryStreamWithByteArray(b.GetBytes())

	if d.DecodeCString() != "HELLO" {
		t.Error("DecodeC HELLO")
	}

	if d.DecodeCString() != "WORLD" {
		t.Error("DecodeC WORLD")
	}

	if d.DecodeCString() != "the go programming language" {
		t.Error("DecodeC the go programming language")
	}
}

func Test_Decode_Length_String(t *testing.T) {
	b := NewBinaryStream()
	b.EncodeUInt16LengthString("HELLO")
	b.EncodeUInt16LengthString("WORLD")
	b.EncodeUInt16LengthString("the go programming language")

	d := NewBinaryStreamWithByteArray(b.GetBytes())

	if d.DecodeUInt16LengthString() != "HELLO" {
		t.Error("Decode HELLO")
	}

	if d.DecodeUInt16LengthString() != "WORLD" {
		t.Error("Decode WORLD")
	}

	if d.DecodeUInt16LengthString() != "the go programming language" {
		t.Error("Decode the go programming language")
	}
}

func Test_Seek(t *testing.T) {
	b := NewBinaryStreamWithSize(32)
	b.EncodeUInt8(4)
	o := b.Offset()
	for i := 0; i < 10; i++ {
		b.EncodeUInt8(4)
	}
	l := b.Offset()
	b.Seek(o)
	b.EncodeUInt8(1)
	b.Seek(l)
	if !checkbyteArrayEquals(b.GetBytes(),
		[]byte{4, 1, 4, 4, 4, 4, 4, 4, 4, 4, 4}) {
		t.Error("seek fail")
	}
}
