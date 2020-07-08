package inventory

import (
	"fmt"
	"testing"
)

func newTestItem(id int64) Item {
	return Item{
		ID:         id,
		ItemID:     3,
		SlotIndex:  2,
		Properties: "Hello",
	}
}
func TestAddItem(t *testing.T) {
	inv := NewInventory()
	inv.Add(&AddItemRequest{newTestItem(3)})
	_, ok := inv.items.storage[3]
	if !ok {
		t.Error("insert fail")
	}
}

func TestRemoveItem(t *testing.T) {
	inv := NewInventory()
	inv.Add(&AddItemRequest{newTestItem(3)})
	_, ok := inv.items.storage[3]
	if !ok {
		t.Error("insert fail")
	}

	inv.RemoveByID(3)
	_, ok = inv.items.storage[3]
	if ok {
		t.Error("remove fail")
	}
}

func TestItemTransaction(t *testing.T) {
	inv := NewInventory()
	inv.Transaction(func(x *X) {
		for i := 0; i < 10; i++ {
			x.Add(&AddItemRequest{newTestItem(int64(i))})
		}
	})
}

func TestItemGetThen(t *testing.T) {
	inv := NewInventory()
	inv.Add(&AddItemRequest{newTestItem(int64(3))})
	inv.GetThen(3, func(item *Item) {
		item.ItemID = 7777
	})

	r, _ := inv.items.storage[3]
	if r.ItemID != 7777 {
		fmt.Println("must modify")
	}

}
