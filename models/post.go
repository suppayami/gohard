package models

import (
	graphql "github.com/graph-gophers/graphql-go"
)

// Post is a post man
type Post struct {
	ID    graphql.ID
	Title string
}

// PostResolver resolves the post fields
type PostResolver struct {
	Record *Post
}

// ID resolves ID field
func (r *PostResolver) ID() graphql.ID {
	return r.Record.ID
}

// Title resolves Title field
func (r *PostResolver) Title() string {
	return r.Record.Title
}
