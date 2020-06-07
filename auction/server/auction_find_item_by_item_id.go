package main

// FindItemByItemIDRequest 아이템 아이디로 아이템을 찾는 요청
type FindItemByItemIDRequest struct {
	ItemID ItemID
	Count  int16
}

// FindItemByItemID 아이템 아이디에 해당하는
func (a *AuctionSevice) FindItemByItemID(req *FindItemByItemIDRequest, res *AuctionItemResponse) error {
	a.lock.RLock()
	defer a.lock.RUnlock()

	rows, err := a.db.Query(
		"select AuctionID, ItemID, BidPrice, ExpireTime, BidUserID from AuctionItem where ItemID = ? limit ?;",
		req.ItemID,
		req.Count,
	)

	if rows != nil {
		defer rows.Close()
	}

	if err != nil {
		return err
	}

	for rows.Next() {
		e := AuctionItem{}
		if rows.Err() != nil {
			return rows.Err()
		}
		e.ReadFromSQL(rows)
		res.Items = append(res.Items, &e)
	}

	return nil
}
