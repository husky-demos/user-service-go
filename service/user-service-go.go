package service

import (
	context "context"
	"github.com/jmoiron/sqlx"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"user-service-go/model"
	"user-service-go/pb/user-service-go"
	"user-service-go/utils"
)

type UserService struct {
	db *sqlx.DB
}

func NewUserService(db *sqlx.DB) *UserService {
	return &UserService{
		db: db,
	}
}

func (s *UserService) Login(ctx context.Context, param *v1.LoginRequest) (*v1.LoginResult, error) {
	if param.LoginName == "" {
		return nil, utils.NewError(1, "登录账号不能为空")
	}
	if param.LoginPass == "" {
		return nil, utils.NewError(2, "登录密码不能为空")
	}
	var user model.User
	if err := s.db.Get(&user, "select * from users where login_name=? and login_pass=?", param.LoginName, param.LoginPass); err != nil {
		return nil, utils.NewError(3, "用户不存在")
	}
	return &v1.LoginResult{
		User: &v1.User{
			Id:        user.Id,
			NickName:  user.NickName,
			LoginName: user.LoginName,
			IsLocking: user.IsLocking,
		},
		Token:      "123",
		ExpireTime: timestamppb.Now(),
	}, nil
}

func (s *UserService) TestResult(ctx context.Context, param *v1.LoginRequest) (*v1.CommonResult, error) {
	if param.LoginName == "" {
		return nil, utils.NewError(1, "登录账号不能为空")
	}
	if param.LoginPass == "" {
		return nil, utils.NewError(2, "登录密码不能为空")
	}
	var user model.User
	if err := s.db.Get(&user, "select * from users where login_name=? and login_pass=?", param.LoginName, param.LoginPass); err != nil {
		return nil, utils.NewError(3, "用户不存在")
	}
	result, err := anypb.New(&v1.LoginResult{
		User: &v1.User{
			Id:        user.Id,
			NickName:  user.NickName,
			LoginName: user.LoginName,
			IsLocking: user.IsLocking,
		},
		Token:      "123",
		ExpireTime: timestamppb.Now(),
	})
	if err != nil {
		return nil, utils.NewError(4, "业务异常")
	}
	return &v1.CommonResult{
		Code:    0,
		Message: "",
		Data:    result,
	}, nil
}
