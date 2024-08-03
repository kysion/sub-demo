package hook

import (
  "context"
  "fmt"
  "github.com/gogf/gf/v2/encoding/gjson"
  "github.com/kysion/base-library/base_hook"
  "github.com/kysion/base-library/base_model"
  "reflect"
)

type sGateway struct {
  userHook *base_hook.BaseHook[int, base_hook.CommonHookFunc]

  // 其他Hook
  otherHook *base_hook.BaseHook[int, base_hook.UserHookFunc]
}

var localGateway *sGateway

func init() {
  Gateway()
}

func Gateway() *sGateway {
  if localGateway != nil {
    return localGateway
  }

  localGateway = &sGateway{
    userHook:  &base_hook.BaseHook[int, base_hook.CommonHookFunc]{},
    otherHook: &base_hook.BaseHook[int, base_hook.UserHookFunc]{},
  }

  //base_hook.BroadcastMessage = localGateway.BroadcastMessage

  return localGateway
}

//
//func init() {
//  //service.RegisterGateway(NewGateway())
//}

func (s *sGateway) UserHook() *base_hook.BaseHook[int, base_hook.CommonHookFunc] {
  return s.userHook
}

func (s *sGateway) OtherHook() *base_hook.BaseHook[int, base_hook.UserHookFunc] {
  return s.otherHook
}

func (s *sGateway) BroadcastMessage(model base_model.HookModel) {

  var xx base_hook.CommonHookFunc
  var ff string = reflect.TypeOf(xx).String()
  fmt.Println(ff)

  if model.BusinessTypeStr == reflect.TypeOf(xx).String() {
    // model.Data 可以获取到消息内容,按实际使用场景按需转换数据类型
    // 广播消息
    s.userHook.Iterator(func(key int, value base_hook.CommonHookFunc) {
      err := value(context.Background(), model.Data)
      if err != nil {
        fmt.Println(err)
      }
    })
  }

  // 业务层通用Hook函数
  var yy base_hook.UserHookFunc
  if model.BusinessTypeStr == reflect.TypeOf(yy).String() {
    // model.Data 可以获取到消息内容,按实际使用场景按需转换数据类型
    user := base_hook.User{}
    gjson.DecodeTo(model.Data, &user)
    s.otherHook.Iterator(func(key int, value base_hook.UserHookFunc) {
      err := value(context.Background(), &user)
      if err != nil {
        fmt.Println(err)
      }
    })
  }

}
