package books

import (
	"Pinjem/exceptions"
	context "context"
	"time"

	"github.com/go-playground/validator/v10"
)

type Usecase struct {
	Repo           DomainRepository
	ContextTimeout time.Duration
}

func NewUsecase(repo DomainRepository, timeout time.Duration) *Usecase {
	return &Usecase{
		Repo:           repo,
		ContextTimeout: timeout,
	}
}

func (u *Usecase) GetAll(ctx context.Context) ([]Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, u.ContextTimeout)
	defer cancel()

	return u.Repo.GetAll(ctx)
}

func (u *Usecase) GetById(ctx context.Context, id string) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, u.ContextTimeout)
	defer cancel()

	if id == "" {
		return Domain{}, exceptions.ErrEmptyInput
	}

	return u.Repo.GetById(ctx, id)
}

func (u *Usecase) GetByISBN(ctx context.Context, isbn string) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, u.ContextTimeout)
	defer cancel()

	if isbn == "" {
		return Domain{}, exceptions.ErrEmptyInput
	}

	return u.Repo.GetByISBN(ctx, isbn)
}

func (u *Usecase) Create(ctx context.Context, domain Domain) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, u.ContextTimeout)
	defer cancel()

	validate := validator.New()
	err := validate.Struct(domain)
	if err != nil {
		return Domain{}, exceptions.ErrValidationFailed
	}

	return u.Repo.Create(ctx, domain)
}

func (u *Usecase) UpdateStatus(ctx context.Context, bookId string, status bool) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, u.ContextTimeout)
	defer cancel()

	if bookId == "" {
		return Domain{}, exceptions.ErrEmptyInput
	}

	return u.Repo.UpdateStatus(ctx, bookId, status)
}
