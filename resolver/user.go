package resolver

import (
	"context"
	"github.com/graph-gophers/graphql-go"
	"github.com/pkg/errors"
	"graphql-go-gin/models"
	"graphql-go-gin/util"
)

// UserResolver contains the database and the user model to resolve against
type UserResolver struct {
	m *models.User
}

// ID resolves the user ID
func (u *UserResolver) ID(ctx context.Context) *graphql.ID {
	return util.GqlIDP(u.m.ID)
}

// Name resolves the Name field for User, it is all caps to avoid name clashes
func (u *UserResolver) Name(ctx context.Context) *string {
	return &u.m.Name
}

// Books resolves the Books field for User
func (u *UserResolver) Books(ctx context.Context) (*[]*BookResolver, error) {
	books, err := u.m.GetUserBooks(ctx, u.m.ID)
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
