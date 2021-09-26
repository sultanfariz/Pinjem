package authors

import (
	"Pinjem/businesses/authors"
	"context"
	"time"

	"gorm.io/gorm"
)

type AuthorRepository struct {
	Conn *gorm.DB
}

func NewAuthorRepository(conn *gorm.DB) authors.DomainRepository {
	return &AuthorRepository{Conn: conn}
}

func (b *AuthorRepository) GetAll(ctx context.Context) ([]authors.Domain, error) {
	var authorsModel []Authors
	if err := b.Conn.Find(&authorsModel).Error; err != nil {
		return nil, err
	}
	var result []authors.Domain
	result = ToListDomain(authorsModel)
	return result, nil
}

func (b *AuthorRepository) GetById(ctx context.Context, id uint) (authors.Domain, error) {
	var book Authors
	if err := b.Conn.Where("id = ?", id).First(&book).Error; err != nil {
		return authors.Domain{}, err
	}
	return book.ToDomain(), nil
}

func (b *AuthorRepository) Create(ctx context.Context, book authors.Domain) (authors.Domain, error) {
	createdBook := Authors{
		AuthorId:     book.AuthorId,
		Name:         book.Name,
		PersonalName: book.PersonalName,
		Birthdate:    book.Birthdate,
		Bio:          book.Bio,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
	insertErr := b.Conn.Create(&createdBook).Error
	if insertErr != nil {
		return authors.Domain{}, insertErr
	}
	return createdBook.ToDomain(), nil
}

// func (b *AuthorRepository) Update(user *User) error {
// 	return b.Conn.Save(user).Error
// }

// func (b *AuthorRepository) Delete(user *User) error {
// 	return b.Conn.Delete(user).Error
// }
