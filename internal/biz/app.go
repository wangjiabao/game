package biz

import (
	v1 "game/api/app/v1"
	"github.com/go-kratos/kratos/v2/log"
)

type UserRepo interface {
}

// AppUsecase is an app usecase.
type AppUsecase struct {
	userRepo UserRepo
	tx       Transaction
	log      *log.Helper
}

// NewAppUsecase new a app usecase.
func NewAppUsecase(userRepo UserRepo, tx Transaction, logger log.Logger) *AppUsecase {
	return &AppUsecase{userRepo: userRepo, tx: tx, log: log.NewHelper(logger)}
}

func (a *AppUsecase) CreateUser(req *v1.CreateUserRequest) (*v1.CreateUserReply, error) {

	return nil, nil
}
