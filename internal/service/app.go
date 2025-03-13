package service

import (
	"context"
	"fmt"
	pb "game/api/app/v1"
	"game/internal/biz"
	"game/internal/conf"
	"game/internal/pkg/middleware/auth"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	jwt2 "github.com/golang-jwt/jwt/v4"
	"regexp"
	"time"
)

type AppService struct {
	pb.UnimplementedAppServer
	log *log.Helper
	ac  *biz.AppUsecase
	ca  *conf.Auth
}

func NewAppService(ac *biz.AppUsecase, logger log.Logger, ca *conf.Auth) *AppService {
	return &AppService{
		ac:  ac,
		log: log.NewHelper(logger),
		ca:  ca,
	}
}

var ethClient *ethclient.Client

func init() {
	var err error
	ethClient, err = ethclient.Dial("https://bsc-dataseed4.binance.org/")
	if err != nil {
		panic("eth client err")
	}
}

func addressCheck(addressParam string) (bool, error) {
	re := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")
	if !re.MatchString(addressParam) {
		return false, nil
	}

	var (
		err      error
		bytecode []byte
	)

	if nil == ethClient {
		ethClient, err = ethclient.Dial("https://bsc-dataseed4.binance.org/")
		if err != nil {
			panic("eth client err")
		}
	}

	// a random user account address
	address := common.HexToAddress(addressParam)
	bytecode, err = ethClient.CodeAt(context.Background(), address, nil) // nil is latest block
	if err != nil {
		return false, err
	}

	if len(bytecode) > 0 {
		return false, nil
	}

	return true, nil
}

// TestSign testSign.
func (a *AppService) TestSign(ctx context.Context, req *pb.TestSignRequest) (*pb.TestSignReply, error) {
	privateKey, err := crypto.HexToECDSA(req.Secret)
	if err != nil {
		return &pb.TestSignReply{Sign: ""}, err
	}

	data := []byte(req.SignContent)
	hash := crypto.Keccak256Hash(data)

	signature, err := crypto.Sign(hash.Bytes(), privateKey)
	if err != nil {
		return &pb.TestSignReply{Sign: ""}, err
	}

	return &pb.TestSignReply{Sign: hexutil.Encode(signature)}, nil
}

func verifySig(sigHex string, msg []byte) (bool, string) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("签名，捕获到异常:")
		}
	}()

	sig := hexutil.MustDecode(sigHex)

	hash := crypto.Keccak256Hash(msg)
	if sig[crypto.RecoveryIDOffset] == 27 || sig[crypto.RecoveryIDOffset] == 28 {
		sig[crypto.RecoveryIDOffset] -= 27 // Transform yellow paper V from 27/28 to 0/1
	}

	recovered, err := crypto.SigToPub(hash.Bytes(), sig)
	if err != nil {
		return false, ""
	}

	recoveredAddr := crypto.PubkeyToAddress(*recovered)

	sigPublicKeyBytes := crypto.FromECDSAPub(recovered)
	signatureNoRecoverID := sig[:len(sig)-1] // remove recovery id
	verified := crypto.VerifySignature(sigPublicKeyBytes, hash.Bytes(), signatureNoRecoverID)
	return verified, recoveredAddr.Hex()
}

// EthAuthorize ethAuthorize.
func (a *AppService) EthAuthorize(ctx context.Context, req *pb.EthAuthorizeRequest) (*pb.EthAuthorizeReply, error) {
	userAddress := req.SendBody.Address // 以太坊账户

	// 验证
	var (
		res bool
		err error
	)
	res, err = addressCheck(userAddress)
	if nil != err {
		return &pb.EthAuthorizeReply{
			Token:  "",
			Status: "地址验证失败",
		}, nil
	}

	if !res {
		return &pb.EthAuthorizeReply{
			Token:  "",
			Status: "地址格式错误",
		}, nil
	}

	if 20 >= len(req.SendBody.Sign) {
		return &pb.EthAuthorizeReply{
			Token:  "",
			Status: "签名错误",
		}, nil
	}

	var (
		addressFromSign string
	)
	res, addressFromSign = verifySig(req.SendBody.Sign, []byte(userAddress))
	if !res || addressFromSign != userAddress {
		return &pb.EthAuthorizeReply{
			Token:  "",
			Status: "地址签名错误",
		}, nil
	}

	// 根据地址查询用户，不存在时则创建
	var (
		user *biz.User
		msg  string
	)
	user, err, msg = a.ac.GetExistUserByAddressOrCreate(ctx, userAddress, req)
	if err != nil {
		return &pb.EthAuthorizeReply{
			Token:  "",
			Status: msg,
		}, nil
	}

	claims := auth.CustomClaims{
		Address: user.Address,
		StandardClaims: jwt2.StandardClaims{
			NotBefore: time.Now().Unix(),              // 签名的生效时间
			ExpiresAt: time.Now().Unix() + 60*60*24*2, // 7天过期
			Issuer:    "game",
		},
	}

	var (
		token string
	)
	token, err = auth.CreateToken(claims, a.ca.JwtKey)
	if err != nil {
		return &pb.EthAuthorizeReply{
			Token:  "",
			Status: "生成token失败",
		}, nil
	}

	return &pb.EthAuthorizeReply{
		Token:  token,
		Status: "ok",
	}, nil
}

// UserInfo userInfo.
func (a *AppService) UserInfo(ctx context.Context, req *pb.UserInfoRequest) (*pb.UserInfoReply, error) {
	// 在上下文 context 中取出 claims 对象
	var (
		address string
	)
	if claims, ok := jwt.FromContext(ctx); ok {
		c := claims.(jwt2.MapClaims)
		if c["Address"] == nil {
			return &pb.UserInfoReply{Status: "无效token"}, nil
		}

		address = c["Address"].(string)

		// 验证
		var (
			res bool
			err error
		)
		res, err = addressCheck(address)
		if nil != err {
			return &pb.UserInfoReply{Status: "无效token"}, nil
		}

		if !res {
			return &pb.UserInfoReply{Status: "无效token"}, nil
		}
	} else {
		return &pb.UserInfoReply{Status: "无效token"}, nil
	}

	return a.ac.UserInfo(ctx, address)
}

// UserRecommend userRecommend.
func (a *AppService) UserRecommend(ctx context.Context, req *pb.UserRecommendRequest) (*pb.UserRecommendReply, error) {
	// 在上下文 context 中取出 claims 对象
	var (
		address string
	)
	if claims, ok := jwt.FromContext(ctx); ok {
		c := claims.(jwt2.MapClaims)
		if c["Address"] == nil {
			return &pb.UserRecommendReply{Status: "无效token"}, nil
		}

		address = c["Address"].(string)

		// 验证
		var (
			res bool
			err error
		)
		res, err = addressCheck(address)
		if nil != err {
			return &pb.UserRecommendReply{Status: "无效token"}, nil
		}

		if !res {
			return &pb.UserRecommendReply{Status: "无效token"}, nil
		}
	} else {
		return &pb.UserRecommendReply{Status: "无效token"}, nil
	}

	return a.ac.UserRecommend(ctx, address, req)
}

// UserRecommendL userRecommendL.
func (a *AppService) UserRecommendL(ctx context.Context, req *pb.UserRecommendLRequest) (*pb.UserRecommendLReply, error) {
	// 在上下文 context 中取出 claims 对象
	var (
		address string
	)
	if claims, ok := jwt.FromContext(ctx); ok {
		c := claims.(jwt2.MapClaims)
		if c["Address"] == nil {
			return &pb.UserRecommendLReply{Status: "无效token"}, nil
		}

		address = c["Address"].(string)

		// 验证
		var (
			res bool
			err error
		)
		res, err = addressCheck(address)
		if nil != err {
			return &pb.UserRecommendLReply{Status: "无效token"}, nil
		}

		if !res {
			return &pb.UserRecommendLReply{Status: "无效token"}, nil
		}
	} else {
		return &pb.UserRecommendLReply{Status: "无效token"}, nil
	}

	return a.ac.UserRecommendL(ctx, address, req)
}

// UserLand userLand.
func (a *AppService) UserLand(ctx context.Context, req *pb.UserLandRequest) (*pb.UserLandReply, error) {
	// 在上下文 context 中取出 claims 对象
	var (
		address string
	)
	if claims, ok := jwt.FromContext(ctx); ok {
		c := claims.(jwt2.MapClaims)
		if c["Address"] == nil {
			return &pb.UserLandReply{Status: "无效token"}, nil
		}

		address = c["Address"].(string)

		// 验证
		var (
			res bool
			err error
		)
		res, err = addressCheck(address)
		if nil != err {
			return &pb.UserLandReply{Status: "无效token"}, nil
		}

		if !res {
			return &pb.UserLandReply{Status: "无效token"}, nil
		}
	} else {
		return &pb.UserLandReply{Status: "无效token"}, nil
	}

	return a.ac.UserLand(ctx, address, req)
}

// UserStakeRewardList userStakeRewardList.
func (a *AppService) UserStakeRewardList(ctx context.Context, req *pb.UserStakeRewardListRequest) (*pb.UserStakeRewardListReply, error) {
	// 在上下文 context 中取出 claims 对象
	var (
		address string
	)
	if claims, ok := jwt.FromContext(ctx); ok {
		c := claims.(jwt2.MapClaims)
		if c["Address"] == nil {
			return &pb.UserStakeRewardListReply{Status: "无效token"}, nil
		}

		address = c["Address"].(string)

		// 验证
		var (
			res bool
			err error
		)
		res, err = addressCheck(address)
		if nil != err {
			return &pb.UserStakeRewardListReply{Status: "无效token"}, nil
		}

		if !res {
			return &pb.UserStakeRewardListReply{Status: "无效token"}, nil
		}
	} else {
		return &pb.UserStakeRewardListReply{Status: "无效token"}, nil
	}

	return a.ac.UserStakeRewardList(ctx, address, req)
}

// UserBoxList userBoxList.
func (a *AppService) UserBoxList(ctx context.Context, req *pb.UserBoxListRequest) (*pb.UserBoxListReply, error) {
	// 在上下文 context 中取出 claims 对象
	var (
		address string
	)
	if claims, ok := jwt.FromContext(ctx); ok {
		c := claims.(jwt2.MapClaims)
		if c["Address"] == nil {
			return &pb.UserBoxListReply{Status: "无效token"}, nil
		}

		address = c["Address"].(string)

		// 验证
		var (
			res bool
			err error
		)
		res, err = addressCheck(address)
		if nil != err {
			return &pb.UserBoxListReply{Status: "无效token"}, nil
		}

		if !res {
			return &pb.UserBoxListReply{Status: "无效token"}, nil
		}
	} else {
		return &pb.UserBoxListReply{Status: "无效token"}, nil
	}

	return a.ac.UserBoxList(ctx, address, req)
}

// UserBackList userBackList.
func (a *AppService) UserBackList(ctx context.Context, req *pb.UserBackListRequest) (*pb.UserBackListReply, error) {
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
