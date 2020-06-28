package main

import (
	"bank/api"
	"bank/migrations"
)

func main() {
	migrations.Migrate()
	api.StartApi()
}
