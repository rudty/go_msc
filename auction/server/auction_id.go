package main

import (
	"sync/atomic"
)

// UniqueID 고유 아이디 인덱스
type UniqueID int64

var auctionID int64

// getAuctionID 새로운 아이디를 만들어서 반환합니다.
func getAuctionID() UniqueID {
	return UniqueID(atomic.AddInt64(&auctionID, 1))
}
