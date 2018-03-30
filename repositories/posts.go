package repositories

import (
	graphql "github.com/graph-gophers/graphql-go"
	"github.com/suppayami/gohard/models"
)

// PostRepository contains Posts
type PostRepository interface {
	Post(id graphql.ID) *models.Post
}

// Posts is a fake repo
type Posts struct {
	posts map[graphql.ID]*models.Post
}

// Post returns a post?
func (repo *Posts) Post(id graphql.ID) *models.Post {
	return repo.posts[id]
}

// NewPosts returns the fake repo
func NewPosts(posts map[graphql.ID]*models.Post) *Posts {
	return &Posts{posts}
}
