package main

import (
	"errors"
	"math/rand"
)

// AuctionItem 관리하는 아이템
type AuctionItem struct {
	AuctionID int64
	ItemID    int32
	BidPrice  int64
}

// AuctionServer 경매 서버
type AuctionServer struct {
	items []*AuctionItem
}

// FindRandomItemRequest 랜덤 하게 아이템을 가져오는 요청
type FindRandomItemRequest struct {
	Count int16
}

// AuctionItemResponse 공용 아이템 응답
type AuctionItemResponse struct {
	Items []*AuctionItem
}

// FindRandomItems 맨 처음에 보여주는 용도로 랜덤한 아이템 몇개를 가져옵니다.
func (a *AuctionServer) FindRandomItems(req *FindRandomItemRequest, res *AuctionItemResponse) error {

	if len(a.items) == 0 {
		return errors.New("empty")
	}

	var count int = int(req.Count)
	for i := 0; i < count; i++ {
		idx := rand.Intn(len(a.items))
		res.Items[i] = a.items[idx]
	}

	return nil
}

// RegisterItem 새로운 아이템을 반환합니다. 반환: 새로운 아이템의 AuctionID
func (a *AuctionServer) RegisterItem(req *AuctionItem, res *int64) error {
	newAuctionID := getAuctionID()
	*res = newAuctionID
	req.AuctionID = newAuctionID
	a.items = append(a.items, req)
	return nil
}
