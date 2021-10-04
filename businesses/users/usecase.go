package users

import (
	"Pinjem/exceptions"

	context "context"
	"time"

	"github.com/go-playground/validator/v10"
)

type Usecase struct {
	Repo           DomainRepository
	contextTimeout time.Duration
}

func NewUsecaseTest(repo DomainRepository) Usecase {
	return Usecase{
		Repo:           repo,
		contextTimeout: time.Hour * 1,
	}
}

func NewUsecase(repo DomainRepository, timeout time.Duration) *Usecase {
	return &Usecase{
		Repo:           repo,
		contextTimeout: timeout,
	}
}

func (u *Usecase) Login(ctx context.Context, email, password string) (Domain, error) {
	validate := validator.New()
	err := validate.Struct(LoginDomain{Email: email, Password: password})
	if err != nil {
		return Domain{}, exceptions.ErrInvalidCredentials
	}

	return u.Repo.Login(ctx, email, password)
}

func (u *Usecase) Register(ctx context.Context, domain Domain) (Domain, error) {
	validate := validator.New()
	err := validate.Struct(domain)
	if err != nil {
		return Domain{}, exceptions.ErrValidationFailed
	}

	return u.Repo.Create(ctx, domain)
}

func (u *Usecase) FindByEmail(ctx context.Context, email string) (Domain, error) {
	if email == "" {
		return Domain{}, exceptions.ErrInvalidCredentials
	}

	return u.Repo.FindByEmail(ctx, email)
}

func (u *Usecase) GetAll(ctx context.Context) ([]Domain, error) {

	return u.Repo.GetAll(ctx)
}

func (u *Usecase) GetById(ctx context.Context, id uint) (Domain, error) {

	return u.Repo.GetById(ctx, id)
}
