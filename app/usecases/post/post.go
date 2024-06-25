package post

import "forum/app/entities"

type PostService interface {
	CreatePost(entities.Post)
	EditPost()
	DeletePost()
	LikePost()
	DisLikePost()
}

type PostRepo interface {
	GetPost()
}
