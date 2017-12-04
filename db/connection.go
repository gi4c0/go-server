package db

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"fmt"
)

var Con, err = sql.Open("mysql", "root:1111@tcp(localhost:3306)/test?multiStatements=true")

func init() {
	if err != nil {
		panic(err)
	}

	createTables := `
	CREATE TABLE if not exists test.Users (
		UserId INT NOT NULL AUTO_INCREMENT,
		Username VARCHAR(80) NOT NULL,
		Password VARCHAR(255) NOT NULL,
		Token VARCHAR(255) NULL,
		PRIMARY KEY (UserId),
		UNIQUE INDEX UserId_UNIQUE (UserId ASC),
		UNIQUE INDEX Username_UNIQUE (Username ASC)
  );CREATE TABLE if not exists test.Articles (
	  ArticleId INT NOT NULL AUTO_INCREMENT,
	  Text TEXT NOT NULL,
	  Title VARCHAR(255) NOT NULL,
	  UserId INT NOT NULL,
	  Approved TINYINT(1) NULL DEFAULT 0,
	  Image VARCHAR(255) NULL,
	  PRIMARY KEY (ArticleId),
	  UNIQUE INDEX ArticleID_UNIQUE (ArticleId ASC),
	  UNIQUE INDEX Title_UNIQUE (Title ASC)
  );
`


	_, err := Con.Exec(createTables)
	fmt.Println(err)
}

