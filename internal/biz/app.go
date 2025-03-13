package biz

import (
	"context"
	pb "game/api/app/v1"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"time"
)

type User struct {
	ID        uint64
	Address   string
	Level     uint64
	Giw       float64
	Git       float64
	Total     float64
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserRecommend struct {
	ID            uint64
	UserId        uint64
	RecommendCode string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type Land struct {
	ID             uint64
	UserId         uint64
	Level          uint64
	OutPutRate     float64
	RentOutPutRate float64
	MaxHealth      uint64
	PerHealth      uint64
	LimitDate      uint64
	Status         uint64
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type Prop struct {
	ID        uint64
	UserId    uint64
	Status    uint64
	PropType  uint64
	OneOne    uint64
	OneTwo    uint64
	TwoOne    uint64
	TwoTwo    float64
	ThreeOne  uint64
	FourOne   uint64
	FiveOne   uint64
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Seed struct {
	ID           uint64
	UserId       uint64
	SeedId       uint64
	Name         string
	OutAmount    float64
	OutOverTime  uint64
	OutMaxAmount uint64
	OutMinAmount uint64
	Status       uint64
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type UserRepo interface {
	GetUserByAddress(ctx context.Context, address string) (*User, error)
	GetUserRecommendByUserId(ctx context.Context, userId uint64) (*UserRecommend, error)
	CreateUser(ctx context.Context, uc *User) (*User, error)
	CreateUserRecommend(ctx context.Context, user *User, recommendUser *UserRecommend) (*UserRecommend, error)
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

func (ac *AppUsecase) GetExistUserByAddressOrCreate(ctx context.Context, address string, req *pb.EthAuthorizeRequest) (*User, error, string) {
	var (
		user            *User
		rURecommendUser *UserRecommend
		err             error
	)

	user, err = ac.userRepo.GetUserByAddress(ctx, address) // 查询用户
	if nil == user || nil != err {
		code := req.SendBody.Code
		if "abf00dd52c08a9213f225827bc3fb100" != code {
			if 1 >= len(code) {
				return nil, errors.New(500, "USER_ERROR", "无效的推荐码1"), "无效的推荐码"
			}

			var (
				rUser *User
			)

			rUser, err = ac.userRepo.GetUserByAddress(ctx, code)
			if nil == rUser || err != nil {
				return nil, errors.New(500, "USER_ERROR", "无效的推荐码1"), "无效的推荐码"
			}

			if 0 >= rUser.Giw {
				return nil, errors.New(500, "USER_ERROR", "推荐人未入金"), "推荐人未入金"
			}

			// 查询推荐人的相关信息
			rURecommendUser, err = ac.userRepo.GetUserRecommendByUserId(ctx, rUser.ID)
			if nil == rURecommendUser || err != nil {
				return nil, errors.New(500, "USER_ERROR", "无效的推荐码3"), "无效的推荐码3"
			}
		}

		if err = ac.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
			user, err = ac.userRepo.CreateUser(ctx, &User{
				Address: address,
			})
			if err != nil {
				return err
			}

			_, err = ac.userRepo.CreateUserRecommend(ctx, user, rURecommendUser) // 创建用户推荐信息
			if err != nil {
				return err
			}

			return nil
		}); err != nil {
			return nil, err, "错误"
		}
	}

	return user, nil, ""
}

func (ac *AppUsecase) UserInfo(ctx context.Context, address string) (*pb.UserInfoReply, error) {
	var (
		user *User
		err  error
	)
	user, err = ac.userRepo.GetUserByAddress(ctx, address) // 查询用户
	if nil != err || nil == user {
		return &pb.UserInfoReply{
			Status: "不存在用户",
		}, nil
	}

	return &pb.UserInfoReply{
		Status:                    "ok",
		MyAddress:                 user.Address,
		Level:                     user.Level,
		Giw:                       user.Giw,
		Git:                       user.Git,
		RecommendTotal:            262,
		RecommendTotalBiw:         464,
		RecommendTotalReward:      33663,
		RecommendTotalBiwOne:      225,
		RecommendTotalRewardOne:   3433,
		RecommendTotalBiwTwo:      34,
		RecommendTotalRewardTwo:   6666,
		RecommendTotalBiwThree:    324,
		RecommendTotalRewardThree: 5324,
		MyStakeGit:                130777,
		TodayRewardSkateGit:       1377774,
		RewardStakeRate:           0.3,
		Box:                       123,
		BoxSell:                   5,
		Start:                     "2025-03-01 00:00:00",
		End:                       "2025-05-01 00:00:00",
		BoxSellAmount:             343432,
		ExchangeRate:              0.21,
		ExchangeFeeRate:           0.04,
		StakeGetTotal:             2134342,
		MyStakeGetTotal:           332,
		StakeGetOverFeeRate:       0.5,
		SellFeeRate:               0.05,
	}, nil
}

func (ac *AppUsecase) UserRecommend(ctx context.Context, address string, req *pb.UserRecommendRequest) (*pb.UserRecommendReply, error) {
	res := make([]*pb.UserRecommendReply_List, 0)
	var (
		user *User
		err  error
	)
	user, err = ac.userRepo.GetUserByAddress(ctx, address) // 查询用户
	if nil != err || nil == user {
		return &pb.UserRecommendReply{
			Status: "不存在用户",
		}, nil
	}

	res = append(res, &pb.UserRecommendReply_List{
		Address:   "aaaa",
		Level:     2,
		CreatedAt: "2023-03-03 00:00:00",
	}, &pb.UserRecommendReply_List{
		Address:   "bbbb",
		Level:     4,
		CreatedAt: "2023-03-03 00:00:00",
	})

	return &pb.UserRecommendReply{
		Status: "ok",
		Count:  2,
		List:   res,
	}, nil
}

func (ac *AppUsecase) UserRecommendL(ctx context.Context, address string, req *pb.UserRecommendLRequest) (*pb.UserRecommendLReply, error) {
	res := make([]*pb.UserRecommendLReply_List, 0)
	var (
		user *User
		err  error
	)
	user, err = ac.userRepo.GetUserByAddress(ctx, address) // 查询用户
	if nil != err || nil == user {
		return &pb.UserRecommendLReply{
			Status: "不存在用户",
		}, nil
	}

	res = append(res, &pb.UserRecommendLReply_List{
		Address:   "aaa",
		Amount:    5.3,
		CreatedAt: "2023-03-03 00:00:00",
	}, &pb.UserRecommendLReply_List{
		Address:   "aaa",
		Amount:    10.2,
		CreatedAt: "2023-03-03 00:00:00",
	})

	return &pb.UserRecommendLReply{
		Status: "ok",
		Count:  2,
		List:   res,
	}, nil
}

func (ac *AppUsecase) UserLand(ctx context.Context, address string, req *pb.UserLandRequest) (*pb.UserLandReply, error) {
	res := make([]*pb.UserLandReply_List, 0)
	var (
		user *User
		err  error
	)
	user, err = ac.userRepo.GetUserByAddress(ctx, address) // 查询用户
	if nil != err || nil == user {
		return &pb.UserLandReply{
			Status: "不存在用户",
		}, nil
	}

	res = append(res, &pb.UserLandReply_List{
		Id:        0,
		Level:     1,
		Health:    120,
		Status:    1,
		OutRate:   120,
		PerHealth: 20,
	}, &pb.UserLandReply_List{
		Id:        0,
		Level:     5,
		Health:    120,
		Status:    4,
		OutRate:   120,
		PerHealth: 20,
	}, &pb.UserLandReply_List{
		Id:        0,
		Level:     5,
		Health:    120,
		Status:    3,
		OutRate:   120,
		PerHealth: 20,
	})

	return &pb.UserLandReply{
		Status: "ok",
		Count:  3,
		List:   res,
	}, nil
}

func (ac *AppUsecase) UserStakeRewardList(ctx context.Context, address string, req *pb.UserStakeRewardListRequest) (*pb.UserStakeRewardListReply, error) {
	res := make([]*pb.UserStakeRewardListReply_List, 0)
	var (
		user *User
		err  error
	)
	user, err = ac.userRepo.GetUserByAddress(ctx, address) // 查询用户
	if nil != err || nil == user {
		return &pb.UserStakeRewardListReply{
			Status: "不存在用户",
		}, nil
	}

	res = append(res, &pb.UserStakeRewardListReply_List{
		Amount:    1000,
		CreatedAt: "2025-03-01 00:00:00",
	}, &pb.UserStakeRewardListReply_List{
		Amount:    23343,
		CreatedAt: "2025-03-01 00:00:00",
	})

	return &pb.UserStakeRewardListReply{
		Status: "ok",
		Count:  2,
		List:   res,
	}, nil
}

func (ac *AppUsecase) UserBoxList(ctx context.Context, address string, req *pb.UserBoxListRequest) (*pb.UserBoxListReply, error) {
	res := make([]*pb.UserBoxListReply_List, 0)
	var (
		user *User
		err  error
	)
	user, err = ac.userRepo.GetUserByAddress(ctx, address) // 查询用户
	if nil != err || nil == user {
		return &pb.UserBoxListReply{
			Status: "不存在用户",
		}, nil
	}

	res = append(res, &pb.UserBoxListReply_List{
		Content:   "1个种子",
		CreatedAt: "2025-03-01 00:00:00",
	}, &pb.UserBoxListReply_List{
		Content:   "1个手套",
		CreatedAt: "2025-03-01 00:00:00",
	})

	return &pb.UserBoxListReply{
		Status: "ok",
		Count:  2,
		List:   res,
	}, nil
}

func (ac *AppUsecase) UserBackList(ctx context.Context, address string, req *pb.UserBackListRequest) (*pb.UserBackListReply, error) {
	res := make([]*pb.UserBackListReply_List, 0)
	var (
		user *User
		err  error
	)
	user, err = ac.userRepo.GetUserByAddress(ctx, address) // 查询用户
	if nil != err || nil == user {
		return &pb.UserBackListReply{
			Status: "不存在用户",
		}, nil
	}

	res = append(res, &pb.UserBackListReply_List{
		Id:     10,
		Type:   1,
		Num:    4,
		UseNum: 0,
		Status: 1,
		OutMax: 1232.4,
	}, &pb.UserBackListReply_List{
		Id:     12,
		Type:   2,
		Num:    13,
		UseNum: 18,
		Status: 2,
		OutMax: 0,
	}, &pb.UserBackListReply_List{
		Id:     12,
		Type:   2,
		Num:    15,
		UseNum: 7,
		Status: 4,
		OutMax: 0,
	})

	return &pb.UserBackListReply{
		Status: "ok",
		Count:  3,
		List:   res,
	}, nil
}
