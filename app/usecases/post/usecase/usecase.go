package usecase

import (
	"forum/app/entities"
	"forum/app/usecases/post"
)

type postUsecase struct {
	repo post.PostRepo
}

func NewPostService(repo post.PostRepo) post.PostService {
	return &postUsecase{
		repo: repo,
	}
}

func (p *postUsecase) CreatePost(post entities.Post) {}

func (p *postUsecase) EditPost() {}

func (p *postUsecase) DeletePost() {}

func (p *postUsecase) LikePost() {}

func (p *postUsecase) DisLikePost() {}
