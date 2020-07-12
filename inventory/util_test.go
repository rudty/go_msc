package inventory

import (
	"testing"
)

func TestRandomID(t *testing.T) {
	s := make(map[int64]bool, 20)
	for i := 0; i < 20; i++ {
		s[newID()] = true
	}

	if len(s) < 15 {
		t.Error("unique id")
	}
}
