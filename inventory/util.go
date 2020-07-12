package inventory

import (
	"crypto/rand"
	"encoding/binary"
)

func newID() int64 {
	var b [8]byte
	rand.Read(b[:])
	return int64(binary.LittleEndian.Uint64(b[:]))
}
