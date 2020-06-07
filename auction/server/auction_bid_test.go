package main

import (
	"testing"
)

func registerItemPrice30(s *AuctionSevice) UniqueID {
	var auctionID UniqueID
	if err := s.RegisterItem(&AuctionRegisterItemRequest{
		ItemID:   ItemID(1),
		BidPrice: int64(30),
	}, &auctionID); err != nil {
		panic(err)
	}
	return auctionID
}

func TestBidOK(t *testing.T) {
	s := NewAuctionService()
	auctionID := registerItemPrice30(s)
	var res bool
	s.Bid(&BidRequest{
		UserID:    "a",
		Price:     31,
		AuctionID: auctionID,
	}, &res)

	if !res {
		t.Fatal("must success")
	}
}

func TestBidFail30(t *testing.T) {
	s := NewAuctionService()
	auctionID := registerItemPrice30(s)
	var res bool
	s.Bid(&BidRequest{
		UserID:    "a",
		Price:     30,
		AuctionID: auctionID,
	}, &res)

	if res {
		t.Fatal("must fail")
	}
}

func TestBidFailNotItem(t *testing.T) {
	s := NewAuctionService()
	var res bool
	s.Bid(&BidRequest{
		UserID:    "a",
		Price:     30,
		AuctionID: 999999999999999999,
	}, &res)

	if res {
		t.Fatal("must fail")
	}
}

func TestBidUserChange(t *testing.T) {
	s := NewAuctionService()
	auctionID := registerItemPrice30(s)
	var res bool
	if err := s.Bid(&BidRequest{
		UserID:    "a",
		Price:     31,
		AuctionID: auctionID,
	}, &res); err != nil {
		t.Error(err)
	}

	if err := s.Bid(&BidRequest{
		UserID:    "b",
		Price:     32,
		AuctionID: auctionID,
	}, &res); err != nil {
		t.Error(err)
	}
}
