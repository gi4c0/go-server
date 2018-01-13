package article

import (
	"go-server/db"
	"github.com/go-sql-driver/mysql"
	"fmt"
)

type NewArticle struct {
	ArticleId int
	Text string
	Title string
	UserId int
	CreatedAt string
	Category string
}

type FetchedArticle struct {
	ArticleId int
	Text string
	Title string
	Username string
	Approved bool
	CreatedAt string
	Category string
}

type ArticlePreview struct {
  ArticleId int
  Title string
  Approved bool
  Username string
}

func Create (article *NewArticle) (bool, string) {
	query := `INSERT INTO test.Articles (Text, Title, Category, UserId) VALUES (?, ?, "New", ?)`

	_, err := db.Con.Exec(query, article.Text, article.Title, article.UserId)
	if err != nil {
		me, ok := err.(*mysql.MySQLError)
		if ok && me.Number == 1062 {
			return false, "This title already exists"
		}
		return false, err.Error()
	}

	return true, ""
}

func GetAll (skip int, limit int) ([]FetchedArticle, error) {
	var fetchedArticles []FetchedArticle

	query := `
		SELECT Articles.ArticleId, Articles.Text, Articles.Title, Articles.Approved, Articles.CreatedAt, Articles.Category, Users.Username
		FROM test.Articles
		JOIN test.Users ON Users.UserId = Articles.UserId
		WHERE Approved = 1
		LIMIT ?, ?
 	`

	res, err := db.Con.Query(query, skip, limit)
	if err != nil {
		return nil, err
	}

	for res.Next() {
		var fa FetchedArticle
		scanErr := res.Scan(&fa.ArticleId, &fa.Text, &fa.Title, &fa.Approved, &fa.CreatedAt, &fa.Category, &fa.Username)

		if scanErr != nil {
			return nil, scanErr
		}

		fetchedArticles = append(fetchedArticles, fa)
	}

	return fetchedArticles, nil
}

func GetSingleArticle (articleId int) (FetchedArticle, error) {
	var fa FetchedArticle

	query := `
		SELECT Articles.ArticleId, Articles.Text, Articles.Title, Articles.Approved, Articles.CreatedAt, Users.Username
  		FROM test.Articles
  		JOIN test.Users ON Users.UserId = Articles.UserId
  		WHERE ArticleId = ?
  `
	scanErr := db.Con.QueryRow(query, articleId).Scan(&fa.ArticleId, &fa.Text, &fa.Title, &fa.Approved, &fa.CreatedAt, &fa.Username)
	if scanErr != nil {
		return fa, scanErr
	}

	return fa, nil
}

func Update (article *NewArticle) (bool, string) {
	query := "UPDATE test.Articles SET Text = ?, Title = ?, WHERE UserId = ? AND ArticleId = ?"

	res, err := db.Con.Exec(query, article.Text, article.Title, article.UserId, article.ArticleId)
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
  res, err := db.Con.Query("SELECT Articles.ArticleId, Articles.Title, Users.Username FROM Articles JOIN Users on Articles.UserId = Users.UserId WHERE Articles.Approved = 0")
  if err != nil {
    fmt.Println(err)
  }

  for res.Next() {
    var article ArticlePreview

    scanErr := res.Scan(&article.ArticleId, &article.Title, &article.Username)
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
