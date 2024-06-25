package repo

import "forum/app/usecases/post"

type postRepo struct{}

func NewPostRepo() post.PostRepo {
	return &postRepo{}
}

func (p *postRepo) GetPost() {}
