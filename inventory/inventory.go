package inventory

import "sync"

type AddItemRequest struct {
	Item
}

//X 아이템을 관리하는 구조체
//Lock으로 관리하지 않습니다
type X struct {
	storage map[int64]*Item
}

//AddItem 아이템을 추가합니다
func (x *X) AddItem(req AddItemRequest) {
	x.storage[req.ID] = &req.Item
}

//RemoveItemByID 아이템을 추가합니다
func (x *X) RemoveItemByID(id int64) {
	delete(x.storage, id)
}

//GetItem 해당 아이디의 아이템을 반환합니다
func (x *X) GetItem(id int64) (*Item, bool) {
	item, ok := x.storage[id]
	return item, ok
}

// Inventory 실제로 저장할 인벤토리 정보
type Inventory struct {
	lock  sync.Mutex
	items X
}

func NewInventory() *Inventory {
	e := &Inventory{}
	e.items.storage = make(map[int64]*Item, 64)
	return e
}

func (inv *Inventory) AddItem(req AddItemRequest) {
	inv.lock.Lock()
	defer inv.lock.Unlock()
	inv.items.AddItem(req)
}

func (inv *Inventory) RemoveItemByID(id int64) {
	inv.lock.Lock()
	defer inv.lock.Unlock()
	inv.items.RemoveItemByID(req)
}

func (inv *Inventory) GetItemThen(id int64, cb func(*Item)) {
	inv.lock.Lock()
	defer inv.lock.Unlock()
	item, ok := inv.items.GetItem(id)
	if ok {
		cb(item)
	}
}

func (inv *Inventory) Transaction(cb func(*X)) {
	inv.lock.Lock()
	defer inv.lock.Unlock()
	cb(inv.items)
}
