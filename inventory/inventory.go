package inventory

import "sync"

// Inventory 실제로 저장할 인벤토리 정보
type Inventory struct {
	lock    sync.Mutex
	storage map[int64]*Item
}

type AddItemRequest struct {
	Item
}

func (inv *Inventory) AddItem(req AddItemRequest) {
	inv.lock.Lock()
	defer inv.lock.Unlock()
	inv.storage[req.ID] = &req.Item
}

func (inv *Inventory) RemoveItemByID(id int64) {
	inv.lock.Lock()
	defer inv.lock.Unlock()
	delete(inv.storage, id)
}

func (inv *Inventory) GetItemThen(id int64, cb func(*Item)) {
	inv.lock.Lock()
	defer inv.lock.Unlock()
	item, ok := inv.storage[id]
	if ok {
		cb(item)
	}
}
