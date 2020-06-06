package main

import (
	"testing"
)

func TestFindByItemID(t *testing.T) {
	s := NewAuctionServer()
	var ignore UniqueID
	for i := 0; i < 100; i++ {
		if err := s.RegisterItem(&AuctionRegisterItemRequest{
			ItemID:   ItemID(1),
			BidPrice: int64(30),
		}, &ignore); err != nil {
			t.Fatal(err)

		}
	}
	res := AuctionItemResponse{}
	s.FindItemByItemID(&FindItemByItemIDRequest{
		ItemID: 1,
		Count:  10,
	}, &res)

	if len(res.Items) != 10 {
		t.Error("must 10")
	}
}

func TestFindByItemIDEmptyResult(t *testing.T) {
	s := NewAuctionServer()
	res := AuctionItemResponse{}
	s.FindItemByItemID(&FindItemByItemIDRequest{
		ItemID: 1,
		Count:  10,
	}, &res)

	if len(res.Items) != 0 {
		t.Error("must 0")
	}
}
