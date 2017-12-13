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
		Permission VARCHAR(40) DEFAULT "user",
		PRIMARY KEY (UserId),
		UNIQUE INDEX UserId_UNIQUE (UserId ASC),
		UNIQUE INDEX Username_UNIQUE (Username ASC)
  );
  CREATE TABLE if not exists test.Articles (
	  ArticleId INT NOT NULL AUTO_INCREMENT,
	  Text TEXT NOT NULL,
	  Title VARCHAR(255) NOT NULL,
	  UserId INT NOT NULL,
	  Approved TINYINT(1) NULL DEFAULT 0,
	  Image VARCHAR(255) NULL,
	  CategoryId INT NOT NULL,
	  CreatedAt DATETIME DEFAULT CURRENT_TIMESTAMP,
	  PRIMARY KEY (ArticleId),
	  UNIQUE INDEX ArticleID_UNIQUE (ArticleId ASC),
	  UNIQUE INDEX Title_UNIQUE (Title ASC)
  );
  CREATE TABLE IF NOT EXISTS test.Comments (
	CommentId INT NOT NULL AUTO_INCREMENT,
	ArticleId INT NOT NULL,
	UserId INT NOT NULL,
	ParentCommentId INT NULL,
	Text TEXT NOT NULL,
	CreatedAt DATETIME DEFAULT CURRENT_TIMESTAMP,
	PRIMARY KEY (CommentId),
	UNIQUE INDEX CommentId_UNIQUE (CommentId ASC),
	INDEX ArticleId_UNIQUE (ArticleId ASC),
	INDEX ParentCommentId_UNIQUE (ParentCommentId ASC)
  );
  CREATE TABLE IF NOT EXISTS test.Categories (
  	CategoryId INT NOT NULL AUTO_INCREMENT,
  	Name VARCHAR(40) NOT NULL,
  	PRIMARY KEY (CategoryId),
  	UNIQUE INDEX CategoryId_UNIQUE (CategoryId ASC),
  	UNIQUE (Name)
  );
`

	_, err := Con.Exec(createTables)
	if err != nil {
		fmt.Println(err)
	}
}

