package biz

import (
	"context"
	"crypto/rand"
	"fmt"
	pb "game/api/app/v1"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"math/big"
	rand2 "math/rand"
	"strconv"
	"sync"
	"time"
)

type Pagination struct {
	PageNum  int
	PageSize int
}

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
	Content   string
}

type UserRecommend struct {
	ID            uint64
	UserId        uint64
	RecommendCode string
	CreatedAt     time.Time
	UpdatedAt     time.Time
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

type Reward struct {
	ID        uint64
	UserId    uint64
	reason    uint64
	One       uint64
	Two       uint64
	Three     float64
	Amount    float64
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
	OutMaxAmount float64
	OutMinAmount float64
	Status       uint64
	SellAmount   float64
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type PropInfo struct {
	ID       uint64
	PropType uint64

	// 化肥相关字段
	OneOne uint64
	OneTwo uint64

	// 铲子相关字段
	TwoOne uint64
	TwoTwo float64

	// 水相关字段
	ThreeOne uint64

	// 除虫剂相关字段
	FourOne uint64

	// 手套相关字段
	FiveOne uint64
	GetRate float64

	// 时间字段
	CreatedAt time.Time
	UpdatedAt time.Time
}

type SeedInfo struct {
	ID           uint64
	Name         string
	OutMinAmount float64
	OutMaxAmount float64
	GetRate      float64
	OutOverTime  uint64
	CreatedAt    time.Time
	UpdatedAt    time.Time
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
	LocationNum    uint64
	CreatedAt      time.Time
	UpdatedAt      time.Time
	One            uint64
	Two            uint64
	Three          uint64
	SellAmount     float64
}

type LandUserUse struct {
	ID           uint64
	LandId       uint64
	Level        uint64
	UserId       uint64
	OwnerUserId  uint64
	SeedId       uint64
	SeedTypeId   uint64
	Status       uint64
	BeginTime    uint64
	TotalTime    uint64
	OverTime     uint64
	OutMaxNum    float64
	OutNum       float64
	InsectStatus uint64
	OutSubNum    float64
	StealNum     float64
	StopStatus   uint64
	StopTime     uint64
	SubTime      uint64
	UseChan      uint64
	CreatedAt    time.Time
	UpdatedAt    time.Time
	One          uint64
	Two          uint64
}

type ExchangeRecord struct {
	ID        uint64
	UserId    int64
	Git       float64
	Giw       float64
	Fee       float64
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Market struct {
	ID        uint64
	UserId    uint64
	GoodId    uint64
	GoodType  int
	Amount    float64
	Status    int
	GetUserId uint64
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Notice struct {
	ID            uint64
	UserId        uint64
	NoticeContent string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type Prop struct {
	ID         uint64
	UserId     uint64
	Status     int
	PropType   int
	OneOne     int
	OneTwo     int
	TwoOne     int
	TwoTwo     float64
	SellAmount float64
	ThreeOne   int
	FourOne    int
	FiveOne    int
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type StakeGet struct {
	ID        uint64
	UserId    uint64
	Status    int
	StakeRate float64
	CreatedAt time.Time
	UpdatedAt time.Time
}

type StakeGetPlayRecord struct {
	ID        uint64
	UserId    uint64
	Amount    float64
	Reward    float64
	Status    int
	CreatedAt time.Time
	UpdatedAt time.Time
}

type StakeGetRecord struct {
	ID        uint64
	UserId    uint64
	Amount    float64
	StakeRate float64
	Total     float64
	StakeType int
	CreatedAt time.Time
	UpdatedAt time.Time
}

type StakeGetTotal struct {
	ID        uint64
	Amount    float64
	CreatedAt time.Time
	UpdatedAt time.Time
}

type StakeGit struct {
	ID        uint64
	UserID    uint64
	Amount    float64
	Reward    float64
	Status    int
	CreatedAt time.Time
	UpdatedAt time.Time
}

type StakeGitRecord struct {
	ID        uint64
	UserID    uint64
	Amount    float64
	StakeType int
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Withdraw struct {
	ID        uint64
	UserID    uint64
	Amount    uint64
	RelAmount uint64
	Status    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type RandomSeed struct {
	ID        uint64
	Scene     uint64
	SeedValue uint64
	UpdatedAt time.Time
	CreatedAt time.Time
}

type UserRepo interface {
	GetAllUsers(ctx context.Context) ([]*User, error)
	GetUserByUserIds(ctx context.Context, userIds []uint64) (map[uint64]*User, error)
	GetUserByAddress(ctx context.Context, address string) (*User, error)
	GetUserRecommendByUserId(ctx context.Context, userId uint64) (*UserRecommend, error)
	GetUserRecommendByCode(ctx context.Context, code string) ([]*UserRecommend, error)
	GetUserRecommends(ctx context.Context) ([]*UserRecommend, error)
	CreateUser(ctx context.Context, uc *User) (*User, error)
	CreateSkateGit(ctx context.Context, sg *SkateGit) (*SkateGit, error)
	CreateUserRecommend(ctx context.Context, user *User, recommendUser *UserRecommend) (*UserRecommend, error)
	GetConfigByKeys(ctx context.Context, keys ...string) ([]*Config, error)
	GetSkateGitByUserId(ctx context.Context, userId uint64) (*SkateGit, error)
	GetBoxRecord(ctx context.Context, num uint64) ([]*BoxRecord, error)
	GetBoxRecordCount(ctx context.Context, num uint64) (int64, error)
	GetUserBoxRecord(ctx context.Context, userId, num uint64, b *Pagination) ([]*BoxRecord, error)
	GetUserBoxRecordOpen(ctx context.Context, userId, num uint64, open bool, b *Pagination) ([]*BoxRecord, error)
	GetSkateGetTotal(ctx context.Context) (*SkateGetTotal, error)
	GetUserSkateGet(ctx context.Context, userId uint64) (*SkateGet, error)
	GetUserRecommendCount(ctx context.Context, code string) (int64, error)
	GetUserRecommendByCodePage(ctx context.Context, code string, b *Pagination) ([]*UserRecommend, error)
	GetLandByUserIDUsing(ctx context.Context, userID uint64, status []uint64) ([]*Land, error)
	GetLandByUserID(ctx context.Context, userID uint64, status []uint64, b *Pagination) ([]*Land, error)
	GetLandByExUserID(ctx context.Context, userID uint64, status []uint64, b *Pagination) ([]*Land, error)
	GetUserRewardPage(ctx context.Context, userId uint64, reason []uint64, b *Pagination) ([]*Reward, error)
	GetUserRewardPageCount(ctx context.Context, userId uint64, reason []uint64) (int64, error)
	GetSeedByUserID(ctx context.Context, userID uint64, status []uint64, b *Pagination) ([]*Seed, error)
	GetSeedByExUserID(ctx context.Context, userID uint64, status []uint64, b *Pagination) ([]*Seed, error)
	GetLandUserUseByUserIDUseing(ctx context.Context, userID uint64, status uint64, b *Pagination) ([]*LandUserUse, error)
	GetExchangeRecordsByUserID(ctx context.Context, userID uint64, b *Pagination) ([]*ExchangeRecord, error)
	GetLandUserUseByID(ctx context.Context, id uint64) (*LandUserUse, error)
	GetMarketRecordsByUserID(ctx context.Context, userID uint64, status uint64, b *Pagination) ([]*Market, error)
	GetNoticesByUserID(ctx context.Context, userID uint64, b *Pagination) ([]*Notice, error)
	GetNoticesCountByUserID(ctx context.Context, userID uint64) (int64, error)
	GetPropsByUserID(ctx context.Context, userID uint64, status []uint64, b *Pagination) ([]*Prop, error)
	GetPropsByExUserID(ctx context.Context, userID uint64, status []uint64, b *Pagination) ([]*Prop, error)
	GetStakeGetsByUserID(ctx context.Context, userID uint64, b *Pagination) ([]*StakeGet, error)
	GetStakeGetPlayRecordsByUserID(ctx context.Context, userID uint64, status uint64, b *Pagination) ([]*StakeGetPlayRecord, error)
	GetStakeGetPlayRecordCount(ctx context.Context, userID uint64, status uint64) (int64, error)
	GetStakeGetRecordsByUserID(ctx context.Context, userID int64, b *Pagination) ([]*StakeGetRecord, error)
	GetStakeGetTotal(ctx context.Context) (*StakeGetTotal, error)
	GetStakeGitByUserID(ctx context.Context, userID int64) (*StakeGit, error)
	GetStakeGitRecordsByUserID(ctx context.Context, userID int64, b *Pagination) ([]*StakeGitRecord, error)
	GetWithdrawRecordsByUserID(ctx context.Context, userID int64, b *Pagination) ([]*Withdraw, error)
	GetUserOrderCount(ctx context.Context) (int64, error)
	GetUserOrder(ctx context.Context, b *Pagination) ([]*User, error)
	GetLandUserUseByLandIDsMapUsing(ctx context.Context, userId uint64, landIDs []uint64) (map[uint64]*LandUserUse, error)
	BuyBox(ctx context.Context, giw float64, originValue, value string, uc *BoxRecord) error
	GetUserBoxRecordById(ctx context.Context, id uint64) (*BoxRecord, error)
	OpenBoxSeed(ctx context.Context, id uint64, content string, seedInfo *Seed) error
	OpenBoxProp(ctx context.Context, id uint64, content string, propInfo *Prop) error
	GetAllSeedInfo(ctx context.Context) ([]*SeedInfo, error)
	GetAllPropInfo(ctx context.Context) ([]*PropInfo, error)
	GetAllRandomSeeds(ctx context.Context) ([]*RandomSeed, error)
	UpdateSeedValue(ctx context.Context, scene uint64, newSeed uint64) error
	GetSeedByID(ctx context.Context, seedID, userId, status uint64) (*Seed, error)
	GetLandByID(ctx context.Context, landID uint64) (*Land, error)
	GetLandByIDTwo(ctx context.Context, landID uint64) (*Land, error)
	GetLandByUserIdLocationNum(ctx context.Context, userId uint64, locationNum uint64) (*Land, error)
	Plant(ctx context.Context, status, originStatus, perHealth uint64, landUserUse *LandUserUse) error
	GetSeedBuyByID(ctx context.Context, seedID, status uint64) (*Seed, error)
	GetPropByID(ctx context.Context, propID, status uint64) (*Prop, error)
	BuySeed(ctx context.Context, git, getGit float64, userId, userIdGet, seedId uint64) error
	BuyLand(ctx context.Context, git, getGit float64, userId, userIdGet, landId uint64) error
	BuyProp(ctx context.Context, git, getGit float64, userId, userIdGet, propId uint64) error
	SellLand(ctx context.Context, propId uint64, userId uint64, sellAmount float64) error
	SellProp(ctx context.Context, propId uint64, userId uint64, sellAmount float64) error
	SellSeed(ctx context.Context, seedId uint64, userId uint64, sellAmount float64) error
	GetPropByIDSell(ctx context.Context, propID, status uint64) (*Prop, error)
	UnSellLand(ctx context.Context, propId uint64, userId uint64) error
	UnSellProp(ctx context.Context, propId uint64, userId uint64) error
	UnSellSeed(ctx context.Context, seedId, userId uint64) error
	RentLand(ctx context.Context, landId uint64, userId uint64, rentRate float64) error
	UnRentLand(ctx context.Context, landId uint64, userId uint64) error
	LandPull(ctx context.Context, landId uint64, userId uint64) error
	LandPush(ctx context.Context, landId uint64, userId, locationNum uint64) error
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

			//_, err = ac.userRepo.CreateSkateGit(ctx, &SkateGit{
			//	UserId: user.ID,
			//})
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
		boxSellNum      uint64
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
		"box_sell_num",
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
		if "box_sell_num" == vConfig.KeyName {
			boxSellNum, _ = strconv.ParseUint(vConfig.Value, 10, 64)
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

	if boxNum > 0 {
		//var (
		//	countBox int64
		//)
		//countBox, err = ac.userRepo.GetBoxRecordCount(ctx, boxNum)
		//if nil != err {
		//	return &pb.UserInfoReply{
		//		Status: "盲盒错误查询",
		//	}, nil
		//}
		//
		//boxSellNum = uint64(countBox)
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

	// 推荐
	var (
		userRecommend  *UserRecommend
		userRecommends []*UserRecommend
		count          int64
	)
	userRecommend, err = ac.userRepo.GetUserRecommendByUserId(ctx, user.ID)
	if nil == userRecommend || nil != err {
		return &pb.UserRecommendReply{
			Status: "推荐错误查询",
		}, nil
	}

	count, err = ac.userRepo.GetUserRecommendCount(ctx, userRecommend.RecommendCode+"D"+strconv.FormatUint(user.ID, 10))
	if nil != err {
		return &pb.UserRecommendReply{
			Status: "推荐错误查询",
		}, nil
	}

	userRecommends, err = ac.userRepo.GetUserRecommendByCodePage(ctx, userRecommend.RecommendCode+"D"+strconv.FormatUint(user.ID, 10), &Pagination{
		PageNum:  int(req.Page),
		PageSize: 20,
	})
	if nil != err {
		return &pb.UserRecommendReply{
			Status: "推荐错误查询",
		}, nil
	}

	userIds := make([]uint64, 0)
	for _, v := range userRecommends {
		userIds = append(userIds, v.UserId)
	}

	usersMap := make(map[uint64]*User)
	usersMap, err = ac.userRepo.GetUserByUserIds(ctx, userIds)
	if nil != err {
		return &pb.UserRecommendReply{
			Status: "推荐用户查询错误",
		}, nil
	}

	for _, v := range usersMap {
		res = append(res, &pb.UserRecommendReply_List{
			Address:   v.Address,
			Level:     v.Level,
			CreatedAt: v.CreatedAt.Add(8 * time.Hour).Format("2006-01-02 15:04:05"),
		})
	}

	return &pb.UserRecommendReply{
		Status: "ok",
		Count:  uint64(count),
		List:   res,
	}, nil
}

func (ac *AppUsecase) UserRecommendL(ctx context.Context, address string, req *pb.UserRecommendLRequest) (*pb.UserRecommendLReply, error) {
	res := make([]*pb.UserRecommendLReply_List, 0)
	var (
		user  *User
		err   error
		count int64
	)
	user, err = ac.userRepo.GetUserByAddress(ctx, address) // 查询用户
	if nil != err || nil == user {
		return &pb.UserRecommendLReply{
			Status: "不存在用户",
		}, nil
	}

	var (
		reward []*Reward
	)

	status := []uint64{}
	if 1 == req.Num {
		status = append(status, 4, 5, 6)
	} else if 2 == req.Num {
		status = append(status, 7, 8, 9)
	} else if 3 == req.Num {
		status = append(status, 10, 11, 12)
	} else {
		return &pb.UserRecommendLReply{
			Status: "参数错误",
		}, nil
	}

	count, err = ac.userRepo.GetUserRewardPageCount(ctx, user.ID, status)
	if nil != err {
		return &pb.UserRecommendLReply{
			Status: "不存在数据L，count",
		}, nil
	}

	reward, err = ac.userRepo.GetUserRewardPage(ctx, user.ID, status, &Pagination{
		PageNum:  int(req.Page),
		PageSize: 20,
	})
	if nil != err {
		return &pb.UserRecommendLReply{
			Status: "不存在数据L",
		}, nil
	}

	userIds := []uint64{}
	for _, v := range reward {
		userIds = append(userIds, v.One)
	}

	usersMap := make(map[uint64]*User)
	usersMap, err = ac.userRepo.GetUserByUserIds(ctx, userIds)
	if nil != err {
		return &pb.UserRecommendLReply{
			Status: "不存在数据L,用户",
		}, nil
	}

	for _, v := range reward {
		tmpAddress := ""
		if _, ok := usersMap[v.One]; ok {
			tmpAddress = usersMap[v.One].Address
		}

		res = append(res, &pb.UserRecommendLReply_List{
			Address:   tmpAddress,
			Amount:    v.Amount,
			CreatedAt: v.CreatedAt.Add(8 * time.Hour).Format("2006-01-02 15:04:05"),
		})
	}

	return &pb.UserRecommendLReply{
		Status: "ok",
		Count:  uint64(count),
		List:   res,
	}, nil
}

func (ac *AppUsecase) UserLand(ctx context.Context, address string, req *pb.UserLandRequest) (*pb.UserLandReply, error) {
	res := make([]*pb.UserLandReply_List, 0)
	var (
		user  *User
		lands []*Land
		err   error
	)
	user, err = ac.userRepo.GetUserByAddress(ctx, address) // 查询用户
	if nil != err || nil == user {
		return &pb.UserLandReply{
			Status: "不存在用户",
		}, nil
	}
	status := []uint64{0, 1, 2, 3, 4, 5, 8}
	lands, err = ac.userRepo.GetLandByUserID(ctx, user.ID, status, nil)
	if nil != err {
		return &pb.UserLandReply{
			Status: "不存在用户",
		}, nil
	}

	for _, v := range lands {
		statusTmp := v.Status
		if 8 == v.Status {
			statusTmp = 3
		}

		res = append(res, &pb.UserLandReply_List{
			Id:        v.ID,
			Level:     v.Level,
			Health:    v.MaxHealth,
			Status:    statusTmp,
			OutRate:   v.OutPutRate,
			PerHealth: v.PerHealth,
			One:       v.One,
			Two:       v.Two,
			Three:     v.Three,
		})
	}

	return &pb.UserLandReply{
		Status: "ok",
		Count:  uint64(len(res)),
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

	var (
		reward []*Reward
		count  int64
	)

	status := []uint64{3}
	count, err = ac.userRepo.GetUserRewardPageCount(ctx, user.ID, status)
	if nil != err {
		return &pb.UserStakeRewardListReply{
			Status: "粮仓查询错误",
		}, nil
	}

	reward, err = ac.userRepo.GetUserRewardPage(ctx, user.ID, status, &Pagination{
		PageNum:  int(req.Page),
		PageSize: 20,
	})
	if nil != err {
		return &pb.UserStakeRewardListReply{
			Status: "粮仓查询错误",
		}, nil
	}

	for _, v := range reward {
		res = append(res, &pb.UserStakeRewardListReply_List{
			Amount:    v.Amount,
			CreatedAt: v.CreatedAt.Add(8 * time.Hour).Format("2006-01-02 15:04:05"),
		})
	}

	return &pb.UserStakeRewardListReply{
		Status: "ok",
		Count:  uint64(count),
		List:   res,
	}, nil
}

func (ac *AppUsecase) UserBoxList(ctx context.Context, address string, req *pb.UserBoxListRequest) (*pb.UserBoxListReply, error) {
	res := make([]*pb.UserBoxListReply_List, 0)
	var (
		boxNum  uint64
		configs []*Config
		user    *User
		err     error
	)
	user, err = ac.userRepo.GetUserByAddress(ctx, address) // 查询用户
	if nil != err || nil == user {
		return &pb.UserBoxListReply{
			Status: "不存在用户",
		}, nil
	}

	configs, err = ac.userRepo.GetConfigByKeys(ctx,
		"box_num",
	)
	if nil != err || nil == configs {
		return &pb.UserBoxListReply{
			Status: "配置错误",
		}, nil
	}

	for _, vConfig := range configs {
		if "box_num" == vConfig.KeyName {
			boxNum, _ = strconv.ParseUint(vConfig.Value, 10, 64)
		}
	}

	if 0 < boxNum {
		var (
			box []*BoxRecord
		)
		box, err = ac.userRepo.GetUserBoxRecordOpen(ctx, user.ID, boxNum, true, &Pagination{
			PageNum:  1,
			PageSize: 20,
		})
		if nil != err {
			return &pb.UserBoxListReply{
				Status: "查询错误",
			}, nil
		}

		for _, v := range box {
			res = append(res, &pb.UserBoxListReply_List{
				Content:   v.Content,
				CreatedAt: v.CreatedAt.Add(8 * time.Hour).Format("2006-01-02 15:04:05"),
			})
		}
	}

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

	var (
		seed []*Seed
	)
	seedStatus := []uint64{0, 4}
	seed, err = ac.userRepo.GetSeedByUserID(ctx, user.ID, seedStatus, nil)
	if nil != err {
		return &pb.UserBackListReply{
			Status: "查询种子错误",
		}, nil
	}

	for _, vSeed := range seed {
		tmpStatus := uint64(1)
		if 4 == vSeed.Status {
			tmpStatus = 4
		}

		res = append(res, &pb.UserBackListReply_List{
			Id:     vSeed.ID,
			Type:   1,
			Num:    vSeed.SeedId,
			UseNum: 0,
			Status: tmpStatus,
			OutMax: vSeed.OutMaxAmount,
		})
	}

	var (
		prop []*Prop
	)
	// 11化肥，12水，13手套，14除虫剂，15铲子，16盲盒，17地契
	propStatus := []uint64{1, 2, 4}
	prop, err = ac.userRepo.GetPropsByUserID(ctx, user.ID, propStatus, nil)
	if nil != err {
		return &pb.UserBackListReply{
			Status: "道具错误",
		}, nil
	}

	for _, vProp := range prop {

		useNum := uint64(0)
		if 12 == vProp.PropType {
			useNum = uint64(vProp.ThreeOne) // 水
		} else if 13 == vProp.PropType {
			useNum = uint64(vProp.FiveOne) // 手套
		} else if 14 == vProp.PropType {
			useNum = uint64(vProp.FourOne) // 除虫剂
		} else if 15 == vProp.PropType {
			useNum = uint64(vProp.TwoOne) // 铲子
		}

		res = append(res, &pb.UserBackListReply_List{
			Id:     vProp.ID,
			Type:   2,
			Num:    uint64(vProp.PropType),
			UseNum: useNum,
			Status: uint64(vProp.Status),
			OutMax: 0,
		})
	}

	var (
		box []*BoxRecord
	)

	box, err = ac.userRepo.GetUserBoxRecordOpen(ctx, user.ID, 0, false, nil)
	if nil != err {
		return &pb.UserBackListReply{
			Status: "查询盒子错误",
		}, nil
	}

	for _, v := range box {
		res = append(res, &pb.UserBackListReply_List{
			Id:     v.ID,
			Type:   2,
			Num:    16,
			UseNum: 0,
			Status: 0,
			OutMax: 0,
		})
	}

	return &pb.UserBackListReply{
		Status: "ok",
		Count:  0,
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
	res := make([]*pb.UserMarketSeedListReply_List, 0)
	var (
		seed []*Seed
	)

	seedStatus := []uint64{4}
	seed, err = ac.userRepo.GetSeedByExUserID(ctx, user.ID, seedStatus, nil)
	if nil != err {
		return &pb.UserMarketSeedListReply{
			Status: "查询错误",
		}, nil
	}

	for _, vSeed := range seed {
		res = append(res, &pb.UserMarketSeedListReply_List{
			Id:     vSeed.ID,
			Num:    vSeed.SeedId,
			Amount: vSeed.SellAmount,
			OutMax: vSeed.OutMaxAmount,
		})
	}

	return &pb.UserMarketSeedListReply{
		Status: "ok",
		Count:  0,
		List:   res,
	}, nil
}

// UserMarketLandList userMarketLandList.
func (ac *AppUsecase) UserMarketLandList(ctx context.Context, address string, req *pb.UserMarketLandListRequest) (*pb.UserMarketLandListReply, error) {
	var (
		user *User
		err  error
	)
	user, err = ac.userRepo.GetUserByAddress(ctx, address) // 查询用户
	if nil != err || nil == user {
		return &pb.UserMarketLandListReply{
			Status: "不存在用户",
		}, nil
	}

	res := make([]*pb.UserMarketLandListReply_List, 0)
	var (
		land []*Land
	)
	landStatus := []uint64{4}
	land, err = ac.userRepo.GetLandByExUserID(ctx, user.ID, landStatus, nil)
	if nil != err {
		return &pb.UserMarketLandListReply{
			Status: "错误查询",
		}, nil
	}

	for _, vLand := range land {
		res = append(res, &pb.UserMarketLandListReply_List{
			Id:        vLand.ID,
			Level:     vLand.Level,
			MaxHealth: vLand.MaxHealth,
			Amount:    vLand.SellAmount,
		})
	}

	return &pb.UserMarketLandListReply{
		Status: "ok",
		Count:  0,
		List:   res,
	}, nil
}

// UserMarketPropList userMarketPropList.
func (ac *AppUsecase) UserMarketPropList(ctx context.Context, address string, req *pb.UserMarketPropListRequest) (*pb.UserMarketPropListReply, error) {
	var (
		user *User
		err  error
	)
	user, err = ac.userRepo.GetUserByAddress(ctx, address) // 查询用户
	if nil != err || nil == user {
		return &pb.UserMarketPropListReply{
			Status: "不存在用户",
		}, nil
	}

	res := make([]*pb.UserMarketPropListReply_List, 0)
	var (
		prop []*Prop
	)
	propStatus := []uint64{4}
	// 11化肥，12水，13手套，14除虫剂，15铲子，16盲盒，17地契
	prop, err = ac.userRepo.GetPropsByExUserID(ctx, user.ID, propStatus, nil)
	if nil != err {
		return &pb.UserMarketPropListReply{
			Status: "错误查询",
		}, nil
	}

	for _, v := range prop {
		res = append(res, &pb.UserMarketPropListReply_List{
			Id:     v.ID,
			Num:    uint64(v.PropType),
			Amount: v.SellAmount,
			UseMax: 0,
		})
	}

	return &pb.UserMarketPropListReply{
		Status: "ok",
		Count:  0,
		List:   res,
	}, nil
}

// UserMarketRentLandList userMarketRentLandList.
func (ac *AppUsecase) UserMarketRentLandList(ctx context.Context, address string, req *pb.UserMarketRentLandListRequest) (*pb.UserMarketRentLandListReply, error) {
	var (
		user *User
		err  error
	)
	user, err = ac.userRepo.GetUserByAddress(ctx, address) // 查询用户
	if nil != err || nil == user {
		return &pb.UserMarketRentLandListReply{
			Status: "不存在用户",
		}, nil
	}

	res := make([]*pb.UserMarketRentLandListReply_List, 0)
	var (
		land []*Land
	)
	landStatus := []uint64{3}
	land, err = ac.userRepo.GetLandByExUserID(ctx, user.ID, landStatus, nil)
	if nil != err {
		return &pb.UserMarketRentLandListReply{
			Status: "错误查询",
		}, nil
	}

	userIds := make([]uint64, 0)
	for _, vLand := range land {
		userIds = append(userIds, vLand.UserId)
	}

	usersMap := make(map[uint64]*User)
	usersMap, err = ac.userRepo.GetUserByUserIds(ctx, userIds)
	if nil != err {
		return &pb.UserMarketRentLandListReply{
			Status: "错误查询",
		}, nil
	}

	for _, vLand := range land {
		addressTmp := ""
		if _, ok := usersMap[vLand.UserId]; ok {
			addressTmp = usersMap[vLand.UserId].Address
		}

		res = append(res, &pb.UserMarketRentLandListReply_List{
			Id:         vLand.ID,
			Level:      vLand.Level,
			MaxHealth:  vLand.MaxHealth,
			RentAmount: vLand.RentOutPutRate,
			Address:    addressTmp,
		})
	}

	return &pb.UserMarketRentLandListReply{
		Status: "ok",
		Count:  0,
		List:   res,
	}, nil
}

// UserMyMarketList userMyMarketList.
func (ac *AppUsecase) UserMyMarketList(ctx context.Context, address string, req *pb.UserMyMarketListRequest) (*pb.UserMyMarketListReply, error) {
	res := make([]*pb.UserMyMarketListReply_List, 0)
	var (
		user *User
		err  error
	)
	user, err = ac.userRepo.GetUserByAddress(ctx, address) // 查询用户
	if nil != err || nil == user {
		return &pb.UserMyMarketListReply{
			Status: "不存在用户",
		}, nil
	}

	var (
		seed []*Seed
	)
	seedStatus := []uint64{4}
	seed, err = ac.userRepo.GetSeedByUserID(ctx, user.ID, seedStatus, nil)
	if nil != err {
		return &pb.UserMyMarketListReply{
			Status: "查询种子错误",
		}, nil
	}

	for _, vSeed := range seed {
		res = append(res, &pb.UserMyMarketListReply_List{
			Id:         vSeed.ID,
			Type:       1,
			Num:        vSeed.SeedId,
			UseNum:     0,
			OutMax:     vSeed.OutMaxAmount,
			Level:      0,
			Status:     0,
			MaxHealth:  0,
			Amount:     vSeed.SellAmount,
			RentAmount: 0,
		})
	}

	var (
		prop []*Prop
	)
	// 11化肥，12水，13手套，14除虫剂，15铲子，16盲盒，17地契
	propStatus := []uint64{4}
	prop, err = ac.userRepo.GetPropsByUserID(ctx, user.ID, propStatus, nil)
	if nil != err {
		return &pb.UserMyMarketListReply{
			Status: "道具错误",
		}, nil
	}

	for _, vProp := range prop {

		useNum := uint64(0)
		if 12 == vProp.PropType {
			useNum = uint64(vProp.ThreeOne) // 水
		} else if 13 == vProp.PropType {
			useNum = uint64(vProp.FiveOne) // 手套
		} else if 14 == vProp.PropType {
			useNum = uint64(vProp.FourOne) // 除虫剂
		} else if 15 == vProp.PropType {
			useNum = uint64(vProp.TwoOne) // 铲子
		}

		res = append(res, &pb.UserMyMarketListReply_List{
			Id:         vProp.ID,
			Type:       2,
			Num:        uint64(vProp.PropType),
			UseNum:     useNum,
			OutMax:     0,
			Level:      0,
			Status:     0,
			MaxHealth:  0,
			Amount:     vProp.SellAmount,
			RentAmount: 0,
		})
	}

	var (
		land []*Land
	)
	landStatus := []uint64{3, 4, 8}
	land, err = ac.userRepo.GetLandByExUserID(ctx, user.ID, landStatus, nil)
	if nil != err {
		return &pb.UserMyMarketListReply{
			Status: "错误查询",
		}, nil
	}

	for _, vLand := range land {
		statusTmp := uint64(1)
		if 4 == vLand.Status {
			statusTmp = 2
		}

		res = append(res, &pb.UserMyMarketListReply_List{
			Id:         vLand.ID,
			Type:       3,
			Num:        0,
			UseNum:     0,
			OutMax:     0,
			Level:      vLand.Level,
			Status:     statusTmp,
			MaxHealth:  vLand.MaxHealth,
			Amount:     vLand.SellAmount,
			RentAmount: vLand.RentOutPutRate,
		})
	}

	return &pb.UserMyMarketListReply{
		Status: "ok",
		Count:  0,
		List:   res,
	}, nil
}

// UserNoticeList NoticeList.
func (ac *AppUsecase) UserNoticeList(ctx context.Context, address string, req *pb.UserNoticeListRequest) (*pb.UserNoticeListReply, error) {
	var (
		user *User
		err  error
	)
	user, err = ac.userRepo.GetUserByAddress(ctx, address) // 查询用户
	if nil != err || nil == user {
		return &pb.UserNoticeListReply{
			Status: "不存在用户",
		}, nil
	}

	res := make([]*pb.UserNoticeListReply_List, 0)

	var (
		notice []*Notice
		count  int64
	)

	count, err = ac.userRepo.GetNoticesCountByUserID(ctx, user.ID)
	if nil != err {
		return &pb.UserNoticeListReply{
			Status: "推荐错误查询",
		}, nil
	}

	notice, err = ac.userRepo.GetNoticesByUserID(ctx, user.ID, &Pagination{
		PageNum:  int(req.Page),
		PageSize: 20,
	})
	if nil != err {
		return &pb.UserNoticeListReply{
			Status: "错误查询",
		}, nil
	}

	for _, vNotice := range notice {
		res = append(res, &pb.UserNoticeListReply_List{
			Content:   vNotice.NoticeContent,
			CreatedAt: vNotice.CreatedAt.Add(8 * time.Hour).Format("2006-01-02 15:04:05"),
		})
	}

	return &pb.UserNoticeListReply{
		Count:  uint64(count),
		List:   res,
		Status: "ok",
	}, nil
}

// UserSkateRewardList userSkateRewardList.
func (ac *AppUsecase) UserSkateRewardList(ctx context.Context, address string, req *pb.UserSkateRewardListRequest) (*pb.UserSkateRewardListReply, error) {
	var (
		user *User
		err  error
	)
	user, err = ac.userRepo.GetUserByAddress(ctx, address) // 查询用户
	if nil != err || nil == user {
		return &pb.UserSkateRewardListReply{
			Status: "不存在用户",
		}, nil
	}

	res := make([]*pb.UserSkateRewardListReply_List, 0)

	var (
		stakeGetPlayRecord []*StakeGetPlayRecord
		count              int64
	)

	count, err = ac.userRepo.GetStakeGetPlayRecordCount(ctx, user.ID, 0)
	if nil != err {
		return &pb.UserSkateRewardListReply{
			Status: "推荐错误查询",
		}, nil
	}

	stakeGetPlayRecord, err = ac.userRepo.GetStakeGetPlayRecordsByUserID(ctx, user.ID, 0, &Pagination{
		PageNum:  int(req.Page),
		PageSize: 20,
	})
	if nil != err {
		return &pb.UserSkateRewardListReply{
			Status: "错误查询",
		}, nil
	}

	userIds := make([]uint64, 0)
	for _, v := range stakeGetPlayRecord {
		userIds = append(userIds, v.UserId)
	}

	usersMap := make(map[uint64]*User)
	usersMap, err = ac.userRepo.GetUserByUserIds(ctx, userIds)
	if nil != err {
		return &pb.UserSkateRewardListReply{
			Status: "错误查询",
		}, nil
	}

	for _, v := range stakeGetPlayRecord {
		addressTmp := ""
		if _, ok := usersMap[v.UserId]; ok {
			addressTmp = usersMap[v.UserId].Address
		}

		res = append(res, &pb.UserSkateRewardListReply_List{
			Address: addressTmp,
			Content: "",
			Amount:  v.Amount,
			Reward:  v.Reward,
			Status:  uint64(v.Status),
		})
	}

	return &pb.UserSkateRewardListReply{
		Count:  uint64(count),
		List:   res,
		Status: "ok",
	}, nil
}

// UserIndexList UserIndexList.
func (ac *AppUsecase) UserIndexList(ctx context.Context, address string, req *pb.UserIndexListRequest) (*pb.UserIndexListReply, error) {
	res := make([]*pb.UserIndexListReply_List, 0)
	var (
		user  *User
		lands []*Land
		err   error
	)
	if 20 < len(req.Address) {
		address = req.Address
	}

	user, err = ac.userRepo.GetUserByAddress(ctx, address) // 查询用户
	if nil != err || nil == user {
		return &pb.UserIndexListReply{
			Status: "不存在用户",
		}, nil
	}

	status := []uint64{1, 2, 3, 8}
	lands, err = ac.userRepo.GetLandByUserIDUsing(ctx, user.ID, status)
	if nil != err {
		return &pb.UserIndexListReply{
			Status: "错误查询",
		}, nil
	}

	landIds := make([]uint64, 0)
	for _, vLand := range lands {
		landIds = append(landIds, vLand.ID)
	}

	var (
		landUserUse map[uint64]*LandUserUse
	)
	landUserUse, err = ac.userRepo.GetLandUserUseByLandIDsMapUsing(ctx, user.ID, landIds)
	if nil != err {
		return &pb.UserIndexListReply{
			Status: "错误查询",
		}, nil
	}

	userIds := make([]uint64, 0)
	for _, vLand := range landUserUse {
		userIds = append(userIds, vLand.UserId)
	}

	usersMap := make(map[uint64]*User)
	usersMap, err = ac.userRepo.GetUserByUserIds(ctx, userIds)
	if nil != err {
		return &pb.UserIndexListReply{
			Status: "错误查询",
		}, nil
	}

	resTmp := make(map[uint64]*pb.UserIndexListReply_List, 0)
	now := time.Now().Unix()
	for _, vLand := range lands {
		plantUserAddressTmp := ""

		if _, ok := landUserUse[vLand.ID]; ok {
			if _, ok2 := usersMap[landUserUse[vLand.ID].UserId]; ok2 {
				plantUserAddressTmp = usersMap[landUserUse[vLand.ID].UserId].Address
			}

			rewardTmp := float64(0)
			statusTmp := uint64(1)
			if 0 != landUserUse[vLand.ID].One {
				statusTmp = 3

			} else if 0 != landUserUse[vLand.ID].Two {
				statusTmp = 2
				// 有虫子但是已经结束
				if landUserUse[vLand.ID].OverTime <= uint64(now) {
					if uint64(now) > landUserUse[vLand.ID].Two {
						tmp := landUserUse[vLand.ID].OutNum * 0.01 * float64(uint64(now)-landUserUse[vLand.ID].Two) / 300

						if tmp >= landUserUse[vLand.ID].OutNum {
							rewardTmp = 0
						} else {
							rewardTmp = landUserUse[vLand.ID].OutNum - tmp
						}
					}
				}
			} else {
				if landUserUse[vLand.ID].OverTime <= uint64(now) {
					rewardTmp = landUserUse[vLand.ID].OutNum
				}
			}

			resTmp[vLand.LocationNum] = &pb.UserIndexListReply_List{
				LocationNum:      vLand.LocationNum,
				LandId:           vLand.ID,
				LandLevel:        vLand.Level,
				Health:           vLand.MaxHealth,
				OutRate:          vLand.OutPutRate,
				PerHealth:        vLand.PerHealth,
				LandUseId:        landUserUse[vLand.ID].ID,
				SeedId:           landUserUse[vLand.ID].SeedTypeId,
				Start:            landUserUse[vLand.ID].BeginTime,
				End:              landUserUse[vLand.ID].OverTime,
				CurrentTime:      uint64(now),
				Status:           statusTmp,
				Reward:           rewardTmp,
				PlantUserAddress: plantUserAddressTmp,
			}
		} else {
			resTmp[vLand.LocationNum] = &pb.UserIndexListReply_List{
				LocationNum:      vLand.LocationNum,
				LandId:           vLand.ID,
				LandLevel:        vLand.Level,
				Health:           vLand.MaxHealth,
				OutRate:          vLand.OutPutRate,
				PerHealth:        vLand.PerHealth,
				LandUseId:        0,
				SeedId:           0,
				Start:            0,
				End:              0,
				CurrentTime:      0,
				Status:           0,
				Reward:           0,
				PlantUserAddress: plantUserAddressTmp,
			}
		}
	}

	for i := uint64(1); i <= 9; i++ {
		if _, ok := resTmp[i]; ok {
			res = append(res, &pb.UserIndexListReply_List{
				LocationNum:      0,
				LandId:           0,
				LandLevel:        0,
				Health:           0,
				OutRate:          0,
				PerHealth:        0,
				LandUseId:        0,
				SeedId:           0,
				Start:            0,
				End:              0,
				CurrentTime:      0,
				Status:           0,
				Reward:           0,
				PlantUserAddress: "",
			})
		} else {
			res = append(res, resTmp[i])
		}
	}

	return &pb.UserIndexListReply{
		Status: "ok",
		Count:  9,
		List:   res,
	}, nil
}

// UserOrderList userOrderList.
func (ac *AppUsecase) UserOrderList(ctx context.Context, address string, req *pb.UserOrderListRequest) (*pb.UserOrderListReply, error) {
	var (
		user *User
		err  error
	)
	user, err = ac.userRepo.GetUserByAddress(ctx, address) // 查询用户
	if nil != err || nil == user {
		return &pb.UserOrderListReply{
			Status: "不存在用户",
		}, nil
	}

	var (
		count int64
		users []*User
	)
	count, err = ac.userRepo.GetUserOrderCount(ctx)
	if nil != err {
		return &pb.UserOrderListReply{
			Status: "查询错误",
		}, nil
	}

	users, err = ac.userRepo.GetUserOrder(ctx, &Pagination{
		PageNum:  int(req.Page),
		PageSize: 20,
	})
	if nil != err {
		return &pb.UserOrderListReply{
			Status: "查询错误",
		}, nil
	}

	res := make([]*pb.UserOrderListReply_List, 0)

	for _, v := range users {
		res = append(res, &pb.UserOrderListReply_List{
			Address: v.Address,
			Git:     v.Git,
		})
	}

	return &pb.UserOrderListReply{
		Count:  uint64(count),
		List:   res,
		Status: "ok",
	}, nil
}

// 随机数生成器的初始化锁
var rngMutexBuyBox sync.Mutex

func (ac *AppUsecase) BuyBox(ctx context.Context, address string, req *pb.BuyBoxRequest) (*pb.BuyBoxReply, error) {
	rngMutexBuyBox.Lock()
	defer rngMutexBuyBox.Unlock()

	var (
		user             *User
		err              error
		boxNum           uint64
		boxSellNum       uint64
		boxSellNumOrigin string
		configs          []*Config
		boxMax           uint64
		boxAmount        float64
		boxStart         string
		boxEnd           string
	)
	user, err = ac.userRepo.GetUserByAddress(ctx, address) // 查询用户
	if nil != err || nil == user {
		return &pb.BuyBoxReply{
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
		"box_sell_num",
		"box_start",
		"box_end",
		"box_amount",
		"skate_over_rate",
		"sell_fee_rate",
		"withdraw_amount_min",
		"withdraw_amount_max",
	)
	if nil != err || nil == configs {
		return &pb.BuyBoxReply{
			Status: "配置错误",
		}, nil
	}

	for _, vConfig := range configs {
		if "box_num" == vConfig.KeyName {
			boxNum, _ = strconv.ParseUint(vConfig.Value, 10, 64)
		}
		if "box_sell_num" == vConfig.KeyName {
			boxSellNum, _ = strconv.ParseUint(vConfig.Value, 10, 64)
		}
		boxSellNumOrigin = vConfig.Value
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
	}
	// 解析时间字符串

	var (
		parsedboxStart time.Time
		parsedboxEnd   time.Time
	)
	parsedboxStart, err = time.Parse("2006-01-02 15:04:05", boxStart)
	if err != nil {
		return &pb.BuyBoxReply{
			Status: "解析时间失败",
		}, nil
	}

	parsedboxEnd, err = time.Parse("2006-01-02 15:04:05", boxEnd)
	if err != nil {
		return &pb.BuyBoxReply{
			Status: "解析时间失败",
		}, nil
	}

	// 获取当前时间
	now := time.Now()

	// 比较时间
	if now.After(parsedboxEnd) {
		return &pb.BuyBoxReply{
			Status: "已结束",
		}, nil
	}

	if now.Before(parsedboxStart) {
		return &pb.BuyBoxReply{
			Status: "未开始",
		}, nil
	}

	if boxSellNum >= boxMax {
		return &pb.BuyBoxReply{
			Status: "已售空",
		}, nil
	}

	if boxAmount >= user.Giw {
		return &pb.BuyBoxReply{
			Status: "余额不足",
		}, nil
	}

	tmpSellNumNew := strconv.FormatUint(boxSellNum+1, 10)
	if err = ac.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
		return ac.userRepo.BuyBox(ctx, boxAmount, boxSellNumOrigin, tmpSellNumNew, &BoxRecord{
			UserId: user.ID,
			Num:    boxNum,
		})
	}); nil != err {
		fmt.Println(err, "buybox", user)
		return &pb.BuyBoxReply{
			Status: "购买失败",
		}, nil
	}

	return &pb.BuyBoxReply{
		Status: "ok",
	}, nil
}

// 随机数生成器
var rngBox *rand2.Rand
var rngPlant *rand2.Rand

// 随机数生成器的初始化锁
var rngMutexBox sync.Mutex

func (ac *AppUsecase) OpenBox(ctx context.Context, address string, req *pb.OpenBoxRequest) (*pb.OpenBoxReply, error) {
	rngMutexBox.Lock()
	defer rngMutexBox.Unlock()

	var (
		user *User
		box  *BoxRecord
		err  error
	)

	user, err = ac.userRepo.GetUserByAddress(ctx, address) // 查询用户
	if nil != err || nil == user {
		return &pb.OpenBoxReply{
			Status: "不存在用户",
		}, nil
	}

	box, err = ac.userRepo.GetUserBoxRecordById(ctx, req.SendBody.Id)
	if nil != err || nil == box {
		return &pb.OpenBoxReply{
			Status: "不存在盲盒",
		}, nil
	}

	if user.ID != box.UserId {
		return &pb.OpenBoxReply{
			Status: "非用户盲盒",
		}, nil
	}

	if 0 != box.GoodId {
		return &pb.OpenBoxReply{
			Status: "已开盲盒",
		}, nil
	}

	// 盲盒道具池
	blindBoxItems := make([]struct {
		Name   uint64
		Weight float64
	}, 0)

	var (
		seedInfos    []*SeedInfo
		seedInfosMap map[uint64]*SeedInfo
	)
	seedInfos, err = ac.userRepo.GetAllSeedInfo(ctx)
	if nil != err {
		return &pb.OpenBoxReply{
			Status: "异常配置",
		}, nil
	}

	seedInfosMap = make(map[uint64]*SeedInfo)
	for _, v := range seedInfos {
		seedInfosMap[v.ID] = v

		blindBoxItems = append(blindBoxItems, struct {
			Name   uint64
			Weight float64
		}{Name: v.ID, Weight: v.GetRate})
	}

	var (
		propInfos    []*PropInfo
		propInfosMap map[uint64]*PropInfo
	)
	propInfos, err = ac.userRepo.GetAllPropInfo(ctx)
	if nil != err {
		return &pb.OpenBoxReply{
			Status: "异常配置",
		}, nil
	}

	propInfosMap = make(map[uint64]*PropInfo)
	for _, v := range propInfos {
		propInfosMap[v.PropType] = v

		blindBoxItems = append(blindBoxItems, struct {
			Name   uint64
			Weight float64
		}{Name: v.PropType, Weight: v.GetRate})
	}

	if 0 >= len(blindBoxItems) {
		return &pb.OpenBoxReply{
			Status: "异常配置概率",
		}, nil
	}

	if nil == rngBox {
		var (
			seedInt     int64
			randomSeeds []*RandomSeed
		)
		randomSeeds, err = ac.userRepo.GetAllRandomSeeds(ctx)
		if nil != err {
			return &pb.OpenBoxReply{
				Status: "异常",
			}, nil
		}

		for _, v := range randomSeeds {
			if 1 == v.Scene {
				seedInt = int64(v.SeedValue)
				break
			}
		}

		if 0 >= seedInt {
			seedInt = time.Now().UnixNano()
			err = ac.userRepo.UpdateSeedValue(ctx, 1, uint64(seedInt))
			if nil != err {
				return &pb.OpenBoxReply{
					Status: "异常",
				}, nil
			}
		}

		rngBox = rand2.New(rand2.NewSource(seedInt))
	}

	r := rngBox.Float64() // 生成 0.0 ~ 1.0 之间的随机数
	// 计算总权重
	var totalWeight float64
	for _, item := range blindBoxItems {
		totalWeight += item.Weight
	}

	// 按照概率随机选择
	result := uint64(0)
	var sum float64
	for _, item := range blindBoxItems {
		sum += item.Weight / totalWeight // 归一化
		if r < sum {
			result = item.Name
			break
		}
	}

	if 0 >= result || 15 < result {
		return &pb.OpenBoxReply{
			Status: "错误盲盒",
		}, nil
	}

	if 1 <= result && result <= 10 {
		if _, ok := seedInfosMap[result]; !ok {
			return &pb.OpenBoxReply{
				Status: "不存在盲盒信息",
			}, nil
		}
		rngTmp := rand2.New(rand2.NewSource(time.Now().UnixNano()))

		outMin := int64(seedInfosMap[result].OutMinAmount)
		outMax := int64(seedInfosMap[result].OutMaxAmount)

		// 计算随机范围
		tmpNum := outMax - outMin
		if tmpNum <= 0 {
			tmpNum = 1 // 避免 Int63n(0) panic
		}

		// 生成随机数
		randomNumber := outMin + rngTmp.Int63n(tmpNum)

		// 种子
		if err = ac.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
			return ac.userRepo.OpenBoxSeed(ctx, box.ID, "", &Seed{
				UserId:       user.ID,
				SeedId:       result,
				Name:         seedInfosMap[result].Name,
				OutOverTime:  seedInfosMap[result].OutOverTime,
				OutMaxAmount: float64(randomNumber),
			})
		}); nil != err {
			fmt.Println(err, "openBox", user)
			return &pb.OpenBoxReply{
				Status: "开启失败",
			}, nil
		}

		return &pb.OpenBoxReply{
			Status:   "ok",
			OpenType: 1,
			Num:      result,
			OutMax:   float64(randomNumber),
		}, nil
	} else if 11 <= result && result <= 15 {
		if _, ok := propInfosMap[result]; !ok {
			return &pb.OpenBoxReply{
				Status: "不存在盲盒信息",
			}, nil
		}

		// 种子
		if err = ac.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
			return ac.userRepo.OpenBoxProp(ctx, box.ID, "", &Prop{
				UserId:   user.ID,
				PropType: int(result),
				OneOne:   int(propInfosMap[result].OneOne),
				OneTwo:   int(propInfosMap[result].OneTwo),
				TwoOne:   int(propInfosMap[result].TwoOne),
				TwoTwo:   propInfosMap[result].TwoTwo,
				ThreeOne: int(propInfosMap[result].ThreeOne),
				FourOne:  int(propInfosMap[result].FourOne),
				FiveOne:  int(propInfosMap[result].FiveOne),
			})
		}); nil != err {
			fmt.Println(err, "openBox", user)
			return &pb.OpenBoxReply{
				Status: "开启失败",
			}, nil
		}

		return &pb.OpenBoxReply{
			Status:   "ok",
			OpenType: 2,
			Num:      result,
			OutMax:   0,
		}, nil
	} else {
		return &pb.OpenBoxReply{
			Status: "开盲盒失败",
		}, nil
	}
}

var rngMutexPlant sync.Mutex

func (ac *AppUsecase) LandPlayOne(ctx context.Context, address string, req *pb.LandPlayOneRequest) (*pb.LandPlayOneReply, error) {
	rngMutexPlant.Lock()
	defer rngMutexPlant.Unlock()

	var (
		user *User
		err  error
	)

	user, err = ac.userRepo.GetUserByAddress(ctx, address) // 查询用户
	if nil != err || nil == user {
		return &pb.LandPlayOneReply{
			Status: "不存在用户",
		}, nil
	}

	var (
		seed *Seed
	)
	seed, err = ac.userRepo.GetSeedByID(ctx, req.SendBody.SeedId, user.ID, 0)
	if nil != err || nil == seed {
		return &pb.LandPlayOneReply{
			Status: "不存种子",
		}, nil
	}

	var (
		land *Land
	)
	land, err = ac.userRepo.GetLandByID(ctx, req.SendBody.LandId)
	if nil != err || nil == land {
		return &pb.LandPlayOneReply{
			Status: "土地信息错误",
		}, nil
	}

	if land.PerHealth > land.MaxHealth {
		return &pb.LandPlayOneReply{
			Status: "肥沃度不足",
		}, nil
	}

	if land.UserId != user.ID {
		if 3 != land.Status {
			return &pb.LandPlayOneReply{
				Status: "未出租土地",
			}, nil
		}
	} else if land.UserId == user.ID {
		if 1 != land.Status {
			return &pb.LandPlayOneReply{
				Status: "未布置土地",
			}, nil
		}
	} else {
		return &pb.LandPlayOneReply{
			Status: "错误参数",
		}, nil
	}

	if 0 == land.LocationNum {
		return &pb.LandPlayOneReply{
			Status: "未布置土地",
		}, nil
	}

	if nil == rngPlant {
		var (
			seedInt     int64
			randomSeeds []*RandomSeed
		)
		randomSeeds, err = ac.userRepo.GetAllRandomSeeds(ctx)
		if nil != err {
			return &pb.LandPlayOneReply{
				Status: "异常",
			}, nil
		}

		for _, v := range randomSeeds {
			if 2 == v.Scene {
				seedInt = int64(v.SeedValue)
				break
			}
		}

		if 0 >= seedInt {
			seedInt = time.Now().UnixNano()
			err = ac.userRepo.UpdateSeedValue(ctx, 2, uint64(seedInt))
			if nil != err {
				return &pb.LandPlayOneReply{
					Status: "异常",
				}, nil
			}
		}

		rngPlant = rand2.New(rand2.NewSource(seedInt))
	}

	now := uint64(time.Now().Unix())
	one := uint64(0)
	two := uint64(0)
	r := rngPlant.Float64() // 生成 0.0 ~ 1.0 之间的随机数
	if r < 0.05 {
		one = now + (seed.OutOverTime / 2)
	} else if r < 0.10 {
		two = now + (seed.OutOverTime / 2)
	}

	originStatusTmp := land.Status
	statusTmp := uint64(1)
	if 3 == originStatusTmp {
		statusTmp = 8
	}

	if err = ac.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
		return ac.userRepo.Plant(ctx, statusTmp, originStatusTmp, land.PerHealth, &LandUserUse{
			LandId:      land.ID,
			Level:       land.Level,
			UserId:      user.ID,
			OwnerUserId: land.UserId,
			SeedId:      seed.ID,
			SeedTypeId:  seed.SeedId,
			Status:      1,
			BeginTime:   now,
			TotalTime:   seed.OutOverTime,
			OverTime:    now + seed.OutOverTime,
			OutMaxNum:   seed.OutMaxAmount * land.OutPutRate,
			One:         one, // 水时间
			Two:         two, // 虫子时间
		})
	}); nil != err {
		fmt.Println(err, "openBox", user)
		return &pb.LandPlayOneReply{
			Status: "种植失败",
		}, nil
	}

	return &pb.LandPlayOneReply{
		Status: "ok",
	}, nil

}

var rngMutexPlantTwo sync.Mutex

func (ac *AppUsecase) LandPlayTwo(ctx context.Context, address string, req *pb.LandPlayTwoRequest) (*pb.LandPlayTwoReply, error) {
	//rngMutexPlantTwo.Lock()
	//defer rngMutexPlantTwo.Unlock()
	//
	//var (
	//	user *User
	//	err  error
	//)
	//
	//user, err = ac.userRepo.GetUserByAddress(ctx, address) // 查询用户
	//if nil != err || nil == user {
	//	return &pb.LandPlayTwoReply{
	//		Status: "不存在用户",
	//	}, nil
	//}
	//
	//var (
	//	landUserUse *LandUserUse
	//)
	//landUserUse, err = ac.userRepo.GetLandUserUseByID(ctx, req.SendBody.LandUseId)
	//if nil != err || nil == landUserUse {
	//	return &pb.LandPlayTwoReply{
	//		Status: "不存在信息",
	//	}, nil
	//}
	//
	//if landUserUse.UserId != user.ID {
	//	return &pb.LandPlayTwoReply{
	//		Status: "非种植用户",
	//	}, nil
	//}
	//
	//if 1 != landUserUse.Status {
	//	return &pb.LandPlayTwoReply{
	//		Status: "状态错误",
	//	}, nil
	//}
	//
	//if 2 == landUserUse.StopStatus {
	//	return &pb.LandPlayTwoReply{
	//		Status: "停止生长状态",
	//	}, nil
	//}
	//
	//current := time.Now().Unix()
	//if uint64(current) < landUserUse.OverTime {
	//	return &pb.LandPlayTwoReply{
	//		Status: "种植未结束",
	//	}, nil
	//}
	//
	//// 已结束
	//
	//var (
	//	seed *Seed
	//)
	//seed, err = ac.userRepo.GetSeedByID(ctx, req.SendBody.SeedId, user.ID, 0)
	//if nil != err || nil == seed {
	//	return &pb.LandPlayOneReply{
	//		Status: "不存种子",
	//	}, nil
	//}
	//
	//var (
	//	land *Land
	//)
	//land, err = ac.userRepo.GetLandByID(ctx, req.SendBody.LandId)
	//if nil != err || nil == land {
	//	return &pb.LandPlayOneReply{
	//		Status: "土地信息错误",
	//	}, nil
	//}
	//
	//if land.PerHealth > land.MaxHealth {
	//	return &pb.LandPlayOneReply{
	//		Status: "肥沃度不足",
	//	}, nil
	//}
	//
	//if land.UserId != user.ID {
	//	if 3 != land.Status {
	//		return &pb.LandPlayOneReply{
	//			Status: "未出租土地",
	//		}, nil
	//	}
	//} else if land.UserId == user.ID {
	//	if 1 != land.Status {
	//		return &pb.LandPlayOneReply{
	//			Status: "未布置土地",
	//		}, nil
	//	}
	//} else {
	//	return &pb.LandPlayOneReply{
	//		Status: "错误参数",
	//	}, nil
	//}
	//
	//if nil == rngPlant {
	//	var (
	//		seedInt     int64
	//		randomSeeds []*RandomSeed
	//	)
	//	randomSeeds, err = ac.userRepo.GetAllRandomSeeds(ctx)
	//	if nil != err {
	//		return &pb.LandPlayOneReply{
	//			Status: "异常",
	//		}, nil
	//	}
	//
	//	for _, v := range randomSeeds {
	//		if 2 == v.Scene {
	//			seedInt = int64(v.SeedValue)
	//			break
	//		}
	//	}
	//
	//	if 0 >= seedInt {
	//		seedInt = time.Now().UnixNano()
	//		err = ac.userRepo.UpdateSeedValue(ctx, 2, uint64(seedInt))
	//		if nil != err {
	//			return &pb.LandPlayOneReply{
	//				Status: "异常",
	//			}, nil
	//		}
	//	}
	//
	//	rngPlant = rand2.New(rand2.NewSource(seedInt))
	//}
	//
	//one := uint64(0)
	//two := uint64(0)
	//r := rngPlant.Float64() // 生成 0.0 ~ 1.0 之间的随机数
	//if r < 0.05 {
	//	one = 1
	//} else if r < 0.10 {
	//	two = 1
	//}
	//
	//originStatusTmp := land.Status
	//statusTmp := uint64(1)
	//if 3 == originStatusTmp {
	//	statusTmp = 8
	//}
	//
	//now := uint64(time.Now().Unix())
	//if err = ac.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
	//	return ac.userRepo.Plant(ctx, statusTmp, originStatusTmp, land.PerHealth, &LandUserUse{
	//		LandId:      land.ID,
	//		Level:       land.Level,
	//		UserId:      user.ID,
	//		OwnerUserId: land.UserId,
	//		SeedId:      seed.ID,
	//		SeedTypeId:  seed.SeedId,
	//		Status:      1,
	//		BeginTime:   now,
	//		TotalTime:   seed.OutOverTime,
	//		OverTime:    now + seed.OutOverTime,
	//		OutMaxNum:   seed.OutMaxAmount * land.OutPutRate,
	//		One:         one,
	//		Two:         two,
	//	})
	//}); nil != err {
	//	fmt.Println(err, "openBox", user)
	//	return &pb.LandPlayOneReply{
	//		Status: "种植失败",
	//	}, nil
	//}
	//
	//return &pb.LandPlayOneReply{
	//	Status: "ok",
	//}, nil

	return &pb.LandPlayTwoReply{
		Status: "ok",
	}, nil
}

var rngMutexBuy sync.Mutex

func (ac *AppUsecase) Buy(ctx context.Context, address string, req *pb.BuyRequest) (*pb.BuyReply, error) {
	rngMutexBuy.Lock()
	defer rngMutexBuy.Unlock()

	var (
		user    *User
		feeRate float64
		configs []*Config
		err     error
	)

	user, err = ac.userRepo.GetUserByAddress(ctx, address) // 查询用户
	if nil != err || nil == user {
		return &pb.BuyReply{
			Status: "不存在用户",
		}, nil
	}

	// 配置
	configs, err = ac.userRepo.GetConfigByKeys(ctx,
		"sell_fee_rate",
	)
	if nil != err || nil == configs {
		return &pb.BuyReply{
			Status: "配置错误",
		}, nil
	}

	for _, vConfig := range configs {
		if "sell_fee_rate" == vConfig.KeyName {
			feeRate, _ = strconv.ParseFloat(vConfig.Value, 10)
		}
	}

	if 1 == req.SendBody.BuyType {
		var (
			seed *Seed
		)
		seed, err = ac.userRepo.GetSeedBuyByID(ctx, req.SendBody.Id, 4)
		if nil != err || nil == seed {
			return &pb.BuyReply{
				Status: "不存种子",
			}, nil
		}

		if user.ID == seed.UserId {
			return &pb.BuyReply{
				Status: "不允许购买自己的",
			}, nil
		}
		if 0 >= seed.SellAmount {
			return &pb.BuyReply{
				Status: "金额错误",
			}, nil
		}

		if user.Git < seed.SellAmount {
			return &pb.BuyReply{
				Status: "余额不足",
			}, nil
		}
		// 种子
		tmpGet := seed.SellAmount - seed.SellAmount*feeRate
		if 0 >= tmpGet {
			return &pb.BuyReply{
				Status: "金额错误",
			}, nil
		}

		if err = ac.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
			return ac.userRepo.BuySeed(ctx, seed.SellAmount, tmpGet, seed.UserId, user.ID, seed.ID)
		}); nil != err {
			fmt.Println(err, "buySeed", user)
			return &pb.BuyReply{
				Status: "购买失败",
			}, nil
		}
	} else if 2 == req.SendBody.BuyType {
		var (
			prop *Prop
		)
		prop, err = ac.userRepo.GetPropByID(ctx, req.SendBody.Id, 4)
		if nil != err || nil == prop {
			return &pb.BuyReply{
				Status: "不存道具",
			}, nil
		}

		if user.ID == prop.UserId {
			return &pb.BuyReply{
				Status: "不允许购买自己的",
			}, nil
		}

		if 0 >= prop.SellAmount {
			return &pb.BuyReply{
				Status: "金额错误",
			}, nil
		}

		if user.Git < prop.SellAmount {
			return &pb.BuyReply{
				Status: "余额不足",
			}, nil
		}

		// 种子
		tmpGet := prop.SellAmount - prop.SellAmount*feeRate
		if 0 >= tmpGet {
			return &pb.BuyReply{
				Status: "金额错误",
			}, nil
		}

		if err = ac.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
			return ac.userRepo.BuyProp(ctx, prop.SellAmount, tmpGet, prop.UserId, user.ID, prop.ID)
		}); nil != err {
			fmt.Println(err, "buyProp", user)
			return &pb.BuyReply{
				Status: "购买失败",
			}, nil
		}
	} else if 3 == req.SendBody.BuyType {
		var (
			land *Land
		)
		land, err = ac.userRepo.GetLandByID(ctx, req.SendBody.Id)
		if nil != err || nil == land {
			return &pb.BuyReply{
				Status: "不存道具",
			}, nil
		}

		if user.ID == land.UserId {
			return &pb.BuyReply{
				Status: "不允许购买自己的",
			}, nil
		}

		if 4 != land.Status {
			return &pb.BuyReply{
				Status: "未出售",
			}, nil
		}

		if 0 >= land.SellAmount {
			return &pb.BuyReply{
				Status: "金额错误",
			}, nil
		}

		if user.Git < land.SellAmount {
			return &pb.BuyReply{
				Status: "余额不足",
			}, nil
		}

		// 土地
		tmpGet := land.SellAmount - land.SellAmount*feeRate
		if 0 >= tmpGet {
			return &pb.BuyReply{
				Status: "金额错误",
			}, nil
		}

		if err = ac.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
			return ac.userRepo.BuyLand(ctx, land.SellAmount, tmpGet, land.UserId, user.ID, land.ID)
		}); nil != err {
			fmt.Println(err, "buyLand", user)
			return &pb.BuyReply{
				Status: "购买失败",
			}, nil
		}
	} else {
		return &pb.BuyReply{
			Status: "参数错误",
		}, nil
	}

	return &pb.BuyReply{
		Status: "ok",
	}, nil
}

func (ac *AppUsecase) Sell(ctx context.Context, address string, req *pb.SellRequest) (*pb.SellReply, error) {
	var (
		user *User
		err  error
	)

	user, err = ac.userRepo.GetUserByAddress(ctx, address) // 查询用户
	if nil != err || nil == user {
		return &pb.SellReply{
			Status: "不存在用户",
		}, nil
	}

	tmpSellAmount := float64(req.SendBody.Amount)
	if 1 == req.SendBody.Num {
		if 0 >= tmpSellAmount {
			return &pb.SellReply{
				Status: "售价不能为0",
			}, nil
		}

		if 1 == req.SendBody.SellType {
			var (
				seed *Seed
			)
			seed, err = ac.userRepo.GetSeedBuyByID(ctx, req.SendBody.Id, 0)
			if nil != err || nil == seed {
				return &pb.SellReply{
					Status: "不存种子",
				}, nil
			}

			if user.ID != seed.UserId {
				return &pb.SellReply{
					Status: "不是自己的种子",
				}, nil
			}

			if err = ac.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
				return ac.userRepo.SellSeed(ctx, seed.ID, user.ID, tmpSellAmount)
			}); nil != err {
				fmt.Println(err, "sellSeed", user)
				return &pb.SellReply{
					Status: "上架失败",
				}, nil
			}
		} else if 2 == req.SendBody.SellType {
			var (
				prop *Prop
			)
			prop, err = ac.userRepo.GetPropByIDSell(ctx, req.SendBody.Id, 2)
			if nil != err || nil == prop {
				return &pb.SellReply{
					Status: "不存道具",
				}, nil
			}

			if user.ID != prop.UserId {
				return &pb.SellReply{
					Status: "不是自己的",
				}, nil
			}

			if err = ac.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
				return ac.userRepo.SellProp(ctx, prop.ID, user.ID, tmpSellAmount)
			}); nil != err {
				fmt.Println(err, "sellProp", user)
				return &pb.SellReply{
					Status: "上架失败",
				}, nil
			}
		} else if 3 == req.SendBody.SellType {
			var (
				land *Land
			)
			land, err = ac.userRepo.GetLandByID(ctx, req.SendBody.Id)
			if nil != err || nil == land {
				return &pb.SellReply{
					Status: "不存道具",
				}, nil
			}

			if user.ID != land.UserId {
				return &pb.SellReply{
					Status: "不是自己的",
				}, nil
			}

			if 0 != land.LocationNum {
				return &pb.SellReply{
					Status: "土地布置中",
				}, nil
			}

			if 0 != land.Status {
				return &pb.SellReply{
					Status: "不符合上架要求",
				}, nil
			}

			if err = ac.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
				return ac.userRepo.SellLand(ctx, land.ID, user.ID, tmpSellAmount)
			}); nil != err {
				fmt.Println(err, "sellProp", user)
				return &pb.SellReply{
					Status: "上架失败",
				}, nil
			}
		} else {
			return &pb.SellReply{
				Status: "参数错误",
			}, nil
		}
	} else {
		if 1 == req.SendBody.SellType {
			var (
				seed *Seed
			)
			seed, err = ac.userRepo.GetSeedBuyByID(ctx, req.SendBody.Id, 4)
			if nil != err || nil == seed {
				return &pb.SellReply{
					Status: "不存种子",
				}, nil
			}

			if user.ID != seed.UserId {
				return &pb.SellReply{
					Status: "不是自己的种子",
				}, nil
			}

			if err = ac.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
				return ac.userRepo.UnSellSeed(ctx, seed.ID, user.ID)
			}); nil != err {
				fmt.Println(err, "sellSeed", user)
				return &pb.SellReply{
					Status: "上架失败",
				}, nil
			}
		} else if 2 == req.SendBody.SellType {
			var (
				prop *Prop
			)
			prop, err = ac.userRepo.GetPropByID(ctx, req.SendBody.Id, 4)
			if nil != err || nil == prop {
				return &pb.SellReply{
					Status: "不存道具",
				}, nil
			}

			if user.ID != prop.UserId {
				return &pb.SellReply{
					Status: "不是自己的",
				}, nil
			}

			if err = ac.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
				return ac.userRepo.UnSellProp(ctx, prop.ID, user.ID)
			}); nil != err {
				fmt.Println(err, "sellProp", user)
				return &pb.SellReply{
					Status: "上架失败",
				}, nil
			}
		} else if 3 == req.SendBody.SellType {
			var (
				land *Land
			)
			land, err = ac.userRepo.GetLandByID(ctx, req.SendBody.Id)
			if nil != err || nil == land {
				return &pb.SellReply{
					Status: "不存道具",
				}, nil
			}

			if user.ID != land.UserId {
				return &pb.SellReply{
					Status: "不是自己的",
				}, nil
			}

			if 0 != land.LocationNum {
				return &pb.SellReply{
					Status: "土地布置中",
				}, nil
			}

			if 4 != land.Status {
				return &pb.SellReply{
					Status: "不符合下架要求",
				}, nil
			}

			if err = ac.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
				return ac.userRepo.UnSellLand(ctx, land.ID, user.ID)
			}); nil != err {
				fmt.Println(err, "sellProp", user)
				return &pb.SellReply{
					Status: "上架失败",
				}, nil
			}
		} else {
			return &pb.SellReply{
				Status: "参数错误",
			}, nil
		}
	}

	return &pb.SellReply{
		Status: "ok",
	}, nil
}

func (ac *AppUsecase) StakeGit(ctx context.Context, address string, req *pb.StakeGitRequest) (*pb.StakeGitReply, error) {
	var (
		user *User
		err  error
	)

	user, err = ac.userRepo.GetUserByAddress(ctx, address) // 查询用户
	if nil != err || nil == user {
		return &pb.StakeGitReply{
			Status: "不存在用户",
		}, nil
	}

	if 1 == req.SendBody.Num {

	} else if 2 == req.SendBody.Num {

	} else {
		return &pb.StakeGitReply{
			Status: "错误参数",
		}, nil
	}

	return &pb.StakeGitReply{Status: "ok"}, nil
}

func (ac *AppUsecase) RentLand(ctx context.Context, address string, req *pb.RentLandRequest) (*pb.RentLandReply, error) {
	var (
		user *User
		err  error
	)

	user, err = ac.userRepo.GetUserByAddress(ctx, address) // 查询用户
	if nil != err || nil == user {
		return &pb.RentLandReply{
			Status: "不存在用户",
		}, nil
	}

	rentRate := 0.05
	if 1 == req.SendBody.Rate {
		rentRate = 0.05
	} else if 2 == req.SendBody.Rate {
		rentRate = 0.1
	} else if 3 == req.SendBody.Rate {
		rentRate = 0.2
	} else if 4 == req.SendBody.Rate {
		rentRate = 0.3
	} else if 5 == req.SendBody.Rate {
		rentRate = 0.4
	} else if 6 == req.SendBody.Rate {
		rentRate = 0.5
	} else {
		return &pb.RentLandReply{
			Status: "比例错误",
		}, nil
	}

	if 1 == req.SendBody.Num {
		var (
			land *Land
		)
		land, err = ac.userRepo.GetLandByID(ctx, req.SendBody.LandId)
		if nil != err || nil == land {
			return &pb.RentLandReply{
				Status: "不存在土地",
			}, nil
		}

		if user.ID != land.UserId {
			return &pb.RentLandReply{
				Status: "不是自己的",
			}, nil
		}

		if 1 != land.Status {
			return &pb.RentLandReply{
				Status: "请将土地布置在农场",
			}, nil
		}

		if land.PerHealth > land.MaxHealth {
			return &pb.RentLandReply{
				Status: "肥沃度不足",
			}, nil
		}

		if err = ac.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
			return ac.userRepo.RentLand(ctx, land.ID, user.ID, rentRate)
		}); nil != err {
			fmt.Println(err, "rendLand", user)
			return &pb.RentLandReply{
				Status: "上架失败",
			}, nil
		}
	} else if 2 == req.SendBody.Num {
		var (
			land *Land
		)
		land, err = ac.userRepo.GetLandByID(ctx, req.SendBody.LandId)
		if nil != err || nil == land {
			return &pb.RentLandReply{
				Status: "不存在土地",
			}, nil
		}

		if user.ID != land.UserId {
			return &pb.RentLandReply{
				Status: "不是自己的",
			}, nil
		}

		if 3 != land.Status {
			return &pb.RentLandReply{
				Status: "土地使用中",
			}, nil
		}

		if err = ac.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
			return ac.userRepo.UnRentLand(ctx, land.ID, user.ID)
		}); nil != err {
			fmt.Println(err, "unRendLand", user)
			return &pb.RentLandReply{
				Status: "上架失败",
			}, nil
		}
	} else {
		return &pb.RentLandReply{
			Status: "错误参数",
		}, nil
	}

	return &pb.RentLandReply{Status: "ok"}, nil
}

func (ac *AppUsecase) LandPlay(ctx context.Context, address string, req *pb.LandPlayRequest) (*pb.LandPlayReply, error) {
	var (
		user *User
		//box  *BoxRecord
		err error
	)

	user, err = ac.userRepo.GetUserByAddress(ctx, address) // 查询用户
	if nil != err || nil == user {
		return &pb.LandPlayReply{
			Status: "不存在用户",
		}, nil
	}

	if 1 == req.SendBody.Num {
		var (
			land *Land
		)
		land, err = ac.userRepo.GetLandByIDTwo(ctx, req.SendBody.LandId)
		if nil != err || nil == land {
			return &pb.LandPlayReply{
				Status: "不存在土地",
			}, nil
		}

		if user.ID != land.UserId {
			return &pb.LandPlayReply{
				Status: "不是自己的",
			}, nil
		}

		if 1 != land.Status {
			return &pb.LandPlayReply{
				Status: "请将土地布置在农场",
			}, nil
		}

		if err = ac.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
			return ac.userRepo.LandPull(ctx, land.ID, user.ID)
		}); nil != err {
			fmt.Println(err, "LandPull", user)
			return &pb.LandPlayReply{
				Status: "上架失败",
			}, nil
		}
	} else if 2 == req.SendBody.Num {
		var (
			land  *Land
			land2 *Land
		)
		land, err = ac.userRepo.GetLandByIDTwo(ctx, req.SendBody.LandId)
		if nil != err || nil == land {
			return &pb.LandPlayReply{
				Status: "不存在土地",
			}, nil
		}

		if 1 <= land.LocationNum && 9 >= land.LocationNum {

		} else {
			return &pb.LandPlayReply{
				Status: "非布置土地",
			}, nil
		}

		land2, err = ac.userRepo.GetLandByID(ctx, req.SendBody.LandTwoId)
		if nil != err || nil == land2 {
			return &pb.LandPlayReply{
				Status: "不存在土地",
			}, nil
		}

		if 0 != land2.LocationNum {
			return &pb.LandPlayReply{
				Status: "已布置土地",
			}, nil
		}

		if user.ID != land.UserId {
			return &pb.LandPlayReply{
				Status: "不是自己的",
			}, nil
		}

		if user.ID != land2.UserId {
			return &pb.LandPlayReply{
				Status: "不是自己的",
			}, nil
		}

		if 1 != land.Status {
			return &pb.LandPlayReply{
				Status: "请将土地布置在农场",
			}, nil
		}

		if 0 != land2.Status {
			return &pb.LandPlayReply{
				Status: "土地使用中",
			}, nil
		}

		if err = ac.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
			err = ac.userRepo.LandPull(ctx, land.ID, user.ID)
			if nil != err {
				return err
			}

			return ac.userRepo.LandPush(ctx, land2.ID, user.ID, land.LocationNum)
		}); nil != err {
			fmt.Println(err, "LandPullPush", user)
			return &pb.LandPlayReply{
				Status: "替换失败",
			}, nil
		}
	} else if 3 == req.SendBody.Num {
		var (
			tmpLand *Land
			land    *Land
		)
		tmpLand, err = ac.userRepo.GetLandByUserIdLocationNum(ctx, user.ID, req.SendBody.LocationNum)
		if nil != err {
			return &pb.LandPlayReply{
				Status: "错误查询",
			}, nil
		}

		if nil != tmpLand {
			return &pb.LandPlayReply{
				Status: "存在布置土地",
			}, nil
		}

		land, err = ac.userRepo.GetLandByID(ctx, req.SendBody.LandId)
		if nil != err || nil == land {
			return &pb.LandPlayReply{
				Status: "不存在土地",
			}, nil
		}

		if user.ID != land.UserId {
			return &pb.LandPlayReply{
				Status: "不是自己的",
			}, nil
		}

		if 0 != land.LocationNum {
			return &pb.LandPlayReply{
				Status: "不是空闲的土地",
			}, nil
		}

		if 0 != land.Status {
			return &pb.LandPlayReply{
				Status: "不是空闲的土地",
			}, nil
		}

		if land.PerHealth > land.MaxHealth {
			return &pb.LandPlayReply{
				Status: "肥沃度不足",
			}, nil
		}

		if err = ac.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
			return ac.userRepo.LandPush(ctx, land.ID, user.ID, req.SendBody.LocationNum)
		}); nil != err {
			fmt.Println(err, "LandPush", user)
			return &pb.LandPlayReply{
				Status: "失败",
			}, nil
		}
	} else {
		return &pb.LandPlayReply{
			Status: "错误参数",
		}, nil
	}

	return &pb.LandPlayReply{Status: "ok"}, nil
}
