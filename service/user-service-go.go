package service

import (
	context "context"
	"user-service-go/pb/user-service-go"
)

type UserService struct {
}

func (s *UserService) Login(ctx context.Context, param *v1.LoginRequest) (*v1.LoginResult, error) {
	panic("implement me")
}
