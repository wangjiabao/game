package biz

import (
	"context"
	"crypto/rand"
	"fmt"
	pb "game/api/app/v1"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"math"
	"math/big"
	rand2 "math/rand"
	"strconv"
	"strings"
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
	GiwTwo           float64
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
	Location         float64
	Recommend        float64
	RecommendTwo     float64
	Area             float64
	AreaTwo          float64
	All              float64
	Amount           float64
	AmountGet        float64
	AmountUsdt       float64
	MyTotalAmount    float64
	UsdtTwo          float64
	OutNum           uint64
	Vip              uint64
	VipAdmin         uint64
	LockUse          uint64
	LockReward       uint64
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

type StakeGit struct {
	ID        uint64
	UserId    uint64
	Status    uint64
	Amount    float64
	Reward    float64
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

type RewardTwo struct {
	ID        uint64
	UserId    uint64
	Reason    uint64
	One       uint64
	Two       uint64
	Three     float64
	Amount    float64
	Four      string
	Five      float64
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

type LandInfo struct {
	ID                uint64
	Level             uint64
	OutPutRateMax     float64
	OutPutRateMin     float64
	RentOutPutRateMax float64
	MaxHealth         uint64
	PerHealth         uint64
	LimitDateMax      uint64
	CreatedAt         time.Time
	UpdatedAt         time.Time
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
	ID               uint64
	UserId           uint64
	NoticeContent    string
	NoticeContentTwo string
	CreatedAt        time.Time
	UpdatedAt        time.Time
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
	Balance   float64
	CreatedAt time.Time
	UpdatedAt time.Time
}

type StakeGitRecord struct {
	ID        uint64
	UserId    uint64
	Amount    float64
	StakeType int
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Withdraw struct {
	ID             uint64
	UserID         uint64
	Amount         uint64
	RelAmount      uint64
	AmountFloat    float64
	RelAmountFloat float64
	Status         string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type BuyLand struct {
	ID        uint64
	Amount    float64
	Status    uint64
	CreatedAt time.Time
	UpdatedAt time.Time
	AmountTwo float64
	Limit     uint64
	Level     uint64
}

type RandomSeed struct {
	ID        uint64
	Scene     uint64
	SeedValue uint64
	UpdatedAt time.Time
	CreatedAt time.Time
}

type BuyLandRecord struct {
	ID        uint64
	BuyLandID uint64
	Amount    float64
	CreatedAt time.Time
	UpdatedAt time.Time
	Status    uint64
	UserID    uint64
}

type UserRepo interface {
	GetAllUsers(ctx context.Context) ([]*User, error)
	GetUserByUserIds(ctx context.Context, userIds []uint64) (map[uint64]*User, error)
	GetUserByAddress(ctx context.Context, address string) (*User, error)
	GetUserById(ctx context.Context, id uint64) (*User, error)
	GetUserRecommendByUserId(ctx context.Context, userId uint64) (*UserRecommend, error)
	GetUserRecommendByCode(ctx context.Context, code string) ([]*UserRecommend, error)
	GetUserRecommendLikeCode(ctx context.Context, code string) ([]*UserRecommend, error)
	GetUserRecommends(ctx context.Context) ([]*UserRecommend, error)
	CreateUser(ctx context.Context, uc *User) (*User, error)
	CreateStakeGet(ctx context.Context, sg *StakeGet) error
	CreateStakeGit(ctx context.Context, sg *StakeGit) error
	CreateNotice(ctx context.Context, userId uint64, content string, contentTwo string) error
	CreateUserRecommend(ctx context.Context, user *User, recommendUser *UserRecommend) (*UserRecommend, error)
	GetConfigByKeys(ctx context.Context, keys ...string) ([]*Config, error)
	GetStakeGitByUserId(ctx context.Context, userId uint64) (*StakeGit, error)
	GetBoxRecord(ctx context.Context, num uint64) ([]*BoxRecord, error)
	GetBoxRecordCount(ctx context.Context, num uint64) (int64, error)
	GetUserBoxRecord(ctx context.Context, userId, num uint64, b *Pagination) ([]*BoxRecord, error)
	GetUserBoxRecordOpen(ctx context.Context, userId, num uint64, open bool, b *Pagination) ([]*BoxRecord, error)
	GetStakeGetTotal(ctx context.Context) (*StakeGetTotal, error)
	GetUserStakeGet(ctx context.Context, userId uint64) (*StakeGet, error)
	GetTotalStakeRate(ctx context.Context) (float64, error)
	GetUserRecommendCount(ctx context.Context, code string) (int64, error)
	GetUserRecommendByCodePage(ctx context.Context, code string, b *Pagination) ([]*UserRecommend, error)
	GetLandByUserIDUsing(ctx context.Context, userID uint64, status []uint64) ([]*Land, error)
	GetLandByUserID(ctx context.Context, userID uint64, status []uint64, b *Pagination) ([]*Land, error)
	GetLandByExUserID(ctx context.Context, userID uint64, status []uint64, b *Pagination) ([]*Land, error)
	GetLandByExUserIDByIds(ctx context.Context, ids []uint64, b *Pagination) ([]*Land, error)
	GetUserRewardPage(ctx context.Context, userId uint64, reason []uint64, b *Pagination) ([]*Reward, error)
	GetUserRewardTwoPage(ctx context.Context, userId uint64, reason uint64, b *Pagination) ([]*RewardTwo, error)
	GetUserRewardPageCount(ctx context.Context, userId uint64, reason []uint64) (int64, error)
	GetUserRewardTwoPageCount(ctx context.Context, userId uint64, reason uint64) (int64, error)
	GetSeedByUserID(ctx context.Context, userID uint64, status []uint64, b *Pagination) ([]*Seed, error)
	GetSeedByExUserID(ctx context.Context, userID uint64, status []uint64, b *Pagination) ([]*Seed, error)
	GetLandUserUseByUserIDUseing(ctx context.Context, userID uint64, status uint64, b *Pagination) ([]*LandUserUse, error)
	GetExchangeRecordsByUserID(ctx context.Context, userID uint64, b *Pagination) ([]*ExchangeRecord, error)
	GetLandUserUseByID(ctx context.Context, id uint64) (*LandUserUse, error)
	GetMarketRecordsByUserID(ctx context.Context, userID uint64, status uint64, b *Pagination) ([]*Market, error)
	GetNoticesByUserID(ctx context.Context, userID uint64, b *Pagination) ([]*Notice, error)
	GetNoticesCountByUserID(ctx context.Context, userID uint64) (int64, error)
	GetPropsByUserID(ctx context.Context, userID uint64, status []uint64, b *Pagination) ([]*Prop, error)
	GetPropsByUserIDPropType(ctx context.Context, userID uint64, propType []uint64) ([]*Prop, error)
	GetPropsByExUserID(ctx context.Context, userID uint64, status []uint64, b *Pagination) ([]*Prop, error)
	GetStakeGetsByUserID(ctx context.Context, userID uint64, b *Pagination) ([]*StakeGet, error)
	GetStakeGetPlayRecordsByUserID(ctx context.Context, userID uint64, status uint64, b *Pagination) ([]*StakeGetPlayRecord, error)
	GetStakeGetPlayRecordCount(ctx context.Context, userID uint64, status uint64) (int64, error)
	GetStakeGetRecordsByUserID(ctx context.Context, userID int64, b *Pagination) ([]*StakeGetRecord, error)
	GetStakeGitByUserID(ctx context.Context, userID int64) (*StakeGit, error)
	GetStakeGitRecordsByUserID(ctx context.Context, userID uint64, b *Pagination) ([]*StakeGitRecord, error)
	GetStakeGitRecordsByID(ctx context.Context, id, userId uint64) (*StakeGitRecord, error)
	GetWithdrawRecordsByUserID(ctx context.Context, userID int64, b *Pagination) ([]*Withdraw, error)
	GetUserOrderCount(ctx context.Context) (int64, error)
	GetUserOrder(ctx context.Context, b *Pagination) ([]*User, error)
	GetLandUserUseByLandIDsMapUsing(ctx context.Context, userId uint64, landIDs []uint64) (map[uint64]*LandUserUse, error)
	GetLandUserUseByLandIDsUsing(ctx context.Context, userId uint64) ([]*LandUserUse, error)
	BuyBox(ctx context.Context, giw float64, originValue, value string, uc *BoxRecord) (uint64, error)
	BuyLandReward(ctx context.Context, userId, landId uint64, giw float64) error
	GetUserBoxRecordById(ctx context.Context, id uint64) (*BoxRecord, error)
	OpenBoxSeed(ctx context.Context, id uint64, content string, seedInfo *Seed) (uint64, error)
	OpenBoxProp(ctx context.Context, id uint64, content string, propInfo *Prop) (uint64, error)
	GetAllSeedInfo(ctx context.Context) ([]*SeedInfo, error)
	GetAllPropInfo(ctx context.Context) ([]*PropInfo, error)
	GetAllRandomSeeds(ctx context.Context) ([]*RandomSeed, error)
	UpdateSeedValue(ctx context.Context, scene uint64, newSeed uint64) error
	GetSeedByID(ctx context.Context, seedID, userId, status uint64) (*Seed, error)
	GetLandByID(ctx context.Context, landID uint64) (*Land, error)
	GetLandByIDTwo(ctx context.Context, landID uint64) (*Land, error)
	GetLandByUserIdLocationNum(ctx context.Context, userId uint64, locationNum uint64) (*Land, error)
	Plant(ctx context.Context, status, originStatus, perHealth uint64, landUserUse *LandUserUse) error
	PlantPlatTwo(ctx context.Context, id, landId uint64, rent bool) error
	PlantPlatThree(ctx context.Context, id, overTime, propId uint64, one, two bool) error
	PlantPlatFour(ctx context.Context, outMax float64, id, propId, propStatus, propNum uint64) error
	PlantPlatFive(ctx context.Context, overTime, id, propId, propStatus, propNum uint64) error
	PlantPlatSix(ctx context.Context, id, propId, propStatus, propNum, landId uint64) error
	PlantPlatSeven(ctx context.Context, outMax, amount float64, subTime, lastTime, id, propId, propStatus, propNum, userId uint64) error
	PlantPlatTwoTwo(ctx context.Context, id, userId, rentUserId uint64, amount, rentAmount float64) error
	PlantPlatTwoTwoL(ctx context.Context, id, userId, lowUserId, num uint64, amount float64) error
	GetSeedBuyByID(ctx context.Context, seedID, status uint64) (*Seed, error)
	GetPropByID(ctx context.Context, propID, status uint64) (*Prop, error)
	GetPropByIDTwo(ctx context.Context, propID uint64) (*Prop, error)
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
	LandPullTwo(ctx context.Context, landId uint64, userId uint64) error
	LandPush(ctx context.Context, landId uint64, userId, locationNum uint64) error
	LandAddOutRate(ctx context.Context, id, landId, userId uint64) error
	PropStatusThree(ctx context.Context, id uint64) error
	CreateLand(ctx context.Context, lc *Land) (*Land, error)
	GetLand(ctx context.Context, id, id2, userId uint64) error
	GetLandInfoByLevels(ctx context.Context) (map[uint64]*LandInfo, error)
	SetGiw(ctx context.Context, address string, giw uint64) error
	SetGit(ctx context.Context, address string, git uint64) error
	SetStakeGetTotal(ctx context.Context, amount, balance float64) error
	SetStakeGetTotalSub(ctx context.Context, amount, balance float64) error
	SetStakeGet(ctx context.Context, userId uint64, git, amount float64) error
	SetStakeGetSub(ctx context.Context, userId uint64, git, amount float64) error
	SetStakeGetPlaySub(ctx context.Context, userId uint64, amount float64) error
	SetStakeGetPlay(ctx context.Context, userId uint64, git, amount float64) error
	SetStakeGit(ctx context.Context, userId uint64, amount float64) error
	SetUnStakeGit(ctx context.Context, id, userId uint64, amount float64) error
	Exchange(ctx context.Context, userId uint64, git, giw float64) error
	ExchangeTwo(ctx context.Context, userId uint64, git, giw float64) error
	ExchangeThree(ctx context.Context, userId uint64, git, giw float64) error
	Withdraw(ctx context.Context, userId uint64, giw, relGiw float64) error
	WithdrawTwo(ctx context.Context, userId uint64, usdt, relUsdt float64) error
	GetAllBuyLandRecords(ctx context.Context, id uint64) ([]*BuyLandRecord, error)
	GetBuyLandById(ctx context.Context) (*BuyLand, error)
	CreateBuyLandRecord(ctx context.Context, limit uint64, bl *BuyLandRecord) error
	CreateBuyLandRecordOne(ctx context.Context, bl *BuyLandRecord) error
	GetSetBuyLandById(ctx context.Context) (*BuyLand, error)
	BackUserGit(ctx context.Context, userId, id uint64, git float64) error
	SetBuyLandOver(ctx context.Context, id uint64) error
	UpdateUserNewTwoNew(ctx context.Context, userId uint64, amount, giw float64, amountUsdt uint64) error
	UpdateUserMyTotalAmount(ctx context.Context, userId uint64, amount float64) error
	UpdateUserMyTotalAmountSub(ctx context.Context, userId int64, amount float64) error
	UpdateUserRewardRecommend2(ctx context.Context, userId uint64, giw, giw2, usdt float64, amountOrigin float64, stop bool, address string) error
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
		if 0 >= len(code) {
			return nil, errors.New(500, "USER_ERROR", "errcode"), "errcode"
		}

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

			//if 0 >= rUser.GiwAdd {
			//	return nil, errors.New(500, "USER_ERROR", "推荐人未入金"), "推荐人未入金"
			//}

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

			err = ac.userRepo.CreateStakeGet(ctx, &StakeGet{
				UserId: user.ID,
			})
			if err != nil {
				return err
			}

			//err = ac.userRepo.CreateStakeGit(ctx, &StakeGit{
			//	UserId: user.ID,
			//})
			//if err != nil {
			//	return err
			//}

			return nil
		}); err != nil {
			return nil, err, "错误"
		}
	}

	return user, nil, ""
}

func (ac *AppUsecase) UserBuy(ctx context.Context, address string) (*pb.UserBuyReply, error) {
	var (
		user *User
		err  error
	)
	user, err = ac.userRepo.GetUserByAddress(ctx, address) // 查询用户
	if nil != err || nil == user {
		return &pb.UserBuyReply{
			Status: "不存在用户",
		}, nil
	}

	if 1 == user.LockUse {
		return &pb.UserBuyReply{
			Status: "锁定用户",
		}, nil
	}

	var (
		configs []*Config
		uPrice  float64
	)

	// 配置
	configs, err = ac.userRepo.GetConfigByKeys(ctx,
		"u_price",
	)
	if nil != err || nil == configs {
		return &pb.UserBuyReply{
			Status: "配置错误",
		}, nil
	}
	for _, vConfig := range configs {
		if "u_price" == vConfig.KeyName {
			uPrice, _ = strconv.ParseFloat(vConfig.Value, 10)
		}
	}

	// 推荐
	var (
		userRecommend   *UserRecommend
		myUserRecommend []*UserRecommend
		team            []*UserRecommend
	)
	userRecommend, err = ac.userRepo.GetUserRecommendByUserId(ctx, user.ID)
	if nil == userRecommend || nil != err {
		return &pb.UserBuyReply{
			Status: "推荐错误查询",
		}, nil
	}

	myUserRecommend, err = ac.userRepo.GetUserRecommendByCode(ctx, userRecommend.RecommendCode+"D"+strconv.FormatUint(user.ID, 10))
	if nil == myUserRecommend || nil != err {
		return &pb.UserBuyReply{
			Status: "推荐错误查询",
		}, nil
	}

	team, err = ac.userRepo.GetUserRecommendLikeCode(ctx, userRecommend.RecommendCode+"D"+strconv.FormatUint(user.ID, 10))
	if nil != err {
		return &pb.UserBuyReply{
			Status: "推荐错误查询",
		}, nil
	}

	var (
		users    []*User
		usersMap map[uint64]*User
	)
	users, err = ac.userRepo.GetAllUsers(ctx)
	if nil == users {
		return &pb.UserBuyReply{
			Status: "错误",
		}, nil
	}

	usersMap = make(map[uint64]*User, 0)
	for _, vUsers := range users {
		usersMap[vUsers.ID] = vUsers
	}

	// 获取业绩
	tmpAreaMax := float64(0)
	tmpAreaMin := float64(0)
	tmpMaxId := uint64(0)
	tmpTwelve := float64(0)
	tmpRecommendNum := uint64(0)
	for _, vMyLowUser := range myUserRecommend {
		if _, ok := usersMap[vMyLowUser.ID]; !ok {
			continue
		}

		tmpTwelve += usersMap[vMyLowUser.UserId].MyTotalAmount
		tmpRecommendNum += 1
		if tmpAreaMax < usersMap[vMyLowUser.UserId].MyTotalAmount+usersMap[vMyLowUser.UserId].Amount {
			tmpAreaMax = usersMap[vMyLowUser.UserId].MyTotalAmount + usersMap[vMyLowUser.UserId].Amount
			tmpMaxId = vMyLowUser.ID
		}
	}

	if 0 < tmpMaxId {
		for _, vMyLowUser := range myUserRecommend {
			if _, ok := usersMap[vMyLowUser.ID]; !ok {
				continue
			}

			if tmpMaxId != vMyLowUser.ID {
				tmpAreaMin += usersMap[vMyLowUser.UserId].MyTotalAmount + usersMap[vMyLowUser.UserId].Amount
			}
		}
	}

	tmpLevel := uint64(0)
	if 0 < user.Vip {
		tmpLevel = user.Vip
	} else {
		if 1000000 <= tmpAreaMin {
			tmpLevel = 5
		} else if 500000 <= tmpAreaMin {
			tmpLevel = 4
		} else if 150000 <= tmpAreaMin {
			tmpLevel = 3
		} else if 50000 <= tmpAreaMin {
			tmpLevel = 2
		} else if 10000 <= tmpAreaMin {
			tmpLevel = 1
		}
	}

	tmpBuyNum := uint64(0)
	for _, v := range team {
		if _, ok := usersMap[v.UserId]; !ok {
			continue
		}

		if 0 >= usersMap[v.UserId].OutNum && 0 >= usersMap[v.UserId].Amount {
			continue
		}

		tmpBuyNum++
	}

	tmpFour := float64(0)
	if user.Amount*2.5 <= user.AmountGet*uPrice {
		tmpFour = 0
	} else {
		tmpFour = user.Amount*2.5 - user.AmountGet*uPrice
	}

	return &pb.UserBuyReply{
		Status:       "ok",
		One:          user.Amount,
		Two:          2.5,
		Three:        user.AmountGet * uPrice,
		Four:         tmpFour,
		Five:         user.Location * uPrice,
		Six:          user.Recommend * uPrice,
		Seven:        user.RecommendTwo * uPrice,
		Eight:        user.Area * uPrice,
		Nine:         user.AreaTwo * uPrice,
		Ten:          user.All * uPrice,
		Elven:        user.MyTotalAmount,
		Twelve:       tmpTwelve,
		Thirteen:     tmpAreaMax,
		Fourteen:     tmpAreaMin,
		RecommendNum: tmpRecommendNum,
		Usdt:         user.AmountUsdt,
		Giw:          user.GiwTwo,
		Price:        uPrice,
		TeamNum:      uint64(len(team)),
		BuyNum:       tmpBuyNum,
		Level:        tmpLevel,
	}, nil
}

func (ac *AppUsecase) UserInfo(ctx context.Context, address string) (*pb.UserInfoReply, error) {
	var (
		user               *User
		boxNum             uint64
		boxSellNum         uint64
		configs            []*Config
		bPrice             float64
		uPrice             float64
		exchangeFeeRate    float64
		exchangeFeeRateTwo float64
		rewardStakeRate    float64
		boxMax             uint64
		boxAmount          float64
		boxStart           string
		boxEnd             string
		stakeOverRate      float64
		sellFeeRate        float64
		withdrawMin        uint64
		withdrawMax        uint64
		withdrawMinTwo     uint64
		withdrawMaxTwo     uint64
		err                error
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
		"u_price",
		"exchange_fee_rate",
		"exchange_fee_rate_two",
		"reward_stake_rate",
		"box_max",
		"box_sell_num",
		"box_start",
		"box_end",
		"box_amount",
		"stake_over_rate",
		"sell_fee_rate",
		"withdraw_amount_min",
		"withdraw_amount_max",
		"withdraw_amount_min_two",
		"withdraw_amount_max_two",
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

		if "withdraw_amount_min_two" == vConfig.KeyName {
			withdrawMinTwo, _ = strconv.ParseUint(vConfig.Value, 10, 64)
		}
		if "withdraw_amount_max_two" == vConfig.KeyName {
			withdrawMaxTwo, _ = strconv.ParseUint(vConfig.Value, 10, 64)
		}
		if "b_price" == vConfig.KeyName {
			bPrice, _ = strconv.ParseFloat(vConfig.Value, 10)
		}
		if "exchange_fee_rate" == vConfig.KeyName {
			exchangeFeeRate, _ = strconv.ParseFloat(vConfig.Value, 10)
		}
		if "exchange_fee_rate_two" == vConfig.KeyName {
			exchangeFeeRateTwo, _ = strconv.ParseFloat(vConfig.Value, 10)
		}
		if "reward_stake_rate" == vConfig.KeyName {
			rewardStakeRate, _ = strconv.ParseFloat(vConfig.Value, 10)
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
		if "stake_over_rate" == vConfig.KeyName {
			stakeOverRate, _ = strconv.ParseFloat(vConfig.Value, 10)
		}
		if "sell_fee_rate" == vConfig.KeyName {
			sellFeeRate, _ = strconv.ParseFloat(vConfig.Value, 10)
		}
		if "u_price" == vConfig.KeyName {
			uPrice, _ = strconv.ParseFloat(vConfig.Value, 10)
		}
	}

	if 0 >= bPrice {
		return &pb.UserInfoReply{
			Status: "币价biw:giw错误",
		}, nil
	}

	if 0 >= uPrice {
		return &pb.UserInfoReply{
			Status: "币价biw:u错误",
		}, nil
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

	var (
		myUserRecommendUserId  uint64
		tmpRecommendUserIdsTmp []string
		addressThree           string
	)

	if 0 < len(userRecommend.RecommendCode) {
		tmpRecommendUserIdsTmp = strings.Split(userRecommend.RecommendCode, "D")
		if 2 <= len(tmpRecommendUserIdsTmp) {
			myUserRecommendUserId, _ = strconv.ParseUint(tmpRecommendUserIdsTmp[len(tmpRecommendUserIdsTmp)-1], 10, 64) // 最后一位是直推人
		}

		if 0 < myUserRecommendUserId {
			var (
				myTopUser *User
			)
			myTopUser, err = ac.userRepo.GetUserById(ctx, myUserRecommendUserId)
			if nil != err {
				return &pb.UserInfoReply{
					Status: "上级查询错误",
				}, nil
			}

			addressThree = myTopUser.Address
		}

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
	RecommendTotalReward := RecommendTotalRewardOne + RecommendTotalRewardTwo + RecommendTotalRewardThree

	var (
		stakeGitRecord []*StakeGitRecord
	)
	stakeGitRecord, err = ac.userRepo.GetStakeGitRecordsByUserID(ctx, user.ID, nil)
	if nil != err {
		return &pb.UserInfoReply{
			Status: "粮仓错误查询",
		}, nil
	}
	stakeGitAmount := float64(0)
	stakeGitAmountToday := float64(0)

	// 获取中国时区（Asia/Shanghai）
	locShanghai, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		return &pb.UserInfoReply{
			Status: "时区查询",
		}, nil
	}
	// 获取当前时间（假设服务器时间是 UTC）
	nowUTC := time.Now()
	// 转换当前时间为中国时区
	nowInShanghai := nowUTC.In(locShanghai)
	for _, v := range stakeGitRecord {
		stakeGitAmount += v.Amount
		// 转换用户注册时间到中国时区
		userRegisterInShanghai := v.CreatedAt.In(locShanghai)
		// 计算用户注册当天在中国时区的 24:00（即第二天 00:00:00）
		nextMidnight := time.Date(userRegisterInShanghai.Year(), userRegisterInShanghai.Month(), userRegisterInShanghai.Day()+1, 0, 0, 0, 0, locShanghai)
		// 判断是否超过注册当天的 24:00
		if nowInShanghai.After(nextMidnight) {
			stakeGitAmountToday += v.Amount
		}
	}

	todayStakeGitAmount := stakeGitAmountToday * rewardStakeRate
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
		stakeGetTotalMy = float64(0)
		stakeGet        *StakeGet

		stakeGetTotalAmount float64
		stakeGetTotal       *StakeGetTotal
	)

	stakeGetTotal, err = ac.userRepo.GetStakeGetTotal(ctx)
	if nil != err || nil == stakeGetTotal {
		return &pb.UserInfoReply{
			Status: "我的放大器错误查询",
		}, nil
	}
	stakeGetTotalAmount = stakeGetTotal.Balance

	stakeGet, err = ac.userRepo.GetUserStakeGet(ctx, user.ID)
	if nil != err {
		return &pb.UserInfoReply{
			Status: "我的放大器错误查询",
		}, nil
	}
	if nil != stakeGet {
		if 0 < stakeGetTotal.Amount {
			// 每份价值
			valuePerShare := stakeGetTotalAmount / stakeGetTotal.Amount
			// 用户最大可提取金额
			stakeGetTotalMy = stakeGet.StakeRate * valuePerShare
		}
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
		MyStakeGit:                stakeGitAmount,
		TodayRewardSkateGit:       todayStakeGitAmount,
		RewardStakeRate:           rewardStakeRate,
		Box:                       boxMax,
		BoxSell:                   boxSellNum,
		Start:                     boxStart,
		End:                       boxEnd,
		BoxSellAmount:             boxAmount / uPrice,
		ExchangeRate:              bPrice / uPrice,
		ExchangeRateTwo:           uPrice,
		BiwPrice:                  uPrice,
		UsdtTwo:                   user.UsdtTwo,
		ExchangeFeeRate:           exchangeFeeRate,
		ExchangeFeeRateTwo:        exchangeFeeRateTwo,
		StakeGetTotal:             stakeGetTotalAmount,
		MyStakeGetTotal:           stakeGetTotalMy,
		StakeGetOverFeeRate:       stakeOverRate,
		SellFeeRate:               sellFeeRate,
		WithdrawMax:               withdrawMax,
		WithdrawMin:               withdrawMin,
		WithdrawMaxTwo:            withdrawMaxTwo,
		WithdrawMinTwo:            withdrawMinTwo,
		Usdt:                      user.AmountUsdt,
		CreatedAt:                 user.CreatedAt.Add(8 * time.Hour).Format("2006-01-02 15:04:05"),
		AddressThree:              addressThree,
		GiwTwo:                    user.GiwTwo,
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

func (ac *AppUsecase) UserBuyL(ctx context.Context, address string, req *pb.UserBuyLRequest) (*pb.UserBuyLReply, error) {
	res := make([]*pb.UserBuyLReply_List, 0)
	var (
		user  *User
		err   error
		count int64
	)
	user, err = ac.userRepo.GetUserByAddress(ctx, address) // 查询用户
	if nil != err || nil == user {
		return &pb.UserBuyLReply{
			Status: "不存在用户",
		}, nil
	}

	var (
		configs []*Config
		uPrice  float64
	)

	// 配置
	configs, err = ac.userRepo.GetConfigByKeys(ctx,
		"u_price",
	)
	if nil != err || nil == configs {
		return &pb.UserBuyLReply{
			Status: "配置错误",
		}, nil
	}
	for _, vConfig := range configs {
		if "u_price" == vConfig.KeyName {
			uPrice, _ = strconv.ParseFloat(vConfig.Value, 10)
		}
	}

	var (
		reward []*RewardTwo
	)
	if 1 <= req.Num && req.Num <= 7 {

	} else {
		return &pb.UserBuyLReply{
			Status: "参数错误",
		}, nil
	}

	count, err = ac.userRepo.GetUserRewardTwoPageCount(ctx, user.ID, req.Num)
	if nil != err {
		return &pb.UserBuyLReply{
			Status: "不存在数据L，count",
		}, nil
	}

	reward, err = ac.userRepo.GetUserRewardTwoPage(ctx, user.ID, req.Num, &Pagination{
		PageNum:  int(req.Page),
		PageSize: 20,
	})
	if nil != err {
		return &pb.UserBuyLReply{
			Status: "不存在数据L",
		}, nil
	}

	for _, v := range reward {

		if 1 == v.Reason {
			res = append(res, &pb.UserBuyLReply_List{
				AmountThree: v.Amount,
				Amount:      v.Five,
				AmountTwo:   v.Three,
				Address:     v.Four,
				Num:         v.One,
				CreatedAt:   v.CreatedAt.Add(8 * time.Hour).Format("2006-01-02 15:04:05"),
			})
		} else {
			res = append(res, &pb.UserBuyLReply_List{
				Amount:    v.Three * uPrice,
				AmountTwo: v.Amount,
				Address:   v.Four,
				Num:       v.One,
				CreatedAt: v.CreatedAt.Add(8 * time.Hour).Format("2006-01-02 15:04:05"),
			})
		}
	}

	return &pb.UserBuyLReply{
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
			Id:         v.ID,
			Level:      v.Level,
			Health:     v.MaxHealth,
			Status:     statusTmp,
			OutRate:    v.OutPutRate,
			PerHealth:  v.PerHealth,
			RentAmount: v.RentOutPutRate,
			One:        v.One,
			Two:        v.Two,
			Three:      v.Three,
			Content:    "在Magic Manor大陆最肥沃的土地，由神秘的地契合成，层叠强大的成长性，任何劣质的种子都可以得到茁壮的成长。",
			EContent:   "The most fertile land in the Magic Manorcontinent is composed of mysterious landdeeds. lt has strong growth potential, andany low-quality seeds can grow vigorously.",
		})
	}

	return &pb.UserLandReply{
		Status: "ok",
		Count:  uint64(len(res)),
		List:   res,
	}, nil
}

func (ac *AppUsecase) UserStakeGitStakeList(ctx context.Context, address string, req *pb.UserStakeGitStakeListRequest) (*pb.UserStakeGitStakeListReply, error) {
	res := make([]*pb.UserStakeGitStakeListReply_List, 0)
	var (
		user *User
		err  error
	)
	user, err = ac.userRepo.GetUserByAddress(ctx, address) // 查询用户
	if nil != err || nil == user {
		return &pb.UserStakeGitStakeListReply{
			Status: "不存在用户",
		}, nil
	}

	var (
		stakeGitRecord []*StakeGitRecord
	)
	stakeGitRecord, err = ac.userRepo.GetStakeGitRecordsByUserID(ctx, user.ID, nil)
	if nil != err {
		return &pb.UserStakeGitStakeListReply{
			Status: "粮仓错误查询",
		}, nil
	}

	for _, v := range stakeGitRecord {
		res = append(res, &pb.UserStakeGitStakeListReply_List{
			Id:        v.ID,
			Stake:     v.Amount,
			CreatedAt: v.CreatedAt.Add(8 * time.Hour).Format("2006-01-02 15:04:05"),
		})
	}

	return &pb.UserStakeGitStakeListReply{
		Status: "ok",
		Count:  0,
		List:   res,
	}, err
}

func (ac *AppUsecase) UserStakeGitRewardList(ctx context.Context, address string, req *pb.UserStakeGitRewardListRequest) (*pb.UserStakeGitRewardListReply, error) {
	res := make([]*pb.UserStakeGitRewardListReply_List, 0)
	var (
		user *User
		err  error
	)
	user, err = ac.userRepo.GetUserByAddress(ctx, address) // 查询用户
	if nil != err || nil == user {
		return &pb.UserStakeGitRewardListReply{
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
		return &pb.UserStakeGitRewardListReply{
			Status: "粮仓查询错误",
		}, nil
	}

	reward, err = ac.userRepo.GetUserRewardPage(ctx, user.ID, status, &Pagination{
		PageNum:  int(req.Page),
		PageSize: 20,
	})
	if nil != err {
		return &pb.UserStakeGitRewardListReply{
			Status: "粮仓查询错误",
		}, nil
	}

	for _, v := range reward {
		res = append(res, &pb.UserStakeGitRewardListReply_List{
			Amount:    v.Amount,
			CreatedAt: v.CreatedAt.Add(8 * time.Hour).Format("2006-01-02 15:04:05"),
		})
	}

	return &pb.UserStakeGitRewardListReply{
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
			Id:       vSeed.ID,
			Type:     1,
			Num:      vSeed.SeedId,
			UseNum:   0,
			Status:   tmpStatus,
			OutMax:   vSeed.OutMaxAmount,
			Time:     vSeed.OutOverTime,
			Amount:   vSeed.SellAmount,
			Content:  "一种来自区块链世界算法加密的种子，打开盲盒或合成获得，每一颗种子都不同，找到适合他的土地后，会有惊人的产出！",
			EContent: "A seed from the blockchain world algorithm encryption, open the blind box or synthetic access, each seed is different, find suitable for his land, there will be amazing output!",
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
		contentTmp := ""
		eContentTmp := ""
		if 12 == vProp.PropType {
			useNum = uint64(vProp.ThreeOne) // 水
			contentTmp = "一种植物在成长过程中也许会用到的特殊用水，缺水的植物会停止成长，使用后您的植物将永远不会缺水。"
			eContentTmp = "A plant may use special water duringits growth. Plants that are short ofwater will stop growing. After using it,your plants will never be short of water."
		} else if 13 == vProp.PropType {
			useNum = uint64(vProp.FiveOne) // 手套
			contentTmp = "带上手套后，可偷取任何邻居家的已经成熟的植物，可获得一定的GIT，但是使用的次数有限。"
			eContentTmp = "After wearing gloves, you can stealany mature plants from your neighbor'shouse and obtain a certain amount ofGlT, but the number of uses is limited."
		} else if 14 == vProp.PropType {
			useNum = uint64(vProp.FourOne) // 除虫剂
			contentTmp = "由Magic Manor大陆中的巫师制作，不除虫子的植物，每5分钟减产1%;直到最后为产量为0，它可以杀死Magic Manor大陆中的任何害虫。"
			eContentTmp = "Made by wizards in the Magic Manorcontinent, plants that do not eliminateinsects will reduce their yield by 1% every5 minutes; until the final yield is O, it can killany pests in the Magic Manor continent."
		} else if 15 == vProp.PropType {
			useNum = uint64(vProp.TwoOne) // 铲子
			contentTmp = "可铲除出租土地上已经成熟的植物，不可铲除自己种植的植物，但是成熟时间必须大于1H。"
			eContentTmp = "Mature plants on the leased land canbe eradicated, but self-grown plantscannot be eradicated, but the maturitytime must be greater than 1H."
		} else if 11 == vProp.PropType {
			contentTmp = "一种通过算法生成的增加产量道具，合成士地和增加土地肥沃度。"
			eContentTmp = "An algorithmically generated item thatincreases yield, synthesizes land, andincreases land fertility."
		} else if 17 == vProp.PropType {
			contentTmp = "在Magic Manor大陆深处埋徵着一张”地契”，它是初创统治者亲手制作，代表着整个Magic Manor最肥沃的土地，找到它的人不仅可以拥有土地，还能解锁谷中隐藏的古老秘密。\n\t\t获取方式:每新增1亿GIT产出业绩，自动获得1张地契:作用:1张地契加5块化肥，可合成一个崭新的1级土地;地契描述"
			eContentTmp = "There is a \"land deed\" buried deep in the Magic Manorcontinent. lt was made by the original ruler himself andrepresents the most fertile land in the entire MagicManor. Whoever finds it will not only claim the land, butalso unlock ancient secrets hidden within the valley.How to obtain: For every 100 milion new GlT outputs,you will automatically obtain a land deed;Function: 1 land deed and 5 fertilizers can be combinedinto a brand new level 1 land."
		}

		res = append(res, &pb.UserBackListReply_List{
			Id:       vProp.ID,
			Type:     2,
			Num:      uint64(vProp.PropType),
			UseNum:   useNum,
			Status:   uint64(vProp.Status),
			OutMax:   0,
			Amount:   vProp.SellAmount,
			Content:  contentTmp,
			EContent: eContentTmp,
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

	if 0 >= len(seed) {
		return &pb.UserMarketSeedListReply{
			Status: "ok",
			Count:  0,
			List:   res,
		}, nil
	}

	userIds := make([]uint64, 0)
	for _, vSeed := range seed {
		userIds = append(userIds, vSeed.UserId)
	}

	usersMap := make(map[uint64]*User)
	usersMap, err = ac.userRepo.GetUserByUserIds(ctx, userIds)
	if nil != err {
		return &pb.UserMarketSeedListReply{
			Status: "查询错误",
		}, nil
	}

	for _, vSeed := range seed {
		addressTmp := ""
		if _, ok := usersMap[vSeed.UserId]; ok {
			addressTmp = usersMap[vSeed.UserId].Address
		}

		res = append(res, &pb.UserMarketSeedListReply_List{
			Id:       vSeed.ID,
			Num:      vSeed.SeedId,
			Amount:   vSeed.SellAmount,
			OutMax:   vSeed.OutMaxAmount,
			Time:     vSeed.OutOverTime,
			Address:  addressTmp,
			Content:  "一种来自区块链世界算法加密的种子，打开盲盒或合成获得，每一颗种子都不同，找到适合他的土地后，会有惊人的产出！",
			EContent: "A seed from the blockchain world algorithm encryption, open the blind box or synthetic access, each seed is different, find suitable for his land, there will be amazing output!",
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

	if 0 >= len(land) {
		return &pb.UserMarketLandListReply{
			Status: "ok",
			Count:  0,
			List:   res,
		}, nil
	}

	userIds := make([]uint64, 0)
	for _, vLand := range land {
		userIds = append(userIds, vLand.UserId)
	}

	usersMap := make(map[uint64]*User)
	usersMap, err = ac.userRepo.GetUserByUserIds(ctx, userIds)
	if nil != err {
		return &pb.UserMarketLandListReply{
			Status: "错误查询",
		}, nil
	}

	for _, vLand := range land {
		addressTmp := ""
		if _, ok := usersMap[vLand.UserId]; ok {
			addressTmp = usersMap[vLand.UserId].Address
		}

		res = append(res, &pb.UserMarketLandListReply_List{
			Id:         vLand.ID,
			Level:      vLand.Level,
			MaxHealth:  vLand.MaxHealth,
			Amount:     vLand.SellAmount,
			PerHealth:  vLand.PerHealth,
			OutPutRate: uint64(vLand.OutPutRate),
			Address:    addressTmp,
			Content:    "在Magic Manor大陆最肥沃的土地，由神秘的地契合成，层叠强大的成长性，任何劣质的种子都可以得到茁壮的成长。",
			EContent:   "The most fertile land in the Magic Manorcontinent is composed of mysterious landdeeds. lt has strong growth potential, andany low-quality seeds can grow vigorously.",
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

	if 0 >= len(prop) {
		return &pb.UserMarketPropListReply{
			Status: "ok",
			Count:  0,
			List:   res,
		}, nil
	}

	userIds := make([]uint64, 0)
	for _, vProp := range prop {
		userIds = append(userIds, vProp.UserId)
	}

	usersMap := make(map[uint64]*User)
	usersMap, err = ac.userRepo.GetUserByUserIds(ctx, userIds)
	if nil != err {
		return &pb.UserMarketPropListReply{
			Status: "错误查询",
		}, nil
	}

	for _, v := range prop {
		addressTmp := ""
		if _, ok := usersMap[v.UserId]; ok {
			addressTmp = usersMap[v.UserId].Address
		}

		useNum := uint64(0)
		contentTmp := ""
		eContentTmp := ""
		if 12 == v.PropType {
			useNum = uint64(v.ThreeOne) // 水
			contentTmp = "一种植物在成长过程中也许会用到的特殊用水，缺水的植物会停止成长，使用后您的植物将永远不会缺水。"
			eContentTmp = "A plant may use special water duringits growth. Plants that are short ofwater will stop growing. After using it,your plants will never be short of water."
		} else if 13 == v.PropType {
			useNum = uint64(v.FiveOne) // 手套
			contentTmp = "带上手套后，可偷取任何邻居家的已经成熟的植物，可获得一定的GIT，但是使用的次数有限。"
			eContentTmp = "After wearing gloves, you can stealany mature plants from your neighbor'shouse and obtain a certain amount ofGlT, but the number of uses is limited."
		} else if 14 == v.PropType {
			useNum = uint64(v.FourOne) // 除虫剂
			contentTmp = "由Magic Manor大陆中的巫师制作，不除虫子的植物，每5分钟减产1%;直到最后为产量为0，它可以杀死Magic Manor大陆中的任何害虫。"
			eContentTmp = "Made by wizards in the Magic Manorcontinent, plants that do not eliminateinsects will reduce their yield by 1% every5 minutes; until the final yield is O, it can killany pests in the Magic Manor continent."
		} else if 15 == v.PropType {
			useNum = uint64(v.TwoOne) // 铲子
			contentTmp = "可铲除出租土地上已经成熟的植物，不可铲除自己种植的植物，但是成熟时间必须大于1H。"
			eContentTmp = "Mature plants on the leased land canbe eradicated, but self-grown plantscannot be eradicated, but the maturitytime must be greater than 1H."
		} else if 11 == v.PropType {
			contentTmp = "一种通过算法生成的增加产量道具，合成士地和增加土地肥沃度。"
			eContentTmp = "An algorithmically generated item thatincreases yield, synthesizes land, andincreases land fertility."
		} else if 17 == v.PropType {
			contentTmp = "在Magic Manor大陆深处埋徵着一张”地契”，它是初创统治者亲手制作，代表着整个Magic Manor最肥沃的土地，找到它的人不仅可以拥有土地，还能解锁谷中隐藏的古老秘密。\n\t\t获取方式:每新增1亿GIT产出业绩，自动获得1张地契:作用:1张地契加5块化肥，可合成一个崭新的1级土地;地契描述"
			eContentTmp = "There is a \"land deed\" buried deep in the Magic Manorcontinent. lt was made by the original ruler himself andrepresents the most fertile land in the entire MagicManor. Whoever finds it will not only claim the land, butalso unlock ancient secrets hidden within the valley.How to obtain: For every 100 milion new GlT outputs,you will automatically obtain a land deed;Function: 1 land deed and 5 fertilizers can be combinedinto a brand new level 1 land."
		}

		res = append(res, &pb.UserMarketPropListReply_List{
			Id:       v.ID,
			Num:      uint64(v.PropType),
			Amount:   v.SellAmount,
			UseMax:   useNum,
			Address:  addressTmp,
			Content:  contentTmp,
			EContent: eContentTmp,
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

	if 2 == req.RentType {
		var (
			landUserUse []*LandUserUse
		)

		landUserUse, err = ac.userRepo.GetLandUserUseByLandIDsUsing(ctx, user.ID)
		if nil != err {
			return &pb.UserMarketRentLandListReply{
				Status: "错误查询",
			}, nil
		}

		// 找出租用的种植信息
		landIds := make([]uint64, 0)
		for _, v := range landUserUse {
			if user.ID == v.OwnerUserId {
				continue
			}

			landIds = append(landIds, v.LandId)
		}

		if 0 >= len(landIds) {
			return &pb.UserMarketRentLandListReply{
				Status: "ok",
				Count:  0,
				List:   res,
			}, nil
		}

		land, err = ac.userRepo.GetLandByExUserIDByIds(ctx, landIds, nil)
		if nil != err {
			return &pb.UserMarketRentLandListReply{
				Status: "错误查询",
			}, nil
		}
	} else {
		landStatus := []uint64{3}
		land, err = ac.userRepo.GetLandByExUserID(ctx, user.ID, landStatus, nil)
		if nil != err {
			return &pb.UserMarketRentLandListReply{
				Status: "错误查询",
			}, nil
		}
	}

	if 0 >= len(land) {
		return &pb.UserMarketRentLandListReply{
			Status: "ok",
			Count:  0,
			List:   res,
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
			OutPutRate: vLand.OutPutRate,
			PerHealth:  vLand.PerHealth,
			Content:    "在Magic Manor大陆最肥沃的土地，由神秘的地契合成，层叠强大的成长性，任何劣质的种子都可以得到茁壮的成长。",
			EContent:   "The most fertile land in the Magic Manorcontinent is composed of mysterious landdeeds. lt has strong growth potential, andany low-quality seeds can grow vigorously.",
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
			Time:       vSeed.OutOverTime,
			Address:    address,
			Content:    "一种来自区块链世界算法加密的种子，打开盲盒或合成获得，每一颗种子都不同，找到适合他的土地后，会有惊人的产出！",
			EContent:   "A seed from the blockchain world algorithm encryption, open the blind box or synthetic access, each seed is different, find suitable for his land, there will be amazing output!",
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
		contentTmp := ""
		eContentTmp := ""
		if 12 == vProp.PropType {
			useNum = uint64(vProp.ThreeOne) // 水
			contentTmp = "一种植物在成长过程中也许会用到的特殊用水，缺水的植物会停止成长，使用后您的植物将永远不会缺水。"
			eContentTmp = "A plant may use special water duringits growth. Plants that are short ofwater will stop growing. After using it,your plants will never be short of water."
		} else if 13 == vProp.PropType {
			useNum = uint64(vProp.FiveOne) // 手套
			contentTmp = "带上手套后，可偷取任何邻居家的已经成熟的植物，可获得一定的GIT，但是使用的次数有限。"
			eContentTmp = "After wearing gloves, you can stealany mature plants from your neighbor'shouse and obtain a certain amount ofGlT, but the number of uses is limited."
		} else if 14 == vProp.PropType {
			useNum = uint64(vProp.FourOne) // 除虫剂
			contentTmp = "由Magic Manor大陆中的巫师制作，不除虫子的植物，每5分钟减产1%;直到最后为产量为0，它可以杀死Magic Manor大陆中的任何害虫。"
			eContentTmp = "Made by wizards in the Magic Manorcontinent, plants that do not eliminateinsects will reduce their yield by 1% every5 minutes; until the final yield is O, it can killany pests in the Magic Manor continent."
		} else if 15 == vProp.PropType {
			useNum = uint64(vProp.TwoOne) // 铲子
			contentTmp = "可铲除出租土地上已经成熟的植物，不可铲除自己种植的植物，但是成熟时间必须大于1H。"
			eContentTmp = "Mature plants on the leased land canbe eradicated, but self-grown plantscannot be eradicated, but the maturitytime must be greater than 1H."
		} else if 11 == vProp.PropType {
			contentTmp = "一种通过算法生成的增加产量道具，合成士地和增加土地肥沃度。"
			eContentTmp = "An algorithmically generated item thatincreases yield, synthesizes land, andincreases land fertility."
		} else if 17 == vProp.PropType {
			contentTmp = "在Magic Manor大陆深处埋徵着一张”地契”，它是初创统治者亲手制作，代表着整个Magic Manor最肥沃的土地，找到它的人不仅可以拥有土地，还能解锁谷中隐藏的古老秘密。\n\t\t获取方式:每新增1亿GIT产出业绩，自动获得1张地契:作用:1张地契加5块化肥，可合成一个崭新的1级土地;地契描述"
			eContentTmp = "There is a \"land deed\" buried deep in the Magic Manorcontinent. lt was made by the original ruler himself andrepresents the most fertile land in the entire MagicManor. Whoever finds it will not only claim the land, butalso unlock ancient secrets hidden within the valley.How to obtain: For every 100 milion new GlT outputs,you will automatically obtain a land deed;Function: 1 land deed and 5 fertilizers can be combinedinto a brand new level 1 land."
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
			Address:    address,
			Content:    contentTmp,
			EContent:   eContentTmp,
		})
	}

	var (
		land []*Land
	)
	landStatus := []uint64{3, 4, 8}
	land, err = ac.userRepo.GetLandByUserID(ctx, user.ID, landStatus, nil)
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
			PerHealth:  vLand.PerHealth,
			OutPutRate: uint64(vLand.OutPutRate),
			Address:    address,
			Content:    "在Magic Manor大陆最肥沃的土地，由神秘的地契合成，层叠强大的成长性，任何劣质的种子都可以得到茁壮的成长。",
			EContent:   "The most fertile land in the Magic Manorcontinent is composed of mysterious landdeeds. lt has strong growth potential, andany low-quality seeds can grow vigorously.",
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
			EContent:  vNotice.NoticeContentTwo,
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

// UserStakeRewardList userStakeRewardList.
func (ac *AppUsecase) UserStakeRewardList(ctx context.Context, address string, req *pb.UserStakeRewardListRequest) (*pb.UserStakeRewardListReply, error) {
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

	res := make([]*pb.UserStakeRewardListReply_List, 0)

	var (
		stakeGetPlayRecord []*StakeGetPlayRecord
		count              int64
	)

	count, err = ac.userRepo.GetStakeGetPlayRecordCount(ctx, user.ID, 0)
	if nil != err {
		return &pb.UserStakeRewardListReply{
			Status: "推荐错误查询",
		}, nil
	}

	stakeGetPlayRecord, err = ac.userRepo.GetStakeGetPlayRecordsByUserID(ctx, user.ID, 0, &Pagination{
		PageNum:  int(req.Page),
		PageSize: 20,
	})
	if nil != err {
		return &pb.UserStakeRewardListReply{
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
		return &pb.UserStakeRewardListReply{
			Status: "错误查询",
		}, nil
	}

	for _, v := range stakeGetPlayRecord {
		addressTmp := ""
		if _, ok := usersMap[v.UserId]; ok {
			addressTmp = usersMap[v.UserId].Address
		}

		res = append(res, &pb.UserStakeRewardListReply_List{
			Address: addressTmp,
			Content: "",
			Amount:  v.Amount,
			Reward:  v.Reward,
			Status:  uint64(v.Status),
		})
	}

	return &pb.UserStakeRewardListReply{
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

			tmpRewardStatus := uint64(2)
			rewardTmp := float64(0)
			statusTmp := uint64(1)
			if 0 != landUserUse[vLand.ID].One {
				if landUserUse[vLand.ID].One <= uint64(now) {
					statusTmp = 3
				}

			} else if 0 != landUserUse[vLand.ID].Two {
				if landUserUse[vLand.ID].Two <= uint64(now) {
					statusTmp = 2
				}

				// 有虫子但是已经结束
				if landUserUse[vLand.ID].OverTime <= uint64(now) {
					if uint64(now) > landUserUse[vLand.ID].Two {
						tmp := landUserUse[vLand.ID].OutMaxNum * 0.01 * float64(uint64(now)-landUserUse[vLand.ID].Two) / 300

						if tmp >= landUserUse[vLand.ID].OutMaxNum {
							rewardTmp = 0
						} else {
							rewardTmp = landUserUse[vLand.ID].OutMaxNum - tmp
						}
					}

					tmpRewardStatus = 1
				}
			} else {
				if landUserUse[vLand.ID].OverTime <= uint64(now) {
					rewardTmp = landUserUse[vLand.ID].OutMaxNum
					tmpRewardStatus = 1
				}
			}

			tmpOneOne := uint64(0)
			if 0 < landUserUse[vLand.ID].SubTime {
				tmpOneOne = 1
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
				RewardStatus:     tmpRewardStatus,
				LandStatus:       vLand.Status,
				OneOne:           tmpOneOne,
			}
		} else {
			// 过期的从放置在移除
			if 1 == vLand.Status || 3 == vLand.Status {
				if vLand.LimitDate <= uint64(time.Now().Unix()) {
					if err = ac.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
						return ac.userRepo.LandPullTwo(ctx, vLand.ID, user.ID)
					}); nil != err {
						fmt.Println(err, "首页触发的，过期土地回收", user)
					}

					continue
				}
			}

			resTmp[vLand.LocationNum] = &pb.UserIndexListReply_List{
				LocationNum:      vLand.LocationNum,
				LandId:           vLand.ID,
				LandLevel:        vLand.Level,
				Health:           vLand.MaxHealth,
				OutRate:          vLand.OutPutRate,
				PerHealth:        vLand.PerHealth,
				LandStatus:       vLand.Status,
				LandUseId:        0,
				SeedId:           0,
				Start:            0,
				End:              0,
				CurrentTime:      0,
				Status:           0,
				Reward:           0,
				PlantUserAddress: plantUserAddressTmp,
				RewardStatus:     2,
			}
		}
	}

	for i := uint64(1); i <= 9; i++ {
		if _, ok := resTmp[i]; !ok {
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
		uPrice           float64
	)
	user, err = ac.userRepo.GetUserByAddress(ctx, address) // 查询用户
	if nil != err || nil == user {
		return &pb.BuyBoxReply{
			Status: "不存在用户",
		}, nil
	}

	if 1 == user.LockUse {
		return &pb.BuyBoxReply{
			Status: "锁定用户",
		}, nil
	}

	rngMutexBuyBox.Lock()
	defer rngMutexBuyBox.Unlock()

	user, err = ac.userRepo.GetUserByAddress(ctx, address) // 查询用户
	if nil != err || nil == user {
		return &pb.BuyBoxReply{
			Status: "不存在用户",
		}, nil
	}

	// 配置
	configs, err = ac.userRepo.GetConfigByKeys(ctx,
		"box_num",
		"box_max",
		"box_sell_num",
		"box_start",
		"box_end",
		"box_amount",
		"u_price",
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
			boxSellNumOrigin = vConfig.Value
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
		if "u_price" == vConfig.KeyName {
			uPrice, _ = strconv.ParseFloat(vConfig.Value, 10)
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

	if 0 >= uPrice {
		return &pb.BuyBoxReply{
			Status: "价格biw:u错误",
		}, nil
	}

	boxAmount = boxAmount / uPrice
	if boxAmount > user.Giw {
		return &pb.BuyBoxReply{
			Status: "余额不足",
		}, nil
	}

	tmpSellNumNew := strconv.FormatUint(boxSellNum+1, 10)
	//fmt.Println(boxSellNum, tmpSellNumNew)
	boxId := uint64(0)
	if err = ac.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
		boxId, err = ac.userRepo.BuyBox(ctx, boxAmount, boxSellNumOrigin, tmpSellNumNew, &BoxRecord{
			UserId: user.ID,
			Num:    boxNum,
		})
		if nil != err {
			return err
		}

		err = ac.userRepo.CreateNotice(
			ctx,
			user.ID,
			"您花费"+fmt.Sprintf("%.2f", boxAmount)+"BIW购买了盲盒",
			"You've used "+fmt.Sprintf("%.2f", boxAmount)+" BIW buy box",
		)
		if nil != err {
			return err
		}

		return nil
	}); nil != err {
		fmt.Println(err, "buybox", user)
		return &pb.BuyBoxReply{
			Status: "购买失败",
		}, nil
	}

	return &pb.BuyBoxReply{
		Status: "ok",
		Id:     boxId,
	}, nil
}

// 随机数生成器
var rngBox *rand2.Rand
var rngPlant *rand2.Rand

// 随机数生成器的初始化锁
var rngMutexBox sync.Mutex

func (ac *AppUsecase) OpenBox(ctx context.Context, address string, req *pb.OpenBoxRequest) (*pb.OpenBoxReply, error) {
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

	if 1 == user.LockUse {
		return &pb.OpenBoxReply{
			Status: "锁定用户",
		}, nil
	}

	rngMutexBox.Lock()
	defer rngMutexBox.Unlock()

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
		boxId := uint64(0)
		if err = ac.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
			boxId, err = ac.userRepo.OpenBoxSeed(ctx, box.ID, "获得种子", &Seed{
				UserId:       user.ID,
				SeedId:       result,
				Name:         seedInfosMap[result].Name,
				OutOverTime:  seedInfosMap[result].OutOverTime,
				OutMaxAmount: float64(randomNumber),
			})
			if nil != err {
				return err
			}

			//err = ac.userRepo.CreateNotice(
			//	ctx,
			//	user.ID,
			//	"您开盒子获得了, 种子",
			//	"You've get seed from box",
			//)
			//if nil != err {
			//	return err
			//}

			return nil
		}); nil != err {
			fmt.Println(err, "openBox", user)
			return &pb.OpenBoxReply{
				Status: "开启失败",
			}, nil
		}

		return &pb.OpenBoxReply{
			Id:       boxId,
			Status:   "ok",
			OpenType: 1,
			Num:      result,
			OutMax:   float64(randomNumber),
			Time:     seedInfosMap[result].OutOverTime,
		}, nil
	} else if 11 <= result && result <= 15 {
		if _, ok := propInfosMap[result]; !ok {
			return &pb.OpenBoxReply{
				Status: "不存在盲盒信息",
			}, nil
		}

		// 种子
		boxId := uint64(0)
		if err = ac.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
			boxId, err = ac.userRepo.OpenBoxProp(ctx, box.ID, "获得道具", &Prop{
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
			if nil != err {
				return err
			}

			//err = ac.userRepo.CreateNotice(
			//	ctx,
			//	user.ID,
			//	"您开盒子获得了, 道具",
			//	"You've get prop from box",
			//)
			//if nil != err {
			//	return err
			//}

			return nil
		}); nil != err {
			fmt.Println(err, "openBox", user)
			return &pb.OpenBoxReply{
				Status: "开启失败",
			}, nil
		}

		useNum := uint64(0)
		if 12 == propInfosMap[result].PropType {
			useNum = propInfosMap[result].ThreeOne // 水
		} else if 13 == propInfosMap[result].PropType {
			useNum = propInfosMap[result].FiveOne // 手套
		} else if 14 == propInfosMap[result].PropType {
			useNum = propInfosMap[result].FourOne // 除虫剂
		} else if 15 == propInfosMap[result].PropType {
			useNum = propInfosMap[result].TwoOne // 铲子
		}

		return &pb.OpenBoxReply{
			Id:       boxId,
			Status:   "ok",
			OpenType: 2,
			Num:      result,
			OutMax:   0,
			UseNum:   useNum,
		}, nil
	} else {
		return &pb.OpenBoxReply{
			Status: "开盲盒失败",
		}, nil
	}
}

var rngMutexPlant sync.Mutex

// LandPlayOne 种植
func (ac *AppUsecase) LandPlayOne(ctx context.Context, address string, req *pb.LandPlayOneRequest) (*pb.LandPlayOneReply, error) {
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

	if 1 == user.LockUse {
		return &pb.LandPlayOneReply{
			Status: "锁定用户",
		}, nil
	}

	rngMutexPlant.Lock()
	defer rngMutexPlant.Unlock()

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
			Status: "土地不存在或已失效，撤销布置",
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
				Status: "已出租土地",
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
	rngTmp := rand2.New(rand2.NewSource(time.Now().UnixNano()))
	outMin := int64(now)
	outMax := int64(now + seed.OutOverTime)

	// 计算随机范围
	tmpNum := outMax - outMin
	if tmpNum <= 0 {
		tmpNum = 1 // 避免 Int63n(0) panic
	}
	// 生成随机数
	randomNumber := outMin + rngTmp.Int63n(tmpNum)

	one := uint64(0)
	two := uint64(0)
	r := rngPlant.Float64() // 生成 0.0 ~ 1.0 之间的随机数
	if r < 0.05 {
		one = uint64(randomNumber)
	} else if r < 0.1 {
		two = uint64(randomNumber)
	}

	originStatusTmp := land.Status
	statusTmp := uint64(2)
	if 3 == originStatusTmp {
		statusTmp = 8
	}

	if err = ac.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
		err = ac.userRepo.Plant(ctx, statusTmp, originStatusTmp, land.PerHealth, &LandUserUse{
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
		if nil != err {
			return err
		}

		err = ac.userRepo.CreateNotice(
			ctx,
			user.ID,
			"您种植了一个产量为"+fmt.Sprintf("%.2f", seed.OutMaxAmount)+"GIW的种子",
			"You've plant a seed with output "+fmt.Sprintf("%.2f", seed.OutMaxAmount)+" GIW",
		)
		if nil != err {
			return err
		}

		return nil
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

// LandPlayTwo 收果实
func (ac *AppUsecase) LandPlayTwo(ctx context.Context, address string, req *pb.LandPlayTwoRequest) (*pb.LandPlayTwoReply, error) {
	var (
		configs    []*Config
		user       *User
		oneRate    float64
		twoRate    float64
		threeRate  float64
		uPrice     float64
		lowRewardU float64
		err        error
	)

	user, err = ac.userRepo.GetUserByAddress(ctx, address) // 查询用户
	if nil != err || nil == user {
		return &pb.LandPlayTwoReply{
			Status: "不存在用户",
		}, nil
	}

	if 1 == user.LockUse {
		return &pb.LandPlayTwoReply{
			Status: "锁定用户",
		}, nil
	}

	rngMutexPlantTwo.Lock()
	defer rngMutexPlantTwo.Unlock()

	user, err = ac.userRepo.GetUserByAddress(ctx, address) // 查询用户
	if nil != err || nil == user {
		return &pb.LandPlayTwoReply{
			Status: "不存在用户",
		}, nil
	}

	// 配置
	configs, err = ac.userRepo.GetConfigByKeys(ctx,
		"one_rate", "two_rate", "three_rate", "u_price", "low_reward_u",
	)
	if nil != err || nil == configs {
		return &pb.LandPlayTwoReply{
			Status: "配置错误",
		}, nil
	}

	for _, vConfig := range configs {
		if "one_rate" == vConfig.KeyName {
			oneRate, _ = strconv.ParseFloat(vConfig.Value, 10)
		}

		if "two_rate" == vConfig.KeyName {
			twoRate, _ = strconv.ParseFloat(vConfig.Value, 10)
		}

		if "three_rate" == vConfig.KeyName {
			threeRate, _ = strconv.ParseFloat(vConfig.Value, 10)
		}

		if "u_price" == vConfig.KeyName {
			uPrice, _ = strconv.ParseFloat(vConfig.Value, 10)
		}

		if "low_reward_u" == vConfig.KeyName {
			lowRewardU, _ = strconv.ParseFloat(vConfig.Value, 10)
		}
	}

	if 0 >= uPrice {
		return &pb.LandPlayTwoReply{
			Status: "币价biw:u错误",
		}, nil
	}

	var (
		landUserUse *LandUserUse
	)
	landUserUse, err = ac.userRepo.GetLandUserUseByID(ctx, req.SendBody.LandUseId)
	if nil != err || nil == landUserUse {
		return &pb.LandPlayTwoReply{
			Status: "不存在信息",
		}, nil
	}

	if landUserUse.UserId != user.ID {
		return &pb.LandPlayTwoReply{
			Status: "非种植用户",
		}, nil
	}

	if 1 != landUserUse.Status {
		return &pb.LandPlayTwoReply{
			Status: "状态错误",
		}, nil
	}

	current := time.Now().Unix()
	if uint64(current) < landUserUse.OverTime {
		return &pb.LandPlayTwoReply{
			Status: "种植未结束",
		}, nil
	}

	if 0 != landUserUse.One {
		return &pb.LandPlayTwoReply{
			Status: "停止生长状态",
		}, nil
	}

	// 已结束
	reward := landUserUse.OutMaxNum
	now := time.Now().Unix()
	// 有虫子 todo 杀虫和浇水更新数量和结束时间 偷的时候注意梳理
	if 0 != landUserUse.Two {
		if uint64(now) > landUserUse.Two {
			tmp := landUserUse.OutMaxNum * 0.01 * float64(uint64(now)-landUserUse.Two) / 300

			if tmp >= landUserUse.OutMaxNum {
				reward = 0
			} else {
				reward = landUserUse.OutMaxNum - tmp
			}
		}
	}

	var (
		land *Land
	)
	land, err = ac.userRepo.GetLandByIDTwo(ctx, landUserUse.LandId)
	if nil != err || nil == land {
		return &pb.LandPlayTwoReply{
			Status: "土地信息错误",
		}, nil
	}

	// 租的
	rentReward := float64(0)
	if landUserUse.UserId != landUserUse.OwnerUserId {
		if 0 < reward {
			rentReward = reward * land.RentOutPutRate
			if reward > rentReward {
				reward = reward - rentReward
			}
		}
	}

	// 推荐
	var (
		userRecommend *UserRecommend
	)
	tmpRecommendUserIds := make([]string, 0)
	userRecommend, err = ac.userRepo.GetUserRecommendByUserId(ctx, landUserUse.UserId)
	if nil == userRecommend || nil != err {
		return &pb.LandPlayTwoReply{
			Status: "查询推荐错误",
		}, nil
	}
	if "" != userRecommend.RecommendCode {
		tmpRecommendUserIds = strings.Split(userRecommend.RecommendCode, "D")
	}

	// 收租推荐
	tmpRecommendUserIdsRent := make([]string, 0)
	tmpRent := false
	rentUserId := uint64(0)
	if landUserUse.UserId != landUserUse.OwnerUserId {
		tmpRent = true
		rentUserId = landUserUse.OwnerUserId
		var (
			userRecommendRent *UserRecommend
		)
		userRecommendRent, err = ac.userRepo.GetUserRecommendByUserId(ctx, landUserUse.OwnerUserId)
		if nil == userRecommendRent || nil != err {
			return &pb.LandPlayTwoReply{
				Status: "查询推荐错误",
			}, nil
		}
		if "" != userRecommendRent.RecommendCode {
			tmpRecommendUserIdsRent = strings.Split(userRecommendRent.RecommendCode, "D")
		}
	}

	userIds := make([]uint64, 0)
	tmpIi := 0
	for i := len(tmpRecommendUserIds) - 1; i >= 0; i-- {
		if 3 <= tmpIi {
			break
		}
		tmpIi++

		tmpUserId, _ := strconv.ParseUint(tmpRecommendUserIds[i], 10, 64) // 最后一位是直推人
		if 0 >= tmpUserId {
			continue
		}

		userIds = append(userIds, tmpUserId)
	}

	tmpIj := 0
	for i := len(tmpRecommendUserIdsRent) - 1; i >= 0; i-- {
		if 3 <= tmpIj {
			break
		}
		tmpIj++

		tmpUserId, _ := strconv.ParseUint(tmpRecommendUserIdsRent[i], 10, 64) // 最后一位是直推人
		if 0 >= tmpUserId {
			continue
		}

		userIds = append(userIds, tmpUserId)
	}

	usersMap := make(map[uint64]*User, 0)
	if 0 < len(userIds) {
		usersMap, err = ac.userRepo.GetUserByUserIds(ctx, userIds)
		if nil != err {
			return &pb.LandPlayTwoReply{
				Status: "查询推荐错误",
			}, nil
		}
	}

	// 分红，状态变更
	if err = ac.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
		// 资源释放
		err = ac.userRepo.PlantPlatTwo(ctx, landUserUse.ID, land.ID, tmpRent)
		if nil != err {
			return err
		}

		// 奖励
		err = ac.userRepo.PlantPlatTwoTwo(ctx, landUserUse.ID, landUserUse.UserId, rentUserId, reward, rentReward)
		if nil != err {
			return err
		}

		err = ac.userRepo.CreateNotice(
			ctx,
			user.ID,
			"您收获了"+fmt.Sprintf("%.2f", reward)+"GIW",
			"You've harvest "+fmt.Sprintf("%.2f", reward)+" GIW",
		)
		if nil != err {
			return err
		}

		if 0 < rentReward {
			err = ac.userRepo.CreateNotice(
				ctx,
				rentUserId,
				"您收获了"+fmt.Sprintf("%.2f", rentReward)+"GIW",
				"You've harvest "+fmt.Sprintf("%.2f", rentReward)+" GIW",
			)
			if nil != err {
				return err
			}
		}

		// l1-l3，奖励发放
		if reward > 0 {
			tmpI := 0
			for i := len(tmpRecommendUserIds) - 1; i >= 0; i-- {
				if 3 <= tmpI {
					break
				}
				tmpI++

				tmpUserId, _ := strconv.ParseUint(tmpRecommendUserIds[i], 10, 64) // 最后一位是直推人
				if 0 >= tmpUserId {
					continue
				}

				if _, ok := usersMap[tmpUserId]; !ok {
					continue
				}

				if lowRewardU > usersMap[tmpUserId].Giw/uPrice {
					continue
				}

				tmpReward := float64(0)

				tmpNum := uint64(4)
				tmpReward = reward * oneRate
				if 1 == tmpI {

				} else if 2 == tmpI {
					tmpReward = reward * twoRate
					tmpNum = 7
				} else if 3 == tmpI {
					tmpReward = reward * threeRate
					tmpNum = 10
				} else {
					break
				}

				// 奖励
				err = ac.userRepo.PlantPlatTwoTwoL(ctx, landUserUse.ID, tmpUserId, landUserUse.UserId, tmpNum, tmpReward)
				if nil != err {
					return err
				}

				err = ac.userRepo.CreateNotice(
					ctx,
					tmpUserId,
					"您收获了"+fmt.Sprintf("%.2f", tmpReward)+"GIW",
					"You've harvest "+fmt.Sprintf("%.2f", tmpReward)+" GIW",
				)
				if nil != err {
					return err
				}
			}
		}

		// l1-l3，奖励发放
		if rentReward > 0 {
			tmpI := 0
			for i := len(tmpRecommendUserIdsRent) - 1; i >= 0; i-- {
				if 3 <= tmpI {
					break
				}
				tmpI++

				tmpUserId, _ := strconv.ParseUint(tmpRecommendUserIdsRent[i], 10, 64) // 最后一位是直推人
				if 0 >= tmpUserId {
					continue
				}

				if _, ok := usersMap[tmpUserId]; !ok {
					continue
				}

				if lowRewardU > usersMap[tmpUserId].Giw/uPrice {
					continue
				}

				tmpReward := float64(0)

				tmpNum := uint64(5)
				tmpReward = rentReward * oneRate
				if 1 == tmpI {

				} else if 2 == tmpI {
					tmpReward = rentReward * twoRate
					tmpNum = 8
				} else if 3 == tmpI {
					tmpReward = rentReward * threeRate
					tmpNum = 11
				} else {
					break
				}

				// 奖励
				err = ac.userRepo.PlantPlatTwoTwoL(ctx, landUserUse.ID, tmpUserId, landUserUse.OwnerUserId, tmpNum, tmpReward)
				if nil != err {
					return err
				}

				err = ac.userRepo.CreateNotice(
					ctx,
					tmpUserId,
					"您收获了"+fmt.Sprintf("%.2f", tmpReward)+"GIW",
					"You've harvest "+fmt.Sprintf("%.2f", tmpReward)+" GIW",
				)
				if nil != err {
					return err
				}
			}
		}

		return nil
	}); nil != err {
		fmt.Println(err, "LandPlayTwo", landUserUse)
		return &pb.LandPlayTwoReply{
			Status: "种植失败",
		}, nil
	}

	return &pb.LandPlayTwoReply{
		Status: "ok",
	}, nil
}

// LandPlayThree 施肥
func (ac *AppUsecase) LandPlayThree(ctx context.Context, address string, req *pb.LandPlayThreeRequest) (*pb.LandPlayThreeReply, error) {
	var (
		user *User
		err  error
	)

	user, err = ac.userRepo.GetUserByAddress(ctx, address) // 查询用户
	if nil != err || nil == user {
		return &pb.LandPlayThreeReply{
			Status: "不存在用户",
		}, nil
	}

	if 1 == user.LockUse {
		return &pb.LandPlayThreeReply{
			Status: "锁定用户",
		}, nil
	}

	var (
		prop *Prop
	)
	prop, err = ac.userRepo.GetPropByID(ctx, req.SendBody.Id, 1)
	if nil != err || nil == prop {
		return &pb.LandPlayThreeReply{
			Status: "不存道具",
		}, nil
	}

	if user.ID != prop.UserId {
		return &pb.LandPlayThreeReply{
			Status: "不是自己的",
		}, nil
	}

	if 11 != prop.PropType {
		return &pb.LandPlayThreeReply{
			Status: "不是化肥",
		}, nil
	}

	var (
		landUserUse *LandUserUse
	)
	landUserUse, err = ac.userRepo.GetLandUserUseByID(ctx, req.SendBody.LandUseId)
	if nil != err || nil == landUserUse {
		return &pb.LandPlayThreeReply{
			Status: "不存在信息",
		}, nil
	}

	if landUserUse.UserId != user.ID {
		return &pb.LandPlayThreeReply{
			Status: "非种植用户",
		}, nil
	}

	if 1 != landUserUse.Status {
		return &pb.LandPlayThreeReply{
			Status: "状态错误",
		}, nil
	}

	current := time.Now().Unix()
	if 0 != landUserUse.One && uint64(current) >= landUserUse.One {
		return &pb.LandPlayThreeReply{
			Status: "停止生长状态",
		}, nil
	}

	if 0 != landUserUse.Two && uint64(current) >= landUserUse.Two {
		return &pb.LandPlayThreeReply{
			Status: "生虫状态",
		}, nil
	}

	if uint64(current) >= landUserUse.OverTime {
		return &pb.LandPlayThreeReply{
			Status: "种植已经结束了",
		}, nil
	}

	if landUserUse.OverTime < uint64(prop.OneTwo) {
		return &pb.LandPlayThreeReply{
			Status: "道具配置错误",
		}, nil
	}

	overTime := landUserUse.OverTime - uint64(prop.OneTwo)
	one := false
	if overTime <= landUserUse.One {
		one = true
	}

	two := false
	if overTime <= landUserUse.Two {
		two = true
	}

	if err = ac.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
		return ac.userRepo.PlantPlatThree(ctx, landUserUse.ID, overTime, prop.ID, one, two)
	}); nil != err {
		fmt.Println(err, "LandPlayThree", user)
		return &pb.LandPlayThreeReply{
			Status: "施肥失败",
		}, nil
	}

	return &pb.LandPlayThreeReply{
		Status: "ok",
	}, nil
}

// LandPlayFour 杀虫
func (ac *AppUsecase) LandPlayFour(ctx context.Context, address string, req *pb.LandPlayFourRequest) (*pb.LandPlayFourReply, error) {
	var (
		user *User
		err  error
	)

	user, err = ac.userRepo.GetUserByAddress(ctx, address) // 查询用户
	if nil != err || nil == user {
		return &pb.LandPlayFourReply{
			Status: "不存在用户",
		}, nil
	}

	if 1 == user.LockUse {
		return &pb.LandPlayFourReply{
			Status: "锁定用户",
		}, nil
	}

	var (
		prop *Prop
	)

	// 11化肥，12水，13手套，14除虫剂，15铲子，16盲盒，17地契
	prop, err = ac.userRepo.GetPropByIDTwo(ctx, req.SendBody.Id)
	if nil != err || nil == prop {
		return &pb.LandPlayFourReply{
			Status: "不存在道具",
		}, nil
	}

	if 14 != prop.PropType {
		return &pb.LandPlayFourReply{
			Status: "无效道具",
		}, nil

	}

	if 2 < prop.Status {
		return &pb.LandPlayFourReply{
			Status: "无效道具",
		}, nil
	}

	if 0 >= prop.FourOne {
		return &pb.LandPlayFourReply{
			Status: "无效道具",
		}, nil
	}

	if user.ID != prop.UserId {
		return &pb.LandPlayFourReply{
			Status: "不是自己的",
		}, nil
	}

	var (
		landUserUse *LandUserUse
	)
	landUserUse, err = ac.userRepo.GetLandUserUseByID(ctx, req.SendBody.LandUseId)
	if nil != err || nil == landUserUse {
		return &pb.LandPlayFourReply{
			Status: "不存在信息",
		}, nil
	}

	if landUserUse.UserId != user.ID {
		return &pb.LandPlayFourReply{
			Status: "非种植用户",
		}, nil
	}

	if 1 != landUserUse.Status {
		return &pb.LandPlayFourReply{
			Status: "状态错误",
		}, nil
	}

	current := time.Now().Unix()
	if 0 >= landUserUse.Two {
		return &pb.LandPlayFourReply{
			Status: "无需杀虫",
		}, nil
	}

	if uint64(current) < landUserUse.Two {
		return &pb.LandPlayFourReply{
			Status: "无需杀虫",
		}, nil
	}

	// 剩余最大产出
	rewardTmp := float64(0)
	if uint64(current) > landUserUse.Two {
		tmp := landUserUse.OutMaxNum * 0.01 * float64(uint64(uint64(current)-landUserUse.Two)/300)
		if tmp < landUserUse.OutMaxNum {
			rewardTmp = landUserUse.OutMaxNum - tmp
		}
	}

	one := uint64(0)
	if 1 <= prop.FourOne {
		one = uint64(prop.FourOne - 1)
	}

	two := uint64(2)
	if 0 >= one {
		two = 3
	}

	if err = ac.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
		return ac.userRepo.PlantPlatFour(ctx, rewardTmp, landUserUse.ID, prop.ID, two, one)
	}); nil != err {
		fmt.Println(err, "LandPlayFour", user)
		return &pb.LandPlayFourReply{
			Status: "杀虫失败",
		}, nil
	}

	return &pb.LandPlayFourReply{
		Status: "ok",
	}, nil
}

// LandPlayFive 浇水
func (ac *AppUsecase) LandPlayFive(ctx context.Context, address string, req *pb.LandPlayFiveRequest) (*pb.LandPlayFiveReply, error) {
	var (
		user *User
		err  error
	)

	user, err = ac.userRepo.GetUserByAddress(ctx, address) // 查询用户
	if nil != err || nil == user {
		return &pb.LandPlayFiveReply{
			Status: "不存在用户",
		}, nil
	}

	if 1 == user.LockUse {
		return &pb.LandPlayFiveReply{
			Status: "锁定用户",
		}, nil
	}

	var (
		prop *Prop
	)

	// 11化肥，12水，13手套，14除虫剂，15铲子，16盲盒，17地契
	prop, err = ac.userRepo.GetPropByIDTwo(ctx, req.SendBody.Id)
	if nil != err || nil == prop {
		return &pb.LandPlayFiveReply{
			Status: "不存在道具",
		}, nil
	}

	if 12 != prop.PropType {
		return &pb.LandPlayFiveReply{
			Status: "无效道具",
		}, nil

	}

	if 2 < prop.Status {
		return &pb.LandPlayFiveReply{
			Status: "无效道具",
		}, nil
	}

	if 0 >= prop.ThreeOne {
		return &pb.LandPlayFiveReply{
			Status: "无效道具",
		}, nil
	}

	if user.ID != prop.UserId {
		return &pb.LandPlayFiveReply{
			Status: "不是自己的",
		}, nil
	}

	var (
		landUserUse *LandUserUse
	)
	landUserUse, err = ac.userRepo.GetLandUserUseByID(ctx, req.SendBody.LandUseId)
	if nil != err || nil == landUserUse {
		return &pb.LandPlayFiveReply{
			Status: "不存在信息",
		}, nil
	}

	if landUserUse.UserId != user.ID {
		return &pb.LandPlayFiveReply{
			Status: "非种植用户",
		}, nil
	}

	if 1 != landUserUse.Status {
		return &pb.LandPlayFiveReply{
			Status: "状态错误",
		}, nil
	}

	current := time.Now().Unix()
	if 0 >= landUserUse.One {
		return &pb.LandPlayFiveReply{
			Status: "无需浇水",
		}, nil
	}

	if uint64(current) < landUserUse.One {
		return &pb.LandPlayFiveReply{
			Status: "无需浇水",
		}, nil
	}

	tmpOverTime := landUserUse.OverTime + uint64(current) - landUserUse.One

	one := uint64(0)
	if 1 <= prop.ThreeOne {
		one = uint64(prop.ThreeOne - 1)
	}

	two := uint64(2)
	if 0 >= one {
		two = 3
	}

	if err = ac.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
		return ac.userRepo.PlantPlatFive(ctx, tmpOverTime, landUserUse.ID, prop.ID, two, one)
	}); nil != err {
		fmt.Println(err, "LandPlayFive", user)
		return &pb.LandPlayFiveReply{
			Status: "浇水失败",
		}, nil
	}

	return &pb.LandPlayFiveReply{
		Status: "ok",
	}, nil
}

// LandPlaySix 铲子
func (ac *AppUsecase) LandPlaySix(ctx context.Context, address string, req *pb.LandPlaySixRequest) (*pb.LandPlaySixReply, error) {
	var (
		user *User
		err  error
	)

	user, err = ac.userRepo.GetUserByAddress(ctx, address) // 查询用户
	if nil != err || nil == user {
		return &pb.LandPlaySixReply{
			Status: "不存在用户",
		}, nil
	}

	if 1 == user.LockUse {
		return &pb.LandPlaySixReply{
			Status: "锁定用户",
		}, nil
	}

	var (
		prop *Prop
	)

	// 11化肥，12水，13手套，14除虫剂，15铲子，16盲盒，17地契
	prop, err = ac.userRepo.GetPropByIDTwo(ctx, req.SendBody.Id)
	if nil != err || nil == prop {
		return &pb.LandPlaySixReply{
			Status: "不存在道具",
		}, nil
	}

	if 15 != prop.PropType {
		return &pb.LandPlaySixReply{
			Status: "无效道具",
		}, nil

	}

	if 2 < prop.Status {
		return &pb.LandPlaySixReply{
			Status: "无效道具",
		}, nil
	}

	if 0 >= prop.TwoOne {
		return &pb.LandPlaySixReply{
			Status: "无效道具",
		}, nil
	}

	if user.ID != prop.UserId {
		return &pb.LandPlaySixReply{
			Status: "不是自己的",
		}, nil
	}

	var (
		landUserUse *LandUserUse
	)
	landUserUse, err = ac.userRepo.GetLandUserUseByID(ctx, req.SendBody.LandUseId)
	if nil != err || nil == landUserUse {
		return &pb.LandPlaySixReply{
			Status: "不存在信息",
		}, nil
	}

	if landUserUse.OwnerUserId != user.ID {
		return &pb.LandPlaySixReply{
			Status: "非土地用户",
		}, nil
	}

	if landUserUse.UserId == user.ID {
		return &pb.LandPlaySixReply{
			Status: "非出租土地",
		}, nil
	}

	if 1 != landUserUse.Status {
		return &pb.LandPlaySixReply{
			Status: "状态错误",
		}, nil
	}

	current := time.Now().Unix()
	// todo
	if uint64(current) < landUserUse.OverTime+3600 {
		return &pb.LandPlaySixReply{
			Status: "成熟1小时后可以铲除",
		}, nil
	}

	one := uint64(0)
	if 1 <= prop.TwoOne {
		one = uint64(prop.TwoOne - 1)
	}

	two := uint64(2)
	if 0 >= one {
		two = 3
	}

	tmpOverMax := float64(0)
	if landUserUse.OutMaxNum > landUserUse.OutMaxNum*prop.TwoTwo {
		tmpOverMax = landUserUse.OutMaxNum - landUserUse.OutMaxNum*prop.TwoTwo
	}
	tmpOverMaxTwo := landUserUse.OutMaxNum * prop.TwoTwo

	if 0 < landUserUse.One {
		tmpOverMax = 0
		tmpOverMaxTwo = 0
	} else if 0 < landUserUse.Two {
		// 剩余最大产出
		rewardTmp := float64(0)
		if uint64(current) > landUserUse.Two {
			tmp := landUserUse.OutMaxNum * 0.01 * float64(uint64(uint64(current)-landUserUse.Two)/300)
			if tmp < landUserUse.OutMaxNum {
				rewardTmp = landUserUse.OutMaxNum - tmp
			}
		}

		if 0 >= rewardTmp {
			tmpOverMax = 0
			tmpOverMaxTwo = 0
		} else {
			if rewardTmp > rewardTmp*prop.TwoTwo {
				tmpOverMax = rewardTmp - rewardTmp*prop.TwoTwo
			}
			tmpOverMaxTwo = rewardTmp * prop.TwoTwo
		}
	}

	// 推荐
	var (
		userRecommend *UserRecommend
	)
	tmpRecommendUserIds := make([]string, 0)
	userRecommend, err = ac.userRepo.GetUserRecommendByUserId(ctx, landUserUse.UserId)
	if nil == userRecommend || nil != err {
		return &pb.LandPlaySixReply{
			Status: "查询推荐错误",
		}, nil
	}
	if "" != userRecommend.RecommendCode {
		tmpRecommendUserIds = strings.Split(userRecommend.RecommendCode, "D")
	}

	// 收租推荐
	tmpRecommendUserIdsRent := make([]string, 0)
	if landUserUse.UserId != landUserUse.OwnerUserId {
		var (
			userRecommendRent *UserRecommend
		)
		userRecommendRent, err = ac.userRepo.GetUserRecommendByUserId(ctx, landUserUse.OwnerUserId)
		if nil == userRecommendRent || nil != err {
			return &pb.LandPlaySixReply{
				Status: "查询推荐错误",
			}, nil
		}
		if "" != userRecommendRent.RecommendCode {
			tmpRecommendUserIdsRent = strings.Split(userRecommendRent.RecommendCode, "D")
		}
	}

	userIds := make([]uint64, 0)
	tmpIi := 0
	for i := len(tmpRecommendUserIds) - 1; i >= 0; i-- {
		if 3 <= tmpIi {
			break
		}
		tmpIi++

		tmpUserId, _ := strconv.ParseUint(tmpRecommendUserIds[i], 10, 64) // 最后一位是直推人
		if 0 >= tmpUserId {
			continue
		}

		userIds = append(userIds, tmpUserId)
	}

	tmpIj := 0
	for i := len(tmpRecommendUserIdsRent) - 1; i >= 0; i-- {
		if 3 <= tmpIj {
			break
		}
		tmpIj++

		tmpUserId, _ := strconv.ParseUint(tmpRecommendUserIdsRent[i], 10, 64) // 最后一位是直推人
		if 0 >= tmpUserId {
			continue
		}

		userIds = append(userIds, tmpUserId)
	}

	usersMap := make(map[uint64]*User, 0)
	if 0 < len(userIds) {
		usersMap, err = ac.userRepo.GetUserByUserIds(ctx, userIds)
		if nil != err {
			return &pb.LandPlaySixReply{
				Status: "查询推荐错误",
			}, nil
		}
	}

	// 配置
	var (
		configs    []*Config
		oneRate    float64
		twoRate    float64
		threeRate  float64
		uPrice     float64
		lowRewardU float64
	)
	configs, err = ac.userRepo.GetConfigByKeys(ctx,
		"one_rate", "two_rate", "three_rate", "u_price", "low_reward_u",
	)
	if nil != err || nil == configs {
		return &pb.LandPlaySixReply{
			Status: "配置错误",
		}, nil
	}

	for _, vConfig := range configs {
		if "one_rate" == vConfig.KeyName {
			oneRate, _ = strconv.ParseFloat(vConfig.Value, 10)
		}

		if "two_rate" == vConfig.KeyName {
			twoRate, _ = strconv.ParseFloat(vConfig.Value, 10)
		}

		if "three_rate" == vConfig.KeyName {
			threeRate, _ = strconv.ParseFloat(vConfig.Value, 10)
		}

		if "u_price" == vConfig.KeyName {
			uPrice, _ = strconv.ParseFloat(vConfig.Value, 10)
		}

		if "low_reward_u" == vConfig.KeyName {
			lowRewardU, _ = strconv.ParseFloat(vConfig.Value, 10)
		}
	}

	if 0 >= uPrice {
		return &pb.LandPlaySixReply{
			Status: "币价biw:u错误",
		}, nil
	}

	if err = ac.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
		err = ac.userRepo.PlantPlatSix(ctx, landUserUse.ID, prop.ID, two, one, landUserUse.LandId)
		if nil != err {
			return err
		}

		// 奖励
		err = ac.userRepo.PlantPlatTwoTwo(ctx, landUserUse.ID, landUserUse.UserId, landUserUse.OwnerUserId, tmpOverMax, tmpOverMaxTwo)
		if nil != err {
			return err
		}

		err = ac.userRepo.CreateNotice(
			ctx,
			landUserUse.UserId,
			"您收获了"+fmt.Sprintf("%.2f", tmpOverMax)+"GIW",
			"You've harvest "+fmt.Sprintf("%.2f", tmpOverMax)+" GIW",
		)
		if nil != err {
			return err
		}

		err = ac.userRepo.CreateNotice(
			ctx,
			landUserUse.OwnerUserId,
			"您收获了"+fmt.Sprintf("%.2f", tmpOverMaxTwo)+"GIW",
			"You've harvest "+fmt.Sprintf("%.2f", tmpOverMaxTwo)+" GIW",
		)

		// l1-l3，奖励发放
		if tmpOverMax > 0 {
			tmpI := 0
			for i := len(tmpRecommendUserIds) - 1; i >= 0; i-- {
				if 3 <= tmpI {
					break
				}
				tmpI++

				tmpUserId, _ := strconv.ParseUint(tmpRecommendUserIds[i], 10, 64) // 最后一位是直推人
				if 0 >= tmpUserId {
					continue
				}

				if _, ok := usersMap[tmpUserId]; !ok {
					continue
				}

				if lowRewardU > usersMap[tmpUserId].Giw/uPrice {
					continue
				}

				tmpReward := float64(0)

				tmpNum := uint64(4)
				tmpReward = tmpOverMax * oneRate
				if 1 == tmpI {

				} else if 2 == tmpI {
					tmpReward = tmpOverMax * twoRate
					tmpNum = 7
				} else if 3 == tmpI {
					tmpReward = tmpOverMax * threeRate
					tmpNum = 10
				} else {
					break
				}

				// 奖励
				err = ac.userRepo.PlantPlatTwoTwoL(ctx, landUserUse.ID, tmpUserId, landUserUse.UserId, tmpNum, tmpReward)
				if nil != err {
					return err
				}

				err = ac.userRepo.CreateNotice(
					ctx,
					tmpUserId,
					"您收获了"+fmt.Sprintf("%.2f", tmpReward)+"GIW",
					"You've harvest "+fmt.Sprintf("%.2f", tmpReward)+" GIW",
				)
				if nil != err {
					return err
				}
			}
		}

		// l1-l3，奖励发放
		if tmpOverMaxTwo > 0 {
			tmpI := 0
			for i := len(tmpRecommendUserIdsRent) - 1; i >= 0; i-- {
				if 3 <= tmpI {
					break
				}
				tmpI++

				tmpUserId, _ := strconv.ParseUint(tmpRecommendUserIdsRent[i], 10, 64) // 最后一位是直推人
				if 0 >= tmpUserId {
					continue
				}

				if _, ok := usersMap[tmpUserId]; !ok {
					continue
				}

				if lowRewardU > usersMap[tmpUserId].Giw/uPrice {
					continue
				}

				tmpReward := float64(0)

				tmpNum := uint64(5)
				tmpReward = tmpOverMaxTwo * oneRate
				if 1 == tmpI {

				} else if 2 == tmpI {
					tmpReward = tmpOverMaxTwo * twoRate
					tmpNum = 8
				} else if 3 == tmpI {
					tmpReward = tmpOverMaxTwo * threeRate
					tmpNum = 11
				} else {
					break
				}

				// 奖励
				err = ac.userRepo.PlantPlatTwoTwoL(ctx, landUserUse.ID, tmpUserId, landUserUse.OwnerUserId, tmpNum, tmpReward)
				if nil != err {
					return err
				}

				err = ac.userRepo.CreateNotice(
					ctx,
					tmpUserId,
					"您收获了"+fmt.Sprintf("%.2f", tmpReward)+"GIW",
					"You've harvest "+fmt.Sprintf("%.2f", tmpReward)+" GIW",
				)
				if nil != err {
					return err
				}
			}
		}

		return nil
	}); nil != err {
		fmt.Println(err, "LandPlaySix", user)
		return &pb.LandPlaySixReply{
			Status: "铲除土地失败",
		}, nil
	}

	return &pb.LandPlaySixReply{
		Status: "ok",
	}, nil
}

// LandPlaySeven 手套
func (ac *AppUsecase) LandPlaySeven(ctx context.Context, address string, req *pb.LandPlaySevenRequest) (*pb.LandPlaySevenReply, error) {
	var (
		user *User
		err  error
	)

	user, err = ac.userRepo.GetUserByAddress(ctx, address) // 查询用户
	if nil != err || nil == user {
		return &pb.LandPlaySevenReply{
			Status: "不存在用户",
		}, nil
	}

	if 1 == user.LockUse {
		return &pb.LandPlaySevenReply{
			Status: "锁定用户",
		}, nil
	}

	var (
		prop *Prop
	)

	// 11化肥，12水，13手套，14除虫剂，15铲子，16盲盒，17地契
	prop, err = ac.userRepo.GetPropByIDTwo(ctx, req.SendBody.Id)
	if nil != err || nil == prop {
		return &pb.LandPlaySevenReply{
			Status: "不存在道具",
		}, nil
	}

	if 13 != prop.PropType {
		return &pb.LandPlaySevenReply{
			Status: "无效道具",
		}, nil

	}

	if 2 < prop.Status {
		return &pb.LandPlaySevenReply{
			Status: "无效道具",
		}, nil
	}

	if 0 >= prop.FiveOne {
		return &pb.LandPlaySevenReply{
			Status: "无效道具",
		}, nil
	}

	if user.ID != prop.UserId {
		return &pb.LandPlaySevenReply{
			Status: "不是自己的",
		}, nil
	}

	var (
		landUserUse *LandUserUse
	)
	landUserUse, err = ac.userRepo.GetLandUserUseByID(ctx, req.SendBody.LandUseId)
	if nil != err || nil == landUserUse {
		return &pb.LandPlaySevenReply{
			Status: "不存在信息",
		}, nil
	}

	if landUserUse.OwnerUserId == user.ID {
		return &pb.LandPlaySevenReply{
			Status: "土地用户不能使用手套",
		}, nil
	}

	if landUserUse.UserId == user.ID {
		return &pb.LandPlaySevenReply{
			Status: "种植用户不能使用手套",
		}, nil
	}

	if 1 != landUserUse.Status {
		return &pb.LandPlaySevenReply{
			Status: "状态错误",
		}, nil
	}

	current := time.Now().Unix()
	if uint64(current) < landUserUse.OverTime {
		return &pb.LandPlaySevenReply{
			Status: "还未成熟",
		}, nil
	}

	if 0 < landUserUse.One {
		return &pb.LandPlaySevenReply{
			Status: "缺水暂停中",
		}, nil
	}

	if 0 < landUserUse.Two {
		return &pb.LandPlaySevenReply{
			Status: "虫蛀减产中",
		}, nil
	}

	lastTime := landUserUse.SubTime
	if 0 < lastTime {
		return &pb.LandPlaySevenReply{
			Status: "偷盗过于频繁",
		}, nil
		//if uint64(current)-600 <= lastTime {
		//	return &pb.LandPlaySevenReply{
		//		Status: "偷盗过于频繁",
		//	}, nil
		//}
	}

	tmpAmount := landUserUse.OutMaxNum * 0.1
	tmpOutMax := float64(0)
	if tmpAmount >= landUserUse.OutMaxNum {
		tmpOutMax = 0
	} else {
		tmpOutMax = landUserUse.OutMaxNum - tmpAmount
	}

	one := uint64(0)
	if 1 <= prop.FiveOne {
		one = uint64(prop.FiveOne - 1)
	}

	two := uint64(2)
	if 0 >= one {
		two = 3
	}

	if err = ac.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
		err = ac.userRepo.PlantPlatSeven(ctx, tmpOutMax, tmpAmount, uint64(current), lastTime, landUserUse.ID, prop.ID, two, one, user.ID)
		if nil != err {
			return err
		}

		err = ac.userRepo.CreateNotice(
			ctx,
			user.ID,
			"您使用手套偷取了"+fmt.Sprintf("%.2f", tmpAmount)+"GIW",
			"You've steal "+fmt.Sprintf("%.2f", tmpAmount)+" GIW use gloves",
		)
		if nil != err {
			return err
		}

		return nil
	}); nil != err {
		fmt.Println(err, "LandPlaySeven", user)
		return &pb.LandPlaySevenReply{
			Status: "偷取失败",
		}, nil
	}

	return &pb.LandPlaySevenReply{
		Status: "ok",
		Amount: tmpAmount,
	}, nil
}

var rngMutexBuy sync.Mutex

func (ac *AppUsecase) Buy(ctx context.Context, address string, req *pb.BuyRequest) (*pb.BuyReply, error) {
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

	if 1 == user.LockUse {
		return &pb.BuyReply{
			Status: "锁定用户",
		}, nil
	}

	rngMutexBuy.Lock()
	defer rngMutexBuy.Unlock()

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
			err = ac.userRepo.BuySeed(ctx, seed.SellAmount, tmpGet, seed.UserId, user.ID, seed.ID)
			if nil != err {
				return err
			}

			err = ac.userRepo.CreateNotice(
				ctx,
				user.ID,
				"您购买了种子，花费"+fmt.Sprintf("%.2f", seed.SellAmount)+"GIW",
				"You've pay "+fmt.Sprintf("%.2f", seed.SellAmount)+" GIW for seed",
			)
			if nil != err {
				return err
			}

			err = ac.userRepo.CreateNotice(
				ctx,
				seed.UserId,
				"您出售了种子，获得"+fmt.Sprintf("%.2f", tmpGet)+"GIW",
				"You've get "+fmt.Sprintf("%.2f", tmpGet)+" GIW for seed",
			)
			if nil != err {
				return err
			}

			return nil
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
			err = ac.userRepo.BuyProp(ctx, prop.SellAmount, tmpGet, prop.UserId, user.ID, prop.ID)
			if nil != err {
				return err
			}
			err = ac.userRepo.CreateNotice(
				ctx,
				user.ID,
				"您购买了道具，花费"+fmt.Sprintf("%.2f", prop.SellAmount)+"GIW",
				"You've pay "+fmt.Sprintf("%.2f", prop.SellAmount)+" GIW for prop",
			)
			if nil != err {
				return err
			}

			err = ac.userRepo.CreateNotice(
				ctx,
				prop.UserId,
				"您出售了道具，获得"+fmt.Sprintf("%.2f", tmpGet)+"GIW",
				"You've get "+fmt.Sprintf("%.2f", tmpGet)+" GIW for prop",
			)
			if nil != err {
				return err
			}

			return nil
		}); nil != err {
			fmt.Println(err, "buyProp", user)
			return &pb.BuyReply{
				Status: "购买失败",
			}, nil
		}
	} else if 3 == req.SendBody.BuyType {
		return &pb.BuyReply{
			Status: "暂未开放",
		}, nil

		//var (
		//	land *Land
		//)
		//land, err = ac.userRepo.GetLandByID(ctx, req.SendBody.Id)
		//if nil != err || nil == land {
		//	return &pb.BuyReply{
		//		Status: "不存道具",
		//	}, nil
		//}
		//
		//if user.ID == land.UserId {
		//	return &pb.BuyReply{
		//		Status: "不允许购买自己的",
		//	}, nil
		//}
		//
		//if 4 != land.Status {
		//	return &pb.BuyReply{
		//		Status: "未出售",
		//	}, nil
		//}
		//
		//if 0 >= land.SellAmount {
		//	return &pb.BuyReply{
		//		Status: "金额错误",
		//	}, nil
		//}
		//
		//if user.Git < land.SellAmount {
		//	return &pb.BuyReply{
		//		Status: "余额不足",
		//	}, nil
		//}
		//
		//// 土地
		//tmpGet := land.SellAmount - land.SellAmount*feeRate
		//if 0 >= tmpGet {
		//	return &pb.BuyReply{
		//		Status: "金额错误",
		//	}, nil
		//}
		//
		//if err = ac.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
		//	err = ac.userRepo.BuyLand(ctx, land.SellAmount, tmpGet, land.UserId, user.ID, land.ID)
		//	if nil != err {
		//		return err
		//	}
		//
		//	err = ac.userRepo.CreateNotice(
		//		ctx,
		//		land.UserId,
		//		"您购买了土地，花费"+strconv.FormatFloat(land.SellAmount, 'f', -1, 64)+"GIT",
		//		"You've pay "+strconv.FormatFloat(land.SellAmount, 'f', -1, 64)+" GIT for land",
		//	)
		//	if nil != err {
		//		return err
		//	}
		//
		//	err = ac.userRepo.CreateNotice(
		//		ctx,
		//		user.ID,
		//		"您出售了土地，获得"+strconv.FormatFloat(tmpGet, 'f', -1, 64)+"GIT",
		//		"You've get "+strconv.FormatFloat(tmpGet, 'f', -1, 64)+" GIT for land",
		//	)
		//	if nil != err {
		//		return err
		//	}
		//
		//	return nil
		//}); nil != err {
		//	fmt.Println(err, "buyLand", user)
		//	return &pb.BuyReply{
		//		Status: "购买失败",
		//	}, nil
		//}
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

	if 1 == user.LockUse {
		return &pb.SellReply{
			Status: "锁定用户",
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
					Status: "不存在种子",
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
					Status: "不存在道具",
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
			return &pb.SellReply{
				Status: "暂未开放",
			}, nil

			//var (
			//	land *Land
			//)
			//land, err = ac.userRepo.GetLandByID(ctx, req.SendBody.Id)
			//if nil != err || nil == land {
			//	return &pb.SellReply{
			//		Status: "不存在土地",
			//	}, nil
			//}
			//
			//if user.ID != land.UserId {
			//	return &pb.SellReply{
			//		Status: "不是自己的",
			//	}, nil
			//}
			//
			//if 0 != land.LocationNum {
			//	return &pb.SellReply{
			//		Status: "土地布置中",
			//	}, nil
			//}
			//
			//if 0 != land.Status {
			//	return &pb.SellReply{
			//		Status: "不符合上架状态",
			//	}, nil
			//}
			//
			//if 1 != land.Two {
			//	return &pb.SellReply{
			//		Status: "不可出售土地",
			//	}, nil
			//}
			//
			//if err = ac.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
			//	return ac.userRepo.SellLand(ctx, land.ID, user.ID, tmpSellAmount)
			//}); nil != err {
			//	fmt.Println(err, "sellProp", user)
			//	return &pb.SellReply{
			//		Status: "上架失败",
			//	}, nil
			//}
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
					Status: "不存在种子",
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
				fmt.Println(err, "unSellSeed", user)
				return &pb.SellReply{
					Status: "下架失败",
				}, nil
			}
		} else if 2 == req.SendBody.SellType {
			var (
				prop *Prop
			)
			prop, err = ac.userRepo.GetPropByID(ctx, req.SendBody.Id, 4)
			if nil != err || nil == prop {
				return &pb.SellReply{
					Status: "不存在道具",
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
				fmt.Println(err, "unSellProp", user)
				return &pb.SellReply{
					Status: "下架失败",
				}, nil
			}
		} else if 3 == req.SendBody.SellType {
			return &pb.SellReply{
				Status: "暂未开放",
			}, nil

			//var (
			//	land *Land
			//)
			//land, err = ac.userRepo.GetLandByIDTwo(ctx, req.SendBody.Id)
			//if nil != err || nil == land {
			//	return &pb.SellReply{
			//		Status: "不存在土地",
			//	}, nil
			//}
			//
			//if user.ID != land.UserId {
			//	return &pb.SellReply{
			//		Status: "不是自己的",
			//	}, nil
			//}
			//
			//if 0 != land.LocationNum {
			//	return &pb.SellReply{
			//		Status: "土地布置中",
			//	}, nil
			//}
			//
			//if 4 != land.Status {
			//	return &pb.SellReply{
			//		Status: "不符合下架要求",
			//	}, nil
			//}
			//
			//if err = ac.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
			//	return ac.userRepo.UnSellLand(ctx, land.ID, user.ID)
			//}); nil != err {
			//	fmt.Println(err, "unSellLand", user)
			//	return &pb.SellReply{
			//		Status: "下架失败",
			//	}, nil
			//}
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

	if 1 == user.LockUse {
		return &pb.StakeGitReply{
			Status: "锁定用户",
		}, nil
	}

	if 1 == req.SendBody.Num {
		if 100 > req.SendBody.Amount {
			return &pb.StakeGitReply{
				Status: "git金额要多于100",
			}, nil
		}

		if req.SendBody.Amount > uint64(user.Git) {
			return &pb.StakeGitReply{
				Status: "git余额不足",
			}, nil
		}

		if err = ac.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
			err = ac.userRepo.SetStakeGit(ctx, user.ID, float64(req.SendBody.Amount))
			if nil != err {
				return err
			}

			err = ac.userRepo.CreateNotice(
				ctx,
				user.ID,
				"您向粮仓质押"+fmt.Sprintf("%.2f", float64(req.SendBody.Amount))+"GIW",
				"You've deposit "+fmt.Sprintf("%.2f", float64(req.SendBody.Amount))+" GIW to granary",
			)
			if nil != err {
				return err
			}
			return nil
		}); nil != err {
			return &pb.StakeGitReply{
				Status: "stakeGit失败",
			}, nil
		}
	} else if 2 == req.SendBody.Num {
		var (
			record *StakeGitRecord
		)
		record, err = ac.userRepo.GetStakeGitRecordsByID(ctx, req.SendBody.Id, user.ID) // 查询用户
		if nil != err || nil == record {
			return &pb.StakeGitReply{
				Status: "不存在记录",
			}, nil
		}

		if err = ac.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
			err = ac.userRepo.SetUnStakeGit(ctx, record.ID, user.ID, record.Amount)
			if nil != err {
				return err
			}

			err = ac.userRepo.CreateNotice(
				ctx,
				user.ID,
				"您从粮仓解押"+fmt.Sprintf("%.2f", record.Amount)+"GIW",
				"You've withdraw "+fmt.Sprintf("%.2f", record.Amount)+" GIW from granary",
			)
			if nil != err {
				return err
			}
			return nil
		}); nil != err {
			return &pb.StakeGitReply{
				Status: "stakeGit失败",
			}, nil
		}
	} else {
		return &pb.StakeGitReply{
			Status: "错误参数",
		}, nil
	}

	return &pb.StakeGitReply{Status: "ok"}, nil
}

var rentLock sync.Mutex

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

	if 1 == user.LockUse {
		return &pb.RentLandReply{
			Status: "锁定用户",
		}, nil
	}

	rentLock.Lock()
	defer rentLock.Unlock()

	user, err = ac.userRepo.GetUserByAddress(ctx, address) // 查询用户
	if nil != err || nil == user {
		return &pb.RentLandReply{
			Status: "不存在用户",
		}, nil
	}

	rentRate := float64(0)
	//if 1 == req.SendBody.Rate {
	//	rentRate = 0.05
	//} else if 2 == req.SendBody.Rate {
	//	rentRate = 0.1
	//} else if 3 == req.SendBody.Rate {
	//	rentRate = 0.2
	//} else if 4 == req.SendBody.Rate {
	//	rentRate = 0.3
	//} else if 5 == req.SendBody.Rate {
	//	rentRate = 0.4
	//} else if 6 == req.SendBody.Rate {
	//	rentRate = 0.5
	//} else {
	//	return &pb.RentLandReply{
	//		Status: "比例错误",
	//	}, nil
	//}

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

		if 1 != land.One {
			return &pb.RentLandReply{
				Status: "不允许出租类型",
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
		land, err = ac.userRepo.GetLandByIDTwo(ctx, req.SendBody.LandId)
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
				Status: "下架失败",
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

	if 1 == user.LockUse {
		return &pb.LandPlayReply{
			Status: "锁定用户",
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

		if 1 > land.LocationNum || 9 < land.LocationNum {
			return &pb.LandPlayReply{
				Status: "位置参数错误",
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
		if 1 > req.SendBody.LocationNum || 9 < req.SendBody.LocationNum {
			return &pb.LandPlayReply{
				Status: "位置参数错误",
			}, nil
		}

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

func (ac *AppUsecase) LandAddOutRate(ctx context.Context, address string, req *pb.LandAddOutRateRequest) (*pb.LandAddOutRateReply, error) {
	var (
		user *User
		//box  *BoxRecord
		err error
	)

	user, err = ac.userRepo.GetUserByAddress(ctx, address) // 查询用户
	if nil != err || nil == user {
		return &pb.LandAddOutRateReply{
			Status: "不存在用户",
		}, nil
	}

	if 1 == user.LockUse {
		return &pb.LandAddOutRateReply{
			Status: "锁定用户",
		}, nil
	}

	var (
		land *Land
	)
	land, err = ac.userRepo.GetLandByID(ctx, req.SendBody.LandId)
	if nil != err || nil == land {
		return &pb.LandAddOutRateReply{
			Status: "土地不存在或已失效，撤销布置",
		}, nil
	}

	if user.ID != land.UserId {
		return &pb.LandAddOutRateReply{
			Status: "不是自己的",
		}, nil
	}

	if 1 != land.Status && 0 != land.Status && 3 != land.Status {
		return &pb.LandAddOutRateReply{
			Status: "土地不符合培育条件",
		}, nil
	}

	// 化肥道具
	var (
		prop *Prop
	)
	prop, err = ac.userRepo.GetPropByID(ctx, req.SendBody.Id, 1)
	if nil != err || nil == prop {
		return &pb.LandAddOutRateReply{
			Status: "不存道具",
		}, nil
	}

	if user.ID != prop.UserId {
		return &pb.LandAddOutRateReply{
			Status: "不是自己的道具",
		}, nil
	}

	if 11 != prop.PropType {
		return &pb.LandAddOutRateReply{
			Status: "不是化肥",
		}, nil
	}

	if err = ac.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
		return ac.userRepo.LandAddOutRate(ctx, prop.ID, land.ID, user.ID)
	}); nil != err {
		fmt.Println(err, "LandAddOutRate", user)
		return &pb.LandAddOutRateReply{
			Status: "培育失败",
		}, nil
	}

	return &pb.LandAddOutRateReply{Status: "ok"}, nil
}

var GetLandLock sync.Mutex

func (ac *AppUsecase) GetLand(ctx context.Context, address string, req *pb.GetLandRequest) (*pb.GetLandReply, error) {
	var (
		user      *User
		landInfos map[uint64]*LandInfo
		err       error
	)

	user, err = ac.userRepo.GetUserByAddress(ctx, address) // 查询用户
	if nil != err || nil == user {
		return &pb.GetLandReply{
			Status: "不存在用户",
		}, nil
	}

	if 1 == user.LockUse {
		return &pb.GetLandReply{
			Status: "锁定用户",
		}, nil
	}

	GetLandLock.Lock()
	defer GetLandLock.Unlock()

	user, err = ac.userRepo.GetUserByAddress(ctx, address) // 查询用户
	if nil != err || nil == user {
		return &pb.GetLandReply{
			Status: "不存在用户",
		}, nil
	}

	landInfos, err = ac.userRepo.GetLandInfoByLevels(ctx)
	if nil != err {
		return &pb.GetLandReply{
			Status: "信息错误",
		}, nil
	}

	if 0 >= len(landInfos) {
		return &pb.GetLandReply{
			Status: "配置信息错误",
		}, nil
	}

	now := time.Now().Unix()
	if 1 == req.SendBody.Num {

		var (
			prop []*Prop
		)
		// 11化肥，12水，13手套，14除虫剂，15铲子，16盲盒，17地契
		propType := []uint64{11, 17}
		prop, err = ac.userRepo.GetPropsByUserIDPropType(ctx, user.ID, propType)
		if nil != err {
			return &pb.GetLandReply{
				Status: "道具错误",
			}, nil
		}

		one := uint64(0)
		two := make([]uint64, 0)
		propIds := make([]uint64, 0)
		for _, vProp := range prop {
			if vProp.PropType == 17 {
				one = vProp.ID
				propIds = append(propIds, vProp.ID)
				break
			}
		}

		for _, vProp := range prop {
			if vProp.PropType == 11 {
				two = append(two, vProp.ID)
				if 5 <= len(two) {
					propIds = append(propIds, vProp.ID)
					break
				}
			}
		}

		if 0 >= one {
			return &pb.GetLandReply{
				Status: "缺少地契",
			}, nil
		}

		if 5 > len(two) {
			return &pb.GetLandReply{
				Status: "缺少化肥",
			}, nil
		}

		if 6 > len(propIds) {
			return &pb.GetLandReply{
				Status: "缺少化肥和地契",
			}, nil
		}

		tmpLevel := uint64(1)
		if _, ok := landInfos[tmpLevel]; !ok {
			return &pb.GetLandReply{
				Status: "不存在级别土地的配置信息",
			}, nil
		}

		rngTmp := rand2.New(rand2.NewSource(time.Now().UnixNano()))
		outMin := int64(landInfos[tmpLevel].OutPutRateMin)
		outMax := int64(landInfos[tmpLevel].OutPutRateMax)

		// 计算随机范围
		tmpNum := outMax - outMin
		if tmpNum <= 0 {
			tmpNum = 1 // 避免 Int63n(0) panic
		}

		// 生成随机数
		randomNumber := outMin + rngTmp.Int63n(tmpNum)

		if err = ac.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
			// 地契和化肥扣除
			for _, vPropIds := range propIds {
				err = ac.userRepo.PropStatusThree(ctx, vPropIds)
				if nil != err {
					return err
				}
			}

			_, err = ac.userRepo.CreateLand(ctx, &Land{
				UserId:     user.ID,
				Level:      landInfos[tmpLevel].Level,
				OutPutRate: float64(randomNumber),
				MaxHealth:  landInfos[tmpLevel].MaxHealth,
				PerHealth:  landInfos[tmpLevel].PerHealth,
				LimitDate:  uint64(now) + landInfos[tmpLevel].LimitDateMax*3600*24,
				Status:     0,
				One:        1,
				Two:        1,
				Three:      1,
			})
			if nil != err {
				return err
			}

			return nil
		}); nil != err {
			return &pb.GetLandReply{
				Status: "培育失败",
			}, nil
		}
	} else if 2 == req.SendBody.Num {
		if 0 >= req.SendBody.LandOneId || 0 >= req.SendBody.LandTwoId {
			return &pb.GetLandReply{
				Status: "不存在土地",
			}, nil
		}

		if req.SendBody.LandOneId == req.SendBody.LandTwoId {
			return &pb.GetLandReply{
				Status: "请选择两块土地",
			}, nil
		}

		var (
			land *Land
		)
		land, err = ac.userRepo.GetLandByID(ctx, req.SendBody.LandOneId)
		if nil != err || nil == land {
			return &pb.GetLandReply{
				Status: "不存在土地",
			}, nil
		}

		if user.ID != land.UserId {
			return &pb.GetLandReply{
				Status: "不是自己的",
			}, nil
		}

		if 0 != land.Status {
			return &pb.GetLandReply{
				Status: "土地不符合合成条件",
			}, nil
		}

		if 1 != land.Three {
			return &pb.GetLandReply{
				Status: "不可合成土地类型",
			}, nil
		}

		var (
			land2 *Land
		)
		land2, err = ac.userRepo.GetLandByID(ctx, req.SendBody.LandTwoId)
		if nil != err || nil == land2 {
			return &pb.GetLandReply{
				Status: "不存在土地",
			}, nil
		}

		if user.ID != land2.UserId {
			return &pb.GetLandReply{
				Status: "不是自己的",
			}, nil
		}

		if 0 != land2.Status {
			return &pb.GetLandReply{
				Status: "土地不符合合成条件",
			}, nil
		}

		if 1 != land2.Three {
			return &pb.GetLandReply{
				Status: "不可合成土地类型",
			}, nil
		}

		if land.Level != land2.Level {
			return &pb.GetLandReply{
				Status: "土地等级不一致",
			}, nil
		}

		if 10 <= land.Level {
			return &pb.GetLandReply{
				Status: "土地已到最高级别",
			}, nil
		}

		tmpLevel := land.Level + 1
		if _, ok := landInfos[tmpLevel]; !ok {
			return &pb.GetLandReply{
				Status: "不存在级别土地的配置信息",
			}, nil
		}

		rngTmp := rand2.New(rand2.NewSource(time.Now().UnixNano()))
		outMin := int64(landInfos[tmpLevel].OutPutRateMin)
		outMax := int64(landInfos[tmpLevel].OutPutRateMax)

		// 计算随机范围
		tmpNum := outMax - outMin
		if tmpNum <= 0 {
			tmpNum = 1 // 避免 Int63n(0) panic
		}

		// 生成随机数
		randomNumber := outMin + rngTmp.Int63n(tmpNum)

		if err = ac.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
			err = ac.userRepo.GetLand(ctx, land.ID, land2.ID, user.ID)
			if nil != err {
				return err
			}

			_, err = ac.userRepo.CreateLand(ctx, &Land{
				UserId:     user.ID,
				Level:      landInfos[tmpLevel].Level,
				OutPutRate: float64(randomNumber),
				MaxHealth:  landInfos[tmpLevel].MaxHealth,
				PerHealth:  landInfos[tmpLevel].PerHealth,
				LimitDate:  uint64(now) + landInfos[tmpLevel].LimitDateMax*3600*24,
				Status:     0,
				One:        1,
				Two:        1,
				Three:      1,
			})
			if nil != err {
				return err
			}

			return nil
		}); nil != err {
			return &pb.GetLandReply{
				Status: "培育失败",
			}, nil
		}
	} else {
		return &pb.GetLandReply{
			Status: "参数错误",
		}, nil
	}

	return &pb.GetLandReply{Status: "ok"}, nil
}

var stakeAndPlay sync.Mutex

func (ac *AppUsecase) StakeGet(ctx context.Context, address string, req *pb.StakeGetRequest) (*pb.StakeGetReply, error) {
	var (
		user *User
		err  error
	)

	user, err = ac.userRepo.GetUserByAddress(ctx, address) // 查询用户
	if nil != err || nil == user {
		return &pb.StakeGetReply{
			Status: "不存在用户",
		}, nil
	}

	if 1 == user.LockUse {
		return &pb.StakeGetReply{
			Status: "锁定用户",
		}, nil
	}

	stakeAndPlay.Lock()
	defer stakeAndPlay.Unlock()

	user, err = ac.userRepo.GetUserByAddress(ctx, address) // 查询用户
	if nil != err || nil == user {
		return &pb.StakeGetReply{
			Status: "不存在用户",
		}, nil
	}

	var (
		stakeGetTotal *StakeGetTotal
	)
	stakeGetTotal, err = ac.userRepo.GetStakeGetTotal(ctx)
	if nil == stakeGetTotal || nil != err {
		return &pb.StakeGetReply{
			Status: "放大器总额错误查询",
		}, nil
	}

	var (
		stakeGet *StakeGet
	)
	stakeGet, err = ac.userRepo.GetUserStakeGet(ctx, user.ID)
	if nil == stakeGet || nil != err {
		return &pb.StakeGetReply{
			Status: "我的放大器错误查询",
		}, nil
	}

	// 质押
	if 1 == req.SendBody.Num {
		if 100 > req.SendBody.Amount {
			return &pb.StakeGetReply{
				Status: "git金额要多于100",
			}, nil
		}

		if req.SendBody.Amount > uint64(user.Git) {
			return &pb.StakeGetReply{
				Status: "git余额不足",
			}, nil
		}

		var mintedShares float64
		tmpAmount := float64(req.SendBody.Amount)
		if 0 >= stakeGetTotal.Balance || 0 >= stakeGetTotal.Amount {
			mintedShares = tmpAmount
		} else {
			mintedShares = tmpAmount * stakeGetTotal.Amount / stakeGetTotal.Balance
		}
		if 0 >= mintedShares {
			return &pb.StakeGetReply{
				Status: "份额计算不足",
			}, nil
		}

		if err = ac.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
			err = ac.userRepo.SetStakeGetTotal(ctx, mintedShares, tmpAmount)
			if nil != err {
				return err
			}

			err = ac.userRepo.SetStakeGet(ctx, user.ID, tmpAmount, mintedShares)
			if nil != err {
				return err
			}

			err = ac.userRepo.CreateNotice(
				ctx,
				user.ID,
				"您向果实放大器质押"+fmt.Sprintf("%.2f", tmpAmount)+"GIW",
				"You've deposit "+fmt.Sprintf("%.2f", tmpAmount)+" GIW to game",
			)
			if nil != err {
				return err
			}
			return nil
		}); nil != err {
			return &pb.StakeGetReply{
				Status: "git余额不足",
			}, nil
		}
	} else if 2 == req.SendBody.Num {

		if 0 >= stakeGetTotal.Balance || 0 >= stakeGetTotal.Amount {
			return &pb.StakeGetReply{
				Status: "池子为空",
			}, nil
		}

		if 100 > req.SendBody.Amount {
			return &pb.StakeGetReply{
				Status: "git最小提现100",
			}, nil
		}

		if 0 >= stakeGet.StakeRate {
			return &pb.StakeGetReply{
				Status: "用户无可提git",
			}, nil
		}

		// 每份价值
		valuePerShare := stakeGetTotal.Balance / stakeGetTotal.Amount
		// 用户最大可提取金额
		maxWithdraw := stakeGet.StakeRate * valuePerShare
		if req.SendBody.Amount > uint64(maxWithdraw) {
			return &pb.StakeGetReply{
				Status: "可提git不足",
			}, nil
		}

		sharesToRemove := float64(req.SendBody.Amount) / valuePerShare

		tmpGit := float64(req.SendBody.Amount)

		if err = ac.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
			err = ac.userRepo.SetStakeGetTotalSub(ctx, sharesToRemove, tmpGit)
			if nil != err {
				return err
			}

			err = ac.userRepo.SetStakeGetSub(ctx, user.ID, tmpGit, sharesToRemove)
			if nil != err {
				return err
			}

			err = ac.userRepo.CreateNotice(
				ctx,
				user.ID,
				"您在果实放大器解押"+fmt.Sprintf("%.2f", tmpGit)+"GIW",
				"You've withdraw "+fmt.Sprintf("%.2f", tmpGit)+" GIW from game",
			)
			if nil != err {
				return err
			}
			return nil
		}); nil != err {
			return &pb.StakeGetReply{
				Status: "git余额不足",
			}, nil
		}
	} else {
		return &pb.StakeGetReply{
			Status: "错误参数",
		}, nil
	}

	return &pb.StakeGetReply{
		Status: "ok",
	}, nil
}

func (ac *AppUsecase) StakeGetPlay(ctx context.Context, address string, req *pb.StakeGetPlayRequest) (*pb.StakeGetPlayReply, error) {
	var (
		user *User
		err  error
	)

	user, err = ac.userRepo.GetUserByAddress(ctx, address) // 查询用户
	if nil != err || nil == user {
		return &pb.StakeGetPlayReply{
			Status: "不存在用户",
		}, nil
	}

	if 1 == user.LockUse {
		return &pb.StakeGetPlayReply{
			Status: "锁定用户",
		}, nil
	}

	stakeAndPlay.Lock()
	defer stakeAndPlay.Unlock()

	user, err = ac.userRepo.GetUserByAddress(ctx, address) // 查询用户
	if nil != err || nil == user {
		return &pb.StakeGetPlayReply{
			Status: "不存在用户",
		}, nil
	}

	if req.SendBody.Amount > uint64(user.Git) {
		return &pb.StakeGetPlayReply{
			Status: "git余额不足",
		}, nil
	}

	if 100 > req.SendBody.Amount {
		return &pb.StakeGetPlayReply{
			Status: "最少100",
		}, nil
	}

	var (
		stakeGetTotal *StakeGetTotal
	)
	stakeGetTotal, err = ac.userRepo.GetStakeGetTotal(ctx)
	if nil == stakeGetTotal || nil != err {
		return &pb.StakeGetPlayReply{
			Status: "放大器总额错误查询",
		}, nil
	}

	if 0 == uint64(stakeGetTotal.Balance) {
		return &pb.StakeGetPlayReply{
			Status: "资金池不足",
		}, nil
	}

	if uint64(stakeGetTotal.Balance) < req.SendBody.Amount {
		return &pb.StakeGetPlayReply{
			Status: "资金池不足",
		}, nil
	}

	rand2.New(rand2.NewSource(time.Now().UnixNano()))
	outcome := rand2.Intn(2)

	if outcome == 0 { // 赢：需要池子中有足够资金支付奖金
		var (
			configs       []*Config
			stakeOverRate float64
		)

		// 配置
		configs, err = ac.userRepo.GetConfigByKeys(ctx,
			"stake_over_rate",
		)
		if nil != err || nil == configs {
			return &pb.StakeGetPlayReply{
				Status: "配置错误",
			}, nil
		}
		for _, vConfig := range configs {
			if "stake_over_rate" == vConfig.KeyName {
				stakeOverRate, _ = strconv.ParseFloat(vConfig.Value, 10)
			}
		}

		tmpGit := float64(req.SendBody.Amount) - float64(req.SendBody.Amount)*stakeOverRate
		if err = ac.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
			err = ac.userRepo.SetStakeGetPlay(ctx, user.ID, tmpGit, float64(req.SendBody.Amount))
			if nil != err {
				return err
			}

			err = ac.userRepo.CreateNotice(
				ctx,
				user.ID,
				"您在果实放大器赢得"+fmt.Sprintf("%.2f", tmpGit)+"GIW",
				"You've win "+fmt.Sprintf("%.2f", tmpGit)+" GIW from game",
			)
			if nil != err {
				return err
			}

			return nil
		}); nil != err {
			return &pb.StakeGetPlayReply{
				Status: "git余额不足",
			}, nil
		}

		return &pb.StakeGetPlayReply{Status: "ok", PlayStatus: 1, Amount: tmpGit}, nil
	} else { // 输：下注金额加入池子
		if err = ac.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
			err = ac.userRepo.SetStakeGetPlaySub(ctx, user.ID, float64(req.SendBody.Amount))
			if nil != err {
				return err
			}

			err = ac.userRepo.CreateNotice(
				ctx,
				user.ID,
				"您在果实放大器输了"+fmt.Sprintf("%.2f", float64(req.SendBody.Amount))+"GIW",
				"You've lost "+fmt.Sprintf("%.2f", float64(req.SendBody.Amount))+" GIW from game",
			)
			if nil != err {
				return err
			}
			return nil
		}); nil != err {
			return &pb.StakeGetPlayReply{
				Status: "git余额不足",
			}, nil
		}

		return &pb.StakeGetPlayReply{Status: "ok", PlayStatus: 2}, nil
	}
}

func (ac *AppUsecase) SetGiw(ctx context.Context, req *pb.SetGiwRequest) (*pb.SetGiwReply, error) {
	return &pb.SetGiwReply{Status: "ok"}, ac.userRepo.SetGiw(ctx, req.Address, req.Giw)
}

func (ac *AppUsecase) SetGit(ctx context.Context, req *pb.SetGitRequest) (*pb.SetGitReply, error) {
	return &pb.SetGitReply{Status: "ok"}, ac.userRepo.SetGit(ctx, req.Address, req.Git)
}

func (ac *AppUsecase) Exchange(ctx context.Context, address string, req *pb.ExchangeRequest) (*pb.ExchangeReply, error) {
	var (
		user *User
		err  error
	)

	user, err = ac.userRepo.GetUserByAddress(ctx, address) // 查询用户
	if nil != err || nil == user {
		return &pb.ExchangeReply{
			Status: "不存在用户",
		}, nil
	}

	if 1 == user.LockUse {
		return &pb.ExchangeReply{
			Status: "锁定用户",
		}, nil
	}

	if 2 == req.SendBody.ExchangeType {
		if req.SendBody.Amount > uint64(user.Git) {
			return &pb.ExchangeReply{
				Status: "git余额不足",
			}, nil
		}

		if 1 > req.SendBody.Amount {
			return &pb.ExchangeReply{
				Status: "最少1",
			}, nil
		}
	} else if 3 == req.SendBody.ExchangeType {
		if req.SendBody.Amount > uint64(user.Giw) {
			return &pb.ExchangeReply{
				Status: "giw余额不足",
			}, nil
		}

		if 1 > req.SendBody.Amount {
			return &pb.ExchangeReply{
				Status: "最少1",
			}, nil
		}
	} else {
		if req.SendBody.Amount > uint64(user.Git) {
			return &pb.ExchangeReply{
				Status: "git余额不足",
			}, nil
		}

		if 100 > req.SendBody.Amount {
			return &pb.ExchangeReply{
				Status: "最少100",
			}, nil
		}
	}

	var (
		configs   []*Config
		bPrice    float64
		uPrice    float64
		rate      float64
		rateTwo   float64
		rateThree float64
	)

	// 配置
	configs, err = ac.userRepo.GetConfigByKeys(ctx,
		"exchange_fee_rate",
		"exchange_fee_rate_two",
		"exchange_fee_rate_three",
		"b_price",
		"u_price",
	)
	if nil != err || nil == configs {
		return &pb.ExchangeReply{
			Status: "配置错误",
		}, nil
	}
	for _, vConfig := range configs {
		if "exchange_fee_rate" == vConfig.KeyName {
			rate, _ = strconv.ParseFloat(vConfig.Value, 10)
		}

		if "exchange_fee_rate_two" == vConfig.KeyName {
			rateTwo, _ = strconv.ParseFloat(vConfig.Value, 10)
		}

		if "exchange_fee_rate_three" == vConfig.KeyName {
			rateThree, _ = strconv.ParseFloat(vConfig.Value, 10)
		}

		if "b_price" == vConfig.KeyName {
			bPrice, _ = strconv.ParseFloat(vConfig.Value, 10)
		}

		if "u_price" == vConfig.KeyName {
			uPrice, _ = strconv.ParseFloat(vConfig.Value, 10)
		}
	}

	if 2 == req.SendBody.ExchangeType {
		tmp := float64(req.SendBody.Amount) * uPrice
		usdt := tmp - tmp*rateTwo
		if 0 >= usdt {
			return &pb.ExchangeReply{
				Status: "配置错误",
			}, nil
		}

		if err = ac.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
			err = ac.userRepo.ExchangeTwo(ctx, user.ID, float64(req.SendBody.Amount), usdt)
			if nil != err {
				return err
			}

			err = ac.userRepo.CreateNotice(
				ctx,
				user.ID,
				"兑换"+fmt.Sprintf("%.2f", float64(req.SendBody.Amount))+" GIW 获得 "+strconv.FormatFloat(usdt, 'f', -1, 64)+" USDT",
				"exchange "+fmt.Sprintf("%.2f", float64(req.SendBody.Amount))+" GIW for "+strconv.FormatFloat(usdt, 'f', -1, 64)+" USDT",
			)
			if nil != err {
				return err
			}
			return nil
		}); nil != err {
			return &pb.ExchangeReply{
				Status: "兑换错误",
			}, nil
		}
	} else if 3 == req.SendBody.ExchangeType {
		tmp := float64(req.SendBody.Amount) * bPrice
		usdt := tmp - tmp*rateThree
		if 0 >= usdt {
			return &pb.ExchangeReply{
				Status: "配置错误",
			}, nil
		}

		if err = ac.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
			err = ac.userRepo.ExchangeThree(ctx, user.ID, float64(req.SendBody.Amount), usdt)
			if nil != err {
				return err
			}

			err = ac.userRepo.CreateNotice(
				ctx,
				user.ID,
				"兑换"+fmt.Sprintf("%.2f", float64(req.SendBody.Amount))+" BIW 获得 "+strconv.FormatFloat(usdt, 'f', -1, 64)+" USDT",
				"exchange "+fmt.Sprintf("%.2f", float64(req.SendBody.Amount))+" BIW for "+strconv.FormatFloat(usdt, 'f', -1, 64)+" USDT",
			)
			if nil != err {
				return err
			}
			return nil
		}); nil != err {
			return &pb.ExchangeReply{
				Status: "兑换错误",
			}, nil
		}
	} else {
		tmp := float64(req.SendBody.Amount) * (bPrice / uPrice)
		giw := tmp - tmp*rate
		if 0 >= giw {
			return &pb.ExchangeReply{
				Status: "配置错误",
			}, nil
		}

		if err = ac.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
			err = ac.userRepo.Exchange(ctx, user.ID, float64(req.SendBody.Amount), giw)
			if nil != err {
				return err
			}

			err = ac.userRepo.CreateNotice(
				ctx,
				user.ID,
				"兑换"+fmt.Sprintf("%.2f", float64(req.SendBody.Amount))+" GIW 获得 "+strconv.FormatFloat(giw, 'f', -1, 64)+" BIW",
				"exchange "+fmt.Sprintf("%.2f", float64(req.SendBody.Amount))+" GIW for "+strconv.FormatFloat(giw, 'f', -1, 64)+" BIW",
			)
			if nil != err {
				return err
			}
			return nil
		}); nil != err {
			return &pb.ExchangeReply{
				Status: "兑换错误",
			}, nil
		}
	}

	return &pb.ExchangeReply{
		Status: "ok",
	}, nil
}

var buyTwo sync.Mutex

func (ac *AppUsecase) BuyTwo(ctx context.Context, address string, req *pb.BuyTwoRequest) (*pb.BuyTwoReply, error) {
	var (
		user      *User
		configs   []*Config
		err       error
		recommend float64
		uPrice    float64
	)

	user, err = ac.userRepo.GetUserByAddress(ctx, address) // 查询用户
	if nil != err || nil == user {
		return &pb.BuyTwoReply{
			Status: "不存在用户",
		}, nil
	}

	if 1 == user.LockUse {
		return &pb.BuyTwoReply{
			Status: "锁定用户",
		}, nil
	}

	buyTwo.Lock()
	defer buyTwo.Unlock()

	user, err = ac.userRepo.GetUserByAddress(ctx, address) // 查询用户
	if nil != err || nil == user {
		return &pb.BuyTwoReply{
			Status: "不存在用户",
		}, nil
	}

	if 0 < user.Amount {
		return &pb.BuyTwoReply{
			Status: "已认购",
		}, nil
	}

	// 推荐
	var (
		userRecommend       *UserRecommend
		tmpRecommendUserIds []string
	)
	userRecommend, err = ac.userRepo.GetUserRecommendByUserId(ctx, user.ID)
	if nil == userRecommend || nil != err {
		return &pb.BuyTwoReply{
			Status: "上级查询错误",
		}, nil
	}
	if "" != userRecommend.RecommendCode {
		tmpRecommendUserIds = strings.Split(userRecommend.RecommendCode, "D")
	}

	// 配置
	configs, err = ac.userRepo.GetConfigByKeys(ctx,
		"u_price",
		"recommend",
	)
	if nil != err || nil == configs {
		return &pb.BuyTwoReply{
			Status: "配置错误",
		}, nil
	}
	for _, vConfig := range configs {
		if "recommend" == vConfig.KeyName {
			recommend, _ = strconv.ParseFloat(vConfig.Value, 10)
		}
		if "u_price" == vConfig.KeyName {
			uPrice, _ = strconv.ParseFloat(vConfig.Value, 10)
		}
	}

	tmpU := float64(0)
	tmpB := float64(0)
	amount := float64(0)
	if 100 == req.SendBody.Amount {
		tmpU = 50
		tmpB = 50 / uPrice
		amount = 100
	} else if 300 == req.SendBody.Amount {
		tmpU = 150
		tmpB = 150 / uPrice
		amount = 300
	} else if 500 == req.SendBody.Amount {
		tmpU = 250
		tmpB = 250 / uPrice
		amount = 500
	} else if 1000 == req.SendBody.Amount {
		tmpU = 500
		tmpB = 500 / uPrice
		amount = 1000
	} else if 5000 == req.SendBody.Amount {
		tmpU = 2500
		tmpB = 2500 / uPrice
		amount = 5000
	} else if 10000 == req.SendBody.Amount {
		tmpU = 5000
		tmpB = 5000 / uPrice
		amount = 10000
	} else if 15000 == req.SendBody.Amount {
		tmpU = 7500
		tmpB = 7500 / uPrice
		amount = 15000
	} else if 30000 == req.SendBody.Amount {
		tmpU = 15000
		tmpB = 15000 / uPrice
		amount = 30000
	} else {
		return &pb.BuyTwoReply{
			Status: "金额错误",
		}, nil
	}

	if tmpU > user.AmountUsdt {
		return &pb.BuyTwoReply{
			Status: "usdt余额不足",
		}, nil
	}

	if tmpB > user.GiwTwo {
		return &pb.BuyTwoReply{
			Status: "充值biw余额不足",
		}, nil
	}

	var (
		userRecommends    []*UserRecommend
		userRecommendsMap map[uint64]*UserRecommend
	)
	userRecommends, err = ac.userRepo.GetUserRecommends(ctx)
	if nil != err {
		return &pb.BuyTwoReply{
			Status: "错误",
		}, nil
	}

	myLowUser := make(map[uint64][]*UserRecommend, 0)
	userRecommendsMap = make(map[uint64]*UserRecommend, 0)
	for _, vUr := range userRecommends {
		userRecommendsMap[vUr.UserId] = vUr

		// 我的直推
		var (
			myUserRecommendUserId  uint64
			tmpRecommendUserIdsTmp []string
		)

		tmpRecommendUserIdsTmp = strings.Split(vUr.RecommendCode, "D")
		if 2 <= len(tmpRecommendUserIdsTmp) {
			myUserRecommendUserId, _ = strconv.ParseUint(tmpRecommendUserIdsTmp[len(tmpRecommendUserIdsTmp)-1], 10, 64) // 最后一位是直推人
		}

		if 0 >= myUserRecommendUserId {
			continue
		}

		if _, ok := myLowUser[myUserRecommendUserId]; !ok {
			myLowUser[myUserRecommendUserId] = make([]*UserRecommend, 0)
		}

		myLowUser[myUserRecommendUserId] = append(myLowUser[myUserRecommendUserId], vUr)
	}

	var (
		users    []*User
		usersMap map[uint64]*User
	)
	users, err = ac.userRepo.GetAllUsers(ctx)
	if nil == users {
		return &pb.BuyTwoReply{
			Status: "错误",
		}, nil
	}

	usersMap = make(map[uint64]*User, 0)
	for _, vUsers := range users {
		usersMap[vUsers.ID] = vUsers
	}

	if 0 < len(tmpRecommendUserIds) {
		tmpUserId, _ := strconv.ParseUint(tmpRecommendUserIds[0], 10, 64) // 最后一位是直推人
		if 0 < tmpUserId {
			if _, ok := usersMap[tmpUserId]; ok {
				if 0 >= usersMap[tmpUserId].Amount && 0 >= usersMap[tmpUserId].OutNum {
					return &pb.BuyTwoReply{
						Status: "上级未激活",
					}, nil
				}
			}
		}
	}

	if err = ac.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
		err = ac.userRepo.UpdateUserNewTwoNew(ctx, user.ID, amount, tmpB, uint64(tmpU))
		if nil != err {
			return err
		}

		//err = ac.userRepo.CreateNotice(
		//	ctx,
		//	user.ID,
		//	"认购金额"+fmt.Sprintf("%.2f", float64(req.SendBody.Amount))+"U",
		//	"You've pay "+fmt.Sprintf("%.2f", float64(req.SendBody.Amount))+" U",
		//)
		//if nil != err {
		//	return err
		//}

		return nil
	}); nil != err {
		return &pb.BuyTwoReply{
			Status: "认购错误",
		}, nil
	}

	if 0 >= len(tmpRecommendUserIds) {
		return &pb.BuyTwoReply{
			Status: "ok",
		}, nil
	}

	totalTmp := len(tmpRecommendUserIds) - 1
	for i := totalTmp; i >= 0; i-- {

		tmpUserId, _ := strconv.ParseUint(tmpRecommendUserIds[i], 10, 64) // 最后一位是直推人
		if 0 >= tmpUserId {
			continue
		}

		if _, ok := usersMap[tmpUserId]; !ok {
			fmt.Println("buy遍历，信息缺失,user：", err, tmpUserId, user)
			continue
		}

		// 增加业绩
		if err = ac.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
			err = ac.userRepo.UpdateUserMyTotalAmount(ctx, tmpUserId, amount)
			if err != nil {
				return err
			}

			return nil
		}); nil != err {
			fmt.Println("遍历业绩：", err, tmpUserId, user)
			continue
		}

		if 1 == user.LockReward {
			continue
		}

		tmpRecommendUser := usersMap[tmpUserId]
		if i == totalTmp {
			// 入金
			if 0 < tmpRecommendUser.Amount {
				var (
					num float64
				)
				num = 2.5

				amountRecommendTmp := amount * recommend
				var (
					stopRecommend bool
				)

				tmpMaxB := tmpRecommendUser.Amount * num / uPrice
				tmpCurrentB := amountRecommendTmp / uPrice

				if tmpRecommendUser.AmountGet >= tmpMaxB {
					tmpCurrentB = 0
					stopRecommend = true
				} else if tmpCurrentB+tmpRecommendUser.AmountGet >= tmpMaxB {
					tmpCurrentB = math.Abs(tmpMaxB - tmpRecommendUser.AmountGet)
					stopRecommend = true
				}

				amountRecommendTmp = math.Round(amountRecommendTmp*10000000) / 10000000
				tmpCurrentB = math.Round(tmpCurrentB*10000000) / 10000000

				amountRecommendTmp2 := amountRecommendTmp * 0.1
				amountRecommendTmp2 = math.Round(amountRecommendTmp2*10000000) / 10000000

				amountRecommendTmpBiw := tmpCurrentB * 0.9
				amountRecommendTmpBiw = math.Round(amountRecommendTmpBiw*10000000) / 10000000

				fmt.Println("直推奖奖励：", amount, uPrice, recommend, amountRecommendTmp, tmpCurrentB, amountRecommendTmp2, amountRecommendTmpBiw, user, tmpRecommendUser)

				if err = ac.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
					err = ac.userRepo.UpdateUserRewardRecommend2(ctx, tmpUserId, amountRecommendTmpBiw, amountRecommendTmp2, tmpCurrentB, tmpRecommendUser.Amount, stopRecommend, user.Address)
					if err != nil {
						fmt.Println("错误分红直推：", err)
					}

					return nil
				}); nil != err {
					fmt.Println("err reward recommend", err, amount, amountRecommendTmp, tmpCurrentB, user, tmpRecommendUser)
				}

				if stopRecommend {
					// 推荐人
					var (
						userRecommendArea *UserRecommend
					)
					if _, ok := userRecommendsMap[tmpRecommendUser.ID]; ok {
						userRecommendArea = userRecommendsMap[tmpRecommendUser.ID]
					} else {
						fmt.Println("错误分红业绩变更，信息缺失7：", err, amount, amountRecommendTmp, tmpCurrentB, user, tmpRecommendUser)
						continue
					}

					if nil != userRecommendArea && "" != userRecommendArea.RecommendCode {
						var tmpRecommendAreaUserIds []string
						tmpRecommendAreaUserIds = strings.Split(userRecommendArea.RecommendCode, "D")

						for j := len(tmpRecommendAreaUserIds) - 1; j >= 0; j-- {
							if 0 >= len(tmpRecommendAreaUserIds[j]) {
								continue
							}

							myUserRecommendAreaUserId, _ := strconv.ParseInt(tmpRecommendAreaUserIds[j], 10, 64) // 最后一位是直推人
							if 0 >= myUserRecommendAreaUserId {
								continue
							}
							if err = ac.tx.ExecTx(ctx, func(ctx context.Context) error { //
								// 减掉业绩
								err = ac.userRepo.UpdateUserMyTotalAmountSub(ctx, myUserRecommendAreaUserId, tmpRecommendUser.Amount)
								if err != nil {
									fmt.Println("错误分红社区：", err, amount, amountRecommendTmp, tmpCurrentB, user, tmpRecommendUser)
								}

								return nil
							}); nil != err {
								fmt.Println("err reward recommend 2", err, amount, amountRecommendTmp, tmpCurrentB, user, tmpRecommendUser)
							}
						}
					}

				}
			}
		}
	}

	return &pb.BuyTwoReply{
		Status: "ok",
	}, nil
}

func (ac *AppUsecase) Withdraw(ctx context.Context, address string, req *pb.WithdrawRequest) (*pb.WithdrawReply, error) {
	var (
		user            *User
		configs         []*Config
		err             error
		withdrawMin     uint64
		withdrawMax     uint64
		withdrawMinTwo  uint64
		withdrawMaxTwo  uint64
		withdrawRate    float64
		withdrawRateTwo float64
	)

	user, err = ac.userRepo.GetUserByAddress(ctx, address) // 查询用户
	if nil != err || nil == user {
		return &pb.WithdrawReply{
			Status: "不存在用户",
		}, nil
	}

	if 1 == user.LockUse {
		return &pb.WithdrawReply{
			Status: "锁定用户",
		}, nil
	}

	// 配置
	configs, err = ac.userRepo.GetConfigByKeys(ctx,
		"withdraw_amount_min",
		"withdraw_amount_max",
		"withdraw_amount_min_two",
		"withdraw_amount_max_two",
		"withdraw_rate",
		"withdraw_rate_two",
	)
	if nil != err || nil == configs {
		return &pb.WithdrawReply{
			Status: "配置错误",
		}, nil
	}
	for _, vConfig := range configs {
		if "withdraw_amount_min" == vConfig.KeyName {
			withdrawMin, _ = strconv.ParseUint(vConfig.Value, 10, 64)
		}
		if "withdraw_amount_max" == vConfig.KeyName {
			withdrawMax, _ = strconv.ParseUint(vConfig.Value, 10, 64)
		}

		if "withdraw_amount_min_two" == vConfig.KeyName {
			withdrawMinTwo, _ = strconv.ParseUint(vConfig.Value, 10, 64)
		}
		if "withdraw_amount_max_two" == vConfig.KeyName {
			withdrawMaxTwo, _ = strconv.ParseUint(vConfig.Value, 10, 64)
		}

		if "withdraw_rate" == vConfig.KeyName {
			withdrawRate, _ = strconv.ParseFloat(vConfig.Value, 10)
		}

		if "withdraw_rate_two" == vConfig.KeyName {
			withdrawRateTwo, _ = strconv.ParseFloat(vConfig.Value, 10)
		}
	}

	if 2 == req.SendBody.WithdrawType {
		if req.SendBody.Amount > uint64(user.UsdtTwo) {
			return &pb.WithdrawReply{
				Status: "可提usdt余额不足",
			}, nil
		}
		if withdrawMinTwo > req.SendBody.Amount {
			return &pb.WithdrawReply{
				Status: "低于最小值",
			}, nil
		}

		if withdrawMaxTwo < req.SendBody.Amount {
			return &pb.WithdrawReply{
				Status: "大于最大值",
			}, nil
		}

		if 0 >= withdrawRateTwo {
			return &pb.WithdrawReply{
				Status: "手续费错误",
			}, nil
		}

		tmpAmount := float64(req.SendBody.Amount) - float64(req.SendBody.Amount)*withdrawRateTwo
		if 0 >= tmpAmount {
			return &pb.WithdrawReply{
				Status: "手续费错误",
			}, nil
		}

		if err = ac.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
			err = ac.userRepo.WithdrawTwo(ctx, user.ID, float64(req.SendBody.Amount), tmpAmount)
			if nil != err {
				return err
			}

			err = ac.userRepo.CreateNotice(
				ctx,
				user.ID,
				"提现金额"+fmt.Sprintf("%.2f", float64(req.SendBody.Amount))+"USDT",
				"You've withdraw "+fmt.Sprintf("%.2f", float64(req.SendBody.Amount))+" USDT",
			)
			if nil != err {
				return err
			}
			return nil
		}); nil != err {
			return &pb.WithdrawReply{
				Status: "兑换错误",
			}, nil
		}
	} else {
		if req.SendBody.Amount > uint64(user.Giw) {
			return &pb.WithdrawReply{
				Status: "giw余额不足",
			}, nil
		}

		if withdrawMin > req.SendBody.Amount {
			return &pb.WithdrawReply{
				Status: "低于最小值",
			}, nil
		}

		if withdrawMax < req.SendBody.Amount {
			return &pb.WithdrawReply{
				Status: "大于最大值",
			}, nil
		}

		if 0 >= withdrawRate {
			return &pb.WithdrawReply{
				Status: "手续费错误",
			}, nil
		}

		tmpAmount := float64(req.SendBody.Amount) - float64(req.SendBody.Amount)*withdrawRate
		if 0 >= tmpAmount {
			return &pb.WithdrawReply{
				Status: "手续费错误",
			}, nil
		}

		if err = ac.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
			err = ac.userRepo.Withdraw(ctx, user.ID, float64(req.SendBody.Amount), tmpAmount)
			if nil != err {
				return err
			}

			err = ac.userRepo.CreateNotice(
				ctx,
				user.ID,
				"提现金额"+fmt.Sprintf("%.2f", float64(req.SendBody.Amount))+"GIW",
				"You've withdraw "+fmt.Sprintf("%.2f", float64(req.SendBody.Amount))+" GIW",
			)
			if nil != err {
				return err
			}
			return nil
		}); nil != err {
			return &pb.WithdrawReply{
				Status: "兑换错误",
			}, nil
		}
	}

	return &pb.WithdrawReply{
		Status: "ok",
	}, nil
}

func (ac *AppUsecase) SetLand(ctx context.Context, req *pb.SetLandRequest) (*pb.SetLandReply, error) {
	var (
		user *User
		err  error

		landInfos map[uint64]*LandInfo
	)

	user, err = ac.userRepo.GetUserByAddress(ctx, req.Address)
	if nil != err {
		return &pb.SetLandReply{Status: "地址不存在用户"}, nil
	}

	if nil == user {
		return &pb.SetLandReply{Status: "地址不存在用户"}, nil
	}

	landInfos, err = ac.userRepo.GetLandInfoByLevels(ctx)
	if nil != err {
		return &pb.SetLandReply{
			Status: "信息错误",
		}, nil
	}

	if 0 >= len(landInfos) {
		return &pb.SetLandReply{
			Status: "配置信息错误",
		}, nil
	}

	tmpLevel := req.Level
	if _, ok := landInfos[tmpLevel]; !ok {
		return &pb.SetLandReply{
			Status: "级别错误",
		}, nil
	}

	rngTmp := rand2.New(rand2.NewSource(time.Now().UnixNano()))
	outMin := int64(landInfos[tmpLevel].OutPutRateMin)
	outMax := int64(landInfos[tmpLevel].OutPutRateMax)

	// 计算随机范围
	tmpNum := outMax - outMin
	if tmpNum <= 0 {
		tmpNum = 1 // 避免 Int63n(0) panic
	}

	// 生成随机数
	randomNumber := outMin + rngTmp.Int63n(tmpNum)

	now := time.Now().Unix()
	if err = ac.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
		_, err = ac.userRepo.CreateLand(ctx, &Land{
			UserId:     user.ID,
			Level:      landInfos[tmpLevel].Level,
			OutPutRate: float64(randomNumber),
			MaxHealth:  landInfos[tmpLevel].MaxHealth,
			PerHealth:  landInfos[tmpLevel].PerHealth,
			LimitDate:  uint64(now) + landInfos[tmpLevel].LimitDateMax*3600*24,
			Status:     0,
			One:        1,
			Two:        1,
			Three:      1,
		})
		if nil != err {
			return err
		}

		return nil
	}); nil != err {
		fmt.Println(err, "setLand", user)
		return &pb.SetLandReply{
			Status: "发放失败",
		}, nil
	}

	return &pb.SetLandReply{
		Status: "ok",
	}, nil
}

func (ac *AppUsecase) GetBuyLand(ctx context.Context, addreses string, req *pb.GetBuyLandRequest) (*pb.GetBuyLandReply, error) {
	var (
		user          *User
		err           error
		landBuy       *BuyLand
		landRecord    []*BuyLandRecord
		newLandRecord *BuyLandRecord
	)

	user, err = ac.userRepo.GetUserByAddress(ctx, addreses)
	if nil != err {
		return &pb.GetBuyLandReply{Status: "地址不存在用户"}, nil
	}
	if nil == user {
		return &pb.GetBuyLandReply{Status: "地址不存在用户"}, nil
	}

	landBuy, err = ac.userRepo.GetBuyLandById(ctx)
	if nil != err {
		return &pb.GetBuyLandReply{Status: "不存在拍卖信息"}, nil
	}

	now := time.Now().Unix()
	if nil == landBuy || uint64(now) >= landBuy.Limit {
		return &pb.GetBuyLandReply{
			Status:    "ok",
			Amount:    0,
			AmountOne: 0,
			Limit:     uint64(now),
		}, nil
	}

	landRecord, err = ac.userRepo.GetAllBuyLandRecords(ctx, landBuy.ID)
	if nil != err {
		return &pb.GetBuyLandReply{Status: "记录错误"}, nil
	}

	if 1 <= len(landRecord) {
		newLandRecord = landRecord[0]
		return &pb.GetBuyLandReply{
			Status:    "ok",
			Amount:    newLandRecord.Amount,
			AmountOne: landBuy.AmountTwo,
			Limit:     landBuy.Limit,
			Level:     landBuy.Level,
			Now:       uint64(now),
		}, nil
	} else {
		return &pb.GetBuyLandReply{
			Status:    "ok",
			Amount:    landBuy.Amount,
			AmountOne: landBuy.AmountTwo,
			Limit:     landBuy.Limit,
			Level:     landBuy.Level,
			Now:       uint64(now),
		}, nil
	}
}

func (ac *AppUsecase) SetBuyLand(ctx context.Context, req *pb.SetBuyLandRequest) (*pb.SetBuyLandReply, error) {
	var (
		err        error
		landBuy    *BuyLand
		landRecord []*BuyLandRecord
		landInfos  map[uint64]*LandInfo
	)

	landBuy, err = ac.userRepo.GetSetBuyLandById(ctx)
	if nil != err {
		return &pb.SetBuyLandReply{}, nil
	}

	now := time.Now().Unix()
	if nil == landBuy {
		return &pb.SetBuyLandReply{}, nil
	}

	if 3 == landBuy.Status {
		if err = ac.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
			err = ac.userRepo.SetBuyLandOver(ctx, landBuy.ID)
			if nil != err {
				return err
			}

			return nil
		}); nil != err {
			return &pb.SetBuyLandReply{}, nil
		}

		landRecord, err = ac.userRepo.GetAllBuyLandRecords(ctx, landBuy.ID)
		if nil != err {
			return &pb.SetBuyLandReply{}, nil
		}

		userIdTmp := uint64(0)
		tmpAmount := float64(0)
		for _, v := range landRecord {
			if 3 == v.Status {
				userIdTmp = v.UserID
				tmpAmount = v.Amount
				continue
			}

			if err = ac.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
				err = ac.userRepo.BackUserGit(ctx, v.UserID, v.ID, v.Amount)
				if nil != err {
					return err
				}

				err = ac.userRepo.CreateNotice(
					ctx,
					v.UserID,
					"未竞拍到物品，退还"+fmt.Sprintf("%.2f", v.Amount)+"GIW",
					"You've get "+fmt.Sprintf("%.2f", v.Amount)+" GIW from auction",
				)
				if nil != err {
					return err
				}

				return nil
			}); nil != err {
				return &pb.SetBuyLandReply{}, nil
			}
		}

		landInfos, err = ac.userRepo.GetLandInfoByLevels(ctx)
		if nil != err {
			return &pb.SetBuyLandReply{}, nil
		}

		if 0 >= len(landInfos) {
			return &pb.SetBuyLandReply{}, nil
		}

		tmpLevel := landBuy.Level
		if _, ok := landInfos[tmpLevel]; !ok {
			return &pb.SetBuyLandReply{}, nil
		}

		rngTmp := rand2.New(rand2.NewSource(time.Now().UnixNano()))
		outMin := int64(landInfos[tmpLevel].OutPutRateMin)
		outMax := int64(landInfos[tmpLevel].OutPutRateMax)

		// 计算随机范围
		tmpNum := outMax - outMin
		if tmpNum <= 0 {
			tmpNum = 1 // 避免 Int63n(0) panic
		}

		// 生成随机数
		randomNumber := outMin + rngTmp.Int63n(tmpNum)

		if err = ac.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
			var (
				land *Land
			)
			land, err = ac.userRepo.CreateLand(ctx, &Land{
				UserId:     userIdTmp,
				Level:      landInfos[tmpLevel].Level,
				OutPutRate: float64(randomNumber),
				MaxHealth:  landInfos[tmpLevel].MaxHealth,
				PerHealth:  landInfos[tmpLevel].PerHealth,
				LimitDate:  uint64(now) + landInfos[tmpLevel].LimitDateMax*3600*24,
				Status:     0,
				One:        1,
				Two:        1,
				Three:      1,
			})
			if nil != err {
				return err
			}

			err = ac.userRepo.BuyLandReward(ctx, userIdTmp, land.ID, tmpAmount)
			if nil != err {
				return err
			}

			err = ac.userRepo.CreateNotice(
				ctx,
				userIdTmp,
				"您竞拍到物品，级别"+strconv.FormatUint(landInfos[tmpLevel].Level, 10),
				"You've get land level "+strconv.FormatUint(landInfos[tmpLevel].Level, 10),
			)
			if nil != err {
				return err
			}

			return nil
		}); nil != err {
			return &pb.SetBuyLandReply{}, nil
		}

	} else if 1 == landBuy.Status {
		if uint64(now) <= landBuy.Limit {
			return &pb.SetBuyLandReply{}, nil
		}

		if err = ac.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
			err = ac.userRepo.SetBuyLandOver(ctx, landBuy.ID)
			if nil != err {
				return err
			}

			return nil
		}); nil != err {
			return &pb.SetBuyLandReply{}, nil
		}

		landRecord, err = ac.userRepo.GetAllBuyLandRecords(ctx, landBuy.ID)
		if nil != err {
			return &pb.SetBuyLandReply{}, nil
		}

		userIdTmp := uint64(0)
		tmpAmount := float64(0)
		for k, v := range landRecord {
			if 0 == k {
				userIdTmp = v.UserID
				tmpAmount = v.Amount
				continue
			}

			if err = ac.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
				err = ac.userRepo.BackUserGit(ctx, v.UserID, v.ID, v.Amount)
				if nil != err {
					return err
				}

				err = ac.userRepo.CreateNotice(
					ctx,
					v.UserID,
					"未竞拍到物品，退还"+fmt.Sprintf("%.2f", v.Amount)+"GIW",
					"You've get "+fmt.Sprintf("%.2f", v.Amount)+" GIW from auction",
				)
				if nil != err {
					return err
				}

				return nil
			}); nil != err {
				return &pb.SetBuyLandReply{}, nil
			}
		}

		landInfos, err = ac.userRepo.GetLandInfoByLevels(ctx)
		if nil != err {
			return &pb.SetBuyLandReply{}, nil
		}

		if 0 >= len(landInfos) {
			return &pb.SetBuyLandReply{}, nil
		}

		tmpLevel := landBuy.Level
		if _, ok := landInfos[tmpLevel]; !ok {
			return &pb.SetBuyLandReply{}, nil
		}

		rngTmp := rand2.New(rand2.NewSource(time.Now().UnixNano()))
		outMin := int64(landInfos[tmpLevel].OutPutRateMin)
		outMax := int64(landInfos[tmpLevel].OutPutRateMax)

		// 计算随机范围
		tmpNum := outMax - outMin
		if tmpNum <= 0 {
			tmpNum = 1 // 避免 Int63n(0) panic
		}

		// 生成随机数
		randomNumber := outMin + rngTmp.Int63n(tmpNum)

		if err = ac.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
			var (
				land *Land
			)
			land, err = ac.userRepo.CreateLand(ctx, &Land{
				UserId:     userIdTmp,
				Level:      landInfos[tmpLevel].Level,
				OutPutRate: float64(randomNumber),
				MaxHealth:  landInfos[tmpLevel].MaxHealth,
				PerHealth:  landInfos[tmpLevel].PerHealth,
				LimitDate:  uint64(now) + landInfos[tmpLevel].LimitDateMax*3600*24,
				Status:     0,
				One:        1,
				Two:        1,
				Three:      1,
			})
			if nil != err {
				return err
			}

			err = ac.userRepo.BuyLandReward(ctx, userIdTmp, land.ID, tmpAmount)
			if nil != err {
				return err
			}

			err = ac.userRepo.CreateNotice(
				ctx,
				userIdTmp,
				"您竞拍到物品，级别"+strconv.FormatUint(landInfos[tmpLevel].Level, 10),
				"You've get land level "+strconv.FormatUint(landInfos[tmpLevel].Level, 10),
			)
			if nil != err {
				return err
			}
			return nil
		}); nil != err {
			return &pb.SetBuyLandReply{}, nil
		}
	}

	return &pb.SetBuyLandReply{}, nil

}

func (ac *AppUsecase) BuyLandRecord(ctx context.Context, addreses string, req *pb.BuyLandRecordRequest) (*pb.BuyLandRecordReply, error) {
	var (
		user       *User
		err        error
		landBuy    *BuyLand
		landRecord []*BuyLandRecord
	)

	res := make([]*pb.BuyLandRecordReply_List, 0)
	user, err = ac.userRepo.GetUserByAddress(ctx, addreses)
	if nil != err {
		return &pb.BuyLandRecordReply{Status: "地址不存在用户"}, nil
	}
	if nil == user {
		return &pb.BuyLandRecordReply{Status: "地址不存在用户"}, nil
	}

	landBuy, err = ac.userRepo.GetBuyLandById(ctx)
	if nil != err {
		return &pb.BuyLandRecordReply{Status: "不存在拍卖信息"}, nil
	}

	now := time.Now().Unix()
	if nil == landBuy || uint64(now) >= landBuy.Limit {
		return &pb.BuyLandRecordReply{Status: "ok", List: res}, nil
	}

	landRecord, err = ac.userRepo.GetAllBuyLandRecords(ctx, landBuy.ID)
	if nil != err {
		return &pb.BuyLandRecordReply{Status: "记录错误"}, nil
	}

	if 0 == len(landRecord) {
		return &pb.BuyLandRecordReply{Status: "ok", List: res}, nil
	}

	userIds := make([]uint64, 0)
	for _, v := range landRecord {
		userIds = append(userIds, v.UserID)
	}

	var (
		usersMap map[uint64]*User
	)
	usersMap, err = ac.userRepo.GetUserByUserIds(ctx, userIds)
	if nil != err {
		return &pb.BuyLandRecordReply{Status: "不存在用户信息"}, nil
	}
	if 0 == len(usersMap) {
		return &pb.BuyLandRecordReply{Status: "ok", List: res}, nil
	}

	for _, v := range landRecord {
		address := ""
		if _, ok := usersMap[v.UserID]; ok {
			address = usersMap[v.UserID].Address
		}

		res = append(res, &pb.BuyLandRecordReply_List{
			Address:   address,
			Amount:    v.Amount,
			CreatedAt: v.CreatedAt.Add(8 * time.Hour).Format("2006-01-02 15:04:05"),
		})
	}

	return &pb.BuyLandRecordReply{Status: "ok", List: res}, nil
}

var BuyLandR sync.Mutex

func (ac *AppUsecase) BuyLand(ctx context.Context, addreses string, req *pb.BuyLandRequest) (*pb.BuyLandReply, error) {
	var (
		user          *User
		err           error
		landBuy       *BuyLand
		landRecord    []*BuyLandRecord
		newLandRecord *BuyLandRecord
	)

	user, err = ac.userRepo.GetUserByAddress(ctx, addreses)
	if nil != err {
		return &pb.BuyLandReply{Status: "地址不存在用户"}, nil
	}
	if nil == user {
		return &pb.BuyLandReply{Status: "地址不存在用户"}, nil
	}

	if 1 == user.LockUse {
		return &pb.BuyLandReply{
			Status: "锁定用户",
		}, nil
	}

	BuyLandR.Lock()
	defer BuyLandR.Unlock()

	user, err = ac.userRepo.GetUserByAddress(ctx, addreses)
	if nil != err {
		return &pb.BuyLandReply{Status: "地址不存在用户"}, nil
	}
	if nil == user {
		return &pb.BuyLandReply{Status: "地址不存在用户"}, nil
	}

	landBuy, err = ac.userRepo.GetBuyLandById(ctx)
	if nil != err {
		return &pb.BuyLandReply{Status: "不存在拍卖信息"}, nil
	}

	if nil == landBuy {
		return &pb.BuyLandReply{Status: "暂无拍卖信息"}, nil
	}

	now := time.Now().Unix()
	if uint64(now) >= landBuy.Limit {
		return &pb.BuyLandReply{Status: "拍卖结束"}, nil
	}

	landRecord, err = ac.userRepo.GetAllBuyLandRecords(ctx, landBuy.ID)
	if nil != err {
		return &pb.BuyLandReply{Status: "记录错误"}, nil
	}

	if 1 <= len(landRecord) {
		newLandRecord = landRecord[0]
	}

	if 1 == req.SendBody.BuyType {
		if 0 >= landBuy.AmountTwo {
			return &pb.BuyLandReply{Status: "未设定一口价"}, nil
		}

		amountTmp := landBuy.AmountTwo
		if nil != newLandRecord {
			if newLandRecord.Amount >= amountTmp {
				return &pb.BuyLandReply{Status: "当前最高价已经高于一口价"}, nil
			}
		}

		if user.Giw < amountTmp {
			return &pb.BuyLandReply{Status: "giw余额不足"}, nil
		}

		if err = ac.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
			err = ac.userRepo.CreateBuyLandRecordOne(ctx, &BuyLandRecord{
				BuyLandID: landBuy.ID,
				Amount:    float64(amountTmp),
				Status:    3,
				UserID:    user.ID,
			})
			if nil != err {
				return err
			}

			err = ac.userRepo.CreateNotice(
				ctx,
				user.ID,
				"您参与了拍卖土地，使用"+fmt.Sprintf("%.2f", amountTmp)+"GIW",
				"You've used "+fmt.Sprintf("%.2f", amountTmp)+" GIW to auction",
			)
			if nil != err {
				return err
			}

			return nil
		}); nil != err {
			return &pb.BuyLandReply{
				Status: "竞拍失败",
			}, nil
		}
	} else {
		amountTmp := req.SendBody.Amount

		if nil != newLandRecord {
			if newLandRecord.Amount >= float64(amountTmp) {
				return &pb.BuyLandReply{Status: "出价低于当前最高价"}, nil
			}
		} else {
			if landBuy.Amount >= float64(amountTmp) {
				return &pb.BuyLandReply{Status: "出价低于当前竞拍价"}, nil
			}
		}

		if uint64(user.Giw) < amountTmp {
			return &pb.BuyLandReply{Status: "git余额不足"}, nil
		}

		if err = ac.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
			err = ac.userRepo.CreateBuyLandRecord(ctx, uint64(now)+1200, &BuyLandRecord{
				BuyLandID: landBuy.ID,
				Amount:    float64(amountTmp),
				Status:    1,
				UserID:    user.ID,
			})
			if nil != err {
				return err
			}

			err = ac.userRepo.CreateNotice(
				ctx,
				user.ID,
				"您参与了拍卖土地，使用"+strconv.FormatUint(amountTmp, 10)+"GIW",
				"You've used "+strconv.FormatUint(amountTmp, 10)+" GIW to auction",
			)
			if nil != err {
				return err
			}

			return nil
		}); nil != err {
			return &pb.BuyLandReply{
				Status: "竞拍失败",
			}, nil
		}
	}

	return &pb.BuyLandReply{
		Status: "ok",
	}, nil
}
