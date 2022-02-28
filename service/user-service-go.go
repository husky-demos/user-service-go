package service

import (
	context "context"
	"database/sql"
	"user-service-go/pb/user-service-go"
)

type UserService struct {
	db *sql.DB
}

func NewUserService(db *sql.DB) *UserService {
	return &UserService{
		db: db,
	}
}

func (s *UserService) Login(ctx context.Context, param *v1.LoginRequest) (*v1.LoginResult, error) {
	panic("implement me")
}
