package deposits

import (
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

func (u *Usecase) GetAll(ctx context.Context) ([]Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, u.ContextTimeout)
	defer cancel()

	return u.Repo.GetAll(ctx)
}

func (u *Usecase) GetByUserId(ctx context.Context, userId uint) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, u.ContextTimeout)
	defer cancel()

	return u.Repo.GetByUserId(ctx, userId)
}

func (u *Usecase) Create(ctx context.Context, d Domain) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, u.ContextTimeout)
	defer cancel()

	return u.Repo.Create(ctx, d)
}
