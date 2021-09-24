package users

import (
	"Pinjem/businesses/users"
	"Pinjem/helpers"
	"context"
	"log"

	"gorm.io/gorm"
)

type UserRepository struct {
	Conn *gorm.DB
}

func NewUserRepository(conn *gorm.DB) users.DomainRepository {
	return &UserRepository{Conn: conn}
}

func (u *UserRepository) Login(ctx context.Context, email string, password string) (users.Domain, error) {
	var user Users
	err := u.Conn.Where("email = ?", email).First(&user).Error
	if err != nil {
		return users.Domain{}, err
	}
	log.Println(user.Password)
	if !helpers.IsMatched(user.Password, password) {
		return users.Domain{}, err
	}

	return user.ToDomain(), nil
}

func (u *UserRepository) Register(ctx context.Context, user users.Domain) (users.Domain, error) {
	var userModel Users
	err := u.Conn.Where("email = ?", user.Email).First(&userModel).Error
	if err != nil {
		log.Println(err)
		log.Println(err.Error())
		return users.Domain{}, err
	}
	password, err := helpers.HashPassword(user.Password)
	if err != nil {
		return users.Domain{}, err
	}
	createdUser := Users{
		Email:       user.Email,
		Password:    password,
		Fullname:    user.Fullname,
		NIK:         user.Nik,
		PhoneNumber: user.PhoneNumber,
		Address:     user.Address,
		Provinsi:    user.Provinsi,
		Kota:        user.Kota,
		Kecamatan:   user.Kecamatan,
		Desa:        user.Desa,
		PostalCode:  user.PostalCode,
		Role:        user.Role,
		Status:      user.Status,
		LinkKTP:     user.LinkKTP,
	}
	u.Conn.Create(&createdUser)
	return createdUser.ToDomain(), nil
}

func (u *UserRepository) FindByEmail(ctx context.Context, email string) (users.Domain, error) {
	var user Users
	if err := u.Conn.Where("email = ?", email).First(&user).Error; err != nil {
		return users.Domain{}, err
	}
	return user.ToDomain(), nil
}

// func (u *UserRepository) FindByID(id uint) (*User, error) {
// 	var user User
// 	if err := u.Conn.Where("id = ?", id).First(&user).Error; err != nil {
// 		return nil, err
// 	}
// 	return &user, nil
// }

func (u *UserRepository) Create(ctx context.Context, user users.Domain) (users.Domain, error) {
	password, err := helpers.HashPassword(user.Password)
	if err != nil {
		return users.Domain{}, err
	}
	createdUser := Users{
		Email:       user.Email,
		Password:    password,
		Fullname:    user.Fullname,
		NIK:         user.Nik,
		PhoneNumber: user.PhoneNumber,
		Address:     user.Address,
		Provinsi:    user.Provinsi,
		Kota:        user.Kota,
		Kecamatan:   user.Kecamatan,
		Desa:        user.Desa,
		PostalCode:  user.PostalCode,
		Role:        user.Role,
		Status:      user.Status,
		LinkKTP:     user.LinkKTP,
	}
	insertErr := u.Conn.Create(&createdUser).Error
	if insertErr != nil {
		return users.Domain{}, insertErr
	}
	return createdUser.ToDomain(), nil
}

// func (u *UserRepository) Update(user *User) error {
// 	return u.Conn.Save(user).Error
// }

// func (u *UserRepository) Delete(user *User) error {
// 	return u.Conn.Delete(user).Error
// }
