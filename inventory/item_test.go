package inventory

import (
	"testing"
)

type AA struct {
	P string
}

func TestItemFromByte(t *testing.T) {
	item := Item{}
	item.ID = 1
	item.ItemID = 2
	item.SlotIndex = 3
	item.Properties = "4"

	b := NewBinaryStream()
	b.EncodeInt64(item.ID)
	b.EncodeInt32(item.ItemID)
	b.EncodeInt32(item.SlotIndex)
	b.EncodeUInt16LengthString(item.Properties)

	newItem, _ := NewItemFromBuffer(b.GetBytes())

	if *newItem != item {
		t.Error("encode decode fail")
	}
}
