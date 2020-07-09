package inventory

import (
	"C"
)

// Item 인벤토리에 들어가는 아이템 정보
type Item struct {
	ID         int64
	ItemID     int32
	SlotIndex  int32
	Properties string
}

func NewItemFromBuffer(buf []byte) (*Item, error) {
	e := Item{}
	b := NewBinaryStreamWithByteArray(buf)

	e.ID = b.DecodeInt64()
	e.ItemID = b.DecodeInt32()
	e.SlotIndex = b.DecodeInt32()
	e.Properties = b.DecodeUInt16LengthString()

	return &e, nil
}
