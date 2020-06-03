package main

import "database/sql"

type helloService struct {
}

func (h *helloService) hello(tx *sql.Tx) map[string]interface{} {
	row := tx.QueryRow("select 1 as value")
	var value int
	row.Scan(&value)
	return map[string]interface{}{
		"a": value,
		"b": "c",
	}
}
