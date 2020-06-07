package main

// FindRandomItemRequest 랜덤 하게 아이템을 가져오는 요청
type FindRandomItemRequest struct {
	Count int16
}

// AuctionItemResponse 공용 아이템 응답
type AuctionItemResponse struct {
	Items []*AuctionItem
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// FindRandomItems 맨 처음에 보여주는 용도로 랜덤한 아이템 몇개를 가져옵니다.
func (a *AuctionSevice) FindRandomItems(req *FindRandomItemRequest, res *AuctionItemResponse) error {
	a.lock.RLock()
	defer a.lock.RUnlock()

	rows, err := a.db.Query(
		"select AuctionID, ItemID, BidPrice, ExpireTime, BidUserID from AuctionItem order by random() limit ?;",
		req.Count)
	if rows != nil {
		defer rows.Close()
	}

	if err != nil {
		return err
	}

	for rows.Next() {
		if rows.Err() != nil {
			return rows.Err()
		}
		e := AuctionItem{}
		if err := e.ReadFromSQL(rows); err != nil {
			return err
		}
		res.Items = append(res.Items, &e)
	}

	return nil
}
