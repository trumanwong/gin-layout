package data

import (
	"context"
	v1 "gin-layout/api/helloworld/v1"
	"github.com/trumanwong/go-tools/log"
)

type GreeterRepo struct {
	data   *Data
	logger *log.Logger
}

func NewGreeterRepo(data *Data, logger *log.Logger) *GreeterRepo {
	return &GreeterRepo{
		data:   data,
		logger: logger,
	}
}

func (r *GreeterRepo) Create(ctx context.Context, req *v1.CreateGreeterRequest) (*v1.CreateGreeterReply, error) {
	// todo create greeter
	return &v1.CreateGreeterReply{}, nil
}
