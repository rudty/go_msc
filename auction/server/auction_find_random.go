package main

import (
	"math/rand"
)

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
func (a *AuctionServer) FindRandomItems(req *FindRandomItemRequest, res *AuctionItemResponse) error {

	if len(a.pkAuctionIDItems) == 0 {
		return nil
	}

	a.lock.RLock()
	defer a.lock.RUnlock()

	auctionItemCount := len(a.pkAuctionIDItems)
	extractCount := minInt(int(req.Count), auctionItemCount)
	extractIndexes := make([]int, extractCount)
	for i := 0; i < extractCount; i++ {
		extractIndexes[i] = rand.Intn(auctionItemCount)
	}

	auctionIt := 0
	extractIt := 0
	for _, v := range a.pkAuctionIDItems {
		if auctionIt == extractIndexes[extractIt] {
			res.Items = append(res.Items, v)
			extractIt++
			if extractIt == extractCount {
				break
			}
		}
		auctionIt++
	}
	return nil
}
