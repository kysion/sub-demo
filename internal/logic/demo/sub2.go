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

  result.OtherHook.InstallHook(1, result.Auth)
  result.OtherHook.InstallHook(2, result.UnAuth)

  return result
}

func init() {
  service.RegisterSub2(NewSub2())
}

// Auth 授权
func (s *sSub2) Auth(ctx context.Context, info *base_hook.User) error {
  fmt.Println("OtherHook_Auth", info.Username)
  fmt.Println(info)

  //// 使用方式1：Latest方式
  //info.Username = "LinFeiFei"
  //base_hook.PublishHookMessage(context.Background(), &s.OtherHook, base_hook.Option{Data: info}) // 发布消息

  return nil
}

// UnAuth 取消授权
func (s *sSub2) UnAuth(ctx context.Context, info *base_hook.User) error {
  fmt.Println("OtherHook_UnAuth", info.Username)
  return nil
}
