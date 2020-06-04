package main

import (
	"sync/atomic"
)

var auctionID int64

// getAuctionID 새로운 아이디를 만들어서 반환합니다.
func getAuctionID() int64 {
	return atomic.AddInt64(&auctionID, 1)
}
