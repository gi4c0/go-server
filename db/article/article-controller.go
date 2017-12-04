package article

import (
	"database/sql"
	"go-server/db"
	"github.com/go-sql-driver/mysql"
)

type NewArticle struct {
	ArticleId int
	Text string
	Title string
	UserId int
	Approved sql.NullBool
	Image string
}

func Create (article *NewArticle) (bool, string) {
	query := "INSERT INTO test.Articles (Text, Title, Approved, Image, UserId) VALUES (?, ?, ?, ?, ?)"

	_, err := db.Con.Exec(query, &article.Text, &article.Title, &article.Approved, &article.Image, &article.UserId)
	if err != nil {
		me, ok := err.(*mysql.MySQLError)
		if ok && me.Number == 1062 {
			return false, "This title already exists"
		}
		return false, err.Error()
	}

	return true, ""
}
