package cmd

import (
  "context"
  "github.com/gogf/gf/v2/frame/g"
  "github.com/gogf/gf/v2/net/ghttp"
  "github.com/gogf/gf/v2/os/gcmd"
  "github.com/kysion/base-library/base_hook"
)

var (
  Main = gcmd.Command{
    Name:  "main",
    Usage: "main",
    Brief: "start http server",
    Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
      s := g.Server()
      s.Group("/", func(group *ghttp.RouterGroup) {
        group.Middleware(ghttp.MiddlewareHandlerResponse)
        //group.Bind(
        //	hello.NewV1(),
        //)
      })

      //hook.Gateway()
      s.BindHandler("/ws", base_hook.HookDistribution) // 接收消息

      //s.BindHandler("/ws", func(r *ghttp.Request) {
      //  ws, err := r.WebSocket()
      //  if err != nil {
      //    glog.Error(ctx, err)
      //    r.Exit()
      //  }
      //  for {
      //    _, msg, err := ws.ReadMessage()
      //    if err != nil {
      //      return
      //    }
      //    fmt.Println(string(msg))
      //
      //    data := base_model.HookModel{}
      //    // 解析订阅数据包
      //    err = gjson.DecodeTo(msg, &data)
      //    if err != nil {
      //      fmt.Println(err)
      //      continue
      //    }
      //
      //    data.Ctx = r.Context()
      //    data.Source = &base_model.HookHostInfo{
      //      Host: r.Host,
      //      Port: 0,
      //    }
      //    hook.Gateway().BroadcastMessage(data)
      //    ws.WriteMessage(websocket.TextMessage, []byte("hello word"))
      //  }
      //})
      s.Run()
      return nil
    },
  }
)
