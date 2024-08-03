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

      // 注册跨进程通信的 websocket 服务，用于接收消息
      wsPath := g.Cfg().MustGet(context.Background(), "service.wsPath", "/ws").String()
      s.BindHandler(wsPath, base_hook.HookDistribution)

      s.Run()
      return nil
    },
  }
)
