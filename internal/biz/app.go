package biz

import (
	"context"
	"crypto/rand"
	pb "game/api/app/v1"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	jwt2 "github.com/golang-jwt/jwt/v4"
	"math/big"
	"strconv"
	"time"
)

type User struct {
	ID               uint64
	Address          string
	Level            uint64
	Giw              float64
	GiwAdd           float64
	Git              float64
	Total            float64
	TotalOne         float64
	TotalTwo         float64
	TotalThree       float64
	RewardOne        float64
	RewardTwo        float64
	RewardThree      float64
	RewardTwoOne     float64
	RewardTwoTwo     float64
	RewardTwoThree   float64
	RewardThreeOne   float64
	RewardThreeTwo   float64
	RewardThreeThree float64
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

type BoxRecord struct {
	ID        uint64
	UserId    uint64
	Num       uint64
	GoodId    uint64
	GoodType  uint64
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

type Config struct {
	ID      int64
	KeyName string
	Name    string
	Value   string
}

type SkateGit struct {
	ID        uint64
	UserId    uint64
	Status    uint64
	Amount    float64
	Reward    float64
	CreatedAt time.Time
	UpdatedAt time.Time
}

type SkateGet struct {
	ID        uint64
	UserId    uint64
	Status    uint64
	SkateRate float64
	CreatedAt time.Time
	UpdatedAt time.Time
}

type SkateGetTotal struct {
	ID        uint64
	Amount    float64
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserRepo interface {
	GetAllUsers(ctx context.Context) ([]*User, error)
	GetUserByUserIds(ctx context.Context, userIds []uint64) (map[uint64]*User, error)
	GetUserByAddress(ctx context.Context, address string) (*User, error)
	GetUserRecommendByUserId(ctx context.Context, userId uint64) (*UserRecommend, error)
	GetUserRecommendByCode(ctx context.Context, code string) ([]*UserRecommend, error)
	GetUserRecommends(ctx context.Context) ([]*UserRecommend, error)
	CreateUser(ctx context.Context, uc *User) (*User, error)
	CreateUserRecommend(ctx context.Context, user *User, recommendUser *UserRecommend) (*UserRecommend, error)
	GetConfigByKeys(ctx context.Context, keys ...string) ([]*Config, error)
	GetSkateGitByUserId(ctx context.Context, userId uint64) (*SkateGit, error)
	GetBoxRecord(ctx context.Context, num uint64) ([]*BoxRecord, error)
	GetSkateGetTotal(ctx context.Context) (*SkateGetTotal, error)
	GetUserSkateGet(ctx context.Context, userId uint64) (*SkateGet, error)
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

func RandomBoolCrypto() (bool, error) {
	n, err := rand.Int(rand.Reader, big.NewInt(100))
	if err != nil {
		return false, err
	}
	return n.Int64() < 5, nil
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

			if 0 >= rUser.GiwAdd {
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
		user            *User
		boxNum          uint64
		configs         []*Config
		bPrice          float64
		exchangeFeeRate float64
		rewardSkateRate float64
		boxMax          uint64
		boxAmount       float64
		boxStart        string
		boxEnd          string
		skateOverRate   float64
		sellFeeRate     float64
		withdrawMin     uint64
		withdrawMax     uint64
		err             error
	)
	user, err = ac.userRepo.GetUserByAddress(ctx, address) // 查询用户
	if nil != err || nil == user {
		return &pb.UserInfoReply{
			Status: "不存在用户",
		}, nil
	}

	// 配置
	configs, err = ac.userRepo.GetConfigByKeys(ctx,
		"box_num",
		"b_price",
		"exchange_fee_rate",
		"reward_skate_rate",
		"box_max",
		"box_start",
		"box_end",
		"box_amount",
		"skate_over_rate",
		"sell_fee_rate",
		"withdraw_amount_min",
		"withdraw_amount_max",
	)
	if nil != err || nil == configs {
		return &pb.UserInfoReply{
			Status: "配置错误",
		}, nil
	}

	for _, vConfig := range configs {
		if "box_num" == vConfig.KeyName {
			boxNum, _ = strconv.ParseUint(vConfig.Value, 10, 64)
		}
		if "withdraw_amount_min" == vConfig.KeyName {
			withdrawMin, _ = strconv.ParseUint(vConfig.Value, 10, 64)
		}
		if "withdraw_amount_max" == vConfig.KeyName {
			withdrawMax, _ = strconv.ParseUint(vConfig.Value, 10, 64)
		}
		if "b_price" == vConfig.KeyName {
			bPrice, _ = strconv.ParseFloat(vConfig.Value, 10)
		}
		if "exchange_fee_rate" == vConfig.KeyName {
			exchangeFeeRate, _ = strconv.ParseFloat(vConfig.Value, 10)
		}
		if "reward_skate_rate" == vConfig.KeyName {
			rewardSkateRate, _ = strconv.ParseFloat(vConfig.Value, 10)
		}
		if "box_start" == vConfig.KeyName {
			boxStart = vConfig.Value
		}
		if "box_end" == vConfig.KeyName {
			boxEnd = vConfig.Value
		}
		if "box_amount" == vConfig.KeyName {
			boxAmount, _ = strconv.ParseFloat(vConfig.Value, 10)
		}
		if "box_max" == vConfig.KeyName {
			boxMax, _ = strconv.ParseUint(vConfig.Value, 10, 64)
		}
		if "skate_over_rate" == vConfig.KeyName {
			skateOverRate, _ = strconv.ParseFloat(vConfig.Value, 10)
		}
		if "sell_fee_rate" == vConfig.KeyName {
			sellFeeRate, _ = strconv.ParseFloat(vConfig.Value, 10)
		}
	}

	// 推荐
	var (
		userRecommend   *UserRecommend
		myUserRecommend []*UserRecommend
	)
	userRecommend, err = ac.userRepo.GetUserRecommendByUserId(ctx, user.ID)
	if nil == userRecommend || nil != err {
		return &pb.UserInfoReply{
			Status: "推荐错误查询",
		}, nil
	}

	myUserRecommend, err = ac.userRepo.GetUserRecommendByCode(ctx, userRecommend.RecommendCode+"D"+strconv.FormatUint(user.ID, 10))
	if nil == myUserRecommend || nil != err {
		return &pb.UserInfoReply{
			Status: "推荐错误查询",
		}, nil
	}

	RecommendTotalRewardOne := user.RewardOne + user.RewardTwo + user.RewardThree
	RecommendTotalRewardTwo := user.RewardTwoOne + user.RewardTwoTwo + user.RewardTwoThree
	RecommendTotalRewardThree := user.RewardThreeOne + user.RewardThreeTwo + user.RewardThreeThree
	RecommendTotalReward := +RecommendTotalRewardOne + RecommendTotalRewardTwo + RecommendTotalRewardThree

	var (
		skateGit *SkateGit
	)
	skateGit, err = ac.userRepo.GetSkateGitByUserId(ctx, user.ID)
	if nil != err {
		return &pb.UserInfoReply{
			Status: "粮仓错误查询",
		}, nil
	}
	skateGitAmount := float64(0)
	if nil != skateGit {
		skateGitAmount = skateGit.Amount
	}

	boxSellNum := uint64(0)
	if boxNum > 0 {
		var (
			boxRecord []*BoxRecord
		)
		boxRecord, err = ac.userRepo.GetBoxRecord(ctx, boxNum)
		if nil == boxRecord || nil != err {
			return &pb.UserInfoReply{
				Status: "盲盒错误查询",
			}, nil
		}

		boxSellNum = uint64(len(boxRecord))
	}

	var (
		skateGetTotal       *SkateGetTotal
		skateGet            *SkateGet
		skateGetTotalAmount float64
		skateGetTotalMy     float64
	)
	skateGetTotal, err = ac.userRepo.GetSkateGetTotal(ctx)
	if nil == skateGetTotal || nil != err {
		return &pb.UserInfoReply{
			Status: "放大器错误查询",
		}, nil
	}

	skateGet, err = ac.userRepo.GetUserSkateGet(ctx, user.ID)
	if nil != err {
		return &pb.UserInfoReply{
			Status: "我的放大器错误查询",
		}, nil
	}
	skateGetTotalAmount = skateGetTotal.Amount
	if nil != skateGet {
		skateGetTotalMy = skateGet.SkateRate * skateGetTotalAmount
	}

	return &pb.UserInfoReply{
		Status:                    "ok",
		MyAddress:                 user.Address,
		Level:                     user.Level,
		Giw:                       user.Giw,
		Git:                       user.Git,
		RecommendTotal:            uint64(len(myUserRecommend)),
		RecommendTotalBiw:         user.Total,
		RecommendTotalReward:      RecommendTotalReward,
		RecommendTotalBiwOne:      user.TotalOne,
		RecommendTotalRewardOne:   RecommendTotalRewardOne,
		RecommendTotalBiwTwo:      user.TotalTwo,
		RecommendTotalRewardTwo:   RecommendTotalRewardTwo,
		RecommendTotalBiwThree:    user.TotalThree,
		RecommendTotalRewardThree: RecommendTotalRewardThree,
		MyStakeGit:                skateGitAmount,
		TodayRewardSkateGit:       skateGitAmount * rewardSkateRate,
		RewardStakeRate:           rewardSkateRate,
		Box:                       boxMax,
		BoxSell:                   boxSellNum,
		Start:                     boxStart,
		End:                       boxEnd,
		BoxSellAmount:             boxAmount,
		ExchangeRate:              bPrice,
		ExchangeFeeRate:           exchangeFeeRate,
		StakeGetTotal:             skateGetTotalAmount,
		MyStakeGetTotal:           skateGetTotalMy,
		StakeGetOverFeeRate:       skateOverRate,
		SellFeeRate:               sellFeeRate,
		WithdrawMax:               withdrawMax,
		WithdrawMin:               withdrawMin,
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

// UserMarketSeedList userMarketSeedList.
func (ac *AppUsecase) UserMarketSeedList(ctx context.Context, address string, req *pb.UserMarketSeedListRequest) (*pb.UserMarketSeedListReply, error) {
	var (
		user *User
		err  error
	)
	user, err = ac.userRepo.GetUserByAddress(ctx, address) // 查询用户
	if nil != err || nil == user {
		return &pb.UserMarketSeedListReply{
			Status: "不存在用户",
		}, nil
	}

}

// UserMarketLandList userMarketLandList.
func (ac *AppUsecase) UserMarketLandList(ctx context.Context, address string, req *pb.UserMarketLandListRequest) (*pb.UserMarketLandListReply, error) {
	// 在上下文 context 中取出 claims 对象
	var (
		address string
	)
	if claims, ok := jwt.FromContext(ctx); ok {
		c := claims.(jwt2.MapClaims)
		if c["Address"] == nil {
			return &pb.UserBackListReply{Status: "无效token"}, nil
		}

		address = c["Address"].(string)

		// 验证
		var (
			res bool
			err error
		)
		res, err = addressCheck(address)
		if nil != err {
			return &pb.UserBackListReply{Status: "无效token"}, nil
		}

		if !res {
			return &pb.UserBackListReply{Status: "无效token"}, nil
		}
	} else {
		return &pb.UserBackListReply{Status: "无效token"}, nil
	}

	return a.ac.UserBackList(ctx, address, req)
}

// UserMarketPropList userMarketPropList.
func (ac *AppUsecase) UserMarketPropList(ctx context.Context, address string, req *pb.UserMarketPropListRequest) (*pb.UserMarketPropListReply, error) {
	// 在上下文 context 中取出 claims 对象
	var (
		address string
	)
	if claims, ok := jwt.FromContext(ctx); ok {
		c := claims.(jwt2.MapClaims)
		if c["Address"] == nil {
			return &pb.UserBackListReply{Status: "无效token"}, nil
		}

		address = c["Address"].(string)

		// 验证
		var (
			res bool
			err error
		)
		res, err = addressCheck(address)
		if nil != err {
			return &pb.UserBackListReply{Status: "无效token"}, nil
		}

		if !res {
			return &pb.UserBackListReply{Status: "无效token"}, nil
		}
	} else {
		return &pb.UserBackListReply{Status: "无效token"}, nil
	}

	return a.ac.UserBackList(ctx, address, req)
}

// UserMarketRentLandList userMarketRentLandList.
func (ac *AppUsecase) UserMarketRentLandList(ctx context.Context, address string, req *pb.UserMarketRentLandListRequest) (*pb.UserMarketRentLandListReply, error) {
	// 在上下文 context 中取出 claims 对象
	var (
		address string
	)
	if claims, ok := jwt.FromContext(ctx); ok {
		c := claims.(jwt2.MapClaims)
		if c["Address"] == nil {
			return &pb.UserBackListReply{Status: "无效token"}, nil
		}

		address = c["Address"].(string)

		// 验证
		var (
			res bool
			err error
		)
		res, err = addressCheck(address)
		if nil != err {
			return &pb.UserBackListReply{Status: "无效token"}, nil
		}

		if !res {
			return &pb.UserBackListReply{Status: "无效token"}, nil
		}
	} else {
		return &pb.UserBackListReply{Status: "无效token"}, nil
	}

	return a.ac.UserBackList(ctx, address, req)
}

// UserMyMarketList userMyMarketList.
func (ac *AppUsecase) UserMyMarketList(ctx context.Context, address string, req *pb.UserMyMarketListRequest) (*pb.UserMyMarketListReply, error) {
	// 在上下文 context 中取出 claims 对象
	var (
		address string
	)
	if claims, ok := jwt.FromContext(ctx); ok {
		c := claims.(jwt2.MapClaims)
		if c["Address"] == nil {
			return &pb.UserBackListReply{Status: "无效token"}, nil
		}

		address = c["Address"].(string)

		// 验证
		var (
			res bool
			err error
		)
		res, err = addressCheck(address)
		if nil != err {
			return &pb.UserBackListReply{Status: "无效token"}, nil
		}

		if !res {
			return &pb.UserBackListReply{Status: "无效token"}, nil
		}
	} else {
		return &pb.UserBackListReply{Status: "无效token"}, nil
	}

	return a.ac.UserBackList(ctx, address, req)
}

// UserNoticeList NoticeList.
func (ac *AppUsecase) UserNoticeList(ctx context.Context, address string, req *pb.UserNoticeListRequest) (*pb.UserNoticeListReply, error) {
	// 在上下文 context 中取出 claims 对象
	var (
		address string
	)
	if claims, ok := jwt.FromContext(ctx); ok {
		c := claims.(jwt2.MapClaims)
		if c["Address"] == nil {
			return &pb.UserBackListReply{Status: "无效token"}, nil
		}

		address = c["Address"].(string)

		// 验证
		var (
			res bool
			err error
		)
		res, err = addressCheck(address)
		if nil != err {
			return &pb.UserBackListReply{Status: "无效token"}, nil
		}

		if !res {
			return &pb.UserBackListReply{Status: "无效token"}, nil
		}
	} else {
		return &pb.UserBackListReply{Status: "无效token"}, nil
	}

	return a.ac.UserBackList(ctx, address, req)
}

// UserSkateRewardList userSkateRewardList.
func (ac *AppUsecase) UserSkateRewardList(ctx context.Context, address string, req *pb.UserSkateRewardListRequest) (*pb.UserSkateRewardListReply, error) {
	// 在上下文 context 中取出 claims 对象
	var (
		address string
	)
	if claims, ok := jwt.FromContext(ctx); ok {
		c := claims.(jwt2.MapClaims)
		if c["Address"] == nil {
			return &pb.UserBackListReply{Status: "无效token"}, nil
		}

		address = c["Address"].(string)

		// 验证
		var (
			res bool
			err error
		)
		res, err = addressCheck(address)
		if nil != err {
			return &pb.UserBackListReply{Status: "无效token"}, nil
		}

		if !res {
			return &pb.UserBackListReply{Status: "无效token"}, nil
		}
	} else {
		return &pb.UserBackListReply{Status: "无效token"}, nil
	}

	return a.ac.UserBackList(ctx, address, req)
}
