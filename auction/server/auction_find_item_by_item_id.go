package main

// FindItemByItemIDRequest 아이템 아이디로 아이템을 찾는 요청
type FindItemByItemIDRequest struct {
	ItemID ItemID
	Count  int16
}

// FindItemByItemID 아이템 아이디에 해당하는
func (a *AuctionServer) FindItemByItemID(req *FindItemByItemIDRequest, res *AuctionItemResponse) error {
	if m, ok := a.indexItemIDitems[req.ItemID]; ok {
		a.lock.RLock()
		defer a.lock.RUnlock()

		count := minInt(int(req.Count), len(m))
		it := 0
		for _, e := range m {
			res.Items = append(res.Items, e)
			it++
			if it == count {
				break
			}
		}
	}
	return nil
}
