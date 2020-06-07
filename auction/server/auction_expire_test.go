package main

import (
	"testing"
	"time"
)

func TestExpireOK(t *testing.T) {
	a := NewAuctionService()
	if _, err := a.db.Exec("insert into AuctionItem values(?,1,2,3,'testuser');",
		getAuctionID(),
	); err != nil {
		t.Fatal(err)
	}

	if handleExpire(a) == 0 {
		t.Error("must remove")
	}
}

func TestExpireNotRemove(t *testing.T) {
	a := NewAuctionService()
	nowSec := time.Now().Unix()
	if _, err := a.db.Exec("insert into AuctionItem values(?,1,2,?,'testuser');",
		getAuctionID(),
		nowSec+3000,
	); err != nil {
		t.Fatal(err)
	}

	if handleExpire(a) != 0 {
		t.Error("must not remove")
	}
}
