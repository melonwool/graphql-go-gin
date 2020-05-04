package resolver

import (
	"context"
	"encoding/base64"
	"fmt"
	"graphql-go-gin/models"
	"graphql-go-gin/util"
	"strconv"
	"strings"

	"github.com/graph-gophers/graphql-go"
	"github.com/pkg/errors"
)

// Resolver root resolver
type Resolver struct{}

// GetUser resolves the User query
func (r *Resolver) GetUser(ctx context.Context, args struct{ ID graphql.ID }) (*UserResolver, error) {
	id, err := util.GqlIDToUint(args.ID)
	if err != nil {
		return nil, errors.Wrap(err, "GetBook")
	}
	user := &models.User{}
	if err = user.GetUser(ctx, id); err != nil {
		return nil, errors.Wrap(err, "GetUser")
	}
	s := UserResolver{
		m: user,
	}

	return &s, nil
}

// GetBook resolves the getBook query
func (r *Resolver) GetBook(ctx context.Context, args struct{ ID graphql.ID }) (*BookResolver, error) {
	id, err := util.GqlIDToUint(args.ID)
	if err != nil {
		return nil, errors.Wrap(err, "GetBook")
	}
	book := &models.Book{}
	if err = book.GetBook(ctx, id); err != nil {
		return nil, err
	}
	s := BookResolver{
		m: book,
	}

	return &s, nil
}

// GetTag resolves the getTag query
func (r *Resolver) GetTag(ctx context.Context, args struct{ Title string }) (*TagResolver, error) {
	tag := &models.Tag{}
	err := tag.GetTagBytTitle(ctx, args.Title)
	if err != nil {
		return nil, errors.Wrap(err, "GetTag")
	}
	s := TagResolver{
		m: *tag,
	}
	return &s, nil
}

// AddBook Resolves the addBook mutation
func (r *Resolver) AddBook(ctx context.Context, args struct{ Book models.BookInput }) (*BookResolver, error) {
	book := &models.Book{}
	err := book.AddBook(ctx, args.Book)
	if err != nil {
		return nil, errors.Wrap(err, "AddBook")
	}

	s := BookResolver{
		m: book,
	}

	return &s, nil
}

// UpdateBook takes care of updating any field on the book
func (r *Resolver) UpdateBook(ctx context.Context, args struct{ Book models.BookInput }) (*BookResolver, error) {
	book := &models.Book{}
	err := book.UpdateBook(ctx, &args.Book)
	if err != nil {
		return nil, errors.Wrap(err, "UpdateBook")
	}

	s := BookResolver{
		m: book,
	}

	return &s, nil
}

// DeleteBook takes care of deleting a book record
func (r *Resolver) DeleteBook(ctx context.Context, args struct{ UserID, BookID graphql.ID }) (*bool, error) {
	bookID, err := util.GqlIDToUint(args.BookID)
	if err != nil {
		return nil, errors.Wrap(err, "DeleteBook")
	}

	userID, err := util.GqlIDToUint(args.UserID)
	if err != nil {
		return nil, errors.Wrap(err, "DeleteBook")
	}
	book := &models.Book{}
	return book.DeleteBook(ctx, userID, bookID)
}

// encode cursor encodes the cursot position in base64
func encodeCursor(i int) graphql.ID {
	return graphql.ID(base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("cursor%d", i))))
}

// decode cursor decodes the base 64 encoded cursor and resturns the integer
func decodeCursor(s string) (int, error) {
	b, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return 0, err
	}

	i, err := strconv.Atoi(strings.TrimPrefix(string(b), "cursor"))
	if err != nil {
		return 0, err
	}

	return i, nil
}
