package resolver

import (
	graphql "github.com/graph-gophers/graphql-go"
	"github.com/suppayami/gohard/models"
	"github.com/suppayami/gohard/repositories"
)

// Resolver resolves graphql queries
type Resolver struct {
	PostRepository repositories.PostRepository
}

// Post resolves a post
func (r *Resolver) Post(params struct{ ID graphql.ID }) *models.PostResolver {
	if record := r.PostRepository.Post(params.ID); record != nil {
		return &models.PostResolver{Record: record}
	}
	return nil
}
