package main

import (
	"log"
)

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
	tx, err := a.db.Begin()

	if tx != nil {
		defer tx.Commit()
	}

	if err != nil {
		return err
	}

	a.lock.Lock()
	defer a.lock.Unlock()

	rows, err := tx.Query(selectBidUserIDByAuctionID, req.AuctionID)

	if rows != nil {
		defer rows.Close()
	}

	if err != nil {
		return err
	}

	if rows.Next() {
		if err := rows.Err(); err != nil {
			return err
		}

		var oldBidUserID string
		rows.Scan(&oldBidUserID)
		if len(oldBidUserID) > 0 {
			log.Println(" bid", oldBidUserID)
		}
	}

	r, err := tx.Exec(
		updateBidPriceAndBidUserIDWhereAuctionIDBidPrice,
		req.Price,
		req.UserID,
		req.AuctionID,
		req.Price)

	modifiedCount, err := r.RowsAffected()
	if err != nil {
		return err
	}

	if modifiedCount == 1 {
		*res = true
	}

	return nil
}
