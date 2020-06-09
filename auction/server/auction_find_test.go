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

func TestFindByItemID(t *testing.T) {
	s := NewAuctionService()
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
	s := NewAuctionService()
	res := AuctionItemResponse{}
	s.FindItemByItemID(&FindItemByItemIDRequest{
		ItemID: 1,
		Count:  10,
	}, &res)

	if len(res.Items) != 0 {
		t.Error("must 0")
	}
}

func TestFindByIdNotFound(t *testing.T) {
	s := NewAuctionService()
	var req UniqueID = 9912
	var res = AuctionItem{}
	if err := s.FindItemByAuctionID(&req, &res); err == nil {
		t.Error("must not found")
	}
}

func TestFindByIdOK(t *testing.T) {
	s := NewAuctionService()
	var uniqueID UniqueID
	if err := s.RegisterItem(&AuctionRegisterItemRequest{
		ItemID:   ItemID(1),
		BidPrice: int64(30),
	}, &uniqueID); err != nil {
		t.Fatal(err)
	}
	var res = AuctionItem{}
	if err := s.FindItemByAuctionID(&uniqueID, &res); err != nil {
		t.Error("must found")
	}
}
