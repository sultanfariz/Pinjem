package book_orders

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

func (u *Usecase) GetById(ctx context.Context, id string) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, u.ContextTimeout)
	defer cancel()

	return u.Repo.GetById(ctx, id)
}

func (u *Usecase) GetByOrderId(ctx context.Context, id uint) ([]Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, u.ContextTimeout)
	defer cancel()

	return u.Repo.GetByOrderId(ctx, id)
}

func (u *Usecase) Create(ctx context.Context, domain Domain) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, u.ContextTimeout)
	defer cancel()

	return u.Repo.Create(ctx, domain)
}

func (u *Usecase) DeleteByOrderId(ctx context.Context, id uint) error {
	ctx, cancel := context.WithTimeout(ctx, u.ContextTimeout)
	defer cancel()
	// log.Println("---------------------------------------------------")

	return u.Repo.DeleteByOrderId(ctx, id)
}
