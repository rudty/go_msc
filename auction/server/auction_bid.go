package main

import (
	"errors"
	"log"
)

// BidRequest 올라간 아이템에 대해서 입찰을 요청합니다.
// 반드시 현재 등록된 금액보다 커야 등록이 가능합니다
type BidRequest struct {
	UserID    string
	Price     int64
	AuctionID UniqueID
}

// BidResponse 입찰 결과를 반환합니다.
// Success 입찰 트랜잭션이 성공했는지
// OldUserID 교체된 유저의 아이디(최초 입찰 시에는 '' 으로 반환)
// OldPrice 교체된 유저의 금액
type BidResponse struct {
	Success bool

	OldUserID string
	OldPrice  int64
}

// Bid 아이템에 대해서 입찰을 요청합니다.
func (a *AuctionSevice) Bid(req *BidRequest, res *BidResponse) error {
	res.Success = false
	tx, err := a.db.Begin()

	if tx != nil {
		defer tx.Commit()
	}

	if err != nil {
		return err
	}

	a.lock.Lock()
	defer a.lock.Unlock()

	item, err := findAuctionItemByAuctionID(tx, req.AuctionID)
	if err != nil {
		return err
	}

	if item.BidPrice > req.Price {
		return errors.New("item.BidPrice > req.Price")
	}

	res.OldPrice = item.BidPrice
	res.OldUserID = item.BidUserID

	log.Println(item, res)

	if _, err := tx.Exec(
		"update AuctionItem set BidPrice = ?, BidUserID = ? where AuctionID = ?;",
		req.Price,
		req.UserID,
		req.AuctionID); err != nil {
		return err
	}

	res.Success = true
	return nil
}
