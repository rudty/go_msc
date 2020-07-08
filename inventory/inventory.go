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

//Add 아이템을 추가합니다
func (x *X) Add(req *AddItemRequest) {
	x.storage[req.ID] = &req.Item
}

//RemoveByID 아이템을 추가합니다
func (x *X) RemoveByID(id int64) {
	delete(x.storage, id)
}

//GetItem 해당 아이디의 아이템을 반환합니다
func (x *X) Get(id int64) (*Item, bool) {
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

func (inv *Inventory) Add(req *AddItemRequest) {
	inv.lock.Lock()
	defer inv.lock.Unlock()
	inv.items.Add(req)
}

func (inv *Inventory) RemoveByID(id int64) {
	inv.lock.Lock()
	defer inv.lock.Unlock()
	inv.items.RemoveByID(id)
}

func (inv *Inventory) GetThen(id int64, cb func(*Item)) {
	inv.lock.Lock()
	defer inv.lock.Unlock()
	item, ok := inv.items.Get(id)
	if ok {
		cb(item)
	}
}

func (inv *Inventory) Transaction(cb func(*X)) {
	inv.lock.Lock()
	defer inv.lock.Unlock()
	cb(&inv.items)
}
