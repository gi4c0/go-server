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
