package comment

import  (
	"go-server/db"
)

type NewComment struct {
	CommentId int
	ArticleId int
	UserId int
	ParentCommentId int
	Text string
	CreatedAt string
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