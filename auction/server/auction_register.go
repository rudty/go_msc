package main

import "time"

// AuctionRegisterItemRequest 아이템 등록 요청
type AuctionRegisterItemRequest struct {
	ItemID   ItemID
	BidPrice int64
}

// RegisterItem 새로운 아이템을 반환합니다. 반환: 새로운 아이템의 AuctionID
func (a *AuctionServer) RegisterItem(req *AuctionRegisterItemRequest, res *UniqueID) error {
	newAuctionID := getAuctionID()
	expireTime := time.Now().Unix() + 3600
	newAuctionItem := &AuctionItem{
		ItemID:     req.ItemID,
		BidPrice:   req.BidPrice,
		ExpireTime: expireTime,
		AuctionID:  newAuctionID,
	}

	*res = newAuctionID

	a.lock.Lock()
	defer a.lock.Unlock()

	a.pkAuctionIDItems[newAuctionID] = newAuctionItem
	a.indexItemIDitems[req.ItemID] = append(a.indexItemIDitems[req.ItemID], newAuctionItem)

	return nil
}
