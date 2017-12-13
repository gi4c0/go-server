package comment

import (
	"fmt"
	"go-server/db"
)

type NewComment struct {
	CommentId       int
	ArticleId       int
	UserId          int
	ParentCommentId int
	Text            string
	CreatedAt       string
	Username 		string
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

func DeleteComment(commentId, userId int) string {
	query := "DELETE FROM test.Comments WHERE CommentId = ? AND UserId = ?"
	res, err := db.Con.Exec(query, commentId, userId)
	if err != nil {
		return err.Error()
	}

	rowsAffected, rowsErr := res.RowsAffected()
	if rowsErr != nil {
		return rowsErr.Error()
	}

	if rowsAffected < 1 {
		return "Wrong comment id or you don't have permission to delete this comment"
	}

	return ""
}

func GetCommentsByArticleId(articleId int) ([]NewComment, error) {
	var comments []NewComment
	query := `
		SELECT CommentId, ArticleId, ParentCommentId, Text, CreatedAt, Username

		FROM test.Comments JOIN test.Users ON Comments.UserId = Users.UserId WHERE ArticleId = ?`
	res, err := db.Con.Query(query, articleId)
	if err != nil {
		return comments, err
	}

	for res.Next() {
		var comment NewComment
		scanErr := res.Scan(&comment.CommentId, &comment.ArticleId, &comment.ParentCommentId, &comment.Text, &comment.CreatedAt, &comment.Username)
		if scanErr != nil {
			fmt.Println(scanErr)
			return comments, scanErr
		}

		comments = append(comments, comment)
	}

	return comments, nil
}
