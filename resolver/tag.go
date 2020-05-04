package resolver

import (
	"context"
	"github.com/graph-gophers/graphql-go"
	"github.com/pkg/errors"
	"graphql-go-gin/models"
	"graphql-go-gin/util"
)

// TagResolver contains the db and the Tag model for resolving
type TagResolver struct {
	m models.Tag
}

// ID resolves the ID for Tag
func (t *TagResolver) ID(ctx context.Context) *graphql.ID {
	return util.GqlIDP(t.m.ID)
}

// Title resolves the title field
func (t *TagResolver) Title(ctx context.Context) *string {
	return &t.m.Title
}

// Books resolves the books field
func (t *TagResolver) Books(ctx context.Context) (*[]*BookResolver, error) {
	books, err := t.m.GetTagBooks(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "Books")
	}

	r := make([]*BookResolver, len(books))
	for i := range books {
		r[i] = &BookResolver{
			m: &books[i],
		}
	}

	return &r, nil
}
