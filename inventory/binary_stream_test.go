package inventory

import (
	"testing"
)

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
