package utils

import (
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ErrNilDB        = errors.New("nil DB error")
	ErrEmptyUUID    = errors.New("empty uuid")
	ErrNotExist     = errors.New("element not exist")
	ErrNotFound     = status.Error(codes.Code(409), "not found")
	ErrInvalidUUID  = status.Error(codes.Code(409), "invalid uuid")
	ErrNotingUpdate = status.Error(codes.Code(409), "nothing to update")
)
