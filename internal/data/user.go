package data

import (
	"context"
	"game/internal/biz"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
	"strconv"
	"time"
)

type User struct {
	ID        uint64    `gorm:"primarykey;type:int"`
	Address   string    `gorm:"type:varchar(100);not null"`
	Level     uint64    `gorm:"type:int;not null"`
	Giw       float64   `gorm:"type:decimal(65,20);not null"`
	Git       float64   `gorm:"type:decimal(65,20);not null"`
	Total     float64   `gorm:"type:decimal(65,20);not null"`
	CreatedAt time.Time `gorm:"type:datetime;not null"`
	UpdatedAt time.Time `gorm:"type:datetime;not null"`
}

type UserRecommend struct {
	ID            uint64    `gorm:"primarykey;type:int"`
	UserId        uint64    `gorm:"type:int;not null"`
	RecommendCode string    `gorm:"type:varchar(1000);not null"`
	CreatedAt     time.Time `gorm:"type:datetime;not null"`
	UpdatedAt     time.Time `gorm:"type:datetime;not null"`
}

type UserRepo struct {
	data *Data
	log  *log.Helper
}

func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &UserRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

// GetUserByAddress .
func (u *UserRepo) GetUserByAddress(ctx context.Context, address string) (*biz.User, error) {
	var user *User
	if err := u.data.DB(ctx).Where("address=?", address).Table("user").First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, errors.New(500, "USER ERROR", err.Error())
	}

	return &biz.User{
		ID:        user.ID,
		Address:   user.Address,
		Level:     user.Level,
		Giw:       user.Giw,
		Git:       user.Git,
		Total:     user.Total,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}

// GetUserRecommendByUserId .
func (u *UserRepo) GetUserRecommendByUserId(ctx context.Context, userId uint64) (*biz.UserRecommend, error) {
	var userRecommend *UserRecommend
	if err := u.data.DB(ctx).Where("user_id=?", userId).Table("user_recommend").First(&userRecommend).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, errors.New(500, "USER RECOMMEND ERROR", err.Error())
	}

	return &biz.UserRecommend{
		ID:            userRecommend.ID,
		UserId:        userRecommend.UserId,
		RecommendCode: userRecommend.RecommendCode,
		CreatedAt:     userRecommend.CreatedAt,
		UpdatedAt:     userRecommend.UpdatedAt,
	}, nil
}

// CreateUser .
func (u *UserRepo) CreateUser(ctx context.Context, uc *biz.User) (*biz.User, error) {
	var user User
	user.Address = uc.Address

	res := u.data.DB(ctx).Table("user").Create(&user)
	if res.Error != nil {
		return nil, errors.New(500, "CREATE_USER_ERROR", "用户创建失败")
	}

	return &biz.User{
		ID:        user.ID,
		Address:   user.Address,
		Level:     user.Level,
		Giw:       user.Giw,
		Git:       user.Git,
		Total:     user.Total,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}

// CreateUserRecommend .
func (u *UserRepo) CreateUserRecommend(ctx context.Context, user *biz.User, recommendUser *biz.UserRecommend) (*biz.UserRecommend, error) {
	var tmpRecommendCode string
	if nil != recommendUser && 0 < recommendUser.UserId {
		tmpRecommendCode = "D" + strconv.FormatUint(recommendUser.UserId, 10)
		if "" != recommendUser.RecommendCode {
			tmpRecommendCode = recommendUser.RecommendCode + tmpRecommendCode
		}
	}

	var userRecommend UserRecommend
	userRecommend.UserId = user.ID
	userRecommend.RecommendCode = tmpRecommendCode

	res := u.data.DB(ctx).Table("user_recommend").Create(&userRecommend)
	if res.Error != nil {
		return nil, errors.New(500, "CREATE_USER_RECOMMEND_ERROR", "用户推荐关系创建失败")
	}

	return &biz.UserRecommend{
		ID:            userRecommend.ID,
		UserId:        userRecommend.UserId,
		RecommendCode: userRecommend.RecommendCode,
		CreatedAt:     userRecommend.CreatedAt,
		UpdatedAt:     userRecommend.UpdatedAt,
	}, nil
}
