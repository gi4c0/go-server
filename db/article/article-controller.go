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
	Image string
}

type FetchedArticle struct {
	ArticleId int
	Text string
	Title string
	Username string
	Approved bool
	Image string
}

func Create (article *NewArticle) (bool, string) {
	query := "INSERT INTO test.Articles (Text, Title, Image, UserId) VALUES (?, ?, ?, ?)"

	_, err := db.Con.Exec(query, article.Text, article.Title, article.Image, article.UserId)
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

	query := `
	SELECT Articles.ArticleId, Articles.Text, Articles.Title, Articles.Approved, Articles.Image, Users.Username
  	FROM test.Articles
  	JOIN test.Users ON Users.UserId = Articles.UserId
  	LIMIT ?, ?
  	`

	res, err := db.Con.Query(query, skip, limit)
	if err != nil {
		return nil, err.Error()
	}

	for res.Next() {
		var fa FetchedArticle
		var fetchedImage sql.NullString
		scanErr := res.Scan(&fa.ArticleId, &fa.Text, &fa.Title, &fa.Approved, &fetchedImage, &fa.Username)


		if scanErr != nil {
			return nil, scanErr.Error()
		}

		fa.Image = fetchedImage.String
		fetchedArticles = append(fetchedArticles, fa)
	}

	return fetchedArticles, ""
}

func Update (article *NewArticle) (bool, string) { // TODO: Delete images
	query := "UPDATE test.Articles SET Text = ?, Title = ?, Image = ? WHERE UserId = ? AND ArticleId = ?"

	_, err := db.Con.Exec(query, article.Text, article.Title, article.Image, article.UserId, article.ArticleId)
	if err != nil {
		return false, err.Error()
	}

	return true, ""
}
