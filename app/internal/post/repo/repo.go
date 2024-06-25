package repo

import (
	"forum/app/entities"
	"forum/app/internal/post"
)

type postRepo struct{}

func NewPostRepo() post.PostRepo {
	return &postRepo{}
}

func (p *postRepo) GetPost() entities.Post {
	//query := `SELECT from POST where id=`
	return entities.Post{}
}

func (p *postRepo) CreatePost(post entities.Post) error {
	return nil
}
