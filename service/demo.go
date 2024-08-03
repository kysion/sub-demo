// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"

	"github.com/kysion/base-library/base_hook"
)

type (
	ISub interface {
		// Auth 授权
		Auth(ctx context.Context, info interface{}) error
		// UnAuth 取消授权
		UnAuth(ctx context.Context, info interface{}) error
	}
	ISub2 interface {
		// Auth 授权
		Auth(ctx context.Context, info *base_hook.User) error
		// UnAuth 取消授权
		UnAuth(ctx context.Context, info *base_hook.User) error
	}
)

var (
	localSub  ISub
	localSub2 ISub2
)

func Sub() ISub {
	if localSub == nil {
		panic("implement not found for interface ISub, forgot register?")
	}
	return localSub
}

func RegisterSub(i ISub) {
	localSub = i
}

func Sub2() ISub2 {
	if localSub2 == nil {
		panic("implement not found for interface ISub2, forgot register?")
	}
	return localSub2
}

func RegisterSub2(i ISub2) {
	localSub2 = i
}
