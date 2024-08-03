package demo

import (
  "context"
  "fmt"
  "github.com/kysion/base-library/base_hook"
  //_ "github.com/kysion/sub-demo/internal/hook"
  "github.com/kysion/sub-demo/service"
)

type sSub struct {
  // 用户授权Hook
  DefaultHook base_hook.BaseHook[int, base_hook.DefaultHookFunc]
}

func NewSub() service.ISub {
  result := &sSub{}

  // 订阅 用户授权Hook 和 取消授权Hook
  result.DefaultHook.InstallHook(1, result.Auth)
  result.DefaultHook.InstallHook(2, result.UnAuth)

  return result
}

func init() {
  service.RegisterSub(NewSub())
}

// Auth 授权
func (s *sSub) Auth(ctx context.Context, info interface{}) error {
  fmt.Println("Auth")

  return nil
}

// UnAuth 取消授权
func (s *sSub) UnAuth(ctx context.Context, info interface{}) error {
  fmt.Println("UnAuth")
  return nil
}
