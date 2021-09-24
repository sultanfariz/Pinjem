package users

import (
	"Pinjem/exceptions"
	context "context"
	"time"
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

func (u *Usecase) Login(ctx context.Context, email, password string) (Domain, error) {
	if email == "" || password == "" {
		return Domain{}, exceptions.ErrInvalidCredentials
	}
	ctx, cancel := context.WithTimeout(ctx, u.ContextTimeout)
	defer cancel()

	return u.Repo.Login(ctx, email, password)
}

func (u *Usecase) Register(ctx context.Context, domain Domain) (Domain, error) {
	if domain.Email == "" || domain.Password == "" {
		return Domain{}, exceptions.ErrInvalidCredentials
	}
	ctx, cancel := context.WithTimeout(ctx, u.ContextTimeout)
	defer cancel()

	user, err := u.Repo.FindByEmail(ctx, domain.Email)
	if err != nil {
		return Domain{}, err
	}
	if user.Id != 0 {
		return Domain{}, exceptions.ErrUserAlreadyExists
	}

	return u.Repo.Create(ctx, domain)
}

func (u *Usecase) FindByEmail(ctx context.Context, email string) (Domain, error) {
	if email == "" {
		return Domain{}, exceptions.ErrInvalidCredentials
	}
	ctx, cancel := context.WithTimeout(ctx, u.ContextTimeout)
	defer cancel()

	return u.Repo.FindByEmail(ctx, email)
}
