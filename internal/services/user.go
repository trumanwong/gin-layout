package services

import (
	"context"
	v1 "gin-layout/api/app/v1"
)

func (s AppService) CreateUser(ctx context.Context, req *v1.CreateUserRequest) (resp *v1.CreateUserReply, err error) {
	var retErr error
	defer func() {
		s.handleError(retErr)
	}()

	resp, err = s.greeterRepo.Create(ctx, req)
	if err != nil {
		retErr = v1.ErrorInternal("failed to create greeter: %s", err.Error())
		return
	}
	return
}

func (s AppService) GetUser(ctx context.Context, req *v1.GetUserRequest) (resp *v1.GetUserReply, err error) {
	resp = &v1.GetUserReply{Name: "hello"}
	return
}
