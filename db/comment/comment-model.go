package comment

import (
	"fmt"
	"go-server/db"
  "strconv"
  "errors"
)

type NewComment struct {
	CommentId       int
	ArticleId       int
	UserId          int
	ParentCommentId int
	Text            string
	CreatedAt       string
	Username 		    string
  Deleted         bool
}

func (comment *NewComment) Save() {
	query := "INSERT INTO test.Comments (ArticleId, UserId, ParentCommentId, Text) VALUES (?, ?, ?, ?)"
	db.Con.Exec(query, comment.ArticleId, comment.UserId, comment.ParentCommentId, comment.Text)
}

func (comment *NewComment) Update() string {
	query := "UPDATE test.Comments SET Text = ? WHERE CommentId = ? AND UserId = ?"
	_, err := db.Con.Exec(query, comment.Text, comment.CommentId, comment.UserId)
	if err != nil {
		return err.Error()
	}

	return ""
}

func Delete(commentId int, userId int, permission string) error {
  var query string

  if permission == "user" {
    query = "UPDATE test.Comments SET Deleted = 1 WHERE CommentId = ? AND UserId = " + strconv.Itoa(userId)
  } else {
    query = "UPDATE test.Comments SET Deleted = 1 WHERE CommentId = ?"
  }

	res, err := db.Con.Exec(query, commentId)
	if err != nil {
		return err
	}

	rowsAffected, rowsErr := res.RowsAffected()
	if rowsErr != nil {
		return rowsErr
	}

	if rowsAffected < 1 {
		return errors.New("Wrong id or you don't have permission to delete this comment")
	}

	return nil
}

func GetCommentsByArticleId(articleId int) ([]NewComment, error) {
	var comments []NewComment
	query := `
		SELECT CommentId, ArticleId, ParentCommentId, Text, CreatedAt, Username, Deleted
		FROM test.Comments JOIN test.Users ON Comments.UserId = Users.UserId WHERE ArticleId = ?`

	res, err := db.Con.Query(query, articleId)
	if err != nil {
		return comments, err
	}

	for res.Next() {
		var comment NewComment
		scanErr := res.Scan(&comment.CommentId, &comment.ArticleId, &comment.ParentCommentId, &comment.Text, &comment.CreatedAt, &comment.Username, &comment.Deleted)
		if scanErr != nil {
			fmt.Println(scanErr)
			return comments, scanErr
		}

		comments = append(comments, comment)
	}

	return comments, nil
}
