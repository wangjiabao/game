package biz

import (
	"context"
	"crypto/rand"
	pb "game/api/app/v1"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"math/big"
	"strconv"
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
			countBox int64
		)
		countBox, err = ac.userRepo.GetBoxRecordCount(ctx, boxNum)
		if nil != err {
			return &pb.UserInfoReply{
				Status: "盲盒错误查询",
			}, nil
		}

		boxSellNum = uint64(countBox)
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
	status := []uint64{0, 1, 2, 3, 4, 5}
	lands, err = ac.userRepo.GetLandByUserID(ctx, user.ID, status, nil)
	if nil != err {
		return &pb.UserLandReply{
			Status: "不存在用户",
		}, nil
	}

	for _, v := range lands {
		res = append(res, &pb.UserLandReply_List{
			Id:        v.ID,
			Level:     v.Level,
			Health:    v.MaxHealth,
			Status:    v.Status,
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
	landStatus := []uint64{3, 4}
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

	status := []uint64{1, 2, 3}
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

	for _, vLand := range lands {
		plantUserAddressTmp := ""

		if _, ok := landUserUse[vLand.ID]; ok {
			if _, ok2 := usersMap[landUserUse[vLand.ID].UserId]; ok2 {
				plantUserAddressTmp = usersMap[landUserUse[vLand.ID].UserId].Address
			}

			rewardTmp := float64(0)
			statusTmp := uint64(1)
			if landUserUse[vLand.ID].OverTime <= uint64(time.Now().Unix()) {
				rewardTmp = landUserUse[vLand.ID].OutNum
			} else {
				if 2 == landUserUse[vLand.ID].StopStatus {
					statusTmp = 3
				} else if 2 == landUserUse[vLand.ID].InsectStatus {
					statusTmp = 2
				}
			}

			res = append(res, &pb.UserIndexListReply_List{
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
				CurrentTime:      uint64(time.Now().Unix()),
				Status:           statusTmp,
				Reward:           rewardTmp,
				PlantUserAddress: plantUserAddressTmp,
			})
		} else {
			res = append(res, &pb.UserIndexListReply_List{
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
			})
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
