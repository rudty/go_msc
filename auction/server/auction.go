package main

import (
	"sync"
)

// ItemID 아이템 아이디 타입
type ItemID int32

// AuctionItem 관리하는 아이템
type AuctionItem struct {
	AuctionID  UniqueID
	ItemID     ItemID
	BidPrice   int64
	ExpireTime int64
}

// AuctionServer 경매 서버
type AuctionServer struct {
	lock             sync.RWMutex
	pkAuctionIDItems map[UniqueID]*AuctionItem
	indexItemIDitems map[ItemID][]*AuctionItem
}

// NewAuctionServer 새로운 경매 서버를 만듭니다.
func NewAuctionServer() *AuctionServer {
	return &AuctionServer{
		pkAuctionIDItems: make(map[UniqueID]*AuctionItem),
		indexItemIDitems: make(map[ItemID][]*AuctionItem),
	}
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

	// if len(a.items) == 0 {
	// return errors.New("empty")
	// }

	var count int = int(req.Count)
	for i := 0; i < count; i++ {
		// 	idx := rand.Intn(len(a.items))
		// 	res.Items[i] = a.items[idx]
	}

	return nil
}
