// Code generated by protoc-gen-gin-http. DO NOT EDIT.
// versions:
// - protoc-gen-gin-http v1.0.0
// - protoc             v4.24.3
// source: api/app/v1/app.proto

package v1

import (
	context "context"
	gin "github.com/gin-gonic/gin"
	transport "github.com/trumanwong/gin-transport/transport"
	errors "github.com/trumanwong/gin-transport/transport/errors"
	reflect "reflect"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the gin package it is being compiled against.
var _ = new(context.Context)
var _ = errors.New
var _ = gin.Version
var _ = reflect.ValueOf
var _ = transport.SupportPackageIsVersion1

const OperationAppCreateUser = "/api.app.v1.App/CreateUser"
const OperationAppDeleteUser = "/api.app.v1.App/DeleteUser"
const OperationAppGetUser = "/api.app.v1.App/GetUser"
const OperationAppGreeter = "/api.app.v1.App/Greeter"
const OperationAppListUser = "/api.app.v1.App/ListUser"
const OperationAppUpdateUser = "/api.app.v1.App/UpdateUser"

type AppHTTPServer interface {
	// CreateUser 创建用户
	CreateUser(context.Context, *CreateUserRequest) (*CreateUserReply, error)
	// DeleteUser 删除用户
	DeleteUser(context.Context, *DeleteUserRequest) (*DeleteUserReply, error)
	// GetUser 查找用户
	GetUser(context.Context, *GetUserRequest) (*GetUserReply, error)
	Greeter(context.Context, *GreeterRequest) (*GreeterReply, error)
	// ListUser 用户列表
	ListUser(context.Context, *ListUserRequest) (*ListUserReply, error)
	// UpdateUser 更新用户
	UpdateUser(context.Context, *UpdateUserRequest) (*UpdateUserReply, error)
}

func RegisterAppHTTPServer(s *transport.Server, srv AppHTTPServer) {
	s.AddMethod("POST", "/api/user", _App_CreateUser0_HTTP_Handler(s, srv, OperationAppCreateUser))
	s.AddMethod("PUT", "/api/user/:Id", _App_UpdateUser0_HTTP_Handler(s, srv, OperationAppUpdateUser))
	s.AddMethod("DELETE", "/api/user/:Id", _App_DeleteUser0_HTTP_Handler(s, srv, OperationAppDeleteUser))
	s.AddMethod("GET", "/api/user/:Id", _App_GetUser0_HTTP_Handler(s, srv, OperationAppGetUser))
	s.AddMethod("GET", "/api/user", _App_ListUser0_HTTP_Handler(s, srv, OperationAppListUser))
	s.AddMethod("GET", "/api/greeter/:Name", _App_Greeter0_HTTP_Handler(s, srv, OperationAppGreeter))
}

func _App_CreateUser0_HTTP_Handler(s *transport.Server, srv AppHTTPServer, operation string) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		for _, f := range s.GetMiddlewares() {
			if err := f(ctx, operation); err != nil {
				s.ResultError(ctx, err)
				return
			}
		}
		var req CreateUserRequest
		if err := ctx.ShouldBindJSON(&req); err != nil {
			s.ResultError(ctx, err)
			return
		}
		if err := ctx.ShouldBindQuery(&req); err != nil {
			s.ResultError(ctx, err)
			return
		}
		if validate := reflect.ValueOf(&req).MethodByName("Validate"); validate.IsValid() {
			if err := validate.Call(nil)[0].Interface(); err != nil {
				s.ResultError(ctx, errors.New(400, "", "validate fail: "+err.(error).Error()))
				return
			}
		}
		reply, err := srv.CreateUser(ctx, &req)
		if err != nil {
			s.ResultError(ctx, err)
			return
		}
		s.Result(ctx, 200, reply)
	}
}

func _App_UpdateUser0_HTTP_Handler(s *transport.Server, srv AppHTTPServer, operation string) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		for _, f := range s.GetMiddlewares() {
			if err := f(ctx, operation); err != nil {
				s.ResultError(ctx, err)
				return
			}
		}
		var req UpdateUserRequest
		if err := ctx.ShouldBindJSON(&req); err != nil {
			s.ResultError(ctx, err)
			return
		}
		if err := ctx.ShouldBindQuery(&req); err != nil {
			s.ResultError(ctx, err)
			return
		}
		if err := ctx.ShouldBindUri(&req); err != nil {
			s.ResultError(ctx, err)
			return
		}
		if validate := reflect.ValueOf(&req).MethodByName("Validate"); validate.IsValid() {
			if err := validate.Call(nil)[0].Interface(); err != nil {
				s.ResultError(ctx, errors.New(400, "", "validate fail: "+err.(error).Error()))
				return
			}
		}
		reply, err := srv.UpdateUser(ctx, &req)
		if err != nil {
			s.ResultError(ctx, err)
			return
		}
		s.Result(ctx, 200, reply)
	}
}

func _App_DeleteUser0_HTTP_Handler(s *transport.Server, srv AppHTTPServer, operation string) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		for _, f := range s.GetMiddlewares() {
			if err := f(ctx, operation); err != nil {
				s.ResultError(ctx, err)
				return
			}
		}
		var req DeleteUserRequest
		if err := ctx.ShouldBindQuery(&req); err != nil {
			s.ResultError(ctx, err)
			return
		}
		if err := ctx.ShouldBindUri(&req); err != nil {
			s.ResultError(ctx, err)
			return
		}
		if validate := reflect.ValueOf(&req).MethodByName("Validate"); validate.IsValid() {
			if err := validate.Call(nil)[0].Interface(); err != nil {
				s.ResultError(ctx, errors.New(400, "", "validate fail: "+err.(error).Error()))
				return
			}
		}
		reply, err := srv.DeleteUser(ctx, &req)
		if err != nil {
			s.ResultError(ctx, err)
			return
		}
		s.Result(ctx, 200, reply)
	}
}

func _App_GetUser0_HTTP_Handler(s *transport.Server, srv AppHTTPServer, operation string) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		for _, f := range s.GetMiddlewares() {
			if err := f(ctx, operation); err != nil {
				s.ResultError(ctx, err)
				return
			}
		}
		var req GetUserRequest
		if err := ctx.ShouldBindQuery(&req); err != nil {
			s.ResultError(ctx, err)
			return
		}
		if err := ctx.ShouldBindUri(&req); err != nil {
			s.ResultError(ctx, err)
			return
		}
		if validate := reflect.ValueOf(&req).MethodByName("Validate"); validate.IsValid() {
			if err := validate.Call(nil)[0].Interface(); err != nil {
				s.ResultError(ctx, errors.New(400, "", "validate fail: "+err.(error).Error()))
				return
			}
		}
		reply, err := srv.GetUser(ctx, &req)
		if err != nil {
			s.ResultError(ctx, err)
			return
		}
		s.Result(ctx, 200, reply)
	}
}

func _App_ListUser0_HTTP_Handler(s *transport.Server, srv AppHTTPServer, operation string) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		for _, f := range s.GetMiddlewares() {
			if err := f(ctx, operation); err != nil {
				s.ResultError(ctx, err)
				return
			}
		}
		var req ListUserRequest
		if err := ctx.ShouldBindQuery(&req); err != nil {
			s.ResultError(ctx, err)
			return
		}
		if validate := reflect.ValueOf(&req).MethodByName("Validate"); validate.IsValid() {
			if err := validate.Call(nil)[0].Interface(); err != nil {
				s.ResultError(ctx, errors.New(400, "", "validate fail: "+err.(error).Error()))
				return
			}
		}
		reply, err := srv.ListUser(ctx, &req)
		if err != nil {
			s.ResultError(ctx, err)
			return
		}
		s.Result(ctx, 200, reply)
	}
}

func _App_Greeter0_HTTP_Handler(s *transport.Server, srv AppHTTPServer, operation string) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		for _, f := range s.GetMiddlewares() {
			if err := f(ctx, operation); err != nil {
				s.ResultError(ctx, err)
				return
			}
		}
		var req GreeterRequest
		if err := ctx.ShouldBindQuery(&req); err != nil {
			s.ResultError(ctx, err)
			return
		}
		if err := ctx.ShouldBindUri(&req); err != nil {
			s.ResultError(ctx, err)
			return
		}
		if validate := reflect.ValueOf(&req).MethodByName("Validate"); validate.IsValid() {
			if err := validate.Call(nil)[0].Interface(); err != nil {
				s.ResultError(ctx, errors.New(400, "", "validate fail: "+err.(error).Error()))
				return
			}
		}
		reply, err := srv.Greeter(ctx, &req)
		if err != nil {
			s.ResultError(ctx, err)
			return
		}
		s.Result(ctx, 200, reply)
	}
}
