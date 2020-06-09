package main

import (
	"database/sql"
	"errors"
	"fmt"
)

// FindRandomItemRequest 랜덤 하게 아이템을 가져오는 요청
type FindRandomItemRequest struct {
	Count int16
}

// AuctionItemResponse 공용 아이템 응답
type AuctionItemResponse struct {
	Items []*AuctionItem
}

// FindRandomItems 맨 처음에 보여주는 용도로 랜덤한 아이템 몇개를 가져옵니다.
func (a *AuctionSevice) FindRandomItems(req *FindRandomItemRequest, res *AuctionItemResponse) error {
	a.lock.RLock()
	defer a.lock.RUnlock()

	rows, err := a.db.Query(
		selectAuctionOrderByRandomLimitX,
		req.Count)
	if rows != nil {
		defer rows.Close()
	}

	if err != nil {
		return err
	}

	res.Items, err = NewAuctionItemListFromSQLRows(rows)
	if err != nil {
		return err
	}

	return nil
}

// FindItemByItemIDRequest 아이템 아이디로 아이템을 찾는 요청
type FindItemByItemIDRequest struct {
	ItemID ItemID
	Count  int16
}

// FindItemByItemID 아이템 아이디에 해당하는
func (a *AuctionSevice) FindItemByItemID(req *FindItemByItemIDRequest, res *AuctionItemResponse) error {
	a.lock.RLock()
	defer a.lock.RUnlock()

	rows, err := a.db.Query(
		selectAuctionWhereItemIDLimitX,
		req.ItemID,
		req.Count,
	)

	if rows != nil {
		defer rows.Close()
	}

	if err != nil {
		return err
	}

	res.Items, err = NewAuctionItemListFromSQLRows(rows)
	if err != nil {
		return err
	}

	return nil
}

// FindItemByAuctionID 해당 유니크 아이디에 해당하는 아이템을 반환합니다.
func (a *AuctionSevice) FindItemByAuctionID(req *UniqueID, res *AuctionItem) error {
	tx, err := a.db.Begin()
	if err != nil {
		return err
	}
	a.lock.RLock()
	defer a.lock.RUnlock()

	item, err := findAuctionItemByAuctionID(tx, *req)

	if err != nil {
		return err
	}

	*res = *item

	return nil
}

func findAuctionItemByAuctionID(tx *sql.Tx, auctionID UniqueID) (*AuctionItem, error) {
	rows, err := tx.Query(selectAuctionWhereAuctionID, auctionID)

	if rows != nil {
		defer rows.Close()
	}

	if err != nil {
		return nil, err
	}

	if !rows.Next() {
		return nil, errors.New(fmt.Sprint("cannot found auctionID:", auctionID))
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	item, err := NewAuctionItemFromSQLRow(rows)
	if err != nil {
		return nil, err
	}

	return item, nil
}
