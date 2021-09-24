package users

import (
	"Pinjem/businesses/users"
	"context"

	"gorm.io/gorm"
)

type UserRepository struct {
	Conn *gorm.DB
}

func NewUserRepository(conn *gorm.DB) users.DomainRepository {
	return &UserRepository{Conn: conn}
}

func (u *UserRepository) Login(ctx context.Context, email string, password string) (users.Domain, error) {
	var user users.Domain
	err := u.Conn.Where("email = ? and password = ?", email, password).First(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

// func (u *UserRepository) FindByEmail(email string) (*User, error) {
// 	var user User
// 	if err := u.Conn.Where("email = ?", email).First(&user).Error; err != nil {
// 		return nil, err
// 	}
// 	return &user, nil
// }

// func (u *UserRepository) FindByID(id uint) (*User, error) {
// 	var user User
// 	if err := u.Conn.Where("id = ?", id).First(&user).Error; err != nil {
// 		return nil, err
// 	}
// 	return &user, nil
// }

// func (u *UserRepository) Store(user *User) error {
// 	return u.Conn.Create(user).Error
// }

// func (u *UserRepository) Update(user *User) error {
// 	return u.Conn.Save(user).Error
// }

// func (u *UserRepository) Delete(user *User) error {
// 	return u.Conn.Delete(user).Error
// }
