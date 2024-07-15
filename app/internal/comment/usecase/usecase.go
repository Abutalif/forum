package usecase

import "forum/app/internal/comment"

type commentUsecase struct {
	repo comment.CommentRepo
}

func NewPostService(repo comment.CommentRepo) comment.CommentService {
	return &commentUsecase{
		repo: repo,
	}
}
