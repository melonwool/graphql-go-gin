package models

import (
	"context"
	"github.com/jinzhu/gorm"
	"graphql-go-gin/provider/sqlite"
)

// Tag is the base type for a book tag to be used by the db and gql
type Tag struct {
	gorm.Model
	Title string
	Books []Book `gorm:"many2many:book_tags"`
}

// GetTagBooks
func (t *Tag) GetTagBooks(ctx context.Context) ([]Book, error) {
	var b []Book
	err := sqlite.DB.Model(t).Related(&b, "Books").Error
	if err != nil {
		return nil, err
	}
	return b, nil
}

// GetTagByTitle
func (t *Tag) GetTagBytTitle(ctx context.Context, title string) error {
	err := sqlite.DB.Where("title = ?", title).First(t).Error
	if err != nil {
		return err
	}
	return nil
}
