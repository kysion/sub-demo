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
  CommonHook base_hook.BaseHook[int, base_hook.CommonHookFunc]
}

func NewSub() service.ISub {
  result := &sSub{}

  // 订阅 用户授权Hook 和 取消授权Hook
  //service.Gateway().UserHook().InstallHook(1, result.Auth)
  //service.Gateway().UserHook().InstallHook(2, result.UnAuth)

  result.CommonHook.InstallHook(1, result.Auth)
  result.CommonHook.InstallHook(2, result.UnAuth)

  //hook.Gateway().UserHook().InstallHook(1, result.Auth)
  //hook.Gateway().UserHook().InstallHook(2, result.UnAuth)

  // 注册网关Hook消息
  base_hook.RegisterHookMessage(&result.CommonHook) // 消息网关，收到消息后，会交给InstallHook()注册的回调函数处理

  // 注册网关Hook消息
  //(result).registerHookMessage()

  return result
}

func init() {
  service.RegisterSub(NewSub())
}

func (s *sSub) registerHookMessage() service.ISub {

  //base_hook.Gateway().RegisterHookMessage(func(model base_model.HookModel) {
  //  if model.BusinessType() == base_enum.Hook.BusinessType.Default {
  //    // model.Data 可以获取到消息内容,按实际使用场景按需转换数据类型
  //    // 广播消息
  //    base_hook.PublishHookMessage(context.Background(), &s.CommonHook, base_hook.Option{Data: model.Data})
  //    // 广播消息
  //    base_hook.PublishHookMessage(context.Background(), &s.OtherHook, base_hook.Option{Data: model.Data})
  //  }
  //})
  return s
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
