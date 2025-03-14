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
	if err := u.data.DB(ctx).Where("user_id=?", userId).Table("skate_git").First(&skateGit).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, errors.New(500, "SKATE GIT ERROR", err.Error())
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
		})
	}

	return res, nil
}

// GetSkateGetTotal .
func (u *UserRepo) GetSkateGetTotal(ctx context.Context) (*biz.SkateGetTotal, error) {
	var skateGetTotal *SkateGetTotal
	if err := u.data.DB(ctx).Table("skate_get_total").First(&skateGetTotal).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, errors.New(500, "skate_get_total ERROR", err.Error())
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
	if err := u.data.DB(ctx).Table("skate_get").Where("user_id=?", userId).First(&skateGet).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, errors.New(500, "SKATE GET ERROR", err.Error())
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
