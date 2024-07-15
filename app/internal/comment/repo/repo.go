package repo

import "forum/app/internal/comment"

type commentRepo struct{}

func NewCommentRepo() comment.CommentRepo {
	return &commentRepo{}
}
