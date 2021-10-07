package users

import (
	"Pinjem/businesses/users"
	"Pinjem/helpers"
	"context"
	"time"

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

	if !helpers.IsMatched(user.Password, password) {
		return users.Domain{}, err
	}

	return user.ToDomain(), nil
}

func (u *UserRepository) FindByEmail(ctx context.Context, email string) (users.Domain, error) {
	var user Users
	if err := u.Conn.Where("email = ?", email).First(&user).Error; err != nil {
		return users.Domain{}, err
	}
	return user.ToDomain(), nil
}

func (u *UserRepository) GetAll(ctx context.Context) ([]users.Domain, error) {
	var usersModel []Users
	if err := u.Conn.Find(&usersModel).Error; err != nil {
		return nil, err
	}
	var result []users.Domain
	result = ToListDomain(usersModel)
	return result, nil
}

func (u *UserRepository) GetById(ctx context.Context, id uint) (users.Domain, error) {
	var user Users
	if err := u.Conn.Where("id = ?", id).First(&user).Error; err != nil {
		return users.Domain{}, err
	}
	return user.ToDomain(), nil
}

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
		Birthdate:   user.Birthdate,
		Address:     user.Address,
		Provinsi:    user.Provinsi,
		Kota:        user.Kota,
		Kecamatan:   user.Kecamatan,
		Desa:        user.Desa,
		PostalCode:  user.PostalCode,
		Role:        user.Role,
		Status:      user.Status,
		LinkKTP:     user.LinkKTP,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
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
