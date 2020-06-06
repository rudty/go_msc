package main

// BidRequest 올라간 아이템에 대해서 입찰을 요청합니다.
// 반드시 현재 등록된 금액보다 커야 등록이 가능합니다
type BidRequest struct {
	UserID    string
	Price     int64
	AuctionID UniqueID
}

// Bid 아이템에 대해서 입찰을 요청합니다.
func (a *AuctionSevice) Bid(req *BidRequest, res *bool) {
	*res = false
	item := a.findItemByUniqueID(req.AuctionID)
	if item == nil {
		return
	}
	if req.Price > item.BidPrice {
		item.BidUserID = &req.UserID
		*res = true
	}
}
