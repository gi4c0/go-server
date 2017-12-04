package article

import (
	"database/sql"
	"go-server/db"
	"github.com/go-sql-driver/mysql"
	//"fmt"
)

type NewArticle struct {
	ArticleId int
	Text string
	Title string
	UserId int
	Image string
}

type FetchedArticle struct {
	ArticleId int
	Text string
	Title string
	UserName string
	Approved sql.NullBool
	Image sql.NullString
}

func Create (article *NewArticle) (bool, string) {
	query := "INSERT INTO test.Articles (Text, Title, Image, UserId) VALUES (?, ?, ?, ?)"

	_, err := db.Con.Exec(query, &article.Text, &article.Title, &article.Image, &article.UserId)
	if err != nil {
		me, ok := err.(*mysql.MySQLError)
		if ok && me.Number == 1062 {
			return false, "This title already exists"
		}
		return false, err.Error()
	}

	return true, ""
}

func GetAll (skip int, limit int) ([]FetchedArticle, string) {
	var fetchedArticles []FetchedArticle

	res, err := db.Con.Query("SELECT * FROM test.Articles LIMIT ?, ?", skip, limit)
	if err != nil {
		return nil, err.Error()
	}

	for res.Next() {
		var fa FetchedArticle
		var userId int
		scanErr := res.Scan(&fa.ArticleId, &fa.Text, &fa.Title, &userId, &fa.Approved, &fa.Image)

		db.Con.QueryRow("SELECT Username FROM test.Users WHERE UserId = ?", userId).Scan(&fa.UserName)

		if scanErr != nil {
			return nil, scanErr.Error()
		}

		fetchedArticles = append(fetchedArticles, fa)
	}

	return fetchedArticles, ""
}