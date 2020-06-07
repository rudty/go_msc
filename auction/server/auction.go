package main

import (
	"database/sql"
	"fmt"
	"sync"

	//sqlite
	_ "github.com/mattn/go-sqlite3"
)

// ItemID 아이템 아이디 타입
type ItemID int32

// AuctionItem 관리하는 아이템
type AuctionItem struct {
	AuctionID  UniqueID
	ItemID     ItemID
	BidPrice   int64
	ExpireTime int64
	BidUserID  string
}

// SQLRowScannable AuctionItem을 SQL로부터 읽을수 있게 하는 공용 인터페이스
type SQLRowScannable interface {
	Scan(dest ...interface{}) error
}

// ReadFromSQL scan interface 로부터 값을 읽습니다.
func (a *AuctionItem) ReadFromSQL(s SQLRowScannable) error {
	if err := s.Scan(
		&a.AuctionID,
		&a.ItemID,
		&a.BidPrice,
		&a.ExpireTime,
		&a.BidUserID,
	); err != nil {
		return err
	}
	return nil
}

func (a *AuctionItem) String() string {
	return fmt.Sprintf("{AuctionID: %v, ItemID:%v, BidPrice: %v, ExpireTime: %v}",
		a.AuctionID,
		a.ItemID,
		a.BidPrice,
		a.ExpireTime)
}

// AuctionSevice 경매 서버
type AuctionSevice struct {
	lock sync.RWMutex
	db   *sql.DB
}

func createInMemoryAuctionTable() *sql.DB {
	db, err := sql.Open("sqlite3", "file::memory:?mode=memory&cache=shared")
	if err != nil {
		return nil
	}

	db.Exec(`
		drop table if exists AuctionItem;
		create table AuctionItem(
			AuctionID integer not null primary key,
			ItemID integer not null,
			BidPrice integer not null,
			ExpireTime integer not null,
			BidUserID varchar(24) not null
		);
	`)
	db.Exec(`create index IDX_ItemPrice on AuctionItem (ItemID, BidPrice);`)
	db.Exec(`create index IDX_Expire on AuctionItem (ExpireTime);`)
	return db
}

// NewAuctionService 새로운 경매 서버를 만듭니다.
func NewAuctionService() *AuctionSevice {
	a := &AuctionSevice{}
	a.db = createInMemoryAuctionTable()
	return a
}

func (a *AuctionSevice) findItemByUniqueID(id UniqueID) *AuctionItem {
	a.lock.Lock()
	defer a.lock.Unlock()
	row := a.db.QueryRow("select AuctionID, ItemID, BidPrice, ExpireTime, BidUserID from AuctionItem where AuctionID = ? limit 1;", id)
	if row == nil {
		return nil
	}
	auctionItem := AuctionItem{}
	if err := auctionItem.ReadFromSQL(row); err != nil {
		return nil
	}
	return &auctionItem
}
