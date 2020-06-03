package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"testing"
	"time"

	// sql
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func init() {
	db, _ = sql.Open("mysql", "selectman:1234@tcp(127.0.0.1:3306)/my_database")
}
func TestDbSelect(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Nanosecond)
	defer cancel()
	fmt.Println(db)
	rows, err := db.QueryContext(ctx, "select 1 as a")
	log.Println(rows)
	log.Println(err)
}

func TestService(t *testing.T) {
	h := helloService{}
	tx, err := db.Begin()
	if err != nil {
		t.Fatal(err)
	}
	log.Println(h.hello(tx))
}
