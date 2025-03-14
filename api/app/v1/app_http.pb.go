// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.7.1
// - protoc             v3.21.7
// source: api/app/v1/app.proto

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationAppEthAuthorize = "/api.app.v1.App/EthAuthorize"
const OperationAppTestSign = "/api.app.v1.App/TestSign"
const OperationAppUserBackList = "/api.app.v1.App/UserBackList"
const OperationAppUserBoxList = "/api.app.v1.App/UserBoxList"
const OperationAppUserInfo = "/api.app.v1.App/UserInfo"
const OperationAppUserLand = "/api.app.v1.App/UserLand"
const OperationAppUserMarketLandList = "/api.app.v1.App/UserMarketLandList"
const OperationAppUserMarketPropList = "/api.app.v1.App/UserMarketPropList"
const OperationAppUserMarketRentLandList = "/api.app.v1.App/UserMarketRentLandList"
const OperationAppUserMarketSeedList = "/api.app.v1.App/UserMarketSeedList"
const OperationAppUserMyMarketList = "/api.app.v1.App/UserMyMarketList"
const OperationAppUserNoticeList = "/api.app.v1.App/UserNoticeList"
const OperationAppUserRecommend = "/api.app.v1.App/UserRecommend"
const OperationAppUserRecommendL = "/api.app.v1.App/UserRecommendL"
const OperationAppUserSkateRewardList = "/api.app.v1.App/UserSkateRewardList"
const OperationAppUserStakeRewardList = "/api.app.v1.App/UserStakeRewardList"

type AppHTTPServer interface {
	EthAuthorize(context.Context, *EthAuthorizeRequest) (*EthAuthorizeReply, error)
	TestSign(context.Context, *TestSignRequest) (*TestSignReply, error)
	// UserBackList 仓库
	UserBackList(context.Context, *UserBackListRequest) (*UserBackListReply, error)
	// UserBoxList 盲盒列表
	UserBoxList(context.Context, *UserBoxListRequest) (*UserBoxListReply, error)
	// UserInfo 用户信息
	UserInfo(context.Context, *UserInfoRequest) (*UserInfoReply, error)
	// UserLand 土地背包列表
	UserLand(context.Context, *UserLandRequest) (*UserLandReply, error)
	// UserMarketLandList 市场 土地
	UserMarketLandList(context.Context, *UserMarketLandListRequest) (*UserMarketLandListReply, error)
	// UserMarketPropList 市场 道具
	UserMarketPropList(context.Context, *UserMarketPropListRequest) (*UserMarketPropListReply, error)
	// UserMarketRentLandList 市场 出租
	UserMarketRentLandList(context.Context, *UserMarketRentLandListRequest) (*UserMarketRentLandListReply, error)
	// UserMarketSeedList 市场 种子
	UserMarketSeedList(context.Context, *UserMarketSeedListRequest) (*UserMarketSeedListReply, error)
	// UserMyMarketList 市场 我的
	UserMyMarketList(context.Context, *UserMyMarketListRequest) (*UserMyMarketListReply, error)
	// UserNoticeList 通知
	UserNoticeList(context.Context, *UserNoticeListRequest) (*UserNoticeListReply, error)
	// UserRecommend 直推列表
	UserRecommend(context.Context, *UserRecommendRequest) (*UserRecommendReply, error)
	// UserRecommendL L1L2L3内容
	UserRecommendL(context.Context, *UserRecommendLRequest) (*UserRecommendLReply, error)
	// UserSkateRewardList 果实放大器 获奖记录
	UserSkateRewardList(context.Context, *UserSkateRewardListRequest) (*UserSkateRewardListReply, error)
	// UserStakeRewardList 粮仓列表
	UserStakeRewardList(context.Context, *UserStakeRewardListRequest) (*UserStakeRewardListReply, error)
}

func RegisterAppHTTPServer(s *http.Server, srv AppHTTPServer) {
	r := s.Route("/")
	r.GET("/api/app_server/test_sign", _App_TestSign0_HTTP_Handler(srv))
	r.POST("/api/app_server/eth_authorize", _App_EthAuthorize0_HTTP_Handler(srv))
	r.GET("/api/app_server/user_info", _App_UserInfo0_HTTP_Handler(srv))
	r.GET("/api/app_server/user_recommend", _App_UserRecommend0_HTTP_Handler(srv))
	r.GET("/api/app_server/user_recommend_l", _App_UserRecommendL0_HTTP_Handler(srv))
	r.GET("/api/app_server/user_land", _App_UserLand0_HTTP_Handler(srv))
	r.GET("/api/app_server/user_stake_reward_list", _App_UserStakeRewardList0_HTTP_Handler(srv))
	r.GET("/api/app_server/user_box_list", _App_UserBoxList0_HTTP_Handler(srv))
	r.GET("/api/app_server/user_back_list", _App_UserBackList0_HTTP_Handler(srv))
	r.GET("/api/app_server/user_market_seed_list", _App_UserMarketSeedList0_HTTP_Handler(srv))
	r.GET("/api/app_server/user_market_land_list", _App_UserMarketLandList0_HTTP_Handler(srv))
	r.GET("/api/app_server/user_market_prop_list", _App_UserMarketPropList0_HTTP_Handler(srv))
	r.GET("/api/app_server/user_market_rent_land_list", _App_UserMarketRentLandList0_HTTP_Handler(srv))
	r.GET("/api/app_server/user_market_my_list", _App_UserMyMarketList0_HTTP_Handler(srv))
	r.GET("/api/app_server/user_notice_list", _App_UserNoticeList0_HTTP_Handler(srv))
	r.GET("/api/app_server/user_skate_reward_list", _App_UserSkateRewardList0_HTTP_Handler(srv))
}

func _App_TestSign0_HTTP_Handler(srv AppHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in TestSignRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAppTestSign)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.TestSign(ctx, req.(*TestSignRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*TestSignReply)
		return ctx.Result(200, reply)
	}
}

func _App_EthAuthorize0_HTTP_Handler(srv AppHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in EthAuthorizeRequest
		if err := ctx.Bind(&in.SendBody); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAppEthAuthorize)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.EthAuthorize(ctx, req.(*EthAuthorizeRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*EthAuthorizeReply)
		return ctx.Result(200, reply)
	}
}

func _App_UserInfo0_HTTP_Handler(srv AppHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UserInfoRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAppUserInfo)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UserInfo(ctx, req.(*UserInfoRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UserInfoReply)
		return ctx.Result(200, reply)
	}
}

func _App_UserRecommend0_HTTP_Handler(srv AppHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UserRecommendRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAppUserRecommend)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UserRecommend(ctx, req.(*UserRecommendRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UserRecommendReply)
		return ctx.Result(200, reply)
	}
}

func _App_UserRecommendL0_HTTP_Handler(srv AppHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UserRecommendLRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAppUserRecommendL)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UserRecommendL(ctx, req.(*UserRecommendLRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UserRecommendLReply)
		return ctx.Result(200, reply)
	}
}

func _App_UserLand0_HTTP_Handler(srv AppHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UserLandRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAppUserLand)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UserLand(ctx, req.(*UserLandRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UserLandReply)
		return ctx.Result(200, reply)
	}
}

func _App_UserStakeRewardList0_HTTP_Handler(srv AppHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UserStakeRewardListRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAppUserStakeRewardList)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UserStakeRewardList(ctx, req.(*UserStakeRewardListRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UserStakeRewardListReply)
		return ctx.Result(200, reply)
	}
}

func _App_UserBoxList0_HTTP_Handler(srv AppHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UserBoxListRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAppUserBoxList)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UserBoxList(ctx, req.(*UserBoxListRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UserBoxListReply)
		return ctx.Result(200, reply)
	}
}

func _App_UserBackList0_HTTP_Handler(srv AppHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UserBackListRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAppUserBackList)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UserBackList(ctx, req.(*UserBackListRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UserBackListReply)
		return ctx.Result(200, reply)
	}
}

func _App_UserMarketSeedList0_HTTP_Handler(srv AppHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UserMarketSeedListRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAppUserMarketSeedList)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UserMarketSeedList(ctx, req.(*UserMarketSeedListRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UserMarketSeedListReply)
		return ctx.Result(200, reply)
	}
}

func _App_UserMarketLandList0_HTTP_Handler(srv AppHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UserMarketLandListRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAppUserMarketLandList)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UserMarketLandList(ctx, req.(*UserMarketLandListRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UserMarketLandListReply)
		return ctx.Result(200, reply)
	}
}

func _App_UserMarketPropList0_HTTP_Handler(srv AppHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UserMarketPropListRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAppUserMarketPropList)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UserMarketPropList(ctx, req.(*UserMarketPropListRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UserMarketPropListReply)
		return ctx.Result(200, reply)
	}
}

func _App_UserMarketRentLandList0_HTTP_Handler(srv AppHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UserMarketRentLandListRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAppUserMarketRentLandList)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UserMarketRentLandList(ctx, req.(*UserMarketRentLandListRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UserMarketRentLandListReply)
		return ctx.Result(200, reply)
	}
}

func _App_UserMyMarketList0_HTTP_Handler(srv AppHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UserMyMarketListRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAppUserMyMarketList)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UserMyMarketList(ctx, req.(*UserMyMarketListRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UserMyMarketListReply)
		return ctx.Result(200, reply)
	}
}

func _App_UserNoticeList0_HTTP_Handler(srv AppHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UserNoticeListRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAppUserNoticeList)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UserNoticeList(ctx, req.(*UserNoticeListRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UserNoticeListReply)
		return ctx.Result(200, reply)
	}
}

func _App_UserSkateRewardList0_HTTP_Handler(srv AppHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UserSkateRewardListRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAppUserSkateRewardList)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UserSkateRewardList(ctx, req.(*UserSkateRewardListRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UserSkateRewardListReply)
		return ctx.Result(200, reply)
	}
}

type AppHTTPClient interface {
	EthAuthorize(ctx context.Context, req *EthAuthorizeRequest, opts ...http.CallOption) (rsp *EthAuthorizeReply, err error)
	TestSign(ctx context.Context, req *TestSignRequest, opts ...http.CallOption) (rsp *TestSignReply, err error)
	UserBackList(ctx context.Context, req *UserBackListRequest, opts ...http.CallOption) (rsp *UserBackListReply, err error)
	UserBoxList(ctx context.Context, req *UserBoxListRequest, opts ...http.CallOption) (rsp *UserBoxListReply, err error)
	UserInfo(ctx context.Context, req *UserInfoRequest, opts ...http.CallOption) (rsp *UserInfoReply, err error)
	UserLand(ctx context.Context, req *UserLandRequest, opts ...http.CallOption) (rsp *UserLandReply, err error)
	UserMarketLandList(ctx context.Context, req *UserMarketLandListRequest, opts ...http.CallOption) (rsp *UserMarketLandListReply, err error)
	UserMarketPropList(ctx context.Context, req *UserMarketPropListRequest, opts ...http.CallOption) (rsp *UserMarketPropListReply, err error)
	UserMarketRentLandList(ctx context.Context, req *UserMarketRentLandListRequest, opts ...http.CallOption) (rsp *UserMarketRentLandListReply, err error)
	UserMarketSeedList(ctx context.Context, req *UserMarketSeedListRequest, opts ...http.CallOption) (rsp *UserMarketSeedListReply, err error)
	UserMyMarketList(ctx context.Context, req *UserMyMarketListRequest, opts ...http.CallOption) (rsp *UserMyMarketListReply, err error)
	UserNoticeList(ctx context.Context, req *UserNoticeListRequest, opts ...http.CallOption) (rsp *UserNoticeListReply, err error)
	UserRecommend(ctx context.Context, req *UserRecommendRequest, opts ...http.CallOption) (rsp *UserRecommendReply, err error)
	UserRecommendL(ctx context.Context, req *UserRecommendLRequest, opts ...http.CallOption) (rsp *UserRecommendLReply, err error)
	UserSkateRewardList(ctx context.Context, req *UserSkateRewardListRequest, opts ...http.CallOption) (rsp *UserSkateRewardListReply, err error)
	UserStakeRewardList(ctx context.Context, req *UserStakeRewardListRequest, opts ...http.CallOption) (rsp *UserStakeRewardListReply, err error)
}

type AppHTTPClientImpl struct {
	cc *http.Client
}

func NewAppHTTPClient(client *http.Client) AppHTTPClient {
	return &AppHTTPClientImpl{client}
}

func (c *AppHTTPClientImpl) EthAuthorize(ctx context.Context, in *EthAuthorizeRequest, opts ...http.CallOption) (*EthAuthorizeReply, error) {
	var out EthAuthorizeReply
	pattern := "/api/app_server/eth_authorize"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationAppEthAuthorize))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in.SendBody, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AppHTTPClientImpl) TestSign(ctx context.Context, in *TestSignRequest, opts ...http.CallOption) (*TestSignReply, error) {
	var out TestSignReply
	pattern := "/api/app_server/test_sign"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationAppTestSign))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AppHTTPClientImpl) UserBackList(ctx context.Context, in *UserBackListRequest, opts ...http.CallOption) (*UserBackListReply, error) {
	var out UserBackListReply
	pattern := "/api/app_server/user_back_list"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationAppUserBackList))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AppHTTPClientImpl) UserBoxList(ctx context.Context, in *UserBoxListRequest, opts ...http.CallOption) (*UserBoxListReply, error) {
	var out UserBoxListReply
	pattern := "/api/app_server/user_box_list"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationAppUserBoxList))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AppHTTPClientImpl) UserInfo(ctx context.Context, in *UserInfoRequest, opts ...http.CallOption) (*UserInfoReply, error) {
	var out UserInfoReply
	pattern := "/api/app_server/user_info"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationAppUserInfo))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AppHTTPClientImpl) UserLand(ctx context.Context, in *UserLandRequest, opts ...http.CallOption) (*UserLandReply, error) {
	var out UserLandReply
	pattern := "/api/app_server/user_land"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationAppUserLand))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AppHTTPClientImpl) UserMarketLandList(ctx context.Context, in *UserMarketLandListRequest, opts ...http.CallOption) (*UserMarketLandListReply, error) {
	var out UserMarketLandListReply
	pattern := "/api/app_server/user_market_land_list"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationAppUserMarketLandList))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AppHTTPClientImpl) UserMarketPropList(ctx context.Context, in *UserMarketPropListRequest, opts ...http.CallOption) (*UserMarketPropListReply, error) {
	var out UserMarketPropListReply
	pattern := "/api/app_server/user_market_prop_list"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationAppUserMarketPropList))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AppHTTPClientImpl) UserMarketRentLandList(ctx context.Context, in *UserMarketRentLandListRequest, opts ...http.CallOption) (*UserMarketRentLandListReply, error) {
	var out UserMarketRentLandListReply
	pattern := "/api/app_server/user_market_rent_land_list"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationAppUserMarketRentLandList))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AppHTTPClientImpl) UserMarketSeedList(ctx context.Context, in *UserMarketSeedListRequest, opts ...http.CallOption) (*UserMarketSeedListReply, error) {
	var out UserMarketSeedListReply
	pattern := "/api/app_server/user_market_seed_list"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationAppUserMarketSeedList))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AppHTTPClientImpl) UserMyMarketList(ctx context.Context, in *UserMyMarketListRequest, opts ...http.CallOption) (*UserMyMarketListReply, error) {
	var out UserMyMarketListReply
	pattern := "/api/app_server/user_market_my_list"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationAppUserMyMarketList))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AppHTTPClientImpl) UserNoticeList(ctx context.Context, in *UserNoticeListRequest, opts ...http.CallOption) (*UserNoticeListReply, error) {
	var out UserNoticeListReply
	pattern := "/api/app_server/user_notice_list"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationAppUserNoticeList))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AppHTTPClientImpl) UserRecommend(ctx context.Context, in *UserRecommendRequest, opts ...http.CallOption) (*UserRecommendReply, error) {
	var out UserRecommendReply
	pattern := "/api/app_server/user_recommend"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationAppUserRecommend))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AppHTTPClientImpl) UserRecommendL(ctx context.Context, in *UserRecommendLRequest, opts ...http.CallOption) (*UserRecommendLReply, error) {
	var out UserRecommendLReply
	pattern := "/api/app_server/user_recommend_l"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationAppUserRecommendL))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AppHTTPClientImpl) UserSkateRewardList(ctx context.Context, in *UserSkateRewardListRequest, opts ...http.CallOption) (*UserSkateRewardListReply, error) {
	var out UserSkateRewardListReply
	pattern := "/api/app_server/user_skate_reward_list"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationAppUserSkateRewardList))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AppHTTPClientImpl) UserStakeRewardList(ctx context.Context, in *UserStakeRewardListRequest, opts ...http.CallOption) (*UserStakeRewardListReply, error) {
	var out UserStakeRewardListReply
	pattern := "/api/app_server/user_stake_reward_list"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationAppUserStakeRewardList))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
