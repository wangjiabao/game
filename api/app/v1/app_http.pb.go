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

const OperationAppCreateUser = "/api.app.v1.App/CreateUser"

type AppHTTPServer interface {
	CreateUser(context.Context, *CreateUserRequest) (*CreateUserReply, error)
}

func RegisterAppHTTPServer(s *http.Server, srv AppHTTPServer) {
	r := s.Route("/")
	r.POST("/api/user", _App_CreateUser0_HTTP_Handler(srv))
}

func _App_CreateUser0_HTTP_Handler(srv AppHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateUserRequest
		if err := ctx.Bind(&in.SendBody); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAppCreateUser)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateUser(ctx, req.(*CreateUserRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreateUserReply)
		return ctx.Result(200, reply)
	}
}

type AppHTTPClient interface {
	CreateUser(ctx context.Context, req *CreateUserRequest, opts ...http.CallOption) (rsp *CreateUserReply, err error)
}

type AppHTTPClientImpl struct {
	cc *http.Client
}

func NewAppHTTPClient(client *http.Client) AppHTTPClient {
	return &AppHTTPClientImpl{client}
}

func (c *AppHTTPClientImpl) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...http.CallOption) (*CreateUserReply, error) {
	var out CreateUserReply
	pattern := "/api/user"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationAppCreateUser))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in.SendBody, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
