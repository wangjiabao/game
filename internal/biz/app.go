package biz

import (
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
