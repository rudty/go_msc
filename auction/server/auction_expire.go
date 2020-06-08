package main

import (
	"log"
	"time"
)

func getExpireItems(a *AuctionSevice, nowSec int64) ([]*AuctionItem, error) {
	a.lock.RLock()
	defer a.lock.RUnlock()

	rows, err := a.db.Query(
		selectAuctionWhereExpire,
		nowSec,
	)

	if rows != nil {
		defer rows.Close()
	}

	if err != nil {
		return nil, err
	}

	expireItems := make([]*AuctionItem, 0, 10)

	for rows.Next() {
		if rows.Err() != nil {
			return nil, rows.Err()
		}

		e := AuctionItem{}

		if err := e.ReadFromSQL(rows); err != nil {
			return nil, err
		}

		expireItems = append(expireItems, &e)
	}

	return expireItems, nil
}

func removeExpireItems(a *AuctionSevice, nowSec int64) error {
	a.lock.Lock()
	defer a.lock.Unlock()
	_, err := a.db.Exec("delete from AuctionItem where ExpireTime < ?", nowSec)
	if err != nil {
		return err
	}
	return nil
}

func handleExpire(a *AuctionSevice) int {
	nowSec := time.Now().Unix()
	expireItems, err := getExpireItems(a, nowSec)
	if err != nil {
		log.Println("expireItem error: ", err)
		return 0
	}

	if err := removeExpireItems(a, nowSec); err != nil {
		log.Println("remove expireItem error: ", err)
		return 0
	}

	for _, e := range expireItems {
		if len(e.BidUserID) > 0 {
			log.Println(e.BidUserID + " get item")
		}
	}
	return len(expireItems)
}

func startExpire(a *AuctionSevice) {
	go func() {
		for {
			time.Sleep(1 * time.Second)
			handleExpire(a)
		}
	}()
}