package books_test

import (
	"Pinjem/businesses/books"
	"Pinjem/businesses/books/mocks"
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var bookRepository mocks.DomainRepository

var bookService books.DomainService
var bookDomain books.Domain

func setup() {
	bookService = books.NewUsecase(&bookRepository, time.Minute*15)
	bookDomain = books.Domain{
		Id:            1,
		BookId:        "5wBQEp6ruIAC",
		ISBN:          "9780132119177",
		Publisher:     "Addison-Wesley Professional",
		PublishDate:   "1999-10-20",
		Title:         "The Pragmatic Programmer",
		Authors:       "Andrew hunt, David Thomas",
		Description:   "What others in the trenches say about The Pragmatic Programmer... “The cool thing about this book is that it’s great for keeping the programming process fresh. The book helps you to continue to grow and clearly comes from people who have been there.” —Kent Beck, author of Extreme Programming Explained: Embrace Change “I found this book to be a great mix of solid advice and wonderful analogies!” —Martin Fowler, author of Refactoring and UML Distilled “I would buy a copy, read it twice, then tell all my colleagues to run out and grab a copy. This is a book I would never loan because I would worry about it being lost.” —Kevin Ruland, Management Science, MSG-Logistics “The wisdom and practical experience of the authors is obvious. The topics presented are relevant and useful.... By far its greatest strength for me has been the outstanding analogies—tracer bullets, broken windows, and the fabulous helicopter-based explanation of the need for orthogonality, especially in a crisis situation. I have little doubt that this book will eventually become an excellent source of useful information for journeymen programmers and expert mentors alike.” —John Lakos, author of Large-Scale C++ Software Design “This is the sort of book I will buy a dozen copies of when it comes out so I can give it to my clients.” —Eric Vought, Software Engineer “Most modern books on software development fail to cover the basics of what makes a great software developer, instead spending their time on syntax or technology where in reality the greatest leverage possible for any software team is in having talented developers who really know their craft well. An excellent book.” —Pete McBreen, Independent Consultant “Since reading this book, I have implemented many of the practical suggestions and tips it contains. Across the board, they have saved my company time and money while helping me get my job done quicker! This should be a desktop reference for everyone who works with code for a living.” —Jared Richardson, Senior Software Developer, iRenaissance, Inc. “I would like to see this issued to every new employee at my company....” —Chris Cleeland, Senior Software Engineer, Object Computing, Inc. “If I’m putting together a project, it’s the authors of this book that I want. . . . And failing that I’d settle for people who’ve read their book.” —Ward Cunningham Straight from the programming trenches, The Pragmatic Programmer cuts through the increasing specialization and technicalities of modern software development to examine the core process--taking a requirement and producing working, maintainable code that delights its users. It covers topics ranging from personal responsibility and career development to architectural techniques for keeping your code flexible and easy to adapt and reuse. Read this book, and you'll learn how to Fight software rot; Avoid the trap of duplicating knowledge; Write flexible, dynamic, and adaptable code; Avoid programming by coincidence; Bullet-proof your code with contracts, assertions, and exceptions; Capture real requirements; Test ruthlessly and effectively; Delight your users; Build teams of pragmatic programmers; and Make your developments more precise with automation. Written as a series of self-contained sections and filled with entertaining anecdotes, thoughtful examples, and interesting analogies, The Pragmatic Programmer illustrates the best practices and major pitfalls of many different aspects of software development. Whether you're a new coder, an experienced programmer, or a manager responsible for software projects, use these lessons daily, and you'll quickly see improvements in personal productivity, accuracy, and job satisfaction. You'll learn skills and develop habits and attitudes that form the foundation for long-term success in your career. You'll become a Pragmatic Programmer.",
		Language:      "en",
		Picture:       "http://books.google.com/books/content?id=5wBQEp6ruIAC&printsec=frontcover&img=1&zoom=1&edge=curl&source=gbs_api",
		NumberOfPages: 352,
		MinDeposit:    120000,
		Status:        true,
	}
}

func TestGetAllBooks(t *testing.T) {
	setup()
	bookRepository.On("GetAll", mock.Anything).Return([]books.Domain{bookDomain}, nil)
	t.Run("Test Case 1 | Get All Books", func(t *testing.T) {
		books, err := bookService.GetAll(context.Background())
		if err != nil {
			t.Errorf("Error: %s", err)
		}
		if len(books) == 0 {
			t.Errorf("Error: %s", "No books found")
		}
		assert.Nil(t, err)
		assert.Equal(t, bookDomain, books[0])
	})

}

func TestGetBookByBookId(t *testing.T) {
	setup()
	bookRepository.On("GetById", mock.Anything, mock.AnythingOfType("string")).Return(bookDomain, nil)
	t.Run("Test Case 1 | Valid Get Book By BookId", func(t *testing.T) {
		book, err := bookService.GetById(context.Background(), "5wBQEp6ruIAC")
		if err != nil {
			t.Errorf("Error: %s", err)
		}
		assert.Nil(t, err)
		assert.Equal(t, bookDomain, book)
	})
	t.Run("Test Case 2 | Invalid Get Book By Empty BookId", func(t *testing.T) {
		book, err := bookService.GetById(context.Background(), "")
		assert.NotNil(t, err)
		assert.NotEqual(t, book, bookDomain)
	})
}

func TestGetBookByISBN(t *testing.T) {
	setup()
	bookRepository.On("GetByISBN", mock.Anything, mock.AnythingOfType("string")).Return(bookDomain, nil)
	t.Run("Test Case 1 | Valid Get Book By ISBN", func(t *testing.T) {
		book, err := bookService.GetByISBN(context.Background(), "9780132119177")
		if err != nil {
			t.Errorf("Error: %s", err)
		}
		assert.Nil(t, err)
		assert.Equal(t, bookDomain, book)
	})
	t.Run("Test Case 2 | Invalid Get Book By Empty ISBN", func(t *testing.T) {
		book, err := bookService.GetByISBN(context.Background(), "")
		assert.NotNil(t, err)
		assert.NotEqual(t, book, bookDomain)
	})
}

func TestCreateNewBook(t *testing.T) {
	setup()
	bookRepository.On("Create", mock.Anything, mock.AnythingOfType("Domain")).Return(bookDomain, nil)
	t.Run("Test Case 1 | Valid Create New Book", func(t *testing.T) {
		book, err := bookService.Create(context.Background(), bookDomain)
		if err != nil {
			t.Errorf("Error: %s", err)
		}
		if book.Id == 0 {
			t.Errorf("Error: %s", "Book Id is empty")
		}
		assert.Nil(t, err)
		assert.Equal(t, bookDomain, book)
	})
	t.Run("Test Case 2 | Invalid Create New Book with Empty Fields", func(t *testing.T) {
		book, err := bookService.Create(context.Background(), books.Domain{
			BookId:        "",
			ISBN:          "",
			Publisher:     "",
			PublishDate:   "",
			Title:         "",
			Authors:       "",
			Description:   "",
			Language:      "",
			Picture:       "",
			NumberOfPages: 0,
			MinDeposit:    0,
			Status:        false,
		})
		assert.NotNil(t, err)
		assert.NotEqual(t, book, bookDomain)
	})
}

func TestUpdateBookStatusByBookId(t *testing.T) {
	setup()
	bookRepository.On("UpdateStatus", mock.Anything, mock.AnythingOfType("string"), mock.AnythingOfType("bool")).Return(bookDomain, nil)
	t.Run("Test Case 1 | Valid Update Book Status", func(t *testing.T) {
		book, err := bookService.UpdateStatus(context.Background(), "5wBQEp6ruIAC", true)
		if err != nil {
			t.Errorf("Error: %s", err)
		}
		assert.Nil(t, err)
		assert.Equal(t, bookDomain, book)
	})
	t.Run("Test Case 2 | Invalid Update Book Status with Empty BookId", func(t *testing.T) {
		book, err := bookService.UpdateStatus(context.Background(), "", true)
		assert.NotNil(t, err)
		assert.NotEqual(t, book, bookDomain)
	})
}
