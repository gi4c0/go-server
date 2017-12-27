package article

import (
	"database/sql"
	"go-server/db"
	"github.com/go-sql-driver/mysql"
	"fmt"
)

type NewArticle struct {
	ArticleId int
	Text string
	Title string
	UserId int
	Image string
	CreatedAt string
	Category string
}

type FetchedArticle struct {
	ArticleId int
	Text string
	Title string
	Username string
	Approved bool
	Image string
	CreatedAt string
	Category string
}

type ArticlePreview struct {
  ArticleId int
  Title string
  Approved bool
}

func Create (article *NewArticle) (bool, string) {
	query := `INSERT INTO test.Articles (Text, Title, Image, Category, UserId) VALUES (?, ?, ?, "New", ?)`

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
		SELECT Articles.ArticleId, Articles.Text, Articles.Title, Articles.Approved, Articles.Image, Articles.CreatedAt, Articles.Category, Users.Username
		FROM test.Articles
		JOIN test.Users ON Users.UserId = Articles.UserId
		WHERE Approved = 1
		LIMIT ?, ?
  	`

	res, err := db.Con.Query(query, skip, limit)
	if err != nil {
		return nil, err.Error()
	}

	for res.Next() {
		var fa FetchedArticle
		var fetchedImage sql.NullString
		scanErr := res.Scan(&fa.ArticleId, &fa.Text, &fa.Title, &fa.Approved, &fetchedImage, &fa.CreatedAt, &fa.Category, &fa.Username)

		if scanErr != nil {
			return nil, scanErr.Error()
		}

		fa.Image = fetchedImage.String
		fetchedArticles = append(fetchedArticles, fa)
	}

	return fetchedArticles, ""
}

func GetSingleArticle (articleId int) (FetchedArticle, error) {
	var fa FetchedArticle
	var fetchedImage sql.NullString

	query := `
		SELECT Articles.ArticleId, Articles.Text, Articles.Title, Articles.Approved, Articles.Image, Articles.CreatedAt, Users.Username
  		FROM test.Articles
  		JOIN test.Users ON Users.UserId = Articles.UserId
  		WHERE ArticleId = ?
`
	scanErr := db.Con.QueryRow(query, articleId).Scan(&fa.ArticleId, &fa.Text, &fa.Title, &fa.Approved, &fetchedImage, &fa.CreatedAt, &fa.Username)
	if scanErr != nil {
		return fa, scanErr
	}

	fa.Image = fetchedImage.String
	return fa, nil
}

func Update (article *NewArticle) (bool, string) {
	query := "UPDATE test.Articles SET Text = ?, Title = ?, Image = ? WHERE UserId = ? AND ArticleId = ?"

	res, err := db.Con.Exec(query, article.Text, article.Title, article.Image, article.UserId, article.ArticleId)
	if err != nil {
		return false, err.Error()
	}

	rowsAffected, rowsErr := res.RowsAffected()
	if rowsErr != nil {
		return false, err.Error()
	}

	if rowsAffected < 1 {
		return false, "Wrong article id, or you do not have permission for this operation"
	}

	return true, ""
}

func DeleteImage (articleId, UserId int) (string, bool) {
	query := "UPDATE test.Articles SET Image = NULL WHERE ArticleId = ? AND UserId = ?"

	res, err := db.Con.Exec(query, articleId, UserId)
	if err != nil {
		return err.Error(), false
	}

	rowsAffected, rowsErr := res.RowsAffected()
	if rowsErr != nil {
		return err.Error(), false
	}

	if rowsAffected < 1 {
		return "Wrong article id, or you do not have permission for this operation", false
	}

	return "", true
}

func Approve(articleId int) error {
	query := "UPDATE test.Articles SET Approved = 1 WHERE ArticleId = ?"
	_, err := db.Con.Exec(query, articleId)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func GetUnapproved() ([]ArticlePreview, error) {
  var articles []ArticlePreview
  res, err := db.Con.Query("SELECT ArticleId, Title FROM test.Articles WHERE Approved = 0")
  if err != nil {
    fmt.Println(err)
  }

  for res.Next() {
    var article ArticlePreview

    scanErr := res.Scan(&article.ArticleId, &article.Title)
    if scanErr != nil {
      fmt.Println(scanErr)
      return articles, scanErr
    }

    articles = append(articles, article)
  }

  return articles, nil
}

func UserArticles (userId int) ([]ArticlePreview, error) {
  var articles []ArticlePreview

  res, dbErr:= db.Con.Query("SELECT ArticleId, Title, Approved FROM test.Articles WHERE UserId = ?", userId)
  if dbErr != nil {
    fmt.Println(dbErr)
    return articles, dbErr
  }

  for res.Next() {
    var article ArticlePreview

    scanErr := res.Scan(&article.ArticleId, &article.Title, &article.Approved)
    if scanErr != nil {
      fmt.Println(scanErr)
      return articles, scanErr
    }

    articles = append(articles, article)
  }

  return articles, nil
}
