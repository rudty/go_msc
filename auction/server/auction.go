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

// SQLRowsScannable []AuctionItem을 SQL로부터 읽을수 있게 하는 공용 인터페이스
type SQLRowsScannable interface {
	SQLRowScannable
	Next() bool
	Err() error
}

// NewAuctionItemFromSQLRow SQL을 scan하여 AuctionItem을 반환합니다.
func NewAuctionItemFromSQLRow(row SQLRowScannable) (*AuctionItem, error) {
	a := AuctionItem{}
	if err := row.Scan(
		&a.AuctionID,
		&a.ItemID,
		&a.BidPrice,
		&a.ExpireTime,
		&a.BidUserID,
	); err != nil {
		return nil, err
	}
	return &a, nil
}

// NewAuctionItemListFromSQLRows SQL을 scan 하여 []AuctionItem을 반환합니다.
func NewAuctionItemListFromSQLRows(rows SQLRowsScannable) ([]*AuctionItem, error) {
	items := make([]*AuctionItem, 0, 32)

	if err := rows.Err(); err != nil {
		return nil, err
	}

	for rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, err
		}

		item, err := NewAuctionItemFromSQLRow(rows)
		if err != nil {
			return nil, err
		}

		items = append(items, item)
	}

	return items, nil
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
