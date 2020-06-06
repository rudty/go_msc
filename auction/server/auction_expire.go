package main

import (
	"fmt"
	"time"
)

func getFirstExpireItem(a *AuctionServer) *AuctionItem {
	a.lock.RLock()
	defer a.lock.RUnlock()
	return a.indexExpireTime.Front().Value.(*AuctionItem)
}

func handleExpire(a *AuctionServer) {
	for {
		if a.indexExpireTime.Len() > 0 {
			nowSec := time.Now().Unix()
			firstItem := getFirstExpireItem(a)

			if firstItem.ExpireTime < nowSec {
				a.lock.Lock()
				defer a.lock.Unlock()
				delete(a.pkAuctionIDItems, firstItem.AuctionID)
				delete(a.indexItemIDitems[firstItem.ItemID], firstItem.AuctionID)
				a.indexExpireTime.Remove(a.indexExpireTime.Front())
				continue
			}
		}
		time.Sleep(1 * time.Second)
	}
}

func onExpireItem(a *AuctionServer, item *AuctionItem) {
	// 여기서 만료되는 아이템을 사용한 사람이 있다면 전달
	if item.BidUserID != nil {
		fmt.Println(item)
	}
}

func startExpire(a *AuctionServer) {
	go handleExpire(a)
}
