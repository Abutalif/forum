package entities

type Comment struct {
	Content  string
	Likes    int
	Dislikes int
	Post     *Post
}
