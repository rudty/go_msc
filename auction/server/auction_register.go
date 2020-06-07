package main

import "time"

// AuctionRegisterItemRequest 아이템 등록 요청
type AuctionRegisterItemRequest struct {
	ItemID   ItemID
	BidPrice int64
}

// RegisterItem 새로운 아이템을 반환합니다. 반환: 새로운 아이템의 AuctionID
func (a *AuctionSevice) RegisterItem(req *AuctionRegisterItemRequest, res *UniqueID) error {
	newAuctionID := getAuctionID()
	expireTime := time.Now().Unix() + 3600
	*res = newAuctionID

	a.lock.Lock()
	defer a.lock.Unlock()

	if _, err := a.db.Exec("insert into AuctionItem values(?,?,?,?,'');",
		newAuctionID,
		req.ItemID,
		req.BidPrice,
		expireTime,
	); err != nil {
		return err
	}
	return nil
}
