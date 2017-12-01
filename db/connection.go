package db

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
)

var Con, err = sql.Open("mysql", "root:1111@tcp(localhost:3306)/test")

func init() {
	if err != nil {
		panic(err)
	}
}

