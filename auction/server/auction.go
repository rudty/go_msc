package main

import (
	"container/list"
	"fmt"
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
	BidUserID  *string
}

func (a *AuctionItem) String() string {
	return fmt.Sprintf("{AuctionID: %v, ItemID:%v, BidPrice: %v, ExpireTime: %v}",
		a.AuctionID,
		a.ItemID,
		a.BidPrice,
		a.ExpireTime)
}

// AuctionServer 경매 서버
type AuctionServer struct {
	lock             sync.RWMutex
	pkAuctionIDItems map[UniqueID]*AuctionItem
	indexItemIDitems map[ItemID]map[UniqueID]*AuctionItem
	indexExpireTime  *list.List
}

// NewAuctionServer 새로운 경매 서버를 만듭니다.
func NewAuctionServer() *AuctionServer {
	return &AuctionServer{
		pkAuctionIDItems: make(map[UniqueID]*AuctionItem),
		indexItemIDitems: make(map[ItemID]map[UniqueID]*AuctionItem),
		indexExpireTime:  list.New(),
	}
}

func (a *AuctionServer) findItemByUniqueID(id UniqueID) *AuctionItem {
	a.lock.Lock()
	defer a.lock.Unlock()
	if e, ok := a.pkAuctionIDItems[id]; ok {
		return e
	}
	return nil
}
