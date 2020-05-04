package resolver

import (
	"context"
	"github.com/graph-gophers/graphql-go"
	"github.com/pkg/errors"
	"graphql-go-gin/models"
	"graphql-go-gin/util"
)

// BookResolver contains the DB and the model for resolving
type BookResolver struct {
	m *models.Book
}

// ID resolves the ID field for Book
func (b *BookResolver) ID(ctx context.Context) *graphql.ID {
	return util.GqlIDP(b.m.ID)
}

// Owner resolves the owner field for Book
func (b *BookResolver) Owner(ctx context.Context) (*UserResolver, error) {
	user, err := b.m.GetBookOwner(ctx, int32(b.m.OwnerID))
	if err != nil {
		return nil, errors.Wrap(err, "Owner")
	}
	r := UserResolver{
		m: user,
	}
	return &r, nil
}

// Name resolves the name field for Book
func (b *BookResolver) Name(ctx context.Context) *string {
	return &b.m.Name
}

// Tags resolves the book tags
func (b *BookResolver) Tags(ctx context.Context) (*[]*TagResolver, error) {
	tags, err := b.m.GetBookTags(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "Tags")
	}

	r := make([]*TagResolver, len(tags))
	for i := range tags {
		r[i] = &TagResolver{
			m: tags[i],
		}
	}

	return &r, nil
}
