package main

import (
	"testing"
)

func TestFindRandomItem(t *testing.T) {
	s := NewAuctionService()
	var ignore UniqueID
	for i := 0; i < 100; i++ {
		if err := s.RegisterItem(&AuctionRegisterItemRequest{
			ItemID:   ItemID(i),
			BidPrice: int64(i * 30),
		}, &ignore); err != nil {
			t.Fatal(err)
		}
	}

	res := AuctionItemResponse{}
	if err := s.FindRandomItems(&FindRandomItemRequest{Count: 2}, &res); err != nil {
		t.Fatal(err)
	}

	if len(res.Items) != 2 {
		t.Fatal("length == 2")
	}
}

func TestFindRandomEmptyItem(t *testing.T) {
	s := NewAuctionService()
	res := AuctionItemResponse{}
	if err := s.FindRandomItems(&FindRandomItemRequest{Count: 2}, &res); err != nil {
		t.Fatal(err)
	}
	if len(res.Items) != 0 {
		t.Fatal("size == 0")
	}
}
