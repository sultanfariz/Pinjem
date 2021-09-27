package books

import (
	"Pinjem/businesses/books"
	"context"
	"time"

	"gorm.io/gorm"
)

type BookRepository struct {
	Conn *gorm.DB
}

func NewBookRepository(conn *gorm.DB) books.DomainRepository {
	return &BookRepository{Conn: conn}
}

func (b *BookRepository) GetAll(ctx context.Context) ([]books.Domain, error) {
	var booksModel []Books
	if err := b.Conn.Find(&booksModel).Error; err != nil {
		return nil, err
	}
	var result []books.Domain
	result = ToListDomain(booksModel)
	return result, nil
}

func (b *BookRepository) GetById(ctx context.Context, id uint) (books.Domain, error) {
	var book Books
	if err := b.Conn.Where("id = ?", id).First(&book).Error; err != nil {
		return books.Domain{}, err
	}
	return book.ToDomain(), nil
}

func (b *BookRepository) GetByISBN(ctx context.Context, isbn string) (books.Domain, error) {
	var book Books
	if err := b.Conn.Where("isbn = ?", isbn).First(&book).Error; err != nil {
		return books.Domain{}, err
	}
	return book.ToDomain(), nil
}

func (b *BookRepository) Create(ctx context.Context, book books.Domain) (books.Domain, error) {
	createdBook := Books{
		BookId:        book.BookId,
		ISBN:          book.ISBN,
		Title:         book.Title,
		Publisher:     book.Publisher,
		PublishDate:   book.PublishDate,
		Description:   book.Description,
		Language:      book.Language,
		Picture:       book.Picture,
		NumberOfPages: book.NumberOfPages,
		MinDeposit:    book.MinDeposit,
		Status:        book.Status,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}
	insertErr := b.Conn.Create(&createdBook).Error
	if insertErr != nil {
		return books.Domain{}, insertErr
	}
	return createdBook.ToDomain(), nil
}

// func (b *BookRepository) Update(user *User) error {
// 	return b.Conn.Save(user).Error
// }

// func (b *BookRepository) Delete(user *User) error {
// 	return b.Conn.Delete(user).Error
// }
