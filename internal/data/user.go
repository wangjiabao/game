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
	ID               uint64    `gorm:"primarykey;type:int"`
	Address          string    `gorm:"type:varchar(100);not null"`
	Level            uint64    `gorm:"type:int;not null"`
	Giw              float64   `gorm:"type:decimal(65,20);not null"`
	GiwAdd           float64   `gorm:"type:decimal(65,20);not null"`
	Git              float64   `gorm:"type:decimal(65,20);not null"`
	Total            float64   `gorm:"type:decimal(65,20);not null"`
	TotalOne         float64   `gorm:"type:decimal(65,20);not null"`
	TotalTwo         float64   `gorm:"type:decimal(65,20);not null"`
	TotalThree       float64   `gorm:"type:decimal(65,20);not null"`
	RewardOne        float64   `gorm:"type:decimal(65,20);not null"`
	RewardTwo        float64   `gorm:"type:decimal(65,20);not null"`
	RewardThree      float64   `gorm:"type:decimal(65,20);not null"`
	RewardTwoOne     float64   `gorm:"type:decimal(65,20);not null"`
	RewardTwoTwo     float64   `gorm:"type:decimal(65,20);not null"`
	RewardTwoThree   float64   `gorm:"type:decimal(65,20);not null"`
	RewardThreeOne   float64   `gorm:"type:decimal(65,20);not null"`
	RewardThreeTwo   float64   `gorm:"type:decimal(65,20);not null"`
	RewardThreeThree float64   `gorm:"type:decimal(65,20);not null"`
	CreatedAt        time.Time `gorm:"type:datetime;not null"`
	UpdatedAt        time.Time `gorm:"type:datetime;not null"`
}

type UserRecommend struct {
	ID            uint64    `gorm:"primarykey;type:int"`
	UserId        uint64    `gorm:"type:int;not null"`
	RecommendCode string    `gorm:"type:varchar(1000);not null"`
	CreatedAt     time.Time `gorm:"type:datetime;not null"`
	UpdatedAt     time.Time `gorm:"type:datetime;not null"`
}

type Config struct {
	ID        int64     `gorm:"primarykey;type:int"`
	Name      string    `gorm:"type:varchar(45);not null"`
	KeyName   string    `gorm:"type:varchar(45);not null"`
	Value     string    `gorm:"type:varchar(1000);not null"`
	CreatedAt time.Time `gorm:"type:datetime;not null"`
	UpdatedAt time.Time `gorm:"type:datetime;not null"`
}

type SkateGit struct {
	ID        uint64    `gorm:"primarykey;type:int"`
	UserId    uint64    `gorm:"type:int;not null"`
	Status    uint64    `gorm:"type:int;not null"`
	Amount    float64   `gorm:"type:decimal(65,20);not null"`
	Reward    float64   `gorm:"type:decimal(65,20);not null"`
	CreatedAt time.Time `gorm:"type:datetime;not null"`
	UpdatedAt time.Time `gorm:"type:datetime;not null"`
}

type BoxRecord struct {
	ID        uint64    `gorm:"primarykey;type:int"`
	UserId    uint64    `gorm:"type:int;not null"`
	Num       uint64    `gorm:"type:int;not null"`
	GoodId    uint64    `gorm:"type:int;not null"`
	GoodType  uint64    `gorm:"type:int;not null"`
	Content   string    `gorm:"type:varchar(200);not null"`
	CreatedAt time.Time `gorm:"type:datetime;not null"`
	UpdatedAt time.Time `gorm:"type:datetime;not null"`
}

type SkateGet struct {
	ID        uint64    `gorm:"primarykey;type:int"`
	UserId    uint64    `gorm:"type:int;not null"`
	Status    uint64    `gorm:"type:int;not null"`
	SkateRate float64   `gorm:"type:decimal(65,20);not null"`
	CreatedAt time.Time `gorm:"type:datetime;not null"`
	UpdatedAt time.Time `gorm:"type:datetime;not null"`
}

type SkateGetTotal struct {
	ID        uint64    `gorm:"primarykey;type:int"`
	Amount    float64   `gorm:"type:decimal(65,20);not null"`
	CreatedAt time.Time `gorm:"type:datetime;not null"`
	UpdatedAt time.Time `gorm:"type:datetime;not null"`
}

type Reward struct {
	ID        uint64    `gorm:"primarykey;type:int"`
	UserId    uint64    `gorm:"type:int;not null"`
	reason    uint64    `gorm:"type:int;not null"`
	One       uint64    `gorm:"type:int;not null"`
	Two       uint64    `gorm:"type:int;not null"`
	Three     float64   `gorm:"type:decimal(65,20);not null"`
	Amount    float64   `gorm:"type:decimal(65,20);not null"`
	CreatedAt time.Time `gorm:"type:datetime;not null"`
	UpdatedAt time.Time `gorm:"type:datetime;not null"`
}

type Seed struct {
	ID           uint64    `gorm:"primarykey;type:int"`
	UserId       uint64    `gorm:"type:int;not null;comment:用户id"`
	SeedId       uint64    `gorm:"type:int;not null;comment:种子信息id"`
	Name         string    `gorm:"type:varchar(45);not null;default:'1';comment:名字"`
	OutAmount    float64   `gorm:"type:decimal(65,20);not null;default:0.00000000000000000000;comment:产出数量"`
	OutOverTime  uint64    `gorm:"type:int;not null;default:0;comment:成熟时间"`
	OutMaxAmount float64   `gorm:"type:decimal(65,20);not null;default:0.00000000000000000000;comment:最大产出数量"`
	OutMinAmount float64   `gorm:"type:decimal(65,20);not null;default:0.00000000000000000000;comment:最小产出数量"`
	Status       uint64    `gorm:"type:int;not null;default:0;comment:状态：0未使用，1使用，出售中4，已售出5，不可用6"`
	SellAmount   float64   `gorm:"type:decimal(65,20);not null;default:0.00000000000000000000;"`
	CreatedAt    time.Time `gorm:"type:datetime;not null"`
	UpdatedAt    time.Time `gorm:"type:datetime;not null"`
}

type Land struct {
	ID             uint64    `gorm:"primarykey;type:int"`
	UserId         uint64    `gorm:"type:int;not null;default:0;comment:用户id"`
	Level          uint64    `gorm:"type:int;not null;default:1;comment:级别"`
	OutPutRate     float64   `gorm:"type:decimal(65,20);not null;default:100.00000000000000000000;comment:增产率"`
	RentOutPutRate float64   `gorm:"type:decimal(65,20);not null;default:0.00000000000000000000;comment:出租产出比率"`
	MaxHealth      uint64    `gorm:"type:int;not null;default:100;comment:最大可消耗肥沃度"`
	PerHealth      uint64    `gorm:"type:int;not null;default:10;comment:每次消耗肥沃度"`
	LimitDate      uint64    `gorm:"type:int;not null;default:0;comment:使用期限"`
	Status         uint64    `gorm:"type:int;not null;default:0;comment:状态"`
	LocationNum    uint64    `gorm:"type:int;not null;default:0;comment:首页位置"`
	One            uint64    `gorm:"type:int;not null;default:1;comment:可出租"`
	Two            uint64    `gorm:"type:int;not null;default:1;comment:可合成"`
	Three          uint64    `gorm:"type:int;not null;default:1;comment:可出售"`
	SellAmount     float64   `gorm:"type:decimal(65,20);not null;default:0.00000000000000000000;"`
	CreatedAt      time.Time `gorm:"type:datetime;not null"`
	UpdatedAt      time.Time `gorm:"type:datetime;not null"`
}

type LandUserUse struct {
	ID           uint64    `gorm:"primarykey;type:int"`
	LandId       uint64    `gorm:"type:int;not null;comment:用户土地id"`
	Level        uint64    `gorm:"type:int;not null;comment:级别"`
	UserId       uint64    `gorm:"type:int;not null;comment:使用用户id"`
	OwnerUserId  uint64    `gorm:"type:int;not null;comment:拥有者用户id"`
	SeedId       uint64    `gorm:"type:int;not null;comment:种子id"`
	SeedTypeId   uint64    `gorm:"type:int;not null;"`
	Status       uint64    `gorm:"type:int;not null;default:1;comment:状态"`
	BeginTime    uint64    `gorm:"type:int;not null;default:0;comment:开始时间戳"`
	TotalTime    uint64    `gorm:"type:int;not null;default:0;comment:成熟总时长"`
	OverTime     uint64    `gorm:"type:int;not null;default:0;comment:结束时间戳"`
	OutMaxNum    float64   `gorm:"type:decimal(65,20);not null;default:0.00000000000000000000;comment:最大产出"`
	OutNum       float64   `gorm:"type:decimal(65,20);not null;default:0.00000000000000000000;comment:已产出"`
	InsectStatus uint64    `gorm:"type:int;not null;default:1;comment:虫子状态"`
	OutSubNum    float64   `gorm:"type:decimal(65,20);not null;default:0.00000000000000000000;comment:减产数"`
	StealNum     float64   `gorm:"type:decimal(65,20);not null;default:0.00000000000000000000;comment:被偷走总数"`
	StopStatus   uint64    `gorm:"type:int;not null;default:1;comment:生长状态"`
	StopTime     uint64    `gorm:"type:int;not null;default:0;comment:暂停时间"`
	SubTime      uint64    `gorm:"type:int;not null;default:0;comment:加速总时长"`
	UseChan      uint64    `gorm:"type:int;not null;default:0;comment:使用铲子次数"`
	CreatedAt    time.Time `gorm:"type:datetime;not null"`
	UpdatedAt    time.Time `gorm:"type:datetime;not null"`
}

type ExchangeRecord struct {
	ID        uint64    `gorm:"primarykey;type:int"`
	UserId    int64     `gorm:"type:int;not null;comment:用户ID"`
	Git       float64   `gorm:"type:decimal(65,20);not null;comment:git数量"`
	Giw       float64   `gorm:"type:decimal(65,20);not null;comment:giw数量"`
	Fee       float64   `gorm:"type:decimal(65,20);not null;comment:手续费"`
	CreatedAt time.Time `gorm:"type:datetime;default:null"`
	UpdatedAt time.Time `gorm:"type:datetime;default:null"`
}

type Market struct {
	ID        uint64    `gorm:"primarykey;type:int"`
	UserId    uint64    `gorm:"type:int;not null;comment:售卖人用户id"`
	GoodId    uint64    `gorm:"type:int;not null;comment:上级的商品各个表中的id"`
	GoodType  int       `gorm:"type:int;not null;comment:商品类型：1土地，2种子，3道具"`
	Amount    float64   `gorm:"type:decimal(65,20);not null;default:0.00000000000000000000;comment:售价"`
	Status    int       `gorm:"type:int;not null;default:0;comment:状态：0下架，1上架，2已出售"`
	GetUserId uint64    `gorm:"type:int;not null;default:0;comment:购买人id"`
	CreatedAt time.Time `gorm:"type:datetime;not null"`
	UpdatedAt time.Time `gorm:"type:datetime;not null"`
}

type Notice struct {
	ID            uint64    `gorm:"primarykey;type:int"`
	UserId        uint64    `gorm:"type:int;not null;comment:用户id"`
	NoticeContent string    `gorm:"type:varchar(500);not null;comment:消息内容"`
	CreatedAt     time.Time `gorm:"type:datetime;not null"`
	UpdatedAt     time.Time `gorm:"type:datetime;not null"`
}

type Prop struct {
	ID         uint64    `gorm:"primarykey;type:int"`
	UserId     uint64    `gorm:"type:int;not null;comment:用户id"`
	Status     int       `gorm:"type:int;not null;default:1;comment:状态：未使用1，使用中2，已使用3，出售中4，已出售5，不可用11"`
	PropType   int       `gorm:"type:int;not null;comment:道具类型：1化肥，2铲子，3水，4除虫剂，5手套，6盲盒"`
	OneOne     int       `gorm:"type:int;not null;default:20;comment:化肥土地肥沃度增加量"`
	OneTwo     int       `gorm:"type:int;not null;default:14400;comment:化肥植物加速成熟时间戳"`
	TwoOne     int       `gorm:"type:int;not null;default:7;comment:铲子最大试用次数"`
	TwoTwo     float64   `gorm:"type:decimal(65,20);not null;default:20.00000000000000000000;comment:铲子使用之后获得百分比"`
	ThreeOne   int       `gorm:"type:int;not null;default:1;comment:水最大试用次数"`
	FourOne    int       `gorm:"type:int;not null;default:10;comment:除虫剂最大试用次数"`
	FiveOne    int       `gorm:"type:int;not null;default:20;comment:手套最大可偷植物次数"`
	SellAmount float64   `gorm:"type:decimal(65,20);not null;default:0.00000000000000000000;"`
	CreatedAt  time.Time `gorm:"type:datetime;not null"`
	UpdatedAt  time.Time `gorm:"type:datetime;not null"`
}

type StakeGet struct {
	ID        uint64    `gorm:"primarykey;type:int"`
	UserId    uint64    `gorm:"type:int;not null;comment:用户id"`
	Status    int       `gorm:"type:int;not null;default:1;comment:状态：1质押，2已解压"`
	StakeRate float64   `gorm:"type:decimal(65,18);not null;comment:质押比率"`
	CreatedAt time.Time `gorm:"type:datetime;not null"`
	UpdatedAt time.Time `gorm:"type:datetime;not null"`
}

type StakeGetPlayRecord struct {
	ID        uint64    `gorm:"primarykey;type:int"`
	UserId    uint64    `gorm:"type:int;not null;comment:用户id"`
	Amount    float64   `gorm:"type:decimal(65,18);not null;default:0.0;comment:玩金额"`
	Reward    float64   `gorm:"type:decimal(65,18);not null;default:0.0;comment:收益"`
	Status    int       `gorm:"type:int;not null;default:0;comment:状态：1赢了，2输"`
	CreatedAt time.Time `gorm:"type:datetime;not null"`
	UpdatedAt time.Time `gorm:"type:datetime;not null"`
}

type StakeGetRecord struct {
	ID        uint64    `gorm:"primarykey;type:int"`
	UserId    uint64    `gorm:"type:int;not null;comment:用户id"`
	Amount    float64   `gorm:"type:decimal(65,18);not null;default:0.0;comment:操作金额"`
	StakeRate float64   `gorm:"type:decimal(65,18);not null;default:0.0;comment:比率"`
	Total     float64   `gorm:"type:decimal(65,18);not null;default:0.0;comment:操作后金额"`
	StakeType int       `gorm:"type:int;not null;default:0;comment:操作类型：1质押，2解压"`
	CreatedAt time.Time `gorm:"type:datetime;not null"`
	UpdatedAt time.Time `gorm:"type:datetime;not null"`
}

type StakeGetTotal struct {
	ID        uint64    `gorm:"primarykey;type:int"`
	Amount    float64   `gorm:"type:decimal(65,18);not null;default:0.0;comment:总数"`
	CreatedAt time.Time `gorm:"type:datetime;not null"`
	UpdatedAt time.Time `gorm:"type:datetime;not null"`
}

type StakeGit struct {
	ID        uint64    `gorm:"primarykey;type:int"`
	UserID    uint64    `gorm:"type:int;not null;comment:用户id"`
	Amount    float64   `gorm:"type:decimal(65,18);not null;default:0.0;comment:质押金额"`
	Reward    float64   `gorm:"type:decimal(65,18);not null;default:0.0;comment:收益总数"`
	Status    int       `gorm:"type:int;not null;default:0;comment:状态：1质押，0解压"`
	CreatedAt time.Time `gorm:"type:datetime;not null"`
	UpdatedAt time.Time `gorm:"type:datetime;not null"`
}

type StakeGitRecord struct {
	ID        uint64    `gorm:"primarykey;type:int"`
	UserID    uint64    `gorm:"type:int;not null;comment:用户id"`
	Amount    float64   `gorm:"type:decimal(65,18);not null;default:0.0;comment:金额"`
	StakeType int       `gorm:"type:int;not null;default:0;comment:操作类型：1质押，2解压"`
	CreatedAt time.Time `gorm:"type:datetime;not null"`
	UpdatedAt time.Time `gorm:"type:datetime;not null"`
}

type Withdraw struct {
	ID        uint64    `gorm:"primarykey;type:int;comment:主键"`
	UserID    uint64    `gorm:"type:int;not null;comment:用户id"`
	Amount    uint64    `gorm:"type:bigint(20);not null;comment:金额"`
	RelAmount uint64    `gorm:"type:bigint(20);not null;comment:实际提现金额"`
	Status    string    `gorm:"type:varchar(45);not null;default:'default';comment:状态"`
	CreatedAt time.Time `gorm:"type:datetime;not null"`
	UpdatedAt time.Time `gorm:"type:datetime;not null"`
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

// GetAllUsers .
func (u *UserRepo) GetAllUsers(ctx context.Context) ([]*biz.User, error) {
	var users []*User

	res := make([]*biz.User, 0)
	if err := u.data.DB(ctx).Table("user").Order("id asc").Find(&users).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, nil
		}

		return nil, errors.New(500, "USER ERROR", err.Error())
	}

	for _, user := range users {
		res = append(res, &biz.User{
			ID:               user.ID,
			Address:          user.Address,
			Level:            user.Level,
			Giw:              user.Giw,
			GiwAdd:           user.GiwAdd,
			Git:              user.Git,
			Total:            user.Total,
			TotalOne:         user.TotalOne,
			TotalTwo:         user.TotalTwo,
			TotalThree:       user.TotalThree,
			CreatedAt:        user.CreatedAt,
			UpdatedAt:        user.UpdatedAt,
			RewardOne:        user.RewardOne,
			RewardTwo:        user.RewardTwo,
			RewardThree:      user.RewardThree,
			RewardTwoOne:     user.RewardTwoOne,
			RewardTwoTwo:     user.RewardTwoTwo,
			RewardTwoThree:   user.RewardTwoThree,
			RewardThreeOne:   user.RewardThreeOne,
			RewardThreeTwo:   user.RewardThreeTwo,
			RewardThreeThree: user.RewardThreeThree,
		})
	}

	return res, nil
}

// GetUserByUserIds .
func (u *UserRepo) GetUserByUserIds(ctx context.Context, userIds []uint64) (map[uint64]*biz.User, error) {
	var users []*User

	res := make(map[uint64]*biz.User, 0)
	if err := u.data.DB(ctx).Table("user").Where("id IN (?)", userIds).Find(&users).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, nil
		}

		return nil, errors.New(500, "USER ERROR", err.Error())
	}

	for _, user := range users {
		res[user.ID] = &biz.User{
			ID:               user.ID,
			Address:          user.Address,
			Level:            user.Level,
			Giw:              user.Giw,
			GiwAdd:           user.GiwAdd,
			Git:              user.Git,
			Total:            user.Total,
			TotalOne:         user.TotalOne,
			TotalTwo:         user.TotalTwo,
			TotalThree:       user.TotalThree,
			CreatedAt:        user.CreatedAt,
			UpdatedAt:        user.UpdatedAt,
			RewardOne:        user.RewardOne,
			RewardTwo:        user.RewardTwo,
			RewardThree:      user.RewardThree,
			RewardTwoOne:     user.RewardTwoOne,
			RewardTwoTwo:     user.RewardTwoTwo,
			RewardTwoThree:   user.RewardTwoThree,
			RewardThreeOne:   user.RewardThreeOne,
			RewardThreeTwo:   user.RewardThreeTwo,
			RewardThreeThree: user.RewardThreeThree,
		}
	}
	return res, nil
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
		ID:               user.ID,
		Address:          user.Address,
		Level:            user.Level,
		Giw:              user.Giw,
		GiwAdd:           user.GiwAdd,
		Git:              user.Git,
		Total:            user.Total,
		TotalOne:         user.TotalOne,
		TotalTwo:         user.TotalTwo,
		TotalThree:       user.TotalThree,
		RewardOne:        user.RewardOne,
		RewardTwo:        user.RewardTwo,
		RewardThree:      user.RewardThree,
		RewardTwoOne:     user.RewardTwoOne,
		RewardTwoTwo:     user.RewardTwoTwo,
		RewardTwoThree:   user.RewardTwoThree,
		RewardThreeOne:   user.RewardThreeOne,
		RewardThreeTwo:   user.RewardThreeTwo,
		RewardThreeThree: user.RewardThreeThree,
		CreatedAt:        user.CreatedAt,
		UpdatedAt:        user.UpdatedAt,
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

// GetUserRecommendByCode .
func (u *UserRepo) GetUserRecommendByCode(ctx context.Context, code string) ([]*biz.UserRecommend, error) {
	var (
		userRecommends []*UserRecommend
	)

	res := make([]*biz.UserRecommend, 0)
	instance := u.data.DB(ctx).Table("user_recommend").Where("recommend_code=?", code)
	if err := instance.Find(&userRecommends).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, nil
		}

		return nil, errors.New(500, "USER RECOMMEND ERROR", err.Error())
	}

	for _, userRecommend := range userRecommends {
		res = append(res, &biz.UserRecommend{
			ID:            userRecommend.ID,
			UserId:        userRecommend.UserId,
			RecommendCode: userRecommend.RecommendCode,
			CreatedAt:     userRecommend.CreatedAt,
			UpdatedAt:     userRecommend.UpdatedAt,
		})
	}

	return res, nil
}

// GetUserRecommends .
func (u *UserRepo) GetUserRecommends(ctx context.Context) ([]*biz.UserRecommend, error) {
	var userRecommends []*UserRecommend
	res := make([]*biz.UserRecommend, 0)
	if err := u.data.DB(ctx).Table("user_recommend").Find(&userRecommends).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, errors.NotFound("USER_RECOMMEND_NOT_FOUND", "user recommend not found")
		}

		return nil, errors.New(500, "USER RECOMMEND ERROR", err.Error())
	}

	for _, userRecommend := range userRecommends {
		res = append(res, &biz.UserRecommend{
			ID:            userRecommend.ID,
			UserId:        userRecommend.UserId,
			RecommendCode: userRecommend.RecommendCode,
			CreatedAt:     userRecommend.CreatedAt,
			UpdatedAt:     userRecommend.UpdatedAt,
		})
	}

	return res, nil
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
		ID:               user.ID,
		Address:          user.Address,
		Level:            user.Level,
		Giw:              user.Giw,
		GiwAdd:           user.GiwAdd,
		Git:              user.Git,
		Total:            user.Total,
		TotalOne:         user.TotalOne,
		TotalTwo:         user.TotalTwo,
		TotalThree:       user.TotalThree,
		RewardOne:        user.RewardOne,
		RewardTwo:        user.RewardTwo,
		RewardThree:      user.RewardThree,
		RewardTwoOne:     user.RewardTwoOne,
		RewardTwoTwo:     user.RewardTwoTwo,
		RewardTwoThree:   user.RewardTwoThree,
		RewardThreeOne:   user.RewardThreeOne,
		RewardThreeTwo:   user.RewardThreeTwo,
		RewardThreeThree: user.RewardThreeThree,
		CreatedAt:        user.CreatedAt,
		UpdatedAt:        user.UpdatedAt,
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

// GetConfigByKeys .
func (u *UserRepo) GetConfigByKeys(ctx context.Context, keys ...string) ([]*biz.Config, error) {
	var configs []*Config
	res := make([]*biz.Config, 0)
	if err := u.data.DB(ctx).Where("key_name IN (?)", keys).Table("config").Find(&configs).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, nil
		}

		return nil, errors.New(500, "Config ERROR", err.Error())
	}

	for _, config := range configs {
		res = append(res, &biz.Config{
			ID:      config.ID,
			KeyName: config.KeyName,
			Name:    config.Name,
			Value:   config.Value,
		})
	}

	return res, nil
}

// GetSkateGitByUserId .
func (u *UserRepo) GetSkateGitByUserId(ctx context.Context, userId uint64) (*biz.SkateGit, error) {
	var skateGit *SkateGit
	if err := u.data.DB(ctx).Where("user_id=?", userId).Table("stake_git").First(&skateGit).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, errors.New(500, "stake GIT ERROR", err.Error())
	}

	return &biz.SkateGit{
		ID:        skateGit.ID,
		UserId:    skateGit.UserId,
		Status:    skateGit.Status,
		Amount:    skateGit.Amount,
		Reward:    skateGit.Reward,
		CreatedAt: skateGit.CreatedAt,
		UpdatedAt: skateGit.UpdatedAt,
	}, nil
}

// GetBoxRecord .
func (u *UserRepo) GetBoxRecord(ctx context.Context, num uint64) ([]*biz.BoxRecord, error) {
	var boxRecord []*BoxRecord
	res := make([]*biz.BoxRecord, 0)
	if err := u.data.DB(ctx).Where("num=?", num).Table("box_record").Find(&boxRecord).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, nil
		}

		return nil, errors.New(500, "box_record ERROR", err.Error())
	}

	for _, v := range boxRecord {
		res = append(res, &biz.BoxRecord{
			ID:        v.ID,
			UserId:    v.UserId,
			Num:       v.Num,
			GoodId:    v.GoodId,
			GoodType:  v.GoodType,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
			Content:   v.Content,
		})
	}

	return res, nil
}

// GetUserBoxRecord .
func (u *UserRepo) GetUserBoxRecord(ctx context.Context, userId, num uint64, b *biz.Pagination) ([]*biz.BoxRecord, error) {
	var boxRecord []*BoxRecord
	res := make([]*biz.BoxRecord, 0)
	instance := u.data.DB(ctx).Where("user_id=?", userId).Table("box_record")

	if 0 < num {
		instance = instance.Where("num=?", num)
	}

	if err := instance.Order("id desc").Scopes(Paginate(b.PageNum, b.PageSize)).Find(&boxRecord).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, nil
		}

		return nil, errors.New(500, "box_record ERROR", err.Error())
	}

	for _, v := range boxRecord {
		res = append(res, &biz.BoxRecord{
			ID:        v.ID,
			UserId:    v.UserId,
			Num:       v.Num,
			GoodId:    v.GoodId,
			GoodType:  v.GoodType,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
			Content:   v.Content,
		})
	}

	return res, nil
}

// GetUserBoxRecordOpen .
func (u *UserRepo) GetUserBoxRecordOpen(ctx context.Context, userId, num uint64, open bool, b *biz.Pagination) ([]*biz.BoxRecord, error) {
	var boxRecord []*BoxRecord
	res := make([]*biz.BoxRecord, 0)
	instance := u.data.DB(ctx).Where("user_id=?", userId).Table("box_record")

	if 0 < num {
		instance = instance.Where("num=?", num)
	}

	if open {
		instance = instance.Where("good_id > ?", 0)
	} else {
		instance = instance.Where("good_id=?", 0)
	}

	if nil != b {
		instance = instance.Scopes(Paginate(b.PageNum, b.PageSize))
	}

	if err := instance.Order("id desc").Find(&boxRecord).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, nil
		}

		return nil, errors.New(500, "box_record ERROR", err.Error())
	}

	for _, v := range boxRecord {
		res = append(res, &biz.BoxRecord{
			ID:        v.ID,
			UserId:    v.UserId,
			Num:       v.Num,
			GoodId:    v.GoodId,
			GoodType:  v.GoodType,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
			Content:   v.Content,
		})
	}

	return res, nil
}

func (u *UserRepo) GetBoxRecordCount(ctx context.Context, num uint64) (int64, error) {
	var count int64

	if err := u.data.DB(ctx).Table("box_record").
		Where("num = ?", num).
		Count(&count).Error; err != nil {
		return 0, errors.New(500, "GetBoxRecordCount ERROR", err.Error())
	}

	return count, nil
}

// GetSkateGetTotal .
func (u *UserRepo) GetSkateGetTotal(ctx context.Context) (*biz.SkateGetTotal, error) {
	var skateGetTotal *SkateGetTotal
	if err := u.data.DB(ctx).Table("stake_get_total").First(&skateGetTotal).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, errors.New(500, "stake_get_total ERROR", err.Error())
	}

	return &biz.SkateGetTotal{
		ID:        skateGetTotal.ID,
		Amount:    skateGetTotal.Amount,
		CreatedAt: skateGetTotal.CreatedAt,
		UpdatedAt: skateGetTotal.UpdatedAt,
	}, nil
}

// GetUserSkateGet .
func (u *UserRepo) GetUserSkateGet(ctx context.Context, userId uint64) (*biz.SkateGet, error) {
	var skateGet *SkateGet
	if err := u.data.DB(ctx).Table("stake_get").Where("user_id=?", userId).First(&skateGet).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, errors.New(500, "stake GET ERROR", err.Error())
	}

	return &biz.SkateGet{
		ID:        skateGet.ID,
		Status:    skateGet.Status,
		UserId:    skateGet.UserId,
		SkateRate: skateGet.SkateRate,
		CreatedAt: skateGet.CreatedAt,
		UpdatedAt: skateGet.UpdatedAt,
	}, nil
}

// GetUserRecommendCount .
func (u *UserRepo) GetUserRecommendCount(ctx context.Context, code string) (int64, error) {
	var count int64
	if err := u.data.DB(ctx).Table("user_recommend").Where("recommend_code=?", code).Count(&count).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return count, nil
		}

		return count, errors.New(500, "USER ERROR", err.Error())
	}

	return count, nil
}

// GetUserRecommendByCodePage .
func (u *UserRepo) GetUserRecommendByCodePage(ctx context.Context, code string, b *biz.Pagination) ([]*biz.UserRecommend, error) {
	var (
		userRecommends []*UserRecommend
	)

	res := make([]*biz.UserRecommend, 0)
	instance := u.data.DB(ctx).Table("user_recommend").
		Where("recommend_code=?", code).
		Order("id asc").
		Scopes(Paginate(b.PageNum, b.PageSize))
	if err := instance.Find(&userRecommends).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, nil
		}

		return nil, errors.New(500, "USER RECOMMEND ERROR", err.Error())
	}

	for _, userRecommend := range userRecommends {
		res = append(res, &biz.UserRecommend{
			ID:            userRecommend.ID,
			UserId:        userRecommend.UserId,
			RecommendCode: userRecommend.RecommendCode,
			CreatedAt:     userRecommend.CreatedAt,
			UpdatedAt:     userRecommend.UpdatedAt,
		})
	}

	return res, nil
}

// GetUserRewardByCodePage .
func (u *UserRepo) GetUserRewardPage(ctx context.Context, userId uint64, reason []uint64, b *biz.Pagination) ([]*biz.Reward, error) {
	var (
		rewards []*Reward
	)

	res := make([]*biz.Reward, 0)
	instance := u.data.DB(ctx).Table("reward").
		Where("user_id=?", userId).
		Where("reason in (?)", reason).
		Order("id desc").
		Scopes(Paginate(b.PageNum, b.PageSize))
	if err := instance.Find(&rewards).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, nil
		}

		return nil, errors.New(500, "GetUserRewardByCodePage ERROR", err.Error())
	}

	for _, v := range rewards {
		res = append(res, &biz.Reward{
			ID:        v.ID,
			UserId:    v.UserId,
			Amount:    v.Amount,
			One:       v.One,
			Two:       v.Two,
			Three:     v.Three,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		})
	}

	return res, nil
}

func (u *UserRepo) GetUserRewardPageCount(ctx context.Context, userId uint64, reason []uint64) (int64, error) {
	var count int64

	instance := u.data.DB(ctx).Table("reward").
		Where("user_id = ?", userId).
		Where("reason IN (?)", reason)

	if err := instance.Count(&count).Error; err != nil {
		return 0, errors.New(500, "GetUserRewardPageCount ERROR", err.Error())
	}

	return count, nil
}

// GetSeedByUserID 查询用户的种子数据
func (u *UserRepo) GetSeedByUserID(ctx context.Context, userID uint64, status []uint64, b *biz.Pagination) ([]*biz.Seed, error) {
	var (
		seeds []*Seed
	)

	res := make([]*biz.Seed, 0)
	instance := u.data.DB(ctx).Table("seed")

	if 0 < userID {
		instance = instance.Where("user_id = ?", userID)
	}

	instance = instance.Where("status in (?)", status).Order("id asc")

	if nil != b {
		instance = instance.Scopes(Paginate(b.PageNum, b.PageSize))
	}

	if err := instance.Find(&seeds).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, nil
		}

		return nil, errors.New(500, "SEED ERROR", err.Error())
	}

	for _, seed := range seeds {
		res = append(res, &biz.Seed{
			ID:           seed.ID,
			UserId:       seed.UserId,
			SeedId:       seed.SeedId,
			Name:         seed.Name,
			OutAmount:    seed.OutAmount,
			OutOverTime:  seed.OutOverTime,
			OutMaxAmount: seed.OutMaxAmount,
			OutMinAmount: seed.OutMinAmount,
			Status:       seed.Status,
			CreatedAt:    seed.CreatedAt,
			UpdatedAt:    seed.UpdatedAt,
			SellAmount:   seed.SellAmount,
		})
	}

	return res, nil
}

// GetSeedByExUserID 查询用户的种子数据
func (u *UserRepo) GetSeedByExUserID(ctx context.Context, userID uint64, status []uint64, b *biz.Pagination) ([]*biz.Seed, error) {
	var (
		seeds []*Seed
	)

	res := make([]*biz.Seed, 0)
	instance := u.data.DB(ctx).Table("seed").Where("user_id != ?", userID).Where("status in (?)", status).Order("id asc")

	if nil != b {
		instance = instance.Scopes(Paginate(b.PageNum, b.PageSize))
	}

	if err := instance.Find(&seeds).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, nil
		}

		return nil, errors.New(500, "SEED ERROR", err.Error())
	}

	for _, seed := range seeds {
		res = append(res, &biz.Seed{
			ID:           seed.ID,
			UserId:       seed.UserId,
			SeedId:       seed.SeedId,
			Name:         seed.Name,
			OutAmount:    seed.OutAmount,
			OutOverTime:  seed.OutOverTime,
			OutMaxAmount: seed.OutMaxAmount,
			OutMinAmount: seed.OutMinAmount,
			Status:       seed.Status,
			CreatedAt:    seed.CreatedAt,
			UpdatedAt:    seed.UpdatedAt,
			SellAmount:   seed.SellAmount,
		})
	}

	return res, nil
}

// GetLandByUserID getLandByUserID
func (u *UserRepo) GetLandByUserID(ctx context.Context, userID uint64, status []uint64, b *biz.Pagination) ([]*biz.Land, error) {
	var (
		lands []*Land
	)

	res := make([]*biz.Land, 0)
	instance := u.data.DB(ctx).Table("land")

	if 0 < userID {
		instance = instance.Where("user_id = ?", userID)
	}

	instance = instance.Where("limit_date>=?", time.Now().Unix()).
		Where("status in (?)", status).
		Order("id asc")

	if nil != b {
		instance = instance.Scopes(Paginate(b.PageNum, b.PageSize))
	}

	if err := instance.Find(&lands).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, nil
		}

		return nil, errors.New(500, "LAND ERROR", err.Error())
	}

	for _, land := range lands {
		res = append(res, &biz.Land{
			ID:             land.ID,
			UserId:         land.UserId,
			Level:          land.Level,
			OutPutRate:     land.OutPutRate,
			RentOutPutRate: land.RentOutPutRate,
			MaxHealth:      land.MaxHealth,
			PerHealth:      land.PerHealth,
			LimitDate:      land.LimitDate,
			Status:         land.Status,
			LocationNum:    land.LocationNum,
			CreatedAt:      land.CreatedAt,
			UpdatedAt:      land.UpdatedAt,
			One:            land.One,
			Two:            land.Two,
			Three:          land.Three,
			SellAmount:     land.SellAmount,
		})
	}

	return res, nil
}

// GetLandByUserIDUsing getLandByUserIDUsing
func (u *UserRepo) GetLandByUserIDUsing(ctx context.Context, userID uint64, status []uint64) ([]*biz.Land, error) {
	var (
		lands []*Land
	)

	res := make([]*biz.Land, 0)
	instance := u.data.DB(ctx).Table("land").
		Where("user_id = ?", userID).
		Where("limit_date>=?", time.Now().Unix()).
		Where("status in (?)", status).
		Where("location_num >?", 0).
		Order("id asc")

	if err := instance.Find(&lands).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, nil
		}

		return nil, errors.New(500, "LAND ERROR", err.Error())
	}

	for _, land := range lands {
		res = append(res, &biz.Land{
			ID:             land.ID,
			UserId:         land.UserId,
			Level:          land.Level,
			OutPutRate:     land.OutPutRate,
			RentOutPutRate: land.RentOutPutRate,
			MaxHealth:      land.MaxHealth,
			PerHealth:      land.PerHealth,
			LimitDate:      land.LimitDate,
			Status:         land.Status,
			LocationNum:    land.LocationNum,
			CreatedAt:      land.CreatedAt,
			UpdatedAt:      land.UpdatedAt,
			One:            land.One,
			Two:            land.Two,
			Three:          land.Three,
			SellAmount:     land.SellAmount,
		})
	}

	return res, nil
}

// GetLandByExUserID getLandByExUserID
func (u *UserRepo) GetLandByExUserID(ctx context.Context, userID uint64, status []uint64, b *biz.Pagination) ([]*biz.Land, error) {
	var (
		lands []*Land
	)

	res := make([]*biz.Land, 0)
	instance := u.data.DB(ctx).Table("land").Where("user_id != ?", userID).Where("limit_date>=?", time.Now().Unix()).
		Where("status in (?)", status).
		Order("id asc")

	if nil != b {
		instance = instance.Scopes(Paginate(b.PageNum, b.PageSize))
	}

	if err := instance.Find(&lands).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, nil
		}

		return nil, errors.New(500, "LAND ERROR", err.Error())
	}

	for _, land := range lands {
		res = append(res, &biz.Land{
			ID:             land.ID,
			UserId:         land.UserId,
			Level:          land.Level,
			OutPutRate:     land.OutPutRate,
			RentOutPutRate: land.RentOutPutRate,
			MaxHealth:      land.MaxHealth,
			PerHealth:      land.PerHealth,
			LimitDate:      land.LimitDate,
			Status:         land.Status,
			LocationNum:    land.LocationNum,
			CreatedAt:      land.CreatedAt,
			UpdatedAt:      land.UpdatedAt,
			One:            land.One,
			Two:            land.Two,
			Three:          land.Three,
			SellAmount:     land.SellAmount,
		})
	}

	return res, nil
}

// GetLandUserUseByUserIDUseing getLandUserUseByUserIDUseing
func (u *UserRepo) GetLandUserUseByUserIDUseing(ctx context.Context, userID uint64, status uint64, b *biz.Pagination) ([]*biz.LandUserUse, error) {
	var (
		landUserUses []*LandUserUse
	)

	res := make([]*biz.LandUserUse, 0)
	instance := u.data.DB(ctx).Table("land_user_use").
		Where("user_id = ?", userID).
		Where("status = ?", status).
		Order("id desc")

	if nil != b {
		instance = instance.Scopes(Paginate(b.PageNum, b.PageSize))
	}

	if err := instance.Find(&landUserUses).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, nil
		}

		return nil, errors.New(500, "LAND USER USE ERROR", err.Error())
	}

	for _, landUserUse := range landUserUses {
		res = append(res, &biz.LandUserUse{
			ID:           landUserUse.ID,
			LandId:       landUserUse.LandId,
			Level:        landUserUse.Level,
			UserId:       landUserUse.UserId,
			OwnerUserId:  landUserUse.OwnerUserId,
			SeedId:       landUserUse.SeedId,
			SeedTypeId:   landUserUse.SeedTypeId,
			Status:       landUserUse.Status,
			BeginTime:    landUserUse.BeginTime,
			TotalTime:    landUserUse.TotalTime,
			OverTime:     landUserUse.OverTime,
			OutMaxNum:    landUserUse.OutMaxNum,
			OutNum:       landUserUse.OutNum,
			InsectStatus: landUserUse.InsectStatus,
			OutSubNum:    landUserUse.OutSubNum,
			StealNum:     landUserUse.StealNum,
			StopStatus:   landUserUse.StopStatus,
			StopTime:     landUserUse.StopTime,
			SubTime:      landUserUse.SubTime,
			UseChan:      landUserUse.UseChan,
			CreatedAt:    landUserUse.CreatedAt,
			UpdatedAt:    landUserUse.UpdatedAt,
		})
	}

	return res, nil
}

func (u *UserRepo) GetExchangeRecordsByUserID(ctx context.Context, userID uint64, b *biz.Pagination) ([]*biz.ExchangeRecord, error) {
	var (
		exchangeRecords []*ExchangeRecord
	)

	res := make([]*biz.ExchangeRecord, 0)
	instance := u.data.DB(ctx).Table("exchange_record").
		Where("user_id = ?", userID).
		Order("id desc")

	if nil != b {
		instance = instance.Scopes(Paginate(b.PageNum, b.PageSize))
	}

	if err := instance.Find(&exchangeRecords).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, nil
		}

		return nil, errors.New(500, "EXCHANGE RECORD ERROR", err.Error())
	}

	for _, exchangeRecord := range exchangeRecords {
		res = append(res, &biz.ExchangeRecord{
			ID:        exchangeRecord.ID,
			UserId:    exchangeRecord.UserId,
			Git:       exchangeRecord.Git,
			Giw:       exchangeRecord.Giw,
			Fee:       exchangeRecord.Fee,
			CreatedAt: exchangeRecord.CreatedAt,
			UpdatedAt: exchangeRecord.UpdatedAt,
		})
	}

	return res, nil
}

func (u *UserRepo) GetMarketRecordsByUserID(ctx context.Context, userID uint64, status uint64, b *biz.Pagination) ([]*biz.Market, error) {
	var (
		marketRecords []*Market
	)

	res := make([]*biz.Market, 0)
	instance := u.data.DB(ctx).Table("market").
		Where("user_id = ?", userID).
		Where("status in (?)", status).
		Order("id asc")

	if nil != b {
		instance = instance.Scopes(Paginate(b.PageNum, b.PageSize))
	}

	if err := instance.Find(&marketRecords).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, nil
		}

		return nil, errors.New(500, "MARKET RECORD ERROR", err.Error())
	}

	for _, marketRecord := range marketRecords {
		res = append(res, &biz.Market{
			ID:        marketRecord.ID,
			UserId:    marketRecord.UserId,
			GoodId:    marketRecord.GoodId,
			GoodType:  marketRecord.GoodType,
			Amount:    marketRecord.Amount,
			Status:    marketRecord.Status,
			GetUserId: marketRecord.GetUserId,
			CreatedAt: marketRecord.CreatedAt,
			UpdatedAt: marketRecord.UpdatedAt,
		})
	}

	return res, nil
}

func (u *UserRepo) GetNoticesByUserID(ctx context.Context, userID uint64, b *biz.Pagination) ([]*biz.Notice, error) {
	var (
		noticeRecords []*Notice
	)

	res := make([]*biz.Notice, 0)
	instance := u.data.DB(ctx).Table("notice").
		Where("user_id = ?", userID).
		Order("id desc")

	if nil != b {
		instance = instance.Scopes(Paginate(b.PageNum, b.PageSize))
	}

	if err := instance.Find(&noticeRecords).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, nil
		}

		return nil, errors.New(500, "NOTICE RECORD ERROR", err.Error())
	}

	for _, noticeRecord := range noticeRecords {
		res = append(res, &biz.Notice{
			ID:            noticeRecord.ID,
			UserId:        noticeRecord.UserId,
			NoticeContent: noticeRecord.NoticeContent,
			CreatedAt:     noticeRecord.CreatedAt,
			UpdatedAt:     noticeRecord.UpdatedAt,
		})
	}

	return res, nil
}

func (u *UserRepo) GetNoticesCountByUserID(ctx context.Context, userID uint64) (int64, error) {
	var count int64
	err := u.data.DB(ctx).Table("notice").
		Where("user_id = ?", userID).
		Count(&count).Error
	if err != nil {
		return 0, errors.New(500, "NOTICE COUNT ERROR", err.Error())
	}
	return count, nil
}

func (u *UserRepo) GetPropsByUserID(ctx context.Context, userID uint64, status []uint64, b *biz.Pagination) ([]*biz.Prop, error) {
	var (
		props []*Prop
	)

	res := make([]*biz.Prop, 0)
	instance := u.data.DB(ctx).Table("prop")

	if 0 < userID {
		instance = instance.Where("user_id = ?", userID)
	}

	instance = instance.Where("status in (?)", status).Order("id asc")

	if nil != b {
		instance = instance.Scopes(Paginate(b.PageNum, b.PageSize))
	}

	if err := instance.Find(&props).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, nil
		}

		return nil, errors.New(500, "PROP RECORD ERROR", err.Error())
	}

	for _, prop := range props {
		res = append(res, &biz.Prop{
			ID:         prop.ID,
			UserId:     prop.UserId,
			Status:     prop.Status,
			PropType:   prop.PropType,
			OneOne:     prop.OneOne,
			OneTwo:     prop.OneTwo,
			TwoOne:     prop.TwoOne,
			TwoTwo:     prop.TwoTwo,
			ThreeOne:   prop.ThreeOne,
			FourOne:    prop.FourOne,
			FiveOne:    prop.FiveOne,
			CreatedAt:  prop.CreatedAt,
			UpdatedAt:  prop.UpdatedAt,
			SellAmount: prop.SellAmount,
		})
	}

	return res, nil
}

func (u *UserRepo) GetPropsByExUserID(ctx context.Context, userID uint64, status []uint64, b *biz.Pagination) ([]*biz.Prop, error) {
	var (
		props []*Prop
	)

	res := make([]*biz.Prop, 0)
	instance := u.data.DB(ctx).Table("prop").Where("user_id != ?", userID).Where("status in (?)", status).Order("id asc")

	if nil != b {
		instance = instance.Scopes(Paginate(b.PageNum, b.PageSize))
	}

	if err := instance.Find(&props).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, nil
		}

		return nil, errors.New(500, "PROP RECORD ERROR", err.Error())
	}

	for _, prop := range props {
		res = append(res, &biz.Prop{
			ID:         prop.ID,
			UserId:     prop.UserId,
			Status:     prop.Status,
			PropType:   prop.PropType,
			OneOne:     prop.OneOne,
			OneTwo:     prop.OneTwo,
			TwoOne:     prop.TwoOne,
			TwoTwo:     prop.TwoTwo,
			ThreeOne:   prop.ThreeOne,
			FourOne:    prop.FourOne,
			FiveOne:    prop.FiveOne,
			CreatedAt:  prop.CreatedAt,
			UpdatedAt:  prop.UpdatedAt,
			SellAmount: prop.SellAmount,
		})
	}

	return res, nil
}

func (u *UserRepo) GetStakeGetsByUserID(ctx context.Context, userID uint64, b *biz.Pagination) ([]*biz.StakeGet, error) {
	var (
		stakeGets []*StakeGet
	)

	res := make([]*biz.StakeGet, 0)
	instance := u.data.DB(ctx).Table("stake_get").
		Where("user_id = ?", userID).
		Order("id desc")

	if nil != b {
		instance = instance.Scopes(Paginate(b.PageNum, b.PageSize))
	}

	if err := instance.Find(&stakeGets).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, nil
		}

		return nil, errors.New(500, "STAKE GET RECORD ERROR", err.Error())
	}

	for _, stakeGet := range stakeGets {
		res = append(res, &biz.StakeGet{
			ID:        stakeGet.ID,
			UserId:    stakeGet.UserId,
			Status:    stakeGet.Status,
			StakeRate: stakeGet.StakeRate,
			CreatedAt: stakeGet.CreatedAt,
			UpdatedAt: stakeGet.UpdatedAt,
		})
	}

	return res, nil
}

func (u *UserRepo) GetStakeGetPlayRecordCount(ctx context.Context, userID uint64, status uint64) (int64, error) {
	var count int64
	instance := u.data.DB(ctx).Table("stake_get_play_record").
		Where("user_id = ?", userID)

	if status > 0 {
		instance = instance.Where("status = ?", status)
	}

	err := instance.Count(&count).Error
	if err != nil {
		return 0, errors.New(500, "STAKE GET PLAY RECORD COUNT ERROR", err.Error())
	}
	return count, nil
}

func (u *UserRepo) GetStakeGetPlayRecordsByUserID(ctx context.Context, userID uint64, status uint64, b *biz.Pagination) ([]*biz.StakeGetPlayRecord, error) {
	var (
		records []*StakeGetPlayRecord
	)

	res := make([]*biz.StakeGetPlayRecord, 0)
	instance := u.data.DB(ctx).Table("stake_get_play_record").
		Where("user_id = ?", userID).
		Order("id desc")

	if 0 < status {
		instance = instance.Where("status=?", status)
	}

	if nil != b {
		instance = instance.Scopes(Paginate(b.PageNum, b.PageSize))
	}

	if err := instance.Find(&records).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, nil
		}

		return nil, errors.New(500, "STAKE GET PLAY RECORD ERROR", err.Error())
	}

	for _, record := range records {
		res = append(res, &biz.StakeGetPlayRecord{
			ID:        record.ID,
			UserId:    record.UserId,
			Amount:    record.Amount,
			Reward:    record.Reward,
			Status:    record.Status,
			CreatedAt: record.CreatedAt,
			UpdatedAt: record.UpdatedAt,
		})
	}

	return res, nil
}

func (u *UserRepo) GetStakeGetRecordsByUserID(ctx context.Context, userID int64, b *biz.Pagination) ([]*biz.StakeGetRecord, error) {
	var (
		records []*StakeGetRecord
	)

	res := make([]*biz.StakeGetRecord, 0)
	instance := u.data.DB(ctx).Table("stake_get_record").
		Where("user_id = ?", userID).
		Order("id desc")

	if nil != b {
		instance = instance.Scopes(Paginate(b.PageNum, b.PageSize))
	}

	if err := instance.Find(&records).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, nil
		}

		return nil, errors.New(500, "STAKE GET RECORD ERROR", err.Error())
	}

	for _, record := range records {
		res = append(res, &biz.StakeGetRecord{
			ID:        record.ID,
			UserId:    record.UserId,
			Amount:    record.Amount,
			StakeRate: record.StakeRate,
			Total:     record.Total,
			StakeType: record.StakeType,
			CreatedAt: record.CreatedAt,
			UpdatedAt: record.UpdatedAt,
		})
	}

	return res, nil
}

func (u *UserRepo) GetStakeGetTotal(ctx context.Context) (*biz.StakeGetTotal, error) {
	var total StakeGetTotal

	err := u.data.DB(ctx).Table("stake_get_total").First(&total).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, errors.New(500, "STAKE GET TOTAL ERROR", err.Error())
	}

	return &biz.StakeGetTotal{
		ID:        total.ID,
		Amount:    total.Amount,
		CreatedAt: total.CreatedAt,
		UpdatedAt: total.UpdatedAt,
	}, nil
}

func (u *UserRepo) GetStakeGitByUserID(ctx context.Context, userID int64) (*biz.StakeGit, error) {
	var stake StakeGit

	err := u.data.DB(ctx).Table("stake_git").
		Where("user_id = ?", userID).
		First(&stake).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, errors.New(500, "STAKE GIT ERROR", err.Error())
	}

	return &biz.StakeGit{
		ID:        stake.ID,
		UserID:    stake.UserID,
		Amount:    stake.Amount,
		Reward:    stake.Reward,
		Status:    stake.Status,
		CreatedAt: stake.CreatedAt,
		UpdatedAt: stake.UpdatedAt,
	}, nil
}

func (u *UserRepo) GetStakeGitRecordsByUserID(ctx context.Context, userID int64, b *biz.Pagination) ([]*biz.StakeGitRecord, error) {
	var (
		records []*StakeGitRecord
	)

	res := make([]*biz.StakeGitRecord, 0)
	instance := u.data.DB(ctx).Table("stake_git_record").
		Where("user_id = ?", userID).
		Order("id desc")

	if nil != b {
		instance = instance.Scopes(Paginate(b.PageNum, b.PageSize))
	}

	if err := instance.Find(&records).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, nil
		}
		return nil, errors.New(500, "STAKE GIT RECORD ERROR", err.Error())
	}

	for _, record := range records {
		res = append(res, &biz.StakeGitRecord{
			ID:        record.ID,
			UserID:    record.UserID,
			Amount:    record.Amount,
			StakeType: record.StakeType,
			CreatedAt: record.CreatedAt,
			UpdatedAt: record.UpdatedAt,
		})
	}

	return res, nil
}

func (u *UserRepo) GetWithdrawRecordsByUserID(ctx context.Context, userID int64, b *biz.Pagination) ([]*biz.Withdraw, error) {
	var (
		records []*Withdraw
	)

	res := make([]*biz.Withdraw, 0)
	instance := u.data.DB(ctx).Table("withdraw").
		Where("user_id = ?", userID).
		Order("id desc")

	if nil != b {
		instance = instance.Scopes(Paginate(b.PageNum, b.PageSize))
	}

	if err := instance.Find(&records).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, nil
		}
		return nil, errors.New(500, "WITHDRAW RECORD ERROR", err.Error())
	}

	for _, record := range records {
		res = append(res, &biz.Withdraw{
			ID:        record.ID,
			UserID:    record.UserID,
			Amount:    record.Amount,
			RelAmount: record.RelAmount,
			Status:    record.Status,
			CreatedAt: record.CreatedAt,
			UpdatedAt: record.UpdatedAt,
		})
	}

	return res, nil
}

func (u *UserRepo) GetUserOrderCount(ctx context.Context) (int64, error) {
	var count int64
	err := u.data.DB(ctx).Table("user").
		Where("git > ?", 0).
		Count(&count).Error

	if err != nil {
		return 0, errors.New(500, "USER COUNT ERROR", err.Error())
	}

	return count, nil
}

func (u *UserRepo) GetUserOrder(ctx context.Context, b *biz.Pagination) ([]*biz.User, error) {
	var users []*User

	res := make([]*biz.User, 0)
	instance := u.data.DB(ctx).Table("user").Where("git>?", 0).Order("git desc")

	if nil != b {
		instance = instance.Scopes(Paginate(b.PageNum, b.PageSize))
	}

	if err := instance.Find(&users).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, nil
		}

		return nil, errors.New(500, "USER ERROR", err.Error())
	}

	for _, user := range users {
		res = append(res, &biz.User{
			ID:               user.ID,
			Address:          user.Address,
			Level:            user.Level,
			Giw:              user.Giw,
			GiwAdd:           user.GiwAdd,
			Git:              user.Git,
			Total:            user.Total,
			TotalOne:         user.TotalOne,
			TotalTwo:         user.TotalTwo,
			TotalThree:       user.TotalThree,
			CreatedAt:        user.CreatedAt,
			UpdatedAt:        user.UpdatedAt,
			RewardOne:        user.RewardOne,
			RewardTwo:        user.RewardTwo,
			RewardThree:      user.RewardThree,
			RewardTwoOne:     user.RewardTwoOne,
			RewardTwoTwo:     user.RewardTwoTwo,
			RewardTwoThree:   user.RewardTwoThree,
			RewardThreeOne:   user.RewardThreeOne,
			RewardThreeTwo:   user.RewardThreeTwo,
			RewardThreeThree: user.RewardThreeThree,
		})
	}

	return res, nil
}

func (u *UserRepo) GetLandUserUseByLandIDsMapUsing(ctx context.Context, userId uint64, landIDs []uint64) (map[uint64]*biz.LandUserUse, error) {
	var landUserUses []*LandUserUse
	res := make(map[uint64]*biz.LandUserUse)

	instance := u.data.DB(ctx).Table("land_user_use").
		Where("owner_user_id = ?", userId).
		Where("land_id IN (?)", landIDs).
		Where("status = ?", 1).
		Order("id desc")

	if err := instance.Find(&landUserUses).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, nil
		}
		return nil, errors.New(500, "LAND USER USE ERROR", err.Error())
	}

	// 归类数据到 map
	for _, landUserUse := range landUserUses {
		res[landUserUse.LandId] = &biz.LandUserUse{
			ID:           landUserUse.ID,
			LandId:       landUserUse.LandId,
			Level:        landUserUse.Level,
			UserId:       landUserUse.UserId,
			OwnerUserId:  landUserUse.OwnerUserId,
			SeedId:       landUserUse.SeedId,
			SeedTypeId:   landUserUse.SeedTypeId,
			Status:       landUserUse.Status,
			BeginTime:    landUserUse.BeginTime,
			TotalTime:    landUserUse.TotalTime,
			OverTime:     landUserUse.OverTime,
			OutMaxNum:    landUserUse.OutMaxNum,
			OutNum:       landUserUse.OutNum,
			InsectStatus: landUserUse.InsectStatus,
			OutSubNum:    landUserUse.OutSubNum,
			StealNum:     landUserUse.StealNum,
			StopStatus:   landUserUse.StopStatus,
			StopTime:     landUserUse.StopTime,
			SubTime:      landUserUse.SubTime,
			UseChan:      landUserUse.UseChan,
			CreatedAt:    landUserUse.CreatedAt,
			UpdatedAt:    landUserUse.UpdatedAt,
		}
	}

	return res, nil
}
