package services

import (
	"context"
	v1 "gin-layout/api/app/v1"
)

func (s AppService) Greeter(ctx context.Context, req *v1.GreeterRequest) (*v1.GreeterReply, error) {
	return &v1.GreeterReply{
		Message: "Hello " + req.Name,
	}, nil
}
