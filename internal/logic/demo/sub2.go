package demo

import (
  "context"
  "fmt"
  "github.com/kysion/base-library/base_hook"
  "github.com/kysion/sub-demo/service"
)

type sSub2 struct {
  // 其他Hook
  OtherHook base_hook.BaseHook[int, base_hook.UserHookFunc]
}

func NewSub2() service.ISub2 {
  result := &sSub2{}

  // 订阅 用户授权Hook 和 取消授权Hook
  //service.Gateway().UserHook().InstallHook(1, result.Auth)
  //service.Gateway().UserHook().InstallHook(2, result.UnAuth)

  // 旧的
  //hook.Gateway().OtherHook().InstallHook(1, result.Auth)
  //hook.Gateway().OtherHook().InstallHook(2, result.UnAuth)

  result.OtherHook.InstallHook(1, result.Auth)
  result.OtherHook.InstallHook(2, result.UnAuth)

  base_hook.RegisterHookMessage(&result.OtherHook)

  return result
}

func init() {
  service.RegisterSub2(NewSub2())
}

// Auth 授权
func (s *sSub2) Auth(ctx context.Context, info *base_hook.User) error {
  fmt.Println("Auth22")
  fmt.Println(info)

  return nil
}

// UnAuth 取消授权
func (s *sSub2) UnAuth(ctx context.Context, info *base_hook.User) error {
  fmt.Println("UnAuth22")
  return nil
}
