package db

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/go-sql-driver/mysql"
)

//import _ "github.com/go-sql-driver/mysql"

var Con, err = sqlx.Connect("mysql", "root:1111@tcp(localhost:3306)/test")

func init() {
	if err != nil {
		panic(err)
	}
}

