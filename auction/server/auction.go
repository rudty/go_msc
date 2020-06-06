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

// AuctionSevice 경매 서버
type AuctionSevice struct {
	lock             sync.RWMutex
	pkAuctionIDItems map[UniqueID]*AuctionItem
	indexItemIDitems map[ItemID]map[UniqueID]*AuctionItem
	indexExpireTime  *list.List
}

// NewAuctionService 새로운 경매 서버를 만듭니다.
func NewAuctionService() *AuctionSevice {
	a := &AuctionSevice{
		pkAuctionIDItems: make(map[UniqueID]*AuctionItem),
		indexItemIDitems: make(map[ItemID]map[UniqueID]*AuctionItem),
		indexExpireTime:  list.New(),
	}
	handleExpire(a)
	return a
}

func (a *AuctionSevice) findItemByUniqueID(id UniqueID) *AuctionItem {
	a.lock.Lock()
	defer a.lock.Unlock()
	if e, ok := a.pkAuctionIDItems[id]; ok {
		return e
	}
	return nil
}
