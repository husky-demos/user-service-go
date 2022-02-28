package utils

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	v1 "user-service-go/pb/common"
)

func NewError(code uint64, message string) error {
	st, err := status.New(codes.Unknown, "").WithDetails(&v1.ErrorResult{
		Code:    code,
		Message: message,
	})
	if err == nil {
		return st.Err()
	}
	panic(err)
}
