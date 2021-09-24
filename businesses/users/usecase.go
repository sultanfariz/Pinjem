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
