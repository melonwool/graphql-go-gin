package models

import (
	"context"
	"github.com/graph-gophers/graphql-go"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"graphql-go-gin/provider/sqlite"
	"graphql-go-gin/util"
)

// Book is the base type for books to be used by the sqlite and gql
type Book struct {
	gorm.Model
	OwnerID uint
	Name    string
	Tags    []Tag `gorm:"many2many:book_tags"`
}

// BookInput has everything needed to do adds and updates on a book
type BookInput struct {
	ID      *graphql.ID
	OwnerID int32
	Name    string
	TagIDs  *[]*int32
}

// GetBook should authorize the user in ctx and return a book or error
func (b *Book) GetBook(ctx context.Context, id uint) error {
	err := sqlite.DB.First(&b, id).Error
	if err != nil {
		return err
	}

	return nil
}

func (b *Book) GetBookOwner(ctx context.Context, id int32) (*User, error) {
	var u User
	err := sqlite.DB.First(&u, id).Error
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func (b *Book) GetBookTags(ctx context.Context) ([]Tag, error) {
	var t []Tag
	err := sqlite.DB.Model(b).Related(&t, "Tags").Error
	if err != nil {
		return nil, err
	}
	return t, nil
}

func GetBooksByID(ctx context.Context, ids []int, from, to int) ([]Book, error) {
	var b []Book
	err := sqlite.DB.Where("id in (?)", ids[from:to]).Find(&b).Error
	if err != nil {
		return nil, err
	}
	return b, nil
}

func (b *Book) UpdateBook(ctx context.Context, args *BookInput) error {
	// get the book to be updated from the sqlite
	bookID, err := util.GqlIDToUint(*args.ID)
	if err != nil {
		return err
	}
	err = sqlite.DB.First(b, bookID).Error
	if err != nil {
		return err
	}

	// so the pointer dereference is safe
	if args.TagIDs == nil {
		return errors.Wrap(err, "UpdateBook")
	}

	// if there are tags to be updated, go through that process
	var newTags []Tag
	if len(*args.TagIDs) > 0 {
		err = sqlite.DB.Where("id in (?)", *args.TagIDs).Find(&newTags).Error
		if err != nil {
			return err
		}

		// replace the old tag set with the new one
		err = sqlite.DB.Model(b).Association("Tags").Replace(newTags).Error
		if err != nil {
			return err
		}
	}

	updated := Book{
		Name:    args.Name,
		OwnerID: uint(args.OwnerID),
	}

	err = sqlite.DB.Model(b).Updates(updated).Error
	if err != nil {
		return err
	}

	err = sqlite.DB.First(b, bookID).Error
	if err != nil {
		return err
	}

	return nil
}

func (b *Book) DeleteBook(ctx context.Context, userID, bookID uint) (*bool, error) {
	// make sure the record exist
	err := sqlite.DB.First(b, bookID).Error
	if err != nil {
		return nil, err
	}

	// delete tags
	err = sqlite.DB.Model(b).Association("Tags").Clear().Error
	if err != nil {
		return nil, err
	}

	// delete record
	err = sqlite.DB.Delete(b).Error
	if err != nil {
		return nil, err
	}

	return util.BoolP(true), err
}

func (b *Book) AddBook(ctx context.Context, input BookInput) error {
	// get the M2M relation tags from the DB and put them in the book to be saved
	var t []Tag
	err := sqlite.DB.Where("id in (?)", input.TagIDs).Find(&t).Error
	if err != nil {
		return err
	}

	b.Name = input.Name
	b.OwnerID = uint(input.OwnerID)
	b.Tags = t

	err = sqlite.DB.Create(b).Error
	if err != nil {
		return err
	}

	return nil
}
