package inventory

import (
	"testing"
)

func Test_Serialize_Int8(t *testing.T) {
	b := NewBinaryStream()
	for i := 1; i < 11; i++ {
		b.EncodeInt8(int8(i))
	}

	if len(b.buf) != 10 {
		t.Error("must not grow")
	}
}
