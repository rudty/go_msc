package main

// BidRequest 올라간 아이템에 대해서 입찰을 요청합니다.
// 반드시 현재 등록된 금액보다 커야 등록이 가능합니다
type BidRequest struct {
	UserID    string
	Price     int64
	AuctionID UniqueID
}

// Bid 아이템에 대해서 입찰을 요청합니다.
func (a *AuctionSevice) Bid(req *BidRequest, res *bool) error {
	*res = false
	r, err := a.db.Exec(
		"update AuctionItem set BidPrice = ?, BidUserID = ? where AuctionID = ? and BidPrice < ?;",
		req.Price,
		req.UserID,
		req.AuctionID,
		req.Price)

	if err != nil {
		return err
	}

	modifiedCount, err := r.RowsAffected()
	if err != nil {
		return err
	}

	if modifiedCount == 1 {
		*res = true
	}

	return nil
}
