package models

import (
	"context"
	"github.com/jinzhu/gorm"
	"graphql-go-gin/provider/sqlite"
)

// User is the base user model to be used throughout the app
type User struct {
	gorm.Model
	Name  string
	Books []Book `gorm:"foreignkey:OwnerID"`
}

func (u *User) GetUserBookIDs(ctx context.Context, userID uint) ([]int, error) {
	var ids []int
	err := sqlite.DB.Where("owner_id = ?", userID).Find(&[]Book{}).Pluck("id", &ids).Error
	if err != nil {
		return nil, err
	}
	return ids, nil
}

func (u *User) GetUser(ctx context.Context, id uint) error {
	err := sqlite.DB.First(u, id).Error
	if err != nil {
		return err
	}

	return nil
}

// GetUserBooks gets books associated with the user
func (u *User) GetUserBooks(ctx context.Context, id uint) ([]Book, error) {
	u.ID = id

	var b []Book
	err := sqlite.DB.Model(&u).Association("Books").Find(&b).Error
	if err != nil {
		return nil, err
	}

	return b, nil
}
